#!/usr/bin/env bash


# Copyright 2019 FinTech Studios, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")
PROJECT_DIR="$(cd -P -- "$(dirname -- "${SCRIPT_ROOT}/../..")" && pwd -P)"
SWAGGER_FILE=ververica-platform-swagger.json
SWAGGER_IGNORE_FILE=.swagger-codegen-ignore
CONFIG_FILE=swagger-gen-config.json
OUT_DIR=ververica-platform-api

rm -rf ${PROJECT_DIR}/${OUT_DIR}
mkdir ${PROJECT_DIR}/${OUT_DIR}
cp ${PROJECT_DIR}/${SWAGGER_IGNORE_FILE} ${PROJECT_DIR}/${OUT_DIR}

docker run --rm \
    -v ${PROJECT_DIR}:/local:rw \
    --user $(id -u):$(id -g) \
    swaggerapi/swagger-codegen-cli generate \
    -i /local/${SWAGGER_FILE} \
    -l go \
    -o /local/${OUT_DIR} \
    -c /local/${CONFIG_FILE}

