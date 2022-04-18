package azure

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

const pluginName = "steampipe-plugin-azure"

// Plugin creates this (azure) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError([]string{"ResourceGroupNotFound"}),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"azure_ad_group":                                              tableAzureAdGroup(ctx),
			"azure_ad_service_principal":                                  tableAzureAdServicePrincipal(ctx),
			"azure_ad_user":                                               tableAzureAdUser(ctx),
			"azure_api_management":                                        tableAzureAPIManagement(ctx),
			"azure_app_configuration":                                     tableAzureAppConfiguration(ctx),
			"azure_app_service_environment":                               tableAzureAppServiceEnvironment(ctx),
			"azure_app_service_function_app":                              tableAzureAppServiceFunctionApp(ctx),
			"azure_app_service_plan":                                      tableAzureAppServicePlan(ctx),
			"azure_app_service_web_app":                                   tableAzureAppServiceWebApp(ctx),
			"azure_application_gateway":                                   tableAzureApplicationGateway(ctx),
			"azure_application_security_group":                            tableAzureApplicationSecurityGroup(ctx),
			"azure_batch_account":                                         tableAzureBatchAccount(ctx),
			"azure_cognitive_account":                                     tableAzureCognitiveAccount(ctx),
			"azure_compute_availability_set":                              tableAzureComputeAvailabilitySet(ctx),
			"azure_compute_disk":                                          tableAzureComputeDisk(ctx),
			"azure_compute_disk_access":                                   tableAzureComputeDiskAccess(ctx),
			"azure_compute_disk_encryption_set":                           tableAzureComputeDiskEncryptionSet(ctx),
			"azure_compute_disk_metric_read_ops":                          tableAzureComputeDiskMetricReadOps(ctx),
			"azure_compute_disk_metric_read_ops_daily":                    tableAzureComputeDiskMetricReadOpsDaily(ctx),
			"azure_compute_disk_metric_read_ops_hourly":                   tableAzureComputeDiskMetricReadOpsHourly(ctx),
			"azure_compute_disk_metric_write_ops":                         tableAzureComputeDiskMetricWriteOps(ctx),
			"azure_compute_disk_metric_write_ops_daily":                   tableAzureComputeDiskMetricWriteOpsDaily(ctx),
			"azure_compute_disk_metric_write_ops_hourly":                  tableAzureComputeDiskMetricWriteOpsHourly(ctx),
			"azure_compute_image":                                         tableAzureComputeImage(ctx),
			"azure_compute_resource_sku":                                  tableAzureResourceSku(ctx),
			"azure_compute_snapshot":                                      tableAzureComputeSnapshot(ctx),
			"azure_compute_virtual_machine":                               tableAzureComputeVirtualMachine(ctx),
			"azure_compute_virtual_machine_metric_cpu_utilization":        tableAzureComputeVirtualMachineMetricCpuUtilization(ctx),
			"azure_compute_virtual_machine_metric_cpu_utilization_daily":  tableAzureComputeVirtualMachineMetricCpuUtilizationDaily(ctx),
			"azure_compute_virtual_machine_metric_cpu_utilization_hourly": tableAzureComputeVirtualMachineMetricCpuUtilizationHourly(ctx),
			"azure_compute_virtual_machine_scale_set":                     tableAzureComputeVirtualMachineScaleSet(ctx),
			"azure_compute_virtual_machine_scale_set_vm":                  tableAzureComputeVirtualMachineScaleSetVm(ctx),
			"azure_container_registry":                                    tableAzureContainerRegistry(ctx),
			"azure_cosmosdb_account":                                      tableAzureCosmosDBAccount(ctx),
			"azure_cosmosdb_mongo_database":                               tableAzureCosmosDBMongoDatabase(ctx),
			"azure_cosmosdb_sql_database":                                 tableAzureCosmosDBSQLDatabase(ctx),
			"azure_data_factory":                                          tableAzureDataFactory(ctx),
			"azure_data_factory_dataset":                                  tableAzureDataFactoryDataset(ctx),
			"azure_data_factory_pipeline":                                 tableAzureDataFactoryPipeline(ctx),
			"azure_data_lake_analytics_account":                           tableAzureDataLakeAnalyticsAccount(ctx),
			"azure_data_lake_store":                                       tableAzureDataLakeStore(ctx),
			"azure_databox_edge_device":                                   tableAzureDataBoxEdgeDevice(ctx),
			"azure_diagnostic_setting":                                    tableAzureDiagnosticSetting(ctx),
			"azure_eventgrid_domain":                                      tableAzureEventGridDomain(ctx),
			"azure_eventgrid_topic":                                       tableAzureEventGridTopic(ctx),
			"azure_eventhub_namespace":                                    tableAzureEventHubNamespace(ctx),
			"azure_express_route_circuit":                                 tableAzureExpressRouteCircuit(ctx),
			"azure_firewall":                                              tableAzureFirewall(ctx),
			"azure_frontdoor":                                             tableAzureFrontDoor(ctx),
			"azure_hdinsight_cluster":                                     tableAzureHDInsightCluster(ctx),
			"azure_healthcare_service":                                    tableAzureHealthcareService(ctx),
			"azure_hpc_cache":                                             tableAzureHPCCache(ctx),
			"azure_hybrid_compute_machine":                                tableAzureHybridComputeMachine(ctx),
			"azure_hybrid_kubernetes_connected_cluster":                   tableAzureHybridKubernetesConnectedCluster(ctx),
			"azure_iothub":                                                tableAzureIotHub(ctx),
			"azure_iothub_dps":                                            tableAzureIotHubDps(ctx),
			"azure_key_vault":                                             tableAzureKeyVault(ctx),
			"azure_key_vault_deleted_vault":                               tableAzureKeyVaultDeletedVault(ctx),
			"azure_key_vault_key":                                         tableAzureKeyVaultKey(ctx),
			"azure_key_vault_managed_hardware_security_module":            tableAzureKeyVaultManagedHardwareSecurityModule(ctx),
			"azure_key_vault_secret":                                      tableAzureKeyVaultSecret(ctx),
			"azure_kubernetes_cluster":                                    tableAzureKubernetesCluster(ctx),
			"azure_kusto_cluster":                                         tableAzureKustoCluster(ctx),
			"azure_lb":                                                    tableAzureLoadBalancer(ctx),
			"azure_lb_backend_address_pool":                               tableAzureLoadBalancerBackendAddressPool(ctx),
			"azure_lb_nat_rule":                                           tableAzureLoadBalancerNatRule(ctx),
			"azure_lb_outbound_rule":                                      tableAzureLoadBalancerOutboundRule(ctx),
			"azure_lb_probe":                                              tableAzureLoadBalancerProbe(ctx),
			"azure_lb_rule":                                               tableAzureLoadBalancerRule(ctx),
			"azure_location":                                              tableAzureLocation(ctx),
			"azure_log_alert":                                             tableAzureLogAlert(ctx),
			"azure_log_profile":                                           tableAzureLogProfile(ctx),
			"azure_logic_app_workflow":                                    tableAzureLogicAppWorkflow(ctx),
			"azure_machine_learning_workspace":                            tableAzureMachineLearningWorkspace(ctx),
			"azure_management_group":                                      tableAzureManagementGroup(ctx),
			"azure_management_lock":                                       tableAzureManagementLock(ctx),
			"azure_mariadb_server":                                        tableAzureMariaDBServer(ctx),
			"azure_mssql_elasticpool":                                     tableAzureMSSQLElasticPool(ctx),
			"azure_mssql_managed_instance":                                tableAzureMSSQLManagedInstance(ctx),
			"azure_mssql_virtual_machine":                                 tableAzureMSSQLVirtualMachine(ctx),
			"azure_mysql_server":                                          tableAzureMySQLServer(ctx),
			"azure_network_interface":                                     tableAzureNetworkInterface(ctx),
			"azure_network_security_group":                                tableAzureNetworkSecurityGroup(ctx),
			"azure_network_watcher":                                       tableAzureNetworkWatcher(ctx),
			"azure_network_watcher_flow_log":                              tableAzureNetworkWatcherFlowLog(ctx),
			"azure_policy_assignment":                                     tableAzurePolicyAssignment(ctx),
			"azure_policy_definition":                                     tableAzurePolicyDefinition(ctx),
			"azure_postgresql_server":                                     tableAzurePostgreSqlServer(ctx),
			"azure_provider":                                              tableAzureProvider(ctx),
			"azure_public_ip":                                             tableAzurePublicIP(ctx),
			"azure_recovery_services_vault":                               tableAzureRecoveryServicesVault(ctx),
			"azure_redis_cache":                                           tableAzureRedisCache(ctx),
			"azure_resource_group":                                        tableAzureResourceGroup(ctx),
			"azure_resource_link":                                         tableAzureResourceLink(ctx),
			"azure_role_assignment":                                       tableAzureIamRoleAssignment(ctx),
			"azure_role_definition":                                       tableAzureIamRoleDefinition(ctx),
			"azure_route_table":                                           tableAzureRouteTable(ctx),
			"azure_search_service":                                        tableAzureSearchService(ctx),
			"azure_security_center_auto_provisioning":                     tableAzureSecurityCenterAutoProvisioning(ctx),
			"azure_security_center_contact":                               tableAzureSecurityCenterContact(ctx),
			"azure_security_center_jit_network_access_policy":             tableAzureSecurityCenterJITNetworkAccessPolicy(ctx),
			"azure_security_center_setting":                               tableAzureSecurityCenterSetting(ctx),
			"azure_security_center_subscription_pricing":                  tableAzureSecurityCenterPricing(ctx),
			"azure_service_fabric_cluster":                                tableAzureServiceFabricCluster(ctx),
			"azure_servicebus_namespace":                                  tableAzureServiceBusNamespace(ctx),
			"azure_signalr_service":                                       tableAzureSignalRService(ctx),
			"azure_spring_cloud_service":                                  tableAzureSpringCloudService(ctx),
			"azure_sql_database":                                          tableAzureSqlDatabase(ctx),
			"azure_sql_server":                                            tableAzureSQLServer(ctx),
			"azure_storage_account":                                       tableAzureStorageAccount(ctx),
			"azure_storage_share_file":                                    tableAzureStorageShareFile(ctx),
			"azure_storage_blob":                                          tableAzureStorageBlob(ctx),
			"azure_storage_blob_service":                                  tableAzureStorageBlobService(ctx),
			"azure_storage_container":                                     tableAzureStorageContainer(ctx),
			"azure_storage_queue":                                         tableAzureStorageQueue(ctx),
			"azure_storage_sync":                                          tableAzureStorageSync(ctx),
			"azure_storage_table":                                         tableAzureStorageTable(ctx),
			"azure_storage_table_service":                                 tableAzureStorageTableService(ctx),
			"azure_stream_analytics_job":                                  tableAzureStreamAnalyticsJob(ctx),
			"azure_subnet":                                                tableAzureSubnet(ctx),
			"azure_subscription":                                          tableAzureSubscription(ctx),
			"azure_synapse_workspace":                                     tableAzureSynapseWorkspace(ctx),
			"azure_tenant":                                                tableAzureTenant(ctx),
			"azure_virtual_network":                                       tableAzureVirtualNetwork(ctx),
			"azure_virtual_network_gateway":                               tableAzureVirtualNetworkGateway(ctx),
		},
	}

	return p
}
