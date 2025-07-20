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
	"strings"

	"github.com/javaheim/launchee/internal/config/frontend"
)

type config struct {
	Title    string
	Programs []*program
}

type program struct {
	Name    string
	Icon    string
	Command string
	Url     string
	Patch   string `yaml:"$patch"`
}

func newConfigWithoutPrograms(title string) *config {
	return &config{
		Title: title,
	}
}

func (yc *config) toFrontendConfig() *frontend.Config {
	frontendConfig := frontend.NewConfig(len(yc.Programs))
	yc.overrideFrontendConfig(frontendConfig)
	return frontendConfig
}

func (yc *config) overrideFrontendConfig(config *frontend.Config) {
	if yc.Title != "" {
		config.UI.Nav.Title = yc.Title
	}
	if frontendPrograms := yc.toFrontendPrograms(); frontendPrograms != nil {
		config.Programs = frontendPrograms
	}
}

func (yc *config) toFrontendPrograms() []*frontend.Program {
	programCount := len(yc.Programs)
	if programCount == 0 {
		return nil
	}
	frontendPrograms := make([]*frontend.Program, programCount)
	for i := range yc.Programs {
		frontendPrograms[i] = yc.toFrontendProgram(i)
	}
	return frontendPrograms
}

func (yc *config) toFrontendProgram(i int) *frontend.Program {
	return &frontend.Program{
		Id:      i,
		Name:    yc.Programs[i].Name,
		Icon:    frontend.NewIcon(yc.Programs[i].Icon),
		Command: strings.Fields(yc.Programs[i].Command),
		Url:     yc.Programs[i].Url,
	}
}

func (p *program) isPatchMode() bool {
	return p.Patch != "" && (p.Patch == patchDelete || p.Patch == patchMerge)
}
