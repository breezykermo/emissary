// Copyright 2020 Datawire.  All rights reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

///////////////////////////////////////////////////////////////////////////
// Important: Run "make generate-fast" to regenerate code after modifying
// this file.
///////////////////////////////////////////////////////////////////////////

package v3alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AdditionalLogHeaders struct {
	HeaderName     string `json:"header_name,omitempty"`
	DuringRequest  *bool  `json:"during_request,omitempty"`
	DuringResponse *bool  `json:"during_response,omitempty"`
	DuringTrailer  *bool  `json:"during_trailer,omitempty"`
}

type DriverConfig struct {
	AdditionalLogHeaders []*AdditionalLogHeaders `json:"additional_log_headers,omitempty"`
}

// LogServiceSpec defines the desired state of LogService
type LogServiceSpec struct {
	AmbassadorID AmbassadorID `json:"ambassador_id,omitempty"`

	Service string `json:"service,omitempty"`

	// ProtocolVersion is the envoy api transport protocol version
	//
	// +kubebuilder:validation:Enum={"v2","v3"}
	ProtocolVersion string `json:"protocol_version,omitempty"`

	// +kubebuilder:validation:Enum={"tcp","http"}
	Driver                string          `json:"driver,omitempty"`
	DriverConfig          *DriverConfig   `json:"driver_config,omitempty"`
	FlushIntervalTime     *SecondDuration `json:"flush_interval_time,omitempty"`
	FlushIntervalByteSize *int            `json:"flush_interval_byte_size,omitempty"`

	// TODO(lukeshu): In v3alpha2, drop this LogService.spec.grpc.  Due to sloppy implementation
	// it is required to be present, and required to be 'true'.  It is silly to have a required
	// field with only one valid value, we should just remove the thing.
	GRPC *bool `json:"grpc,omitempty"`

	StatsName string `json:"stats_name,omitempty"`
}

// LogService is the Schema for the logservices API
//
// +kubebuilder:object:root=true
type LogService struct {
	metav1.TypeMeta   `json:""`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec LogServiceSpec `json:"spec,omitempty"`
}

// LogServiceList contains a list of LogServices.
//
// +kubebuilder:object:root=true
type LogServiceList struct {
	metav1.TypeMeta `json:""`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LogService{}, &LogServiceList{})
}
