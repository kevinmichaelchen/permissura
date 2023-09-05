package cli

import (
	"github.com/charmbracelet/log"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

type RuleFile struct {
	Rules []PermissionsRule `json:"rules"`
}

type PermissionsRule struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Draft       bool      `json:"draft"`
	Rule        string    `json:"rule"`
	Role        string    `json:"role"`
	Type        string    `json:"type"`
	Table       TableName `json:"table"`
	Permission  any       `json:"permission"`
}

// parse permissions rules from local YAML files
func parseRules() []PermissionsRule {
	// Read YAML files in the provided directory.
	// These are the Permissions Rules we'll parse and feed into Hasura.
	paths := getYamlPaths()

	var rules []PermissionsRule

	for _, path := range paths {
		f, err := os.ReadFile(path)
		if err != nil {
			log.Fatal("unable to read file", "err", err, "path", path)
		}

		t := RuleFile{}
		err = yaml.Unmarshal(f, &t)
		if err != nil {
			log.Fatal("unable to parse yaml file", "err", err, "path", path)
		}

		rules = append(rules, t.Rules...)

		log.Debug("parsed permissions rule", "rule", t)
	}

	return rules
}

// absolute paths
func getYamlPaths() []string {
	var paths []string

	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Only consider YAML files
			if !strings.HasSuffix(path, ".yml") && !strings.HasSuffix(path, ".yaml") {
				return nil
			}

			log.Debug("path", "path", path, "size", info.Size())

			paths = append(paths, path)

			return nil
		})
	if err != nil {
		log.Fatal("unable to read yaml files in directory", "dir", dir, "err", err)
	}

	return paths
}
