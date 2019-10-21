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

package haproxy

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "haproxy", asset.ModuleFieldsPri, AssetHaproxy); err != nil {
		panic(err)
	}
}

// AssetHaproxy returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/haproxy.
func AssetHaproxy() string {
	return "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"
}
