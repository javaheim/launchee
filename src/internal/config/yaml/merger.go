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
	"github.com/javaheim/launchee/internal/lctx"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const (
	patchDelete  = "delete"
	patchMerge   = "merge"
	patchReplace = "replace"
)

func (yc *config) sanitize() *config {
	sanitizedPrograms := make([]*program, 0, len(yc.Programs))
	processed := make(map[string]bool)
	for _, program := range yc.Programs {
		if !processed[program.Name] && program.Patch != patchDelete && program.Patch != patchMerge {
			processed[program.Name] = true
			sanitizedPrograms = append(sanitizedPrograms, program)
		}
	}
	yc.Programs = sanitizedPrograms
	return yc
}

func (yc *config) merge(other *config) *config {
	if other == nil {
		return yc
	}
	merged := newConfigWithoutPrograms(yc.Title)
	if other.Title != "" {
		merged.Title = other.Title
	}
	if len(other.Programs) != 0 {
		merged.Programs = yc.mergePrograms(other)
	} else {
		merged.Programs = yc.Programs
	}
	return merged
}

func (yc *config) mergePrograms(other *config) []*program {
	mergedPrograms := make([]*program, 0, len(yc.Programs)+len(other.Programs))
	otherPrograms := toProgramMapByName(other.Programs)
	processed := make(map[string]bool)

	runtime.LogInfo(lctx.GetContext(), "----- Configuration merge started -----")
	for _, program := range yc.Programs {
		if otherProgram, found := otherPrograms[program.Name]; found {
			processed[program.Name] = true
			if otherProgram.Patch == patchDelete {
				runtime.LogInfof(lctx.GetContext(), "Deleting %+v", otherProgram)
				continue
			} else if otherProgram.Patch == patchMerge {
				mergedProgram := program.merge(otherProgram)
				runtime.LogInfof(lctx.GetContext(), "Overridding with %+v", mergedProgram)
				mergedPrograms = append(mergedPrograms, mergedProgram)
				continue
			}
			runtime.LogInfof(lctx.GetContext(), "Replacing with %+v", otherProgram)
			mergedPrograms = append(mergedPrograms, otherProgram)
		} else {
			runtime.LogInfof(lctx.GetContext(), "Adding %+v", program)
			mergedPrograms = append(mergedPrograms, program)
		}
	}

	for _, otherProgram := range other.Programs {
		if !processed[otherProgram.Name] && otherProgram.Patch != patchDelete && otherProgram.Patch != patchMerge {
			runtime.LogInfof(lctx.GetContext(), "Adding not processed one %+v", otherProgram)
			processed[otherProgram.Name] = true
			mergedPrograms = append(mergedPrograms, otherProgram)
		}
	}
	runtime.LogInfo(lctx.GetContext(), "----- Configuration merge finished -----")
	return mergedPrograms
}

func toProgramMapByName(programs []*program) map[string]*program {
	programMap := make(map[string]*program)
	for _, program := range programs {
		programMap[program.Name] = program
	}
	return programMap
}

func (p *program) merge(other *program) *program {
	p.Patch = patchMerge
	if other.Icon != "" {
		p.Icon = other.Icon
	}
	if other.Command != "" {
		p.Command = other.Command
	} else if other.Url != "" {
		p.Url = other.Url
	}
	return p
}
