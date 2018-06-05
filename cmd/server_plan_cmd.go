package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/cloud"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpServerPlan prepares common resources to send request to Concerto API
func WireUpServerPlan(c *cli.Context) (ds *cloud.ServerPlanService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = cloud.NewServerPlanService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up serverPlan service", err)
	}

	return ds, f
}

// ServerPlanList subcommand function
func ServerPlanList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverPlanSvc, formatter := WireUpServerPlan(c)

	serverPlans, err := serverPlanSvc.GetServerPlanList(c.String("cloud_provider_id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive serverPlan data", err)
	}

	cloudProvidersSvc, formatter := WireUpCloudProvider(c)
	cloudProviders, err := cloudProvidersSvc.GetCloudProviderList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive cloudProvider data", err)
	}
	cloudProviderMap := make(map[string]string)
	for _, cloudProvider := range cloudProviders {
		cloudProviderMap[cloudProvider.Id] = cloudProvider.Name
	}

	locationSvc, formatter := WireUpLocation(c)
	locations, err := locationSvc.GetLocationList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive location data", err)
	}
	locationsMap := make(map[string]string)
	for _, location := range locations {
		locationsMap[location.Id] = location.Name
	}
	for id, sp := range serverPlans {
		serverPlans[id].CloudProviderName = cloudProviderMap[sp.CloudProviderId]
		serverPlans[id].LocationName = locationsMap[sp.LocationId]
	}

	if err = formatter.PrintList(serverPlans); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// ServerPlanShow subcommand function
func ServerPlanShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverPlanSvc, formatter := WireUpServerPlan(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	serverPlan, err := serverPlanSvc.GetServerPlan(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive serverPlan data", err)
	}

	cloudProvidersSvc, formatter := WireUpCloudProvider(c)
	cloudProviders, err := cloudProvidersSvc.GetCloudProviderList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive cloudProvider data", err)
	}
	cloudProviderMap := make(map[string]string)
	for _, cloudProvider := range cloudProviders {
		cloudProviderMap[cloudProvider.Id] = cloudProvider.Name
	}

	locationSvc, formatter := WireUpLocation(c)
	locations, err := locationSvc.GetLocationList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive location data", err)
	}
	locationsMap := make(map[string]string)
	for _, location := range locations {
		locationsMap[location.Id] = location.Name
	}

	serverPlan.CloudProviderName = cloudProviderMap[serverPlan.CloudProviderId]
	serverPlan.LocationName = locationsMap[serverPlan.LocationId]

	if err = formatter.PrintItem(*serverPlan); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}
