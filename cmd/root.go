package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"helicopter-hr/cmd/migration"
	"helicopter-hr/internal/app_rest"
	"os"
)

func Start() {
	var rootCmd = &cobra.Command{Use: "sukha command"}

	migrateCmd := &cobra.Command{
		Use:   "db:migrate",
		Short: "database migration",
		Run: func(c *cobra.Command, args []string) {
			migration.MigrateDatabase()
		},
	}

	migrateCmd.Flags().BoolP("version", "", false, "print version")
	migrateCmd.Flags().StringP("dir", "", "database/migration/", "directory with migration files")
	migrateCmd.Flags().StringP("table", "", "db", "migrations table name")
	migrateCmd.Flags().BoolP("verbose", "", false, "enable verbose mode")
	migrateCmd.Flags().BoolP("guide", "", false, "print help")

	var allCmd = &cobra.Command{
		Use:   "http",
		Short: "Run HTTP",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				err := errors.New("missing args")
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			configPath := args[0]
			app_rest.Run(configPath)
		},
	}

	rootCmd.AddCommand(allCmd, migrateCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
