package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/open-falcon/falcon-plus/opsultra-monitor/cmd"
	"os"
)

var RootCmd = &cobra.Command{
	Use: "opsultra-monitor",
	RunE: func(c *cobra.Command, args []string) error {
		return c.Usage()
	},
}

func init(){
	RootCmd.AddCommand(cmd.Check)
	RootCmd.AddCommand(cmd.Start)
	RootCmd.AddCommand(cmd.Stop)
	RootCmd.AddCommand(cmd.Restart)
	RootCmd.AddCommand(cmd.Monitor)
	RootCmd.AddCommand(cmd.Reload)

//	RootCmd.Flags().

}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		
}
