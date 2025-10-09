package seeder

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/ent/permission"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/internal/infra/database/seeder"
	"github.com/umardev500/laundry/pkg/types"
)

// PermissionSeeder seeds permissions tied to features
type PermissionSeeder struct {
	client *entdb.Client
}

var _ seeder.Seeder = (*PermissionSeeder)(nil)

func NewPermissionSeeder(client *entdb.Client) *PermissionSeeder {
	return &PermissionSeeder{client: client}
}

func (s *PermissionSeeder) Seed(ctx context.Context) error {
	log.Info().Msg("🌿 Seeding permissions...")

	conn := s.client.GetConn(ctx)

	// Map feature name -> feature ID (from FeatureSeeder)
	featureMap := map[string]uuid.UUID{
		"users":          uuid.MustParse("22222222-1111-1111-1111-111111111111"),
		"roles":          uuid.MustParse("22222222-1111-1111-1111-222222222222"),
		"permissions":    uuid.MustParse("22222222-1111-1111-1111-333333333333"),
		"tenants":        uuid.MustParse("22222222-1111-1111-1111-444444444444"),
		"laundry_orders": uuid.MustParse("22222222-1111-1111-1111-555555555555"),
	}

	permissions := []struct {
		ID          uuid.UUID
		Name        string
		DisplayName string
		Description string
		FeatureName string
	}{
		// Users feature
		{uuid.MustParse("aaaaaaaa-1111-1111-1111-aaaaaaaaaaaa"), "create_user", "Create User", "Ability to create users", "users"},
		{uuid.MustParse("aaaaaaaa-2222-2222-2222-aaaaaaaaaaaa"), "update_user", "Update User", "Ability to update users", "users"},
		{uuid.MustParse("aaaaaaaa-3333-3333-3333-aaaaaaaaaaaa"), "delete_user", "Delete User", "Ability to delete users", "users"},

		// Roles feature
		{uuid.MustParse("bbbbbbbb-1111-1111-1111-bbbbbbbbbbbb"), "create_role", "Create Role", "Ability to create roles", "roles"},
		{uuid.MustParse("bbbbbbbb-2222-2222-2222-bbbbbbbbbbbb"), "update_role", "Update Role", "Ability to update roles", "roles"},

		// Permissions feature
		{uuid.MustParse("cccccccc-1111-1111-1111-cccccccccccc"), "create_permission", "Create Permission", "Ability to create permissions", "permissions"},
		{uuid.MustParse("cccccccc-2222-2222-2222-cccccccccccc"), "update_permission", "Update Permission", "Ability to update permissions", "permissions"},

		// Tenants feature
		{uuid.MustParse("dddddddd-1111-1111-1111-dddddddddddd"), "create_tenant", "Create Tenant", "Ability to create tenants", "tenants"},
		{uuid.MustParse("dddddddd-2222-2222-2222-dddddddddddd"), "update_tenant", "Update Tenant", "Ability to update tenants", "tenants"},

		// Laundry Orders feature
		{uuid.MustParse("eeeeeeee-1111-1111-1111-eeeeeeeeeeee"), "create_order", "Create Order", "Ability to create laundry orders", "laundry_orders"},
		{uuid.MustParse("eeeeeeee-2222-2222-2222-eeeeeeeeeeee"), "update_order", "Update Order", "Ability to update laundry orders", "laundry_orders"},
	}

	for _, p := range permissions {
		featureID, ok := featureMap[p.FeatureName]
		if !ok {
			return fmt.Errorf("feature %s not found for permission %s", p.FeatureName, p.Name)
		}

		err := conn.Permission.
			Create().
			SetID(p.ID). // ✅ static ID
			SetName(p.Name).
			SetDisplayName(p.DisplayName).
			SetDescription(p.Description).
			SetFeatureID(featureID). // ✅ tie to feature
			SetStatus(permission.Status(types.StatusActive)).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			OnConflict(
				sql.ConflictColumns(permission.FieldName),
			).
			UpdateNewValues().
			Exec(ctx)

		if err != nil {
			return err
		}
	}

	log.Info().Msg("✅ Permissions seeded successfully.")
	return nil
}
