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

package dropwizard

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "dropwizard", asset.ModuleFieldsPri, AssetDropwizard); err != nil {
		panic(err)
	}
}

// AssetDropwizard returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/dropwizard.
func AssetDropwizard() string {
	return "eJxsjk2KAjEQhfc5xSPr6TlAFrOaG7gUkZBUuoPpJFSVSHt6aRX/sJaveO/7BhxocYjc+imfPUcDaNZCDvb/EVoDRJLAuWtu1eHPAMBGvQpCK4WCUkTiNuPZ+jUAUyEv5DB6Awip5jqKw9aKFPsDO6l2u1t/U2Pdh1ZTHh2SL0IGSJlKFHfFDah+pg/Z9XTpK4Hbsd+TL7LvW7e9V7tLAAAA///zhlJc"
}
