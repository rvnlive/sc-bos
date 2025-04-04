// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: queries.sql

package queries

import (
	"context"
	"database/sql"
)

const addRolePermission = `-- name: AddRolePermission :exec
INSERT INTO role_permissions (role_id, permission)
VALUES (?1, ?2)
ON CONFLICT (role_id, permission) DO NOTHING
`

type AddRolePermissionParams struct {
	RoleID     int64
	Permission string
}

func (q *Queries) AddRolePermission(ctx context.Context, arg AddRolePermissionParams) error {
	_, err := q.db.ExecContext(ctx, addRolePermission, arg.RoleID, arg.Permission)
	return err
}

const clearRolePermissions = `-- name: ClearRolePermissions :execrows
DELETE FROM role_permissions
WHERE role_id = ?1
`

func (q *Queries) ClearRolePermissions(ctx context.Context, roleID int64) (int64, error) {
	result, err := q.db.ExecContext(ctx, clearRolePermissions, roleID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const countAccounts = `-- name: CountAccounts :one
SELECT COUNT(*) AS count
FROM accounts
`

func (q *Queries) CountAccounts(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countAccounts)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countRoleAssignments = `-- name: CountRoleAssignments :one
SELECT COUNT(*)
FROM role_assignments
`

func (q *Queries) CountRoleAssignments(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countRoleAssignments)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countRoleAssignmentsForAccount = `-- name: CountRoleAssignmentsForAccount :one
SELECT COUNT(*) AS count
FROM role_assignments
WHERE account_id = ?1
`

func (q *Queries) CountRoleAssignmentsForAccount(ctx context.Context, accountID int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, countRoleAssignmentsForAccount, accountID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countRoleAssignmentsForRole = `-- name: CountRoleAssignmentsForRole :one
SELECT COUNT(*) AS count
FROM role_assignments
WHERE role_id = ?1
`

func (q *Queries) CountRoleAssignmentsForRole(ctx context.Context, roleID int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, countRoleAssignmentsForRole, roleID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countRoles = `-- name: CountRoles :one
SELECT COUNT(*) AS count
FROM roles
`

func (q *Queries) CountRoles(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countRoles)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countServiceCredentialsForAccount = `-- name: CountServiceCredentialsForAccount :one
SELECT COUNT(*) AS count
FROM service_credentials
WHERE account_id = ?1
`

func (q *Queries) CountServiceCredentialsForAccount(ctx context.Context, accountID int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, countServiceCredentialsForAccount, accountID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createRole = `-- name: CreateRole :one
INSERT INTO roles (display_name, description)
VALUES (?1, ?2)
RETURNING id, display_name, description
`

type CreateRoleParams struct {
	DisplayName string
	Description sql.NullString
}

func (q *Queries) CreateRole(ctx context.Context, arg CreateRoleParams) (Role, error) {
	row := q.db.QueryRowContext(ctx, createRole, arg.DisplayName, arg.Description)
	var i Role
	err := row.Scan(&i.ID, &i.DisplayName, &i.Description)
	return i, err
}

const createRoleAssignment = `-- name: CreateRoleAssignment :one
INSERT INTO role_assignments (account_id, role_id, scope_type, scope_resource)
VALUES (?1, ?2, ?3, ?4)
RETURNING id, account_id, role_id, scope_type, scope_resource
`

type CreateRoleAssignmentParams struct {
	AccountID     int64
	RoleID        int64
	ScopeKind     sql.NullString
	ScopeResource sql.NullString
}

func (q *Queries) CreateRoleAssignment(ctx context.Context, arg CreateRoleAssignmentParams) (RoleAssignment, error) {
	row := q.db.QueryRowContext(ctx, createRoleAssignment,
		arg.AccountID,
		arg.RoleID,
		arg.ScopeKind,
		arg.ScopeResource,
	)
	var i RoleAssignment
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.RoleID,
		&i.ScopeType,
		&i.ScopeResource,
	)
	return i, err
}

const createServiceAccount = `-- name: CreateServiceAccount :one
INSERT INTO accounts (display_name, description, type, create_time)
VALUES (?1, ?2, 'SERVICE_ACCOUNT', datetime('now', 'subsec'))
RETURNING id, username, display_name, description, type, create_time
`

type CreateServiceAccountParams struct {
	DisplayName string
	Description sql.NullString
}

func (q *Queries) CreateServiceAccount(ctx context.Context, arg CreateServiceAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createServiceAccount, arg.DisplayName, arg.Description)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.DisplayName,
		&i.Description,
		&i.Type,
		&i.CreateTime,
	)
	return i, err
}

const createServiceCredential = `-- name: CreateServiceCredential :one
INSERT INTO service_credentials (account_id, display_name, description, secret_hash, create_time, expire_time)
VALUES (?1, ?2, ?3, ?4, datetime('now', 'subsec'), ?5)
RETURNING id, account_id, display_name, description, secret_hash, create_time, expire_time
`

type CreateServiceCredentialParams struct {
	AccountID   int64
	DisplayName string
	Description sql.NullString
	SecretHash  []byte
	ExpireTime  sql.NullTime
}

func (q *Queries) CreateServiceCredential(ctx context.Context, arg CreateServiceCredentialParams) (ServiceCredential, error) {
	row := q.db.QueryRowContext(ctx, createServiceCredential,
		arg.AccountID,
		arg.DisplayName,
		arg.Description,
		arg.SecretHash,
		arg.ExpireTime,
	)
	var i ServiceCredential
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.DisplayName,
		&i.Description,
		&i.SecretHash,
		&i.CreateTime,
		&i.ExpireTime,
	)
	return i, err
}

const createUserAccount = `-- name: CreateUserAccount :one
INSERT INTO accounts (username, display_name, description, type, create_time)
VALUES (?1, ?2, ?3, 'USER_ACCOUNT', datetime('now', 'subsec'))
RETURNING id, username, display_name, description, type, create_time
`

type CreateUserAccountParams struct {
	Username    sql.NullString
	DisplayName string
	Description sql.NullString
}

func (q *Queries) CreateUserAccount(ctx context.Context, arg CreateUserAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createUserAccount, arg.Username, arg.DisplayName, arg.Description)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.DisplayName,
		&i.Description,
		&i.Type,
		&i.CreateTime,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :execrows
DELETE FROM accounts
WHERE id = ?1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int64) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteAccount, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const deleteRole = `-- name: DeleteRole :execrows
DELETE FROM roles
WHERE id = ?1
`

