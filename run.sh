#!/bin/sh
# Copyright 2020 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


# Replace the value here with the full path and filename of your Service Account
# JSON file that would have been previously downloaded from GCP.
SERVICE_ACCOUNT_LOCATION="$HOME/path/to/local/service/account.json"

docker build . \
    --tag tenantjwt \
    --quiet

docker run \
    -it \
    --rm \
    -v "$SERVICE_ACCOUNT_LOCATION:/sa.json:ro" \
    tenantjwt --service-account /sa.json
