package cli

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Syncs permissions to Hasura GraphQL Engine",
	Long:  `Syncs permissions to Hasura GraphQL Engine`,
	Run:   fnSync,
}

func fnSync(cmd *cobra.Command, args []string) {
	log.Info("Syncing permissions...")

	// parse permissions rules from local YAML files
	rules := parseRules()

	for _, rule := range rules {
		source := defaultSource

		table := TableName{
			Name:   rule.Table.Name,
			Schema: rule.Table.Schema,
		}

		switch rule.Type {
		case "select":
			dropPermission("pg_drop_select_permission", table, rule.Role, source)

			log.Debug("Loaded a permission rule from file system", "rule", rule)

			if rule.Draft {
				log.Info("Rule is still being drafted... Skipping creation.")
				break
			}

			req := buildNewSelectPermissions(table, source, rule)
			err := executeRequest(req)
			if err != nil {
				log.Fatal("unable to create select rule", "err", err)
			}
		}
	}
}
