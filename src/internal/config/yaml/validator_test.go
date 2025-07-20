/*
 * © 2025-2025 JDHeim.com
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package yaml

import (
	"os"
	"path/filepath"
	"testing"
)

func TestValidateShortcutCommandWithInvalid(t *testing.T) {
	err := validateShortcutCommand(&shortcut{Command: "invalid"})
	if err == nil {
		t.Error("Validation Error expected")
	}
}

func TestValidateShortcutCommandWithEmpty(t *testing.T) {
	err := validateShortcutCommand(&shortcut{Command: ""})
	if err != nil {
		t.Errorf("Validation Error not expected, got %v", err)
	}
}

func TestValidateShortcutCommandWithExecOnPath(t *testing.T) {
	err := validateShortcutCommand(&shortcut{Command: "echo"})
	if err != nil {
		t.Errorf("Validation Error not expected, got %v", err)
	}
}

func TestValidateShortcutCommandWithPathToExec(t *testing.T) {
	err := validateShortcutCommand(&shortcut{Command: "/usr/bin/echo"})
	if err != nil {
		t.Errorf("Validation Error not expected, got %v", err)
	}
}

func TestValidateShortcutCommandWithSpaces(t *testing.T) {
	pathWithSpace := filepath.Join(t.TempDir(), "test with spaces")
	if err := os.Mkdir(pathWithSpace, 0755); err != nil {
		t.Fatalf("Failed to create %s: %v", pathWithSpace, err)
	}

	execPath := filepath.Join(pathWithSpace, "echo")
	if err := os.WriteFile(execPath, []byte("#!/bin/sh\necho hi\n"), 0755); err != nil {
		t.Fatalf("Failed to write %s: %v", execPath, err)
	}

	err := validateShortcutCommand(&shortcut{Command: execPath})
	if err != nil {
		t.Errorf("Validation Error not expected, got %v", err)
	}
}
