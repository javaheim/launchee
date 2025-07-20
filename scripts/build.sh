#!/bin/bash
# BUILDS LAUNCHEE

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
Usage: $(basename "$0") [OPTION]

Builds Launchee

OPTIONS:
  -d                     Dry-run JReleaser release
EOF
  exit 1
}

main() {
  cd ..
  readOptions "$@"
  local appVersion
  appVersion="$(getAppVersion)"
  scripts/common/updateVersion.sh "${appVersion}"
  scripts/common/updateCopyright.sh
  cd src
  build "${appVersion}"
}

readOptions() {
  while getopts ":hd" option; do
    case "${option}" in
      d) dryRunJReleaserRelease;;
      h|?) usage;;
    esac
  done
}

dryRunJReleaserRelease() {
  step "Dry-run JReleaser release"
  if [[ -z "${GITHUB_TOKEN-}" ]]; then
    echo -e "${ERROR} The GITHUB_TOKEN env variable is not set"
    exit 1
  fi
  JRELEASER_GITHUB_TOKEN=${GITHUB_TOKEN-} run jreleaser release --dry-run --output-directory=target
  exit $?
}

build() {
  local appVersion="${1}"
  step "Build Launchee ${appVersion}"
  run wails build -clean \
    -tags "webkit2_41" \
    -ldflags "-X github.com/javaheim/launchee/cmd.appVersion=${appVersion}"
}

main "$@"
