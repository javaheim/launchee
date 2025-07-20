#!/bin/bash
# UPDATES C0PYRIGHT

#
# © 2025-2025 Javaheim
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

[[ -f "$(dirname "${BASH_SOURCE[0]}")/functions.sh" ]] && . "$(dirname "${BASH_SOURCE[0]}")/functions.sh"

readonly COPYRIGHTS_START_YEAR="2025"

main() {
  step "Update copyright"
  updateCopyrightInFile "README.md"
  updateCopyrightInFile "LICENSE"
  updateCopyrightInFile "src/wails.json"
  if [[ "${isUpdated:-false}" == false ]]; then
    echo -e "${WARNING} Nothing to update"
  fi
}

updateCopyrightInFile() {
  local name="${1}"
  sed -i "s/\(©\).*\(Javaheim\)/\1 ${COPYRIGHTS_START_YEAR}-$(date +%Y) \2/" "${name}"
}

findAndUpdateCopyright() {
  local name="${1}"
  find . -name "${name}" -type f -not -path "./**/target/*" -exec sed -i "s/\(©\).*\(Javaheim\)/\1 ${COPYRIGHTS_START_YEAR}-$(date +%Y) \2/" {} +
}

main "$@"
