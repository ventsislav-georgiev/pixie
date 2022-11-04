#!/usr/bin/env bash

# Copyright 2018- The Pixie Authors.
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
#
# SPDX-License-Identifier: Apache-2.0

set -ex

printenv

repo_path=$(pwd)
release_tag=custom.2
versions_file="$(pwd)/src/utils/artifacts/artifact_db_updater/VERSIONS.json"

echo "The release tag is: ${release_tag}"

bazel run -c opt //src/utils/artifacts/versions_gen:versions_gen -- \
      --repo_path "${repo_path}" --artifact_name vizier --versions_file "${versions_file}"

public="True"
extra_bazel_args=()

bazel run --stamp -c opt --define BUNDLE_VERSION="${release_tag}" \
    --stamp --define public="${public}" //k8s/vizier:vizier_images_push "${extra_bazel_args[@]}"
bazel build --stamp -c opt --define BUNDLE_VERSION="${release_tag}" \
    --stamp --define public="${public}" //k8s/vizier:vizier_yamls "${extra_bazel_args[@]}"

ls -la "${repo_path}/bazel-bin/k8s/vizier/"
