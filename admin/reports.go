package admin

import (
	"encoding/json"
	"fmt"
	// log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
	"os"
	"text/tabwriter"
	"time"
)

type Report struct {
	Id             string          `json:"id"`
	Year           int             `json:"year"`
	Month          time.Month      `json:"month"`
	Start_time     time.Time       `json:"start_time"`
	End_time       time.Time       `json:"end_time"`
	Server_seconds float32         `json:"server_seconds"`
	Closed         bool            `json:"closed"`
	Lines          json.RawMessage `json:"lines"`
	Account_group  AccountGroup    `json:"account_group"`
}

type AccountGroup struct {
	Id   string `json:"_id"`
	Name string `json:"name"`
}

func cmdList(c *cli.Context) {
	var reports []Report

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/admin/reports")
	utils.CheckError(err)

	err = json.Unmarshal(data, &reports)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "REPORT ID\tYEAR\tMONTH\tSTART TIME\tEND TIME\tSERVER SECONDS\tCLOSED\r")

	for _, report := range reports {
		fmt.Fprintf(w, "%s\t%d\t%s\t%s\t%s\t%g\t%t\n", report.Id, report.Year, report.Month, report.Start_time, report.End_time, report.Server_seconds, report.Closed)
	}

	w.Flush()
}

func cmdShow(c *cli.Context) {
	var vals Report

	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/admin/reports/%s", c.String("id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &vals)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)

	fmt.Fprintln(w, "REPORT ID\tYEAR\tMONTH\tSTART TIME\tEND TIME\tSERVER SECONDS\tCLOSED\tLINES\tACCOUNT GROUP ID\tACCOUNT GROUP NAME\r")
	fmt.Fprintf(w, "%s\t%d\t%s\t%s\t%s\t%g\t%t\t%s\t%s\t%s\n", vals.Id, vals.Year, vals.Month, vals.Start_time, vals.End_time, vals.Server_seconds, vals.Closed, vals.Lines, vals.Account_group.Id, vals.Account_group.Name)
	w.Flush()

}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Returns information about the reports related to all the account groups of the tenant. The authenticated user must be an admin.",
			Action: cmdList,
		},
		{
			Name:   "show",
			Usage:  "Returns details about a particular report associated to any account group of the tenant. The authenticated user must be an admin.",
			Action: cmdShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Report Identifier",
				},
			},
		},
	}
}
