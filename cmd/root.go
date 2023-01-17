package cmd

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
	"golang.org/x/mod/module"

	"github.com/cloudsmith-io/gopack/archive"
	"github.com/cloudsmith-io/gopack/modfile"
)

func Execute() error {
	return rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use:   "gopack <version> [<path>]",
	Short: "gopack is a tool for packing Go modules for distribution",
	Long:  "A simple tool for packaging Go modules suitable for distribution via a private module registry.",
	Args: func(cmd *cobra.Command, args []string) error {
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
		modPath, err := modfile.GetModuleNameFromModfile(dir)
		if err != nil {
			fmt.Printf("Unable to determine module path: %s", err.Error())
			os.Exit(1)
		}

		if err := module.Check(modPath, version); err != nil {
			fmt.Printf("Unable to validate module metadata: %s", err.Error())
			os.Exit(1)
		}

		var filter archive.FileFilter
		if filterRegex != "" {
			r, err := regexp.Compile(filterRegex)

			if err != nil {
				fmt.Printf("Invalid filter regex: %s", err.Error())
				os.Exit(1)
			}

			filter = func(s string) bool {
				return r.MatchString(s)
			}
		}

		if err := archive.CreateModuleArchive(dir, modPath, version, filter); err != nil {
			fmt.Printf("Unable to create module archive: %s", err.Error())
			os.Exit(1)
		}
	},
}

var filterRegex string

func init() {
	rootCmd.PersistentFlags().StringVar(
		&filterRegex,
		"filter",
		"",
		"filters which files are added to the resulting archive",
	)
}
