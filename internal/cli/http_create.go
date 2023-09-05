package cli

func buildNewSelectPermissions(
	table TableName, source string, rule PermissionsRule,
) CreatePermissionRequest {

	return CreatePermissionRequest{
		Type: "pg_create_select_permission",
		Args: PermissionArgs{
			Table:      table,
			Role:       rule.Role,
			Source:     source,
			Comment:    rule.Description,
			Permission: rule.Permission,
		},
	}
}

// CreatePermissionRequest - Request for creating select permissions.
// https://hasura.io/docs/latest/api-reference/metadata-api/permission/#metadata-pg-create-select-permission
type CreatePermissionRequest struct {
	Type string         `json:"type"`
	Args PermissionArgs `json:"args"`
}

type PermissionArgs struct {
	Table      TableName `json:"table"`
	Role       string    `json:"role"`
	Source     string    `json:"source"`
	Comment    string    `json:"comment"`
	Permission any       `json:"permission"`
}
