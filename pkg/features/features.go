// Copyright (c) 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package features

import (
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/component-base/featuregate"
)

const (
	// NonPrivilegedCalicoNode runs the long-lived calico-node container in non-privileged and non-root mode.
	// owner @ialidzhikov
	// alpha: v1.27.0
	NonPrivilegedCalicoNode featuregate.Feature = "NonPrivilegedCalicoNode"
)

var (
	// FeatureGate is a shared global FeatureGate for networking-calico extension flags.
	FeatureGate  = featuregate.NewFeatureGate()
	featureGates = map[featuregate.Feature]featuregate.FeatureSpec{
		NonPrivilegedCalicoNode: {Default: false, PreRelease: featuregate.Alpha},
	}
)

// RegisterFeatureGates registers the feature gates of the networking-calico extension.
func RegisterFeatureGates() {
	runtime.Must(FeatureGate.Add(featureGates))
}
