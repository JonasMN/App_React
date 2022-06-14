package main

import (
	"fmt"
	"time"
)

type ResponseFapiCache struct {
	Metadata struct {
		SecretsWithSub              bool        `json:"secrets_with_sub"`
		ApplicationConfigVersion    interface{} `json:"application_config_version"`
		ScaleToZero                 bool        `json:"scale_to_zero"`
		LoadbalancerProvider        string      `json:"loadbalancer_provider"`
		RequestedEphemeralStorage   string      `json:"requested_ephemeral_storage"`
		LimitVcpu                   int         `json:"limit_vcpu"`
		KairosClusterID             int         `json:"kairos_cluster_id"`
		ApplicationConfigURL        interface{} `json:"application_config_url"`
		Branch                      string      `json:"branch"`
		MiddlewareMigrationStatus   interface{} `json:"middleware_migration_status"`
		PolicyAgentVersion          interface{} `json:"policy_agent_version"`
		DNSSidecarEnabled           bool        `json:"dns_sidecar_enabled"`
		PolicyAgentCPUUsage         string      `json:"policy_agent_cpu_usage"`
		ConfigServiceSidecarEnabled bool        `json:"config_service_sidecar_enabled"`
		LimitMemory                 string      `json:"limit_memory"`
		LimitEphemeralStorage       string      `json:"limit_ephemeral_storage"`
		ProxysqlEnabled             bool        `json:"proxysql_enabled"`
		LocalapiConfigVersion       interface{} `json:"localapi_config_version"`
		NodeAffinity                string      `json:"node_affinity"`
		GranulateEnabled            bool        `json:"granulate_enabled"`
		HTTPMiddlewareConfigVersion interface{} `json:"http_middleware_config_version"`
		DNSSidecarVersion           interface{} `json:"dns_sidecar_version"`
		ConfigServiceSidecarVersion string      `json:"config_service_sidecar_version"`
		ClusterName                 string      `json:"cluster_name"`
		PolicyAgentEnvironment      string      `json:"policy_agent_environment"`
		Test                        bool        `json:"test"`
		TrafficSidecarVersion       interface{} `json:"traffic_sidecar_version"`
		ServerlessCluster           string      `json:"serverless_cluster"`
		SandboxMetadata             interface{} `json:"sandbox_metadata"`
		DesiredInstances            int         `json:"desired_instances"`
		Weight                      int         `json:"weight"`
		HTTPMiddlewareConfigURL     interface{} `json:"http_middleware_config_url"`
		NewrelicID                  int         `json:"newrelic_id"`
		PoolName                    string      `json:"pool_name"`
		TerminationGracePeriod      int         `json:"termination_grace_period"`
		TrafficSidecarEnabled       bool        `json:"traffic_sidecar_enabled"`
		ApdexT                      float64     `json:"apdex_t"`
		RequestedVcpu               int         `json:"requested_vcpu"`
		Build                       string      `json:"build"`
		LocalapiConfigURL           interface{} `json:"localapi_config_url"`
		ComputeClusterProvider      string      `json:"compute_cluster_provider"`
		OtelAgentEnabled            bool        `json:"otel_agent_enabled"`
		TrafficSidecarMode          string      `json:"traffic_sidecar_mode"`
		ConfigService               interface{} `json:"config_service"`
		MeshMigrationIncentive      bool        `json:"mesh_migration_incentive"`
		RequestedMemory             string      `json:"requested_memory"`
		OtelAgentVersion            interface{} `json:"otel_agent_version"`
	} `json:"metadata"`
	LastActionAt      string `json:"last_action_at"`
	Criticality       string `json:"criticality"`
	CreatedAt         string `json:"created_at"`
	Type              string `json:"type"`
	ApplicationRegion struct {
		RegistryURL            interface{} `json:"registry_url"`
		UpdatedAt              string      `json:"updated_at"`
		CreationBlocked        bool        `json:"creation_blocked"`
		ForceEnableDeployments bool        `json:"force_enable_deployments"`
		RegionID               int         `json:"region_id"`
		Name                   string      `json:"name"`
		CreatedAt              string      `json:"created_at"`
		ID                     int         `json:"id"`
		ApplicationID          int         `json:"application_id"`
		UsesNat                bool        `json:"uses_nat"`
		Status                 string      `json:"status"`
	} `json:"application_region"`
	Assets struct {
		Build struct {
			Version string `json:"version"`
			Branch  string `json:"branch"`
		} `json:"build"`
		DeploymentID interface{} `json:"deployment_id"`
		Region       struct {
			Vendor string `json:"vendor"`
			Name   string `json:"name"`
		} `json:"region"`
	} `json:"assets"`
	UpdatedAt                    string        `json:"updated_at"`
	ScalingPolicies              []interface{} `json:"scaling_policies"`
	OncallName                   interface{}   `json:"oncall_name"`
	ServiceID                    interface{}   `json:"service_id"`
	AlertRecipientType           string        `json:"alert_recipient_type"`
	ID                           int           `json:"id"`
	DefaultAlertMonitorTemplates interface{}   `json:"default_alert_monitor_templates"`
	SandboxService               bool          `json:"sandbox_service"`
	ApplicationID                int           `json:"application_id"`
	AdditionalSecurityGroups     []interface{} `json:"additional_security_groups"`
	Tags                         []interface{} `json:"tags"`
	ServiceType                  string        `json:"service_type"`
	Application                  string        `json:"application"`
	Productive                   bool          `json:"productive"`
	LastIndexed                  time.Time     `json:"last_indexed"`
	Name                         string        `json:"name"`
	Config                       struct {
		MiddlewareStatus        string      `json:"middleware_status"`
		TrafficMiddlewareTag    interface{} `json:"traffic_middleware_tag"`
		EnableTrafficMiddleware bool        `json:"enable_traffic_middleware"`
	} `json:"config"`
	FanaticalSupport bool   `json:"fanatical_support"`
	Status           string `json:"status"`
}

func main() {

	var scopes []ResponseFapiCache

        respuestaFapiCache := `[{"metadata":{"secrets_with_sub":true,"application_config_version":"ejmp1"}},{"metadata":{"secrets_with_sub":false,"application_config_version":"test"}},{"metadata":{"secrets_with_sub":false,"application_config_version":"test"}}]`

	json.Unmarshal([]byte(respuestaFapiCache), &scopes)
	
	for _, scope := range scopes {

		fmt.Printf("scope correcto", scope.LastIndexed)

	}

}
