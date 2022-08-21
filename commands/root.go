package commands

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "url-monitor",
	Short: "url-monitor",
	Long: `url-monitor`,
 }

 func init() {
	rootCmd.AddCommand(checkInstalledVersion, setVersion)
 }
 
 func Execute() {
	if err := rootCmd.Execute(); err != nil {
	   log.Fatal(err)
	}
 }