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

package ui

import (
	"fmt"
	"strings"
)

var _ UI = &BashUI{}

type BashUI struct {
	callback     func(text string) error
	instructions string
}

func NewBashUI(instructions string) UI {
	return &BashUI{instructions: instructions}
}

func (u *BashUI) Run() error {
	lines := strings.Split(u.instructions, "\n\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		if err := u.callback(line); err != nil {
			return fmt.Errorf("error running callback: %w", err)
		}
	}
	return nil
}

func (u *BashUI) SetCallback(callback func(text string) error) {
	u.callback = callback
}

func (u *BashUI) AddLLMOutput(output *LLMOutput) {
	fmt.Printf("%v\n", output.Text)
}
