package cli

import (
	"fmt"
	"github.com/charmbracelet/log"
)

type DropPermissionRequest struct {
	Type string `json:"type"`
	Args struct {
		Table  TableName `json:"table"`
		Role   string    `json:"role"`
		Source string    `json:"source"`
	} `json:"args"`
}

// drops the permission if it exists.
// this is an idempotent operation.
func dropPermission(typeStr string, table TableName, role, source string) {
	r := DropPermissionRequest{
		Type: typeStr,
		Args: struct {
			Table  TableName `json:"table"`
			Role   string    `json:"role"`
			Source string    `json:"source"`
		}(struct {
			Table  TableName
			Role   string
			Source string
		}{
			Table:  table,
			Role:   role,
			Source: source,
		}),
	}

	err := executeRequest(r,
		fmt.Sprintf("drop-%s-%s-%s", source, table, role),
		false,
	)

	// idempotence means not aborting if the permission has already been dropped
	if err != nil {
		log.Warn("unable to drop permission", "err", err)
	}
}
