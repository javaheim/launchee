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
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"

	"github.com/pkg/errors"

	"github.com/javaheim/launchee/internal/config/frontend"
)

const (
	configFileTestSystem = "/home/dev/projects/launchee/src/test/config/launchee-7.yml"
	configFileTestUser   = "/home/dev/projects/launchee/src/test/config/launchee-override.yml"
	configFileDir        = "launchee"
	configFile           = "launchee.yml"
)

type unmarshalResult struct {
	config *config
	err    error
}

func UnmarshalConfigs() (*frontend.Config, error) {
	systemConfigResult, userConfigResult := unmarshalConfigsAsync()
	if systemConfigResult.err != nil {
		return frontend.NewConfig(0), systemConfigResult.err
	}
	if userConfigResult.err != nil {
		return frontend.NewConfig(0), userConfigResult.err
	}
	if systemConfigResult.config != nil {
		return systemConfigResult.config.sanitize().merge(userConfigResult.config).toFrontendConfig(), nil
	} else if userConfigResult.config != nil {
		return userConfigResult.config.sanitize().toFrontendConfig(), nil
	}
	return frontend.NewConfig(0), nil
}

func unmarshalConfigsAsync() (*unmarshalResult, *unmarshalResult) {
	var systemConfigResult *unmarshalResult
	var userConfigResult *unmarshalResult
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		systemConfigResult = unmarshalConfigFile(getSystemConfigPath())
	}()
	go func() {
		defer wg.Done()
		userConfigResult = unmarshalConfigFile(getUserConfigPath())
	}()
	wg.Wait()
	return systemConfigResult, userConfigResult
}

func unmarshalConfigFile(configFile string) *unmarshalResult {
	if bytes, err := os.ReadFile(configFile); err == nil {
		var config config
		if err = yaml.Unmarshal(bytes, &config); err != nil {
			return &unmarshalResult{nil, errors.WithMessagef(err, "Could not parse %s", configFile)}
		}
		return &unmarshalResult{&config, validate(config.trim())}
	}
	return &unmarshalResult{nil, nil}
}

func (c *config) trim() *config {
	c.Title = strings.TrimSpace(c.Title)
	for _, program := range c.Programs {
		program.Name = strings.TrimSpace(program.Name)
		program.Icon = strings.TrimSpace(program.Icon)
		program.Command = strings.TrimSpace(program.Command)
		program.Url = strings.TrimSpace(program.Url)
		program.Patch = strings.TrimSpace(program.Patch)
	}
	return c
}

func getSystemConfigPath() string {
	switch runtime.GOOS {
	case "windows":
		if systemConfigPath := os.Getenv("PROGRAMDATA"); systemConfigPath != "" {
			return filepath.Join(systemConfigPath, configFileDir, configFile)
		}
	case "linux":
		return filepath.Join("/etc", configFileDir, configFile)
	case "darwin":
		return filepath.Join("/Library/Application Support", configFileDir, configFile)
	}
	return ""
}

func getUserConfigPath() string {
	if userConfigPath, _ := os.UserConfigDir(); userConfigPath != "" {
		return filepath.Join(userConfigPath, configFileDir, configFile)
	}
	return ""
}
