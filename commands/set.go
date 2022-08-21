package commands

import (
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var setVersion = &cobra.Command{
	Use:   "set",
	Short: "set",
	Long:  `Set the version of PHP`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			log.Fatal("subcommand set dont take any argument")
		}
		err := showAndSetInstalledVersion()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func showAndSetInstalledVersion() error {
	cmd := exec.Command("update-alternatives", "--config", "php")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return nil
}

