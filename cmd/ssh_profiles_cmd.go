package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/cloud"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpSSHProfile prepares common resources to send request to Concerto API
func WireUpSSHProfile(c *cli.Context) (ds *cloud.SSHProfileService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = cloud.NewSSHProfileService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up sshProfile service", err)
	}

	return ds, f
}

// SSHProfileList subcommand function
func SSHProfileList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	sshProfileSvc, formatter := WireUpSSHProfile(c)

	sshProfiles, err := sshProfileSvc.GetSSHProfileList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive sshProfile data", err)
	}
	if err = formatter.PrintList(sshProfiles); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// SSHProfileShow subcommand function
func SSHProfileShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	sshProfileSvc, formatter := WireUpSSHProfile(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	sshProfile, err := sshProfileSvc.GetSSHProfile(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive sshProfile data", err)
	}
	if err = formatter.PrintItem(*sshProfile); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// SSHProfileCreate subcommand function
func SSHProfileCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	sshProfileSvc, formatter := WireUpSSHProfile(c)

	checkRequiredFlags(c, []string{"name", "public_key"}, formatter)
	sshProfile, err := sshProfileSvc.CreateSSHProfile(utils.FlagConvertParams(c))
	if err != nil {
		formatter.PrintFatal("Couldn't create sshProfile", err)
	}
	if err = formatter.PrintItem(*sshProfile); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// SSHProfileUpdate subcommand function
func SSHProfileUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	sshProfileSvc, formatter := WireUpSSHProfile(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	sshProfile, err := sshProfileSvc.UpdateSSHProfile(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update sshProfile", err)
	}
	if err = formatter.PrintItem(*sshProfile); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// SSHProfileDelete subcommand function
func SSHProfileDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	sshProfileSvc, formatter := WireUpSSHProfile(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := sshProfileSvc.DeleteSSHProfile(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete sshProfile", err)
	}
	return nil
}
