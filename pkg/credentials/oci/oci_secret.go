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
	v1 "k8s.io/api/core/v1"
)

const (
	OCICredentialFileName        = "oci-api-key.pem"             // #nosec G101, File name for OCI credentials
	OCICredentialVolumeName      = "user-oci-sa"                 // #nosec G101, Volume name for OCI credentials
	OCICredentialVolumeMountPath = "/var/secrets/oci"            // #nosec G101, Mount path for OCI credentials
	OCICredentialEnvKey          = "OCI_APPLICATION_CREDENTIALS" // #nosec G101, Environment variable key for OCI credentials.
)

// OCIConfig holds configuration for accessing OCI resources
type OCIConfig struct {
	OCICredentialFileName string `json:"ociCredentialFileName,omitempty"`
}

// BuildSecretVolume constructs a Kubernetes Volume and VolumeMount for OCI credentials based on a Kubernetes Secret
func BuildSecretVolume(secret *v1.Secret) (v1.Volume, v1.VolumeMount) {
	volume := v1.Volume{
		Name: OCICredentialVolumeName,
		VolumeSource: v1.VolumeSource{
			Secret: &v1.SecretVolumeSource{
				SecretName: secret.Name,
			},
		},
	}
	volumeMount := v1.VolumeMount{
		MountPath: OCICredentialVolumeMountPath,
		Name:      OCICredentialVolumeName,
		ReadOnly:  true,
	}
	return volume, volumeMount
}
