// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"cloud.google.com/go/vertexai/genai"
	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codebot"
	"k8s.io/klog"

	"context"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

type Options struct {
	BaseDir       string
	Group         string
	ProtoResource string
	Retries       int
}

func run(ctx context.Context) error {
	var o Options

	klog.InitFlags(nil)
	flag.StringVar(&o.BaseDir, "base-dir", o.BaseDir, "base directory for the project code")
	flag.StringVar(&o.Group, "group", o.Group, "the Config Connector service name")
	flag.StringVar(&o.ProtoResource, "proto", o.ProtoResource, "the proto name")
	flag.IntVar(&o.Retries, "retries", o.Retries, "the max number of attempts to rerun the prediction for better result")

	flag.Parse()

	if o.BaseDir == "" {
		return fmt.Errorf("base-dir is required")
	}
	if o.Group == "" {
		return fmt.Errorf("group is required")
	}
	fpath := filepath.Join("./pkg/controller/direct/", o.Group, o.ProtoResource+"_controller.go")

	includeDirs := []string{
		"./apis/bigqueryconnection",
		"./apis/cloudbuild",
		"./apis/kms",
		"./apis/" + o.Group,
		"./pkg/controller/direct/common",
		"./pkg/controller/direct/bigqueryconnection",
		"./pkg/controller/direct/cloudbuild",
		"./pkg/controller/direct/kms",
		fpath,
		"./pkg/controller/direct/directbase",
		"./pkg/controller/direct/registry",
	}
	contextFiles, err := readGoSourceFromSubdirs(o.BaseDir, includeDirs)
	if err != nil {
		return fmt.Errorf("read Go source directory %s: %w", o.BaseDir, err)
	}

	// Construct the prompt
	prompt := fmt.Sprintf("File %q:\n", fpath)

	for range o.Retries {
		content := generateContent(ctx, o.BaseDir, prompt, contextFiles)

	}
}

func readGoSourceFromSubdirs(rootDir string, subdirs []string) (map[string]*codebot.FileInfo, error) {
	contextFiles := make(map[string]*codebot.FileInfo)

	for _, subdir := range subdirs {
		// Construct the full path to the subdirectory
		fullSubdirPath := filepath.Join(rootDir, subdir)

		// Walk through the subdirectory
		err := filepath.WalkDir(fullSubdirPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !d.IsDir() && strings.HasSuffix(path, ".go") {
				content, err := os.ReadFile(path)
				if err != nil {
					return err
				}
				contextFiles[path] = &codebot.FileInfo{
					Path:    path,
					Content: string(content),
				}
			}
			return nil
		})

		if err != nil {
			return nil, err
		}
	}

	return contextFiles, nil
}

// Placeholder function for interacting with the Gemini API
func generateContent(ctx context.Context, basedir, prompt string, contextFiles map[string]*codebot.FileInfo) error {
	funcDecl := []*genai.FunctionDeclaration{}
	chat, err := codebot.NewChat(ctx, basedir, contextFiles, funcDecl)
	if err != nil {
		return err
	}
	defer chat.Close()

	if err := chat.SendMessage(ctx, genai.Text(prompt)); err != nil {
		return err
	}
	return nil
}