func (q *Queries) DeleteRole(ctx context.Context, id int64) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteRole, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const deleteRoleAssignment = `-- name: DeleteRoleAssignment :execrows
DELETE FROM role_assignments
WHERE id = ?1
`

func (q *Queries) DeleteRoleAssignment(ctx context.Context, id int64) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteRoleAssignment, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const deleteRolePermission = `-- name: DeleteRolePermission :execrows
DELETE FROM role_permissions
WHERE role_id = ?1 AND permission = ?2
`

type DeleteRolePermissionParams struct {
	RoleID     int64
	Permission string
}

func (q *Queries) DeleteRolePermission(ctx context.Context, arg DeleteRolePermissionParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteRolePermission, arg.RoleID, arg.Permission)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const deleteServiceCredential = `-- name: DeleteServiceCredential :execrows
DELETE FROM service_credentials
WHERE id = ?1
`

func (q *Queries) DeleteServiceCredential(ctx context.Context, id int64) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteServiceCredential, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const getAccount = `-- name: GetAccount :one
SELECT id, username, display_name, description, type, create_time
FROM accounts
WHERE id = ?1
`

func (q *Queries) GetAccount(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.DisplayName,
		&i.Description,
		&i.Type,
		&i.CreateTime,
	)
	return i, err
}

const getAccountByUsername = `-- name: GetAccountByUsername :one
SELECT id, username, display_name, description, type, create_time
FROM accounts
WHERE username = ?1
`

func (q *Queries) GetAccountByUsername(ctx context.Context, username sql.NullString) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByUsername, username)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.DisplayName,
		&i.Description,
		&i.Type,
		&i.CreateTime,
	)
	return i, err
}

const getAccountPasswordHash = `-- name: GetAccountPasswordHash :one
SELECT password_hash
FROM password_credentials
WHERE account_id = ?1
`

func (q *Queries) GetAccountPasswordHash(ctx context.Context, accountID int64) ([]byte, error) {
	row := q.db.QueryRowContext(ctx, getAccountPasswordHash, accountID)
	var password_hash []byte
	err := row.Scan(&password_hash)
	return password_hash, err
}

const getRole = `-- name: GetRole :one
SELECT id, display_name, description
FROM roles
WHERE id = ?1
`

func (q *Queries) GetRole(ctx context.Context, id int64) (Role, error) {
	row := q.db.QueryRowContext(ctx, getRole, id)
	var i Role
	err := row.Scan(&i.ID, &i.DisplayName, &i.Description)
	return i, err
}

const getRoleAssignment = `-- name: GetRoleAssignment :one
SELECT id, account_id, role_id, scope_type, scope_resource
FROM role_assignments
WHERE id = ?1
`

