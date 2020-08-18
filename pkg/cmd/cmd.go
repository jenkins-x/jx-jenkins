package cmd

import (
	"github.com/jenkins-x/jx-jenkins/pkg/cmd/add"
	"github.com/jenkins-x/jx-jenkins/pkg/cmd/deletecmd"
	"github.com/jenkins-x/jx-jenkins/pkg/cmd/jobs"
	"github.com/jenkins-x/jx-jenkins/pkg/cmd/list"
	"github.com/jenkins-x/jx-jenkins/pkg/cmd/trigger"
	"github.com/jenkins-x/jx-jenkins/pkg/common"
	"github.com/jenkins-x/jx-jenkins/pkg/rootcmd"
	"github.com/jenkins-x/jx-logging/pkg/log"
	"github.com/spf13/cobra"
)

// NewCmd creates the new command
func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   rootcmd.TopLevelCommand,
		Short: "a command line tool to work with Jenkins servers in kubernetes",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				log.Logger().Errorf(err.Error())
			}
		},
	}
	cmd.AddCommand(common.SplitCommand(trigger.NewCmdTrigger()))
	cmd.AddCommand(common.SplitCommand(add.NewCmdAdd()))
	cmd.AddCommand(common.SplitCommand(deletecmd.NewCmdDelete()))
	cmd.AddCommand(common.SplitCommand(jobs.NewCmdJobs()))
	cmd.AddCommand(common.SplitCommand(list.NewCmdList()))
	return cmd
}
