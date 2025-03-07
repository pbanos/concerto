package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/cloud"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpCloudProvider prepares common resources to send request to Concerto API
func WireUpCloudProvider(c *cli.Context) (cs *cloud.CloudProviderService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	cs, err = cloud.NewCloudProviderService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up cloudProvider service", err)
	}

	return cs, f
}

// CloudProviderList subcommand function
func CloudProviderList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	cloudProviderSvc, formatter := WireUpCloudProvider(c)

	cloudProviders, err := cloudProviderSvc.GetCloudProviderList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive cloudProvider data", err)
	}
	if err = formatter.PrintList(cloudProviders); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// LoadcloudProvidersMapping retrieves Cloud Providers and create a map between ID and Name
func LoadcloudProvidersMapping(c *cli.Context) map[string]string {
	debugCmdFuncInfo(c)

	cloudProvidersSvc, formatter := WireUpCloudProvider(c)
	cloudProviders, err := cloudProvidersSvc.GetCloudProviderList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive cloudProvider data", err)
	}
	cloudProvidersMap := make(map[string]string)
	for _, cloudProvider := range cloudProviders {
		cloudProvidersMap[cloudProvider.ID] = cloudProvider.Name
	}

	return cloudProvidersMap
}
