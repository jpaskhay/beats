// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by x-pack/dev-tools/cmd/buildspec/buildspec.go - DO NOT EDIT.

package program

import (
	"strings"

	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/packer"
)

var Supported []Spec
var SupportedMap map[string]Spec

func init() {
	// Packed Files
	// spec/apm-server.yml
	// spec/endpoint.yml
	// spec/filebeat.yml
	// spec/fleet-server.yml
	// spec/heartbeat.yml
	// spec/metricbeat.yml
	// spec/osquerybeat.yml
	// spec/packetbeat.yml
	unpacked := packer.MustUnpack("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")
	SupportedMap = make(map[string]Spec)

	for f, v := range unpacked {
		s, err := NewSpecFromBytes(v)
		if err != nil {
			panic("Cannot read spec from " + f)
		}
		Supported = append(Supported, s)
		SupportedMap[strings.ToLower(s.Cmd)] = s
	}
}
