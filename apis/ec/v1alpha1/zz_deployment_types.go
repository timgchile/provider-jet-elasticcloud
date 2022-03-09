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

type ApmObservation struct {
	HTTPEndpoint *string `json:"httpEndpoint,omitempty" tf:"http_endpoint,omitempty"`

	HTTPSEndpoint *string `json:"httpsEndpoint,omitempty" tf:"https_endpoint,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`
}

type ApmParameters struct {

	// Optionally define the Apm configuration options for the APM Server
	// +kubebuilder:validation:Optional
	Config []ConfigParameters `json:"config,omitempty" tf:"config,omitempty"`

	// +kubebuilder:validation:Optional
	ElasticsearchClusterRefID *string `json:"elasticsearchClusterRefId,omitempty" tf:"elasticsearch_cluster_ref_id,omitempty"`

	// +kubebuilder:validation:Optional
	RefID *string `json:"refId,omitempty" tf:"ref_id,omitempty"`

	// +kubebuilder:validation:Optional
	Topology []TopologyParameters `json:"topology,omitempty" tf:"topology,omitempty"`
}

type AutoscalingObservation struct {
	PolicyOverrideJSON *string `json:"policyOverrideJson,omitempty" tf:"policy_override_json,omitempty"`
}

type AutoscalingParameters struct {

	// Maximum size value for the maximum autoscaling setting.
	// +kubebuilder:validation:Optional
	MaxSize *string `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	// Maximum resource type for the maximum autoscaling setting.
	// +kubebuilder:validation:Optional
	MaxSizeResource *string `json:"maxSizeResource,omitempty" tf:"max_size_resource,omitempty"`

	// Minimum size value for the minimum autoscaling setting.
	// +kubebuilder:validation:Optional
	MinSize *string `json:"minSize,omitempty" tf:"min_size,omitempty"`

	// Minimum resource type for the minimum autoscaling setting.
	// +kubebuilder:validation:Optional
	MinSizeResource *string `json:"minSizeResource,omitempty" tf:"min_size_resource,omitempty"`
}

type ConfigObservation struct {
}

type ConfigParameters struct {

	// Optionally enable debug mode for APM servers - defaults to false
	// +kubebuilder:validation:Optional
	DebugEnabled *bool `json:"debugEnabled,omitempty" tf:"debug_enabled,omitempty"`

	// Optionally override the docker image the APM nodes will use. Note that this field will only work for internal users only.
	// +kubebuilder:validation:Optional
	DockerImage *string `json:"dockerImage,omitempty" tf:"docker_image,omitempty"`

	// An arbitrary JSON object allowing (non-admin) cluster owners to set their parameters (only one of this and 'user_settings_yaml' is allowed), provided they are on the whitelist ('user_settings_whitelist') and not on the blacklist ('user_settings_blacklist'). (This field together with 'user_settings_override*' and 'system_settings' defines the total set of resource settings)
	// +kubebuilder:validation:Optional
	UserSettingsJSON *string `json:"userSettingsJson,omitempty" tf:"user_settings_json,omitempty"`

	// An arbitrary JSON object allowing ECE admins owners to set clusters' parameters (only one of this and 'user_settings_override_yaml' is allowed), ie in addition to the documented 'system_settings'. (This field together with 'system_settings' and 'user_settings*' defines the total set of resource settings)
	// +kubebuilder:validation:Optional
	UserSettingsOverrideJSON *string `json:"userSettingsOverrideJson,omitempty" tf:"user_settings_override_json,omitempty"`

	// An arbitrary YAML object allowing (non-admin) cluster owners to set their parameters (only one of this and 'user_settings_json' is allowed), provided they are on the whitelist ('user_settings_whitelist') and not on the blacklist ('user_settings_blacklist'). (These field together with 'user_settings_override*' and 'system_settings' defines the total set of resource settings)
	// +kubebuilder:validation:Optional
	UserSettingsOverrideYaml *string `json:"userSettingsOverrideYaml,omitempty" tf:"user_settings_override_yaml,omitempty"`

	// An arbitrary YAML object allowing ECE admins owners to set clusters' parameters (only one of this and 'user_settings_override_json' is allowed), ie in addition to the documented 'system_settings'. (This field together with 'system_settings' and 'user_settings*' defines the total set of resource settings)
	// +kubebuilder:validation:Optional
	UserSettingsYaml *string `json:"userSettingsYaml,omitempty" tf:"user_settings_yaml,omitempty"`
}

type DeploymentObservation struct {
	ElasticsearchUsername *string `json:"elasticsearchUsername,omitempty" tf:"elasticsearch_username,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DeploymentParameters struct {

	// Optional deployment alias that affects the format of the resource URLs
	// +kubebuilder:validation:Optional
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// Optional APM resource definition
	// +kubebuilder:validation:Optional
	Apm []ApmParameters `json:"apm,omitempty" tf:"apm,omitempty"`

	// Required Deployment Template identifier to create the deployment from
	// +kubebuilder:validation:Required
	DeploymentTemplateID *string `json:"deploymentTemplateId" tf:"deployment_template_id,omitempty"`

	// Required Elasticsearch resource definition
	// +kubebuilder:validation:Required
	Elasticsearch []ElasticsearchParameters `json:"elasticsearch" tf:"elasticsearch,omitempty"`

	// Optional Enterprise Search resource definition
	// +kubebuilder:validation:Optional
	EnterpriseSearch []EnterpriseSearchParameters `json:"enterpriseSearch,omitempty" tf:"enterprise_search,omitempty"`

	// Optional Integrations Server resource definition
	// +kubebuilder:validation:Optional
	IntegrationsServer []IntegrationsServerParameters `json:"integrationsServer,omitempty" tf:"integrations_server,omitempty"`

	// Optional Kibana resource definition
	// +kubebuilder:validation:Optional
	Kibana []KibanaParameters `json:"kibana,omitempty" tf:"kibana,omitempty"`

	// Optional observability settings. Ship logs and metrics to a dedicated deployment.
	// +kubebuilder:validation:Optional
	Observability []ObservabilityParameters `json:"observability,omitempty" tf:"observability,omitempty"`

	// Required ESS region where to create the deployment, for ECE environments "ece-region" must be set
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// Optional request_id to set on the create operation, only use when previous create attempts return with an error and a request_id is returned as part of the error
	// +kubebuilder:validation:Optional
	RequestID *string `json:"requestId,omitempty" tf:"request_id,omitempty"`

	// Optional map of deployment tags
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Optional list of traffic filters to apply to this deployment.
	// +kubebuilder:validation:Optional
	TrafficFilter []*string `json:"trafficFilter,omitempty" tf:"traffic_filter,omitempty"`

	// Required Elastic Stack version to use for all of the deployment resources
	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type ElasticsearchConfigObservation struct {
}

type ElasticsearchConfigParameters struct {

	// Optionally override the docker image the Elasticsearch nodes will use. Note that this field will only work for internal users only.
	// +kubebuilder:validation:Optional
	DockerImage *string `json:"dockerImage,omitempty" tf:"docker_image,omitempty"`

	// List of Elasticsearch supported plugins, which vary from version to version. Check the Stack Pack version to see which plugins are supported for each version. This is currently only available from the UI and [ecctl](https://www.elastic.co/guide/en/ecctl/master/ecctl_stack_list.html)
	// +kubebuilder:validation:Optional
	Plugins []*string `json:"plugins,omitempty" tf:"plugins,omitempty"`

	// JSON-formatted user level "elasticsearch.yml" setting overrides
	// +kubebuilder:validation:Optional
	UserSettingsJSON *string `json:"userSettingsJson,omitempty" tf:"user_settings_json,omitempty"`

	// JSON-formatted admin (ECE) level "elasticsearch.yml" setting overrides
	// +kubebuilder:validation:Optional
	UserSettingsOverrideJSON *string `json:"userSettingsOverrideJson,omitempty" tf:"user_settings_override_json,omitempty"`

	// YAML-formatted admin (ECE) level "elasticsearch.yml" setting overrides
	// +kubebuilder:validation:Optional
	UserSettingsOverrideYaml *string `json:"userSettingsOverrideYaml,omitempty" tf:"user_settings_override_yaml,omitempty"`

	// YAML-formatted user level "elasticsearch.yml" setting overrides
	// +kubebuilder:validation:Optional
	UserSettingsYaml *string `json:"userSettingsYaml,omitempty" tf:"user_settings_yaml,omitempty"`
}

type ElasticsearchObservation struct {
	CloudID *string `json:"cloudId,omitempty" tf:"cloud_id,omitempty"`

	HTTPEndpoint *string `json:"httpEndpoint,omitempty" tf:"http_endpoint,omitempty"`

	HTTPSEndpoint *string `json:"httpsEndpoint,omitempty" tf:"https_endpoint,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`
}

type ElasticsearchParameters struct {

	// Enable or disable autoscaling. Defaults to the setting coming from the deployment template. Accepted values are "true" or "false".
	// +kubebuilder:validation:Optional
	Autoscale *string `json:"autoscale,omitempty" tf:"autoscale,omitempty"`

	// Optional Elasticsearch settings which will be applied to all topologies unless overridden on the topology element
	// +kubebuilder:validation:Optional
	Config []ElasticsearchConfigParameters `json:"config,omitempty" tf:"config,omitempty"`

	// Optional Elasticsearch extensions such as custom bundles or plugins.
	// +kubebuilder:validation:Optional
	Extension []ExtensionParameters `json:"extension,omitempty" tf:"extension,omitempty"`

	// Optional ref_id to set on the Elasticsearch resource
	// +kubebuilder:validation:Optional
	RefID *string `json:"refId,omitempty" tf:"ref_id,omitempty"`

	// Optional Elasticsearch remote clusters to configure for the Elasticsearch resource, can be set multiple times
	// +kubebuilder:validation:Optional
	RemoteCluster []RemoteClusterParameters `json:"remoteCluster,omitempty" tf:"remote_cluster,omitempty"`

	// Optional snapshot source settings. Restore data from a snapshot of another deployment.
	// +kubebuilder:validation:Optional
	SnapshotSource []SnapshotSourceParameters `json:"snapshotSource,omitempty" tf:"snapshot_source,omitempty"`

	// Optional topology element which must be set once but can be set multiple times to compose complex topologies
	// +kubebuilder:validation:Optional
	Topology []ElasticsearchTopologyParameters `json:"topology,omitempty" tf:"topology,omitempty"`

	// Optional Elasticsearch account trust settings.
	// +kubebuilder:validation:Optional
	TrustAccount []TrustAccountParameters `json:"trustAccount,omitempty" tf:"trust_account,omitempty"`

	// Optional Elasticsearch external trust settings.
	// +kubebuilder:validation:Optional
	TrustExternal []TrustExternalParameters `json:"trustExternal,omitempty" tf:"trust_external,omitempty"`
}

type ElasticsearchTopologyObservation struct {
	Config []TopologyConfigObservation `json:"config,omitempty" tf:"config,omitempty"`

	InstanceConfigurationID *string `json:"instanceConfigurationId,omitempty" tf:"instance_configuration_id,omitempty"`

	NodeRoles []*string `json:"nodeRoles,omitempty" tf:"node_roles,omitempty"`
}

type ElasticsearchTopologyParameters struct {

	// Optional Elasticsearch autoscaling settings, such a maximum and minimum size and resources.
	// +kubebuilder:validation:Optional
	Autoscaling []AutoscalingParameters `json:"autoscaling,omitempty" tf:"autoscaling,omitempty"`

	// Required topology ID from the deployment template
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// The node type for the Elasticsearch Topology element (data node)
	// +kubebuilder:validation:Optional
	NodeTypeData *string `json:"nodeTypeData,omitempty" tf:"node_type_data,omitempty"`

	// The node type for the Elasticsearch Topology element (ingest node)
	// +kubebuilder:validation:Optional
	NodeTypeIngest *string `json:"nodeTypeIngest,omitempty" tf:"node_type_ingest,omitempty"`

	// The node type for the Elasticsearch Topology element (machine learning node)
	// +kubebuilder:validation:Optional
	NodeTypeML *string `json:"nodeTypeMl,omitempty" tf:"node_type_ml,omitempty"`

	// The node type for the Elasticsearch Topology element (master node)
	// +kubebuilder:validation:Optional
	NodeTypeMaster *string `json:"nodeTypeMaster,omitempty" tf:"node_type_master,omitempty"`

	// Optional amount of memory per node in the "<size in GB>g" notation
	// +kubebuilder:validation:Optional
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// Optional size type, defaults to "memory".
	// +kubebuilder:validation:Optional
	SizeResource *string `json:"sizeResource,omitempty" tf:"size_resource,omitempty"`

	// Optional number of zones that the Elasticsearch cluster will span. This is used to set HA
	// +kubebuilder:validation:Optional
	ZoneCount *float64 `json:"zoneCount,omitempty" tf:"zone_count,omitempty"`
}

type EnterpriseSearchConfigObservation struct {
}

type EnterpriseSearchConfigParameters struct {

	// Optionally override the docker image the Enterprise Search nodes will use. Note that this field will only work for internal users only.
	// +kubebuilder:validation:Optional
	DockerImage *string `json:"dockerImage,omitempty" tf:"docker_image,omitempty"`

	// An arbitrary JSON object allowing (non-admin) cluster owners to set their parameters (only one of this and 'user_settings_yaml' is allowed), provided they are on the whitelist ('user_settings_whitelist') and not on the blacklist ('user_settings_blacklist'). (This field together with 'user_settings_override*' and 'system_settings' defines the total set of resource settings)
	// +kubebuilder:validation:Optional
	UserSettingsJSON *string `json:"userSettingsJson,omitempty" tf:"user_settings_json,omitempty"`

	// An arbitrary JSON object allowing ECE admins owners to set clusters' parameters (only one of this and 'user_settings_override_yaml' is allowed), ie in addition to the documented 'system_settings'. (This field together with 'system_settings' and 'user_settings*' defines the total set of resource settings)
	// +kubebuilder:validation:Optional
	UserSettingsOverrideJSON *string `json:"userSettingsOverrideJson,omitempty" tf:"user_settings_override_json,omitempty"`

	// An arbitrary YAML object allowing (non-admin) cluster owners to set their parameters (only one of this and 'user_settings_json' is allowed), provided they are on the whitelist ('user_settings_whitelist') and not on the blacklist ('user_settings_blacklist'). (These field together with 'user_settings_override*' and 'system_settings' defines the total set of resource settings)
	// +kubebuilder:validation:Optional
	UserSettingsOverrideYaml *string `json:"userSettingsOverrideYaml,omitempty" tf:"user_settings_override_yaml,omitempty"`

	// An arbitrary YAML object allowing ECE admins owners to set clusters' parameters (only one of this and 'user_settings_override_json' is allowed), ie in addition to the documented 'system_settings'. (This field together with 'system_settings' and 'user_settings*' defines the total set of resource settings)
	// +kubebuilder:validation:Optional
	UserSettingsYaml *string `json:"userSettingsYaml,omitempty" tf:"user_settings_yaml,omitempty"`
}

type EnterpriseSearchObservation struct {
	HTTPEndpoint *string `json:"httpEndpoint,omitempty" tf:"http_endpoint,omitempty"`

	HTTPSEndpoint *string `json:"httpsEndpoint,omitempty" tf:"https_endpoint,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`
}

type EnterpriseSearchParameters struct {

	// Optionally define the Enterprise Search configuration options for the Enterprise Search Server
	// +kubebuilder:validation:Optional
	Config []EnterpriseSearchConfigParameters `json:"config,omitempty" tf:"config,omitempty"`

	// +kubebuilder:validation:Optional
	ElasticsearchClusterRefID *string `json:"elasticsearchClusterRefId,omitempty" tf:"elasticsearch_cluster_ref_id,omitempty"`

	// +kubebuilder:validation:Optional
	RefID *string `json:"refId,omitempty" tf:"ref_id,omitempty"`

	// +kubebuilder:validation:Optional
	Topology []EnterpriseSearchTopologyParameters `json:"topology,omitempty" tf:"topology,omitempty"`
}

type EnterpriseSearchTopologyObservation struct {
	NodeTypeAppserver *bool `json:"nodeTypeAppserver,omitempty" tf:"node_type_appserver,omitempty"`

	NodeTypeConnector *bool `json:"nodeTypeConnector,omitempty" tf:"node_type_connector,omitempty"`

	NodeTypeWorker *bool `json:"nodeTypeWorker,omitempty" tf:"node_type_worker,omitempty"`
}

type EnterpriseSearchTopologyParameters struct {

	// +kubebuilder:validation:Optional
	InstanceConfigurationID *string `json:"instanceConfigurationId,omitempty" tf:"instance_configuration_id,omitempty"`

	// +kubebuilder:validation:Optional
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// Optional size type, defaults to "memory".
	// +kubebuilder:validation:Optional
	SizeResource *string `json:"sizeResource,omitempty" tf:"size_resource,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneCount *float64 `json:"zoneCount,omitempty" tf:"zone_count,omitempty"`
}

type ExtensionObservation struct {
}

type ExtensionParameters struct {

	// Extension name.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Extension type, only `bundle` or `plugin` are supported.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Bundle or plugin URL, the extension URL can be obtained from the `ec_deployment_extension.<name>.url` attribute or the API and cannot be a random HTTP address that is hosted elsewhere.
	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`

	// Elasticsearch compatibility version. Bundles should specify major or minor versions with wildcards, such as `7.*` or `*` but **plugins must use full version notation down to the patch level**, such as `7.10.1` and wildcards are not allowed.
	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type IntegrationsServerConfigObservation struct {
}

type IntegrationsServerConfigParameters struct {

	// Optionally enable debug mode for IntegrationsServer servers - defaults to false
	// +kubebuilder:validation:Optional
	DebugEnabled *bool `json:"debugEnabled,omitempty" tf:"debug_enabled,omitempty"`

	// Optionally override the docker image the IntegrationsServer nodes will use. Note that this field will only work for internal users only.
	// +kubebuilder:validation:Optional
	DockerImage *string `json:"dockerImage,omitempty" tf:"docker_image,omitempty"`

	// An arbitrary JSON object allowing (non-admin) cluster owners to set their parameters (only one of this and 'user_settings_yaml' is allowed), provided they are on the whitelist ('user_settings_whitelist') and not on the blacklist ('user_settings_blacklist'). (This field together with 'user_settings_override*' and 'system_settings' defines the total set of resource settings)
	// +kubebuilder:validation:Optional
	UserSettingsJSON *string `json:"userSettingsJson,omitempty" tf:"user_settings_json,omitempty"`

	// An arbitrary JSON object allowing ECE admins owners to set clusters' parameters (only one of this and 'user_settings_override_yaml' is allowed), ie in addition to the documented 'system_settings'. (This field together with 'system_settings' and 'user_settings*' defines the total set of resource settings)
	// +kubebuilder:validation:Optional
	UserSettingsOverrideJSON *string `json:"userSettingsOverrideJson,omitempty" tf:"user_settings_override_json,omitempty"`

	// An arbitrary YAML object allowing (non-admin) cluster owners to set their parameters (only one of this and 'user_settings_json' is allowed), provided they are on the whitelist ('user_settings_whitelist') and not on the blacklist ('user_settings_blacklist'). (These field together with 'user_settings_override*' and 'system_settings' defines the total set of resource settings)
	// +kubebuilder:validation:Optional
	UserSettingsOverrideYaml *string `json:"userSettingsOverrideYaml,omitempty" tf:"user_settings_override_yaml,omitempty"`

	// An arbitrary YAML object allowing ECE admins owners to set clusters' parameters (only one of this and 'user_settings_override_json' is allowed), ie in addition to the documented 'system_settings'. (This field together with 'system_settings' and 'user_settings*' defines the total set of resource settings)
	// +kubebuilder:validation:Optional
	UserSettingsYaml *string `json:"userSettingsYaml,omitempty" tf:"user_settings_yaml,omitempty"`
}

type IntegrationsServerObservation struct {
	HTTPEndpoint *string `json:"httpEndpoint,omitempty" tf:"http_endpoint,omitempty"`

	HTTPSEndpoint *string `json:"httpsEndpoint,omitempty" tf:"https_endpoint,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`
}

type IntegrationsServerParameters struct {

	// Optionally define the IntegrationsServer configuration options for the IntegrationsServer Server
	// +kubebuilder:validation:Optional
	Config []IntegrationsServerConfigParameters `json:"config,omitempty" tf:"config,omitempty"`

	// +kubebuilder:validation:Optional
	ElasticsearchClusterRefID *string `json:"elasticsearchClusterRefId,omitempty" tf:"elasticsearch_cluster_ref_id,omitempty"`

	// +kubebuilder:validation:Optional
	RefID *string `json:"refId,omitempty" tf:"ref_id,omitempty"`

	// +kubebuilder:validation:Optional
	Topology []IntegrationsServerTopologyParameters `json:"topology,omitempty" tf:"topology,omitempty"`
}

type IntegrationsServerTopologyObservation struct {
}

type IntegrationsServerTopologyParameters struct {

	// +kubebuilder:validation:Optional
	InstanceConfigurationID *string `json:"instanceConfigurationId,omitempty" tf:"instance_configuration_id,omitempty"`

	// +kubebuilder:validation:Optional
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// Optional size type, defaults to "memory".
	// +kubebuilder:validation:Optional
	SizeResource *string `json:"sizeResource,omitempty" tf:"size_resource,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneCount *float64 `json:"zoneCount,omitempty" tf:"zone_count,omitempty"`
}

type KibanaConfigObservation struct {
}

type KibanaConfigParameters struct {

	// Optionally override the docker image the Kibana nodes will use. Note that this field will only work for internal users only.
	// +kubebuilder:validation:Optional
	DockerImage *string `json:"dockerImage,omitempty" tf:"docker_image,omitempty"`

	// An arbitrary JSON object allowing (non-admin) cluster owners to set their parameters (only one of this and 'user_settings_yaml' is allowed), provided they are on the whitelist ('user_settings_whitelist') and not on the blacklist ('user_settings_blacklist'). (This field together with 'user_settings_override*' and 'system_settings' defines the total set of resource settings)
	// +kubebuilder:validation:Optional
	UserSettingsJSON *string `json:"userSettingsJson,omitempty" tf:"user_settings_json,omitempty"`

	// An arbitrary JSON object allowing ECE admins owners to set clusters' parameters (only one of this and 'user_settings_override_yaml' is allowed), ie in addition to the documented 'system_settings'. (This field together with 'system_settings' and 'user_settings*' defines the total set of resource settings)
	// +kubebuilder:validation:Optional
	UserSettingsOverrideJSON *string `json:"userSettingsOverrideJson,omitempty" tf:"user_settings_override_json,omitempty"`

	// An arbitrary YAML object allowing (non-admin) cluster owners to set their parameters (only one of this and 'user_settings_json' is allowed), provided they are on the whitelist ('user_settings_whitelist') and not on the blacklist ('user_settings_blacklist'). (These field together with 'user_settings_override*' and 'system_settings' defines the total set of resource settings)
	// +kubebuilder:validation:Optional
	UserSettingsOverrideYaml *string `json:"userSettingsOverrideYaml,omitempty" tf:"user_settings_override_yaml,omitempty"`

	// An arbitrary YAML object allowing ECE admins owners to set clusters' parameters (only one of this and 'user_settings_override_json' is allowed), ie in addition to the documented 'system_settings'. (This field together with 'system_settings' and 'user_settings*' defines the total set of resource settings)
	// +kubebuilder:validation:Optional
	UserSettingsYaml *string `json:"userSettingsYaml,omitempty" tf:"user_settings_yaml,omitempty"`
}

type KibanaObservation struct {
	HTTPEndpoint *string `json:"httpEndpoint,omitempty" tf:"http_endpoint,omitempty"`

	HTTPSEndpoint *string `json:"httpsEndpoint,omitempty" tf:"https_endpoint,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`
}

type KibanaParameters struct {

	// Optionally define the Kibana configuration options for the Kibana Server
	// +kubebuilder:validation:Optional
	Config []KibanaConfigParameters `json:"config,omitempty" tf:"config,omitempty"`

	// +kubebuilder:validation:Optional
	ElasticsearchClusterRefID *string `json:"elasticsearchClusterRefId,omitempty" tf:"elasticsearch_cluster_ref_id,omitempty"`

	// +kubebuilder:validation:Optional
	RefID *string `json:"refId,omitempty" tf:"ref_id,omitempty"`

	// +kubebuilder:validation:Optional
	Topology []KibanaTopologyParameters `json:"topology,omitempty" tf:"topology,omitempty"`
}

type KibanaTopologyObservation struct {
}

type KibanaTopologyParameters struct {

	// +kubebuilder:validation:Optional
	InstanceConfigurationID *string `json:"instanceConfigurationId,omitempty" tf:"instance_configuration_id,omitempty"`

	// +kubebuilder:validation:Optional
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// Optional size type, defaults to "memory".
	// +kubebuilder:validation:Optional
	SizeResource *string `json:"sizeResource,omitempty" tf:"size_resource,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneCount *float64 `json:"zoneCount,omitempty" tf:"zone_count,omitempty"`
}

type ObservabilityObservation struct {
}

type ObservabilityParameters struct {

	// +kubebuilder:validation:Required
	DeploymentID *string `json:"deploymentId" tf:"deployment_id,omitempty"`

	// +kubebuilder:validation:Optional
	Logs *bool `json:"logs,omitempty" tf:"logs,omitempty"`

	// +kubebuilder:validation:Optional
	Metrics *bool `json:"metrics,omitempty" tf:"metrics,omitempty"`

	// +kubebuilder:validation:Optional
	RefID *string `json:"refId,omitempty" tf:"ref_id,omitempty"`
}

type RemoteClusterObservation struct {
}

type RemoteClusterParameters struct {

	// Alias for this Cross Cluster Search binding
	// +kubebuilder:validation:Required
	Alias *string `json:"alias" tf:"alias,omitempty"`

	// Remote deployment ID
	// +kubebuilder:validation:Required
	DeploymentID *string `json:"deploymentId" tf:"deployment_id,omitempty"`

	// Remote elasticsearch "ref_id", it is best left to the default value
	// +kubebuilder:validation:Optional
	RefID *string `json:"refId,omitempty" tf:"ref_id,omitempty"`

	// If true, skip the cluster during search when disconnected
	// +kubebuilder:validation:Optional
	SkipUnavailable *bool `json:"skipUnavailable,omitempty" tf:"skip_unavailable,omitempty"`
}

type SnapshotSourceObservation struct {
}

type SnapshotSourceParameters struct {

	// Name of the snapshot to restore. Use '__latest_success__' to get the most recent successful snapshot.
	// +kubebuilder:validation:Optional
	SnapshotName *string `json:"snapshotName,omitempty" tf:"snapshot_name,omitempty"`

	// ID of the Elasticsearch cluster that will be used as the source of the snapshot
	// +kubebuilder:validation:Required
	SourceElasticsearchClusterID *string `json:"sourceElasticsearchClusterId" tf:"source_elasticsearch_cluster_id,omitempty"`
}

type TopologyConfigObservation struct {
}

type TopologyConfigParameters struct {

	// +kubebuilder:validation:Required
	Plugins []*string `json:"plugins" tf:"plugins,omitempty"`

	// +kubebuilder:validation:Required
	UserSettingsJSON *string `json:"userSettingsJson" tf:"user_settings_json,omitempty"`

	// +kubebuilder:validation:Required
	UserSettingsOverrideJSON *string `json:"userSettingsOverrideJson" tf:"user_settings_override_json,omitempty"`

	// +kubebuilder:validation:Required
	UserSettingsOverrideYaml *string `json:"userSettingsOverrideYaml" tf:"user_settings_override_yaml,omitempty"`

	// +kubebuilder:validation:Required
	UserSettingsYaml *string `json:"userSettingsYaml" tf:"user_settings_yaml,omitempty"`
}

type TopologyObservation struct {
}

type TopologyParameters struct {

	// +kubebuilder:validation:Optional
	InstanceConfigurationID *string `json:"instanceConfigurationId,omitempty" tf:"instance_configuration_id,omitempty"`

	// +kubebuilder:validation:Optional
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// Optional size type, defaults to "memory".
	// +kubebuilder:validation:Optional
	SizeResource *string `json:"sizeResource,omitempty" tf:"size_resource,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneCount *float64 `json:"zoneCount,omitempty" tf:"zone_count,omitempty"`
}

type TrustAccountObservation struct {
}

type TrustAccountParameters struct {

	// The ID of the Account.
	// +kubebuilder:validation:Required
	AccountID *string `json:"accountId" tf:"account_id,omitempty"`

	// If true, all clusters in this account will by default be trusted and the `trust_allowlist` is ignored.
	// +kubebuilder:validation:Required
	TrustAll *bool `json:"trustAll" tf:"trust_all,omitempty"`

	// The list of clusters to trust. Only used when `trust_all` is false.
	// +kubebuilder:validation:Optional
	TrustAllowlist []*string `json:"trustAllowlist,omitempty" tf:"trust_allowlist,omitempty"`
}

type TrustExternalObservation struct {
}

type TrustExternalParameters struct {

	// The ID of the external trust relationship.
	// +kubebuilder:validation:Required
	RelationshipID *string `json:"relationshipId" tf:"relationship_id,omitempty"`

	// If true, all clusters in this account will by default be trusted and the `trust_allowlist` is ignored.
	// +kubebuilder:validation:Required
	TrustAll *bool `json:"trustAll" tf:"trust_all,omitempty"`

	// The list of clusters to trust. Only used when `trust_all` is false.
	// +kubebuilder:validation:Optional
	TrustAllowlist []*string `json:"trustAllowlist,omitempty" tf:"trust_allowlist,omitempty"`
}

// DeploymentSpec defines the desired state of Deployment
type DeploymentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeploymentParameters `json:"forProvider"`
}

// DeploymentStatus defines the observed state of Deployment.
type DeploymentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeploymentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Deployment is the Schema for the Deployments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,elasticcloudjet}
type Deployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeploymentSpec   `json:"spec"`
	Status            DeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentList contains a list of Deployments
type DeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Deployment `json:"items"`
}

// Repository type metadata.
var (
	Deployment_Kind             = "Deployment"
	Deployment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Deployment_Kind}.String()
	Deployment_KindAPIVersion   = Deployment_Kind + "." + CRDGroupVersion.String()
	Deployment_GroupVersionKind = CRDGroupVersion.WithKind(Deployment_Kind)
)

func init() {
	SchemeBuilder.Register(&Deployment{}, &DeploymentList{})
}
