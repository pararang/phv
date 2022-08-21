package commands

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "phv",
	Short: "phv",
	Long: `phv is a tool to manage installed PHP version.`,
 }

 func init() {
	rootCmd.AddCommand(checkInstalledVersion, setVersion)
 }
 
 func Execute() {
	if err := rootCmd.Execute(); err != nil {
	   log.Fatal(err)
	}
 }