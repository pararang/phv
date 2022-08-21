package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var checkInstalledVersion = &cobra.Command{
	Use:   "list",
	Short: "list",
	Long:  `Check installed PHP version`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			log.Fatal("subcommand list dont take any argument")
		}
		err := printInstalledVersion()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func printInstalledVersion() error {
	cmd := exec.Command("update-alternatives", "--config", "php")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
		
	/**
		Selection    Path             Priority   Status
	------------------------------------------------------------
		0            /usr/bin/php8.1   81        auto mode
	  * 1            /usr/bin/php7.2   72        manual mode
		2            /usr/bin/php8.1   81        manual mode
	**/
	lines := strings.Split(string(output), "\n")
	for i := 1; i < len(lines)-1; i++ {
		if lines[i] == "" {
			continue
		}
		fmt.Printf("%v\n", lines[i])
	}

	return nil
}
