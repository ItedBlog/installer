package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = getVersion()
)

func main() {
	var rootCmd = &cobra.Command{
		Use:     "installer",
		Version: version,
		Short:   "itedctl for Initializing a New Server",
		Long: `itedctl is a command-line tool for initializing a new server for Ited Blog.
It provides various commands for initializing a new server, such as installation, uninstallation, and management of software packages.`,
	}

	rootCmd.AddCommand(newInstallCmd())
	rootCmd.AddCommand(newUninstallCmd())
	rootCmd.AddCommand(newListCmd())
	rootCmd.AddCommand(newStatusCmd())
	rootCmd.AddCommand(newManageCmd())
	rootCmd.AddCommand(newUpgradeCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
