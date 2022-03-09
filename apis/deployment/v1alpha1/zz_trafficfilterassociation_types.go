/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TrafficFilterAssociationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TrafficFilterAssociationParameters struct {

	// Required deployment ID where the traffic filter will be associated
	// +kubebuilder:validation:Required
	DeploymentID *string `json:"deploymentId" tf:"deployment_id,omitempty"`

	// Required traffic filter ruleset ID to tie to a deployment
	// +kubebuilder:validation:Required
	TrafficFilterID *string `json:"trafficFilterId" tf:"traffic_filter_id,omitempty"`
}

// TrafficFilterAssociationSpec defines the desired state of TrafficFilterAssociation
type TrafficFilterAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrafficFilterAssociationParameters `json:"forProvider"`
}

// TrafficFilterAssociationStatus defines the observed state of TrafficFilterAssociation.
type TrafficFilterAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrafficFilterAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficFilterAssociation is the Schema for the TrafficFilterAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,elasticcloudjet}
type TrafficFilterAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrafficFilterAssociationSpec   `json:"spec"`
	Status            TrafficFilterAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficFilterAssociationList contains a list of TrafficFilterAssociations
type TrafficFilterAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficFilterAssociation `json:"items"`
}

// Repository type metadata.
var (
	TrafficFilterAssociation_Kind             = "TrafficFilterAssociation"
	TrafficFilterAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TrafficFilterAssociation_Kind}.String()
	TrafficFilterAssociation_KindAPIVersion   = TrafficFilterAssociation_Kind + "." + CRDGroupVersion.String()
	TrafficFilterAssociation_GroupVersionKind = CRDGroupVersion.WithKind(TrafficFilterAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&TrafficFilterAssociation{}, &TrafficFilterAssociationList{})
}
