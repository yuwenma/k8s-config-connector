// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	"k8s.io/client-go/util/homedir"
	"k8s.io/klog"
)

func main() {
	timeout := 50
	ctx := context.Background()
	if timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()
	}

	if err := Compile(ctx); err != nil {
		fmt.Printf("ERROR %s", err.Error())
	}
	fmt.Println("Done")

}

func Compile(ctx context.Context) error {
	path := filepath.Join(homedir.HomeDir(), "go/src/github.com/GoogleCloudPlatform/k8s-config-connector/")
	cmd := exec.CommandContext(ctx, "go", "run", "./cmd/manager/main.go")
	cmd.Dir = path

	_, compileErr := cmd.CombinedOutput()
	// geminiOut := &bytes.Buffer{}

	if compileErr != nil {
		fmt.Printf("Should fix the compile error: %s\n", compileErr)
		/*
			if err := RunGemini(ctx, compileErr.Error(), geminiOut); err != nil {
				return err
			}
		*/
	} else {
		klog.Info("Pass compile check")
	}
	return nil
}

/*
// RunGemini runs a prompt against Gemini, generating context based on the source code.
func RunGemini(ctx context.Context, input string, out io.Writer) error {
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return fmt.Errorf("building gemini client: %w", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-pro-002")

	// Some values that are recommended by aistudio
	model.SetTemperature(1)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "text/plain"

	var parts []genai.Part
	parts = append(parts, genai.Text("how to fix the error"))
	parts = append(parts, genai.Text(input))

	fmt.Println("Running gemini....")

	resp, err := model.GenerateContent(ctx, parts...)
	if err != nil {
		return fmt.Errorf("generating content with gemini: %w", err)
	}

	// Print the usage metadata (includes token count i.e. cost)
	fmt.Printf(": %+v", resp.UsageMetadata)

	for _, candidate := range resp.Candidates {
		content := candidate.Content

		for _, part := range content.Parts {
			if text, ok := part.(genai.Text); ok {
				fmt.Printf("TEXT: %+v", text)
				out.Write([]byte(text))
			} else {
				fmt.Printf("UNKNOWN: %T %+v", part, part)
			}
		}
	}
	return nil
}
*/
