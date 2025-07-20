/*
 * © 2025-2025 Javaheim
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
	"os/exec"
	"strings"
	"unicode/utf8"

	"github.com/pkg/errors"

	"github.com/javaheim/launchee/internal/config/frontend"
)

const maxIconSize = 1 << 20 // 1 MB

func validate(config *config) error {
	if err := validateTitle(config); err != nil {
		return err
	}
	if err := validatePrograms(config); err != nil {
		return err
	}
	return nil
}

func validateTitle(config *config) error {
	titleLength := utf8.RuneCountInString(config.Title)
	if titleLength != 0 && (titleLength < 3 || titleLength > 30) {
		return errors.Errorf("Title \"%s\" must be between 3 and 30 characters long (got %d)", config.Title, titleLength)
	}
	return nil
}

func validatePrograms(config *config) error {
	for _, program := range config.Programs {
		if err := validateProgram(program); err != nil {
			return err
		}
	}
	return nil
}

func validateProgram(program *program) error {
	if err := validateProgramName(program); err != nil {
		return err
	}
	if err := validateProgramPatch(program); err != nil {
		return err
	}
	if err := validateProgramIcon(program); err != nil {
		return err
	}
	if err := validateProgramCommandAndUrl(program); err != nil {
		return err
	}
	if err := validateProgramCommand(program); err != nil {
		return err
	}
	if err := validateProgramUrl(program); err != nil {
		return err
	}
	return nil
}

func validateProgramName(program *program) error {
	nameLength := utf8.RuneCountInString(program.Name)
	if nameLength < 3 || nameLength > 30 {
		return errors.Errorf("Name of \"%s\" Program must be between 3 and 30 characters long (got %d)", program.Name, nameLength)
	}
	return nil
}

func validateProgramPatch(program *program) error {
	if program.Patch != "" && program.Patch != patchDelete && program.Patch != patchMerge && program.Patch != patchReplace {
		return errors.Errorf("Patch of \"%s\" Program must be either \"%s\", \"%s\" or \"%s\" (got \"%s\")",
			program.Name, patchDelete, patchMerge, patchReplace, program.Patch)
	}
	return nil
}

func validateProgramIcon(program *program) error {
	if !program.isPatchMode() && program.Icon == "" {
		return errors.Errorf("Icon of \"%s\" Program must be set", program.Name)
	}
	if program.Icon != "" {
		if !fileExists(program.Icon) {
			return errors.Errorf("Icon of \"%s\" Program does not exist under: \"%s\"", program.Name, program.Icon)
		}
		if !frontend.IsValidIcon(program.Icon) {
			return errors.Errorf("Icon of \"%s\" Program is not a valid Icon: \"%s\". Supported extensions: %s",
				program.Name, program.Icon, frontend.SupportedExtensions())
		}
		if isFileLargerThan1MB(program.Icon) {
			return errors.Errorf("Icon of \"%s\" Program is larger than 1 MB", program.Name)
		}
	}
	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func isFileLargerThan1MB(path string) bool {
	info, _ := os.Stat(path)
	return info.Size() > maxIconSize
}

func validateProgramCommandAndUrl(program *program) error {
	if !program.isPatchMode() && program.Command == "" && program.Url == "" {
		return errors.Errorf("Either Command or URL of \"%s\" Program must be set", program.Name)
	}
	if program.Command != "" && program.Url != "" {
		return errors.Errorf("\"%s\" Program cannot have both a Command and a URL set - choose one (got Command: \"%s\" and URL: \"%s\")",
			program.Name, program.Command, program.Url)
	}
	return nil
}

func validateProgramCommand(program *program) error {
	if program.Command != "" && !isExec(program.Command) {
		return errors.Errorf("Command of \"%s\" Program is not a valid Command (got \"%s\")", program.Name, program.Command)
	}
	return nil
}

func isExec(command string) bool {
	commandParts := strings.Fields(command)
	if len(commandParts) == 0 {
		return false
	}
	_, err := exec.LookPath(commandParts[0])
	return err == nil
}

func validateProgramUrl(program *program) error {
	if program.Url != "" {
		programUrl := strings.ToLower(program.Url)
		if !strings.HasPrefix(programUrl, "https://") && !strings.HasPrefix(programUrl, "http://") {
			return errors.Errorf("URL of \"%s\" Program must start with \"http://\" or \"https://\" (got \"%s\")", program.Name, program.Url)
		}
	}
	return nil
}
