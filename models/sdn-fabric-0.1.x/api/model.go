// Code generated by model-compiler. DO NOT EDIT.

// SPDX-FileCopyrightText: 2022-present Intel Corporation
// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"github.com/openconfig/gnmi/proto/gnmi"
)

var modelData = []*gnmi.ModelData{
	{Name: "onf-switch", Organization: "Open Networking Foundation", Version: "2022-05-25"},
	{Name: "onf-switch-model", Organization: "Open Networking Foundation", Version: "2022-05-25"},
	{Name: "onf-route", Organization: "Open Networking Foundation", Version: "2022-05-25"},
	{Name: "onf-dhcp-server", Organization: "Open Networking Foundation", Version: "2022-05-25"},
}

var encodings = []gnmi.Encoding{gnmi.Encoding_JSON_IETF}

func ModelData() []*gnmi.ModelData {
	return modelData
}

func Encodings() []gnmi.Encoding {
	return encodings
}
