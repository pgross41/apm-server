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

package package_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/apm-server/config"
	"github.com/elastic/apm-server/processor/sourcemap"
	"github.com/elastic/apm-server/tests"
	"github.com/elastic/apm-server/tests/loader"
)

// ensure all valid documents pass through the whole validation and transformation process
func TestSourcemapProcessorOK(t *testing.T) {
	requestInfo := []tests.RequestInfo{
		{Name: "TestProcessSourcemapFull", Path: "data/valid/sourcemap/payload.json"},
		{Name: "TestProcessSourcemapMinimalPayload", Path: "data/valid/sourcemap/minimal_payload.json"},
	}
	tests.TestProcessRequests(t, sourcemap.NewProcessor(), config.Config{}, requestInfo, map[string]string{"@timestamp": "***IGNORED***"})
}

// ensure invalid documents fail the json schema validation already
func TestSourcemapProcessorValidationFailed(t *testing.T) {
	data, err := loader.LoadInvalidData("sourcemap")
	assert.Nil(t, err)
	p := sourcemap.NewProcessor()
	err = p.Validate(data)
	assert.NotNil(t, err)
}
