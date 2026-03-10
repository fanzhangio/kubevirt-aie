/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright The KubeVirt Authors.
 *
 */

package admitters

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type graceVirtualizationConfig struct {
	HostDevices        *bool `json:"hostDevices,omitempty"`
	SMMUv3             *bool `json:"smmuv3,omitempty"`
	VCMDQ              *bool `json:"vcmdq,omitempty"`
	EGM                *bool `json:"egm,omitempty"`
	NUMAStrictLocality *bool `json:"numaStrictLocality,omitempty"`
}

func parseGraceVirtualizationConfig(raw string) (*graceVirtualizationConfig, error) {
	decoder := json.NewDecoder(strings.NewReader(raw))
	decoder.DisallowUnknownFields()

	cfg := &graceVirtualizationConfig{}
	if err := decoder.Decode(cfg); err != nil {
		return nil, err
	}

	var trailing struct{}
	if err := decoder.Decode(&trailing); err != io.EOF {
		return nil, fmt.Errorf("invalid trailing content")
	}

	return cfg, nil
}

func isEnabled(value *bool) bool {
	return value != nil && *value
}
