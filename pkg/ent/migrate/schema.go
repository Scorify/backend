// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ChecksColumns holds the columns for the "checks" table.
	ChecksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "source", Type: field.TypeString},
		{Name: "weight", Type: field.TypeInt},
		{Name: "default_config", Type: field.TypeJSON},
	}
	// ChecksTable holds the schema information for the "checks" table.
	ChecksTable = &schema.Table{
		Name:       "checks",
		Columns:    ChecksColumns,
		PrimaryKey: []*schema.Column{ChecksColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "check_name",
				Unique:  false,
				Columns: []*schema.Column{ChecksColumns[3]},
			},
		},
	}
	// CheckConfigsColumns holds the columns for the "check_configs" table.
	CheckConfigsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "config", Type: field.TypeJSON},
		{Name: "check_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// CheckConfigsTable holds the schema information for the "check_configs" table.
	CheckConfigsTable = &schema.Table{
		Name:       "check_configs",
		Columns:    CheckConfigsColumns,
		PrimaryKey: []*schema.Column{CheckConfigsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "check_configs_checks_check",
				Columns:    []*schema.Column{CheckConfigsColumns[4]},
				RefColumns: []*schema.Column{ChecksColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "check_configs_users_user",
				Columns:    []*schema.Column{CheckConfigsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "checkconfig_check_id_user_id",
				Unique:  false,
				Columns: []*schema.Column{CheckConfigsColumns[4], CheckConfigsColumns[5]},
			},
		},
	}
	// RoundsColumns holds the columns for the "rounds" table.
	RoundsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "number", Type: field.TypeInt, Unique: true},
		{Name: "complete", Type: field.TypeBool, Default: false},
		{Name: "points", Type: field.TypeInt},
	}
	// RoundsTable holds the schema information for the "rounds" table.
	RoundsTable = &schema.Table{
		Name:       "rounds",
		Columns:    RoundsColumns,
		PrimaryKey: []*schema.Column{RoundsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "round_number",
				Unique:  false,
				Columns: []*schema.Column{RoundsColumns[3]},
			},
		},
	}
	// ScoreCachesColumns holds the columns for the "score_caches" table.
	ScoreCachesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "points", Type: field.TypeInt},
		{Name: "round_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// ScoreCachesTable holds the schema information for the "score_caches" table.
	ScoreCachesTable = &schema.Table{
		Name:       "score_caches",
		Columns:    ScoreCachesColumns,
		PrimaryKey: []*schema.Column{ScoreCachesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "score_caches_rounds_round",
				Columns:    []*schema.Column{ScoreCachesColumns[4]},
				RefColumns: []*schema.Column{RoundsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "score_caches_users_user",
				Columns:    []*schema.Column{ScoreCachesColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "scorecache_round_id_user_id",
				Unique:  false,
				Columns: []*schema.Column{ScoreCachesColumns[4], ScoreCachesColumns[5]},
			},
		},
	}
	// StatusColumns holds the columns for the "status" table.
	StatusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "error", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"up", "down", "unknown"}, Default: "unknown"},
		{Name: "points", Type: field.TypeInt},
		{Name: "check_id", Type: field.TypeUUID},
		{Name: "round_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// StatusTable holds the schema information for the "status" table.
	StatusTable = &schema.Table{
		Name:       "status",
		Columns:    StatusColumns,
		PrimaryKey: []*schema.Column{StatusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "status_checks_check",
				Columns:    []*schema.Column{StatusColumns[6]},
				RefColumns: []*schema.Column{ChecksColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "status_rounds_round",
				Columns:    []*schema.Column{StatusColumns[7]},
				RefColumns: []*schema.Column{RoundsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "status_users_user",
				Columns:    []*schema.Column{StatusColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "status_check_id_round_id_user_id",
				Unique:  false,
				Columns: []*schema.Column{StatusColumns[6], StatusColumns[7], StatusColumns[8]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"admin", "user"}, Default: "user"},
		{Name: "number", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_username",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[3]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChecksTable,
		CheckConfigsTable,
		RoundsTable,
		ScoreCachesTable,
		StatusTable,
		UsersTable,
	}
)

func init() {
	CheckConfigsTable.ForeignKeys[0].RefTable = ChecksTable
	CheckConfigsTable.ForeignKeys[1].RefTable = UsersTable
	ScoreCachesTable.ForeignKeys[0].RefTable = RoundsTable
	ScoreCachesTable.ForeignKeys[1].RefTable = UsersTable
	StatusTable.ForeignKeys[0].RefTable = ChecksTable
	StatusTable.ForeignKeys[1].RefTable = RoundsTable
	StatusTable.ForeignKeys[2].RefTable = UsersTable
}
