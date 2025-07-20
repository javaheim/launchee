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

package frontend

import "github.com/javaheim/launchee/build"

type UI struct {
	Nav     *Nav
	Content *Content
}

type Nav struct {
	Title      string
	AppIcon    *Icon
	IconSize   int
	IconUrl    string
	MenuHeight int
}

type Content struct {
	IconColumns int
	IconsPerRow int
	IconSize    int
	Margin      int
}

const (
	defaultTitle          = "Launchee"
	defaultAppIconSize    = 23
	defaultAppIconUrl     = "https://launchee.javaheim.com"
	defaultMenuHeight     = 8
	defaultIconColumns    = 1
	defaultMinIconsPerRow = 5
	defaultMaxIconsPerRow = 20
	defaultIconSize       = 8
	defaultMargin         = 5
	spacingScale          = 4
)

func NewUI(programCount int) *UI {
	return &UI{
		Nav:     NewNav(),
		Content: NewContent(programCount),
	}
}

func NewNav() *Nav {
	return &Nav{
		Title:      defaultTitle,
		AppIcon:    NewIconWithBytes(build.AppIconBytes),
		IconSize:   defaultAppIconSize,
		IconUrl:    defaultAppIconUrl,
		MenuHeight: defaultMenuHeight,
	}
}

func NewContent(programCount int) *Content {
	iconColumns, iconsPerRow := determineIconLayout(programCount)
	return &Content{
		IconColumns: iconColumns,
		IconsPerRow: iconsPerRow,
		IconSize:    defaultIconSize * spacingScale,
		Margin:      defaultMargin,
	}
}

func determineIconLayout(programCount int) (int, int) {
	iconColumns, iconsPerRow := programCount, programCount
	if programCount == 0 {
		iconColumns = defaultIconColumns
		iconsPerRow = defaultMinIconsPerRow
	} else if programCount < defaultMinIconsPerRow {
		iconColumns = programCount
		iconsPerRow = defaultMinIconsPerRow
	} else if programCount > defaultMaxIconsPerRow {
		iconColumns = defaultMaxIconsPerRow
		iconsPerRow = defaultMaxIconsPerRow
	}
	return iconColumns, iconsPerRow
}

func (u *UI) Width() int {
	return u.Content.Margin*spacingScale*2 +
		u.Content.IconSize*u.Content.IconsPerRow +
		u.Content.Margin*spacingScale*(u.Content.IconsPerRow-1)
}

func (u *UI) Height(programCount int) int {
	programCount = max(defaultMinIconsPerRow, programCount)
	rows := (programCount + u.Content.IconsPerRow - 1) / u.Content.IconsPerRow
	return u.Content.Margin*spacingScale*(rows+1) +
		u.Content.IconSize*rows +
		u.Nav.MenuHeight*spacingScale
}
