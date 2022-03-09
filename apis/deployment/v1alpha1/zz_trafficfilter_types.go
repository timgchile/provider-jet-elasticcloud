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

type RuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RuleParameters struct {

	// Optional Azure endpoint GUID
	// +kubebuilder:validation:Optional
	AzureEndpointGUID *string `json:"azureEndpointGuid,omitempty" tf:"azure_endpoint_guid,omitempty"`

	// Optional Azure endpoint name
	// +kubebuilder:validation:Optional
	AzureEndpointName *string `json:"azureEndpointName,omitempty" tf:"azure_endpoint_name,omitempty"`

	// Optional rule description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Required traffic filter source: IP address, CIDR mask, or VPC endpoint ID, not required when the type is azure_private_endpoint
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type TrafficFilterObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TrafficFilterParameters struct {

	// Optional ruleset description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Should the ruleset be automatically included in the new deployments (Defaults to false)
	// +kubebuilder:validation:Optional
	IncludeByDefault *bool `json:"includeByDefault,omitempty" tf:"include_by_default,omitempty"`

	// Required filter region, the ruleset can only be attached to deployments in the specific region
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// Required list of rules, which the ruleset is made of.
	// +kubebuilder:validation:Required
	Rule []RuleParameters `json:"rule" tf:"rule,omitempty"`

	// Required type of the ruleset ("ip", "vpce" or "azure_private_endpoint")
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// TrafficFilterSpec defines the desired state of TrafficFilter
type TrafficFilterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrafficFilterParameters `json:"forProvider"`
}

// TrafficFilterStatus defines the observed state of TrafficFilter.
type TrafficFilterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrafficFilterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficFilter is the Schema for the TrafficFilters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,elasticcloudjet}
type TrafficFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrafficFilterSpec   `json:"spec"`
	Status            TrafficFilterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrafficFilterList contains a list of TrafficFilters
type TrafficFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficFilter `json:"items"`
}

// Repository type metadata.
var (
	TrafficFilter_Kind             = "TrafficFilter"
	TrafficFilter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TrafficFilter_Kind}.String()
	TrafficFilter_KindAPIVersion   = TrafficFilter_Kind + "." + CRDGroupVersion.String()
	TrafficFilter_GroupVersionKind = CRDGroupVersion.WithKind(TrafficFilter_Kind)
)

func init() {
	SchemeBuilder.Register(&TrafficFilter{}, &TrafficFilterList{})
}
