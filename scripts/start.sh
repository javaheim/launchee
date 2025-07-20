#!/bin/bash
# STARTS LAUNCHEE

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

[[ -f "$(dirname "${BASH_SOURCE[0]}")/common/functions.sh" ]] && . "$(dirname "${BASH_SOURCE[0]}")/common/functions.sh"

usage() {
  cat << EOF
Usage: $(basename "$0")

Starts Launchee
EOF
  exit 1
}

main() {
  cd ..
  readOptions "$@"
  scripts/common/updateVersion.sh
  scripts/common/updateCopyright.sh
  local appVersion
  appVersion="$(getAppVersion)"
  cd src
  start "${appVersion}"
}

readOptions() {
  while getopts ":h" option; do
    case "${option}" in
      h|?) usage ;;
    esac
  done
}

start() {
  local appVersion="${1}"
  step "Start Launchee ${appVersion}"
  run wails dev -extensions "go,yml" \
    -tags "webkit2_41" \
    -ldflags "-X github.com/javaheim/launchee/cmd.appVersion=${appVersion}"
}

main "$@"
