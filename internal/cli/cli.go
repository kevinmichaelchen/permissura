package cli

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "permissura",
	Short: "Utility for managing Hasura permission rules",
	Long:  `Utility for managing Hasura permission rules`,
	Run: func(cmd *cobra.Command, args []string) {
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if debug {
			log.SetLevel(log.DebugLevel)
		}
	},
}

var (
	ldFlags       LDFlags
	endpoint      string
	project       string
	dir           string
	defaultSource string
	adminKey      string
	debug         bool
	dryRun        bool
)

// LDFlags contain fields that get linked and compiled into the final binary
// program at build time.
type LDFlags struct {
	Version string
	Commit  string
	Date    string
}

func init() {
	rootCmd.AddCommand(syncCmd)
	rootCmd.AddCommand(versionCmd)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("unable to read working directory", "err", err)
	}

	log.Debug("working dir", "dir", wd)

	rootCmd.PersistentFlags().StringVarP(&endpoint, "endpoint", "", "http://localhost:8080/v1/metadata", "hasura metadata url")
	rootCmd.PersistentFlags().StringVarP(&project, "project", "", "", "hasura project dir")
	rootCmd.PersistentFlags().StringVarP(&dir, "dir", "", wd, "where to read yaml files")
	rootCmd.PersistentFlags().StringVarP(&defaultSource, "default-source", "", "default", "default hasura data source - see databases.yaml")
	rootCmd.PersistentFlags().StringVarP(&adminKey, "admin-key", "", "myadminsecretkey", "hasura admin key")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "", false, "verbose debug logging")
	rootCmd.PersistentFlags().BoolVarP(&dryRun, "dry-run", "", false, "dry run")
}

func Main(ldf LDFlags) {
	ldFlags = ldf
	if err := rootCmd.Execute(); err != nil {
		log.Error("execution failed", "err", err)
		os.Exit(1)
	}
}
