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

// Package v1alpha1 defines the v1alpha1 API types for the DataPipeline operator.
package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// ConditionStatus is the status of a DataPipeline condition.
type ConditionStatus string

const (
	ConditionTrue    ConditionStatus = "True"
	ConditionFalse   ConditionStatus = "False"
	ConditionUnknown ConditionStatus = "Unknown"
)

// DataPipelineConditionType is the type of a DataPipeline condition.
type DataPipelineConditionType string

const (
	// DataPipelineAvailable indicates the pipeline is fully operational.
	DataPipelineAvailable DataPipelineConditionType = "Available"
	// DataPipelineProgressing indicates an update is being rolled out.
	DataPipelineProgressing DataPipelineConditionType = "Progressing"
	// DataPipelineDegraded indicates the pipeline is running but not at full capacity.
	DataPipelineDegraded DataPipelineConditionType = "Degraded"
)

// DataPipelineCondition describes the state of a DataPipeline at a certain point.
type DataPipelineCondition struct {
	Type               DataPipelineConditionType `json:"type"`
	Status             ConditionStatus           `json:"status"`
	LastTransitionTime metav1.Time               `json:"lastTransitionTime,omitempty"`
	Reason             string                    `json:"reason,omitempty"`
	Message            string                    `json:"message,omitempty"`
}

// DataPipelineSpec defines the desired state of a DataPipeline.
type DataPipelineSpec struct {
	// Source is the URI of the upstream data source.
	Source string `json:"source"`
	// Sink is the URI of the downstream data sink.
	Sink string `json:"sink"`
	// Replicas is the number of pipeline worker replicas.
	// +kubebuilder:default=1
	Replicas int32 `json:"replicas,omitempty"`
}

// DataPipelineStatus defines the observed state of a DataPipeline.
type DataPipelineStatus struct {
	// Conditions holds the conditions that describe the current state of the pipeline.
	Conditions []DataPipelineCondition `json:"conditions,omitempty"`
	// ReadyReplicas is the number of pipeline workers currently ready.
	ReadyReplicas int32 `json:"readyReplicas,omitempty"`
	// ObservedGeneration is the most recent generation observed by the controller.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.readyReplicas"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +metrics:conditions:path=.status.conditions,resourceType=data_pipeline
type DataPipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataPipelineSpec   `json:"spec,omitempty"`
	Status DataPipelineStatus `json:"status,omitempty"`
}

// DataPipelineList contains a list of DataPipeline resources.
// +kubebuilder:object:root=true
type DataPipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataPipeline `json:"items"`
}