func (q *Queries) GetRoleAssignment(ctx context.Context, id int64) (RoleAssignment, error) {
	row := q.db.QueryRowContext(ctx, getRoleAssignment, id)
	var i RoleAssignment
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.RoleID,
		&i.ScopeType,
		&i.ScopeResource,
	)
	return i, err
}

const getServiceCredential = `-- name: GetServiceCredential :one
SELECT id, account_id, display_name, description, secret_hash, create_time, expire_time
FROM service_credentials
WHERE id = ?1
`

func (q *Queries) GetServiceCredential(ctx context.Context, id int64) (ServiceCredential, error) {
	row := q.db.QueryRowContext(ctx, getServiceCredential, id)
	var i ServiceCredential
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.DisplayName,
		&i.Description,
		&i.SecretHash,
		&i.CreateTime,
		&i.ExpireTime,
	)
	return i, err
}

const listAccountServiceCredentials = `-- name: ListAccountServiceCredentials :many
SELECT id, account_id, display_name, description, secret_hash, create_time, expire_time
FROM service_credentials
WHERE account_id = ?1
ORDER BY id
`

func (q *Queries) ListAccountServiceCredentials(ctx context.Context, accountID int64) ([]ServiceCredential, error) {
	rows, err := q.db.QueryContext(ctx, listAccountServiceCredentials, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ServiceCredential
	for rows.Next() {
		var i ServiceCredential
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.DisplayName,
			&i.Description,
			&i.SecretHash,
			&i.CreateTime,
			&i.ExpireTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, username, display_name, description, type, create_time
FROM accounts
WHERE id > ?1
ORDER BY id
LIMIT ?2
`

type ListAccountsParams struct {
	AfterID int64
	Limit   int64
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts, arg.AfterID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.DisplayName,
			&i.Description,
			&i.Type,
			&i.CreateTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRoleAssignments = `-- name: ListRoleAssignments :many
SELECT id, account_id, role_id, scope_type, scope_resource
FROM role_assignments
WHERE id > ?1
ORDER BY id
LIMIT ?2
`

type ListRoleAssignmentsParams struct {
	AfterID int64
	Limit   int64
}

func (q *Queries) ListRoleAssignments(ctx context.Context, arg ListRoleAssignmentsParams) ([]RoleAssignment, error) {
	rows, err := q.db.QueryContext(ctx, listRoleAssignments, arg.AfterID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RoleAssignment
	for rows.Next() {
		var i RoleAssignment
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.RoleID,
			&i.ScopeType,
			&i.ScopeResource,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRoleAssignmentsForAccount = `-- name: ListRoleAssignmentsForAccount :many
SELECT id, account_id, role_id, scope_type, scope_resource
FROM role_assignments
WHERE account_id = ?1
  AND id > ?2
ORDER BY id
LIMIT ?3
`

type ListRoleAssignmentsForAccountParams struct {
	AccountID int64
	AfterID   int64
	Limit     int64
}

func (q *Queries) ListRoleAssignmentsForAccount(ctx context.Context, arg ListRoleAssignmentsForAccountParams) ([]RoleAssignment, error) {
	rows, err := q.db.QueryContext(ctx, listRoleAssignmentsForAccount, arg.AccountID, arg.AfterID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RoleAssignment
	for rows.Next() {
		var i RoleAssignment
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.RoleID,
			&i.ScopeType,
			&i.ScopeResource,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRoleAssignmentsForRole = `-- name: ListRoleAssignmentsForRole :many
SELECT id, account_id, role_id, scope_type, scope_resource
FROM role_assignments
WHERE role_id = ?1
  AND id > ?2
ORDER BY id
LIMIT ?3
`

type ListRoleAssignmentsForRoleParams struct {
	RoleID  int64
	AfterID int64
	Limit   int64
}

func (q *Queries) ListRoleAssignmentsForRole(ctx context.Context, arg ListRoleAssignmentsForRoleParams) ([]RoleAssignment, error) {
	rows, err := q.db.QueryContext(ctx, listRoleAssignmentsForRole, arg.RoleID, arg.AfterID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RoleAssignment
	for rows.Next() {
		var i RoleAssignment
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.RoleID,
			&i.ScopeType,
			&i.ScopeResource,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRolePermissions = `-- name: ListRolePermissions :many
SELECT permission
FROM role_permissions
WHERE role_id = ?1
ORDER BY permission
`

func (q *Queries) ListRolePermissions(ctx context.Context, roleID int64) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listRolePermissions, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var permission string
		if err := rows.Scan(&permission); err != nil {
			return nil, err
		}
		items = append(items, permission)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRoles = `-- name: ListRoles :many
SELECT id, display_name, description
FROM roles
WHERE id > ?1
ORDER BY id
LIMIT ?2
`

type ListRolesParams struct {
	AfterID int64
	Limit   int64
}

func (q *Queries) ListRoles(ctx context.Context, arg ListRolesParams) ([]Role, error) {
	rows, err := q.db.QueryContext(ctx, listRoles, arg.AfterID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Role
	for rows.Next() {
		var i Role
		if err := rows.Scan(&i.ID, &i.DisplayName, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRolesAndPermissions = `-- name: ListRolesAndPermissions :many
SELECT roles.id, roles.display_name, roles.description, group_concat(coalesce(role_permissions.permission, ''), ',') AS permissions
FROM roles
LEFT OUTER JOIN role_permissions ON roles.id = role_permissions.role_id
WHERE roles.id > ?1
GROUP BY roles.id
ORDER BY roles.id
LIMIT ?2
`

type ListRolesAndPermissionsParams struct {
	AfterID int64
	Limit   int64
}

type ListRolesAndPermissionsRow struct {
	Role        Role
	Permissions string
}

func (q *Queries) ListRolesAndPermissions(ctx context.Context, arg ListRolesAndPermissionsParams) ([]ListRolesAndPermissionsRow, error) {
	rows, err := q.db.QueryContext(ctx, listRolesAndPermissions, arg.AfterID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListRolesAndPermissionsRow
	for rows.Next() {
		var i ListRolesAndPermissionsRow
		if err := rows.Scan(
			&i.Role.ID,
			&i.Role.DisplayName,
			&i.Role.Description,
			&i.Permissions,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccountDescription = `-- name: UpdateAccountDescription :exec
UPDATE accounts
SET description = ?1
WHERE id = ?2
`

type UpdateAccountDescriptionParams struct {
	Description sql.NullString
	ID          int64
}

func (q *Queries) UpdateAccountDescription(ctx context.Context, arg UpdateAccountDescriptionParams) error {
	_, err := q.db.ExecContext(ctx, updateAccountDescription, arg.Description, arg.ID)
	return err
}

const updateAccountDisplayName = `-- name: UpdateAccountDisplayName :exec
UPDATE accounts
SET display_name = ?1
WHERE id = ?2
`

type UpdateAccountDisplayNameParams struct {
	DisplayName string
	ID          int64
}

func (q *Queries) UpdateAccountDisplayName(ctx context.Context, arg UpdateAccountDisplayNameParams) error {
	_, err := q.db.ExecContext(ctx, updateAccountDisplayName, arg.DisplayName, arg.ID)
	return err
}

const updateAccountPasswordHash = `-- name: UpdateAccountPasswordHash :exec
INSERT INTO password_credentials (account_id, password_hash)
VALUES (?1, ?2)
ON CONFLICT (account_id) DO UPDATE
SET password_hash = excluded.password_hash
`

type UpdateAccountPasswordHashParams struct {
	AccountID    int64
	PasswordHash []byte
}

func (q *Queries) UpdateAccountPasswordHash(ctx context.Context, arg UpdateAccountPasswordHashParams) error {
	_, err := q.db.ExecContext(ctx, updateAccountPasswordHash, arg.AccountID, arg.PasswordHash)
	return err
}

const updateAccountUsername = `-- name: UpdateAccountUsername :exec
UPDATE accounts
SET username = ?1
WHERE id = ?2
`

type UpdateAccountUsernameParams struct {
	Username sql.NullString
	ID       int64
}

func (q *Queries) UpdateAccountUsername(ctx context.Context, arg UpdateAccountUsernameParams) error {
	_, err := q.db.ExecContext(ctx, updateAccountUsername, arg.Username, arg.ID)
	return err
}

const updateRoleDescription = `-- name: UpdateRoleDescription :execrows
UPDATE roles
SET description = ?1
WHERE id = ?2
`

type UpdateRoleDescriptionParams struct {
	Description sql.NullString
	ID          int64
}

func (q *Queries) UpdateRoleDescription(ctx context.Context, arg UpdateRoleDescriptionParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, updateRoleDescription, arg.Description, arg.ID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const updateRoleDisplayName = `-- name: UpdateRoleDisplayName :execrows
UPDATE roles
SET display_name = ?1
WHERE id = ?2
`

type UpdateRoleDisplayNameParams struct {
	DisplayName string
	ID          int64
}

func (q *Queries) UpdateRoleDisplayName(ctx context.Context, arg UpdateRoleDisplayNameParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, updateRoleDisplayName, arg.DisplayName, arg.ID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
