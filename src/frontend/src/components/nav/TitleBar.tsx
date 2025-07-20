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

import type {CSSProperties} from "react";
import {ChevronDown, X} from "lucide-react";
import {Quit, WindowMinimise} from "../../../wailsjs/runtime";
import {AppIconWithTooltip} from "@/components/nav/AppIconWithTooltip.tsx";
import {frontend} from "../../../wailsjs/go/models.ts";

export function TitleBar({nav}: Readonly<{ nav: frontend.Nav | null }>) {
    const defaultAppIconSize = 23;
    const defaultAppIconUrl = "https://launchee.javaheim.com";
    const defaultMenuHeight = 8;

    const appIconSize = nav?.IconSize ?? defaultAppIconSize;
    const appIconUrl = nav?.IconUrl ?? defaultAppIconUrl;
    const menuHeightClass = `h-${nav?.MenuHeight ?? defaultMenuHeight}`;

    return (
        <div className={`flex flex-row justify-between items-center-safe ${menuHeightClass} bg-[#1e1f22]`}
             style={{"--wails-draggable": "drag"} as CSSProperties}>
            <div className="flex flex-row mx-1 ml-1.5" style={{"--wails-draggable": "no-drag"} as CSSProperties}>
                <AppIconWithTooltip appIcon={nav?.AppIcon ?? null}
                                    iconSize={appIconSize}
                                    url={appIconUrl}/>
            </div>
            <div className="absolute left-1/2 transform -translate-x-1/2 text-gray-200 text-[13px] truncate max-w-[calc(100%-(20px+20px+5px)*2-4px-4px)]">
                {nav?.Title && (
                    <span>{nav.Title}</span>
                )}
            </div>
            <div className="flex flex-row mx-1 gap-1" style={{"--wails-draggable": "no-drag"} as CSSProperties}>
                <ChevronDown className="size-5 text-gray-400 hover:text-gray-100" onClick={() => WindowMinimise()}/>
                <X className="size-5 text-red-700 hover:text-red-500" onClick={() => Quit()}/>
            </div>
        </div>
    );
}
