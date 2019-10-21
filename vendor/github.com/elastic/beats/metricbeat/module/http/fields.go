// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package http

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "http", asset.ModuleFieldsPri, AssetHttp); err != nil {
		panic(err)
	}
}

// AssetHttp returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/http.
func AssetHttp() string {
	return "eJzMlDFv6jAUhff8iqPMD+kNTBne/KanN7BVHdz4AIbETn1vaPn3lYEgkhokSofe0Tbf+Y6MM8OW+wpr1a4A1GnDCuXfxeJ/WQCWUkfXqQu+wp8CANIW2mD7hgUQ2dAIK6xMAQhVnV9JhadSpCl/oUzg8rkAlo6NlerAmMGblufUNLrvEiWGfljJZI8pl6TI156i5/Uc8Cp0mEO1EwnOL0NsTTp5cWyaP2pDYxllAj16hJcNa51s3ZQBFmselU5gCL1mmksXvPBbqh9RP7F7ZE23o82m18EyG73l/i1E+8VsUaO95Oh8N22XXsr89zxr1K2jkXuczsR/QbEMvX9QeiIweG1kdKV3/1PS79FSo6uFl3c6/hAMc+25CuOO8RGRT4TbAh8BAAD//zeEPwk="
}
