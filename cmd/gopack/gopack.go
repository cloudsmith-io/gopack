package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/mod/module"

	"github.com/cloudsmith-io/gopack"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "gopack <version> [<path>]",
		Short: "gopack is a tool for packing Go modules for distribution",
		Long:  "A simple tool for packaging Go modules suitable for distribution via a private module registry.",
		Args: func(cmd *cobra.Command, args []string) error {
			// validate required version argument is present
			if len(args) < 1 {
				return errors.New("requires a version argument")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			version := args[0]

			var dir string
			if len(args) >= 2 {
				dir = args[1]
			}
			modPath, err := gopack.GetModuleNameFromModfile(dir)
			if err != nil {
				fmt.Printf("Unable to determine module path: %s", err.Error())
				os.Exit(1)
			}

			if err := module.Check(modPath, version); err != nil {
				fmt.Printf("Unable to validate module metadata: %s", err.Error())
				os.Exit(1)
			}

			if err := gopack.CreateModuleArchive(dir, modPath, version); err != nil {
				fmt.Printf("Unable to create module archive: %s", err.Error())
				os.Exit(1)
			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
