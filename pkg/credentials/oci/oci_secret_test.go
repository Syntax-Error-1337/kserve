/*
Copyright 2021 The KServe Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package oci

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestOCISecret(t *testing.T) {
	scenarios := map[string]struct {
		secret              *v1.Secret
		expectedVolume      v1.Volume
		expectedVolumeMount v1.VolumeMount
	}{
		"OCISecretVolume": {
			secret: &v1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name: "user-oci-sa",
				},
				Data: map[string][]byte{
					OCICredentialFileName: {},
				},
			},
			expectedVolumeMount: v1.VolumeMount{
				Name:      OCICredentialVolumeName,
				ReadOnly:  true,
				MountPath: OCICredentialVolumeMountPath,
			},
			expectedVolume: v1.Volume{
				Name: OCICredentialVolumeName,
				VolumeSource: v1.VolumeSource{
					Secret: &v1.SecretVolumeSource{
						SecretName: "user-oci-sa",
					},
				},
			},
		},
	}

	for name, scenario := range scenarios {
		volume, volumeMount := BuildSecretVolume(scenario.secret)

		if diff := cmp.Diff(scenario.expectedVolume, volume); diff != "" {
			t.Errorf("Test %q unexpected volume (-want +got):\n%s", name, diff)
		}

		if diff := cmp.Diff(scenario.expectedVolumeMount, volumeMount); diff != "" {
			t.Errorf("Test %q unexpected volumeMount (-want +got):\n%s", name, diff)
		}
	}
}
