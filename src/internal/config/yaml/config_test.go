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

import "testing"

func TestParseCommandWithEmpty(t *testing.T) {
	shortcut := &shortcut{CommandArgs: ""}
	commandArgs := shortcut.parseCommandArgs()
	if len(commandArgs) != 0 {
		t.Errorf("Parsed Command Args not expected, got %s", commandArgs)
	}
}

func TestParseCommandWithSpacesAndArgs(t *testing.T) {
	shortcut := &shortcut{CommandArgs: "-d /tmp/dummy -f foo.txt"}
	commandArgs := shortcut.parseCommandArgs()
	if len(commandArgs) != 4 {
		t.Errorf("Parsed Command Args expected, got %s", commandArgs)
	}
	if commandArgs[0] != "-d" {
		t.Errorf("Parsed Command Arg mismatch, got %s", commandArgs[0])
	}
	if commandArgs[1] != "/tmp/dummy" {
		t.Errorf("Parsed Command Arg mismatch, got %s", commandArgs[1])
	}
	if commandArgs[2] != "-f" {
		t.Errorf("Parsed Command Arg mismatch, got %s", commandArgs[2])
	}
	if commandArgs[3] != "foo.txt" {
		t.Errorf("Parsed Command Arg mismatch, got %s", commandArgs[3])
	}
}
