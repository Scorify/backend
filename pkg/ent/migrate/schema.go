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
		{Name: "config", Type: field.TypeJSON},
		{Name: "editable_fields", Type: field.TypeJSON},
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
				Symbol:     "check_configs_checks_configs",
				Columns:    []*schema.Column{CheckConfigsColumns[4]},
				RefColumns: []*schema.Column{ChecksColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "check_configs_users_configs",
				Columns:    []*schema.Column{CheckConfigsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
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
	// InjectsColumns holds the columns for the "injects" table.
	InjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString, Unique: true},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "end_time", Type: field.TypeTime},
		{Name: "files", Type: field.TypeJSON},
		{Name: "rubric", Type: field.TypeJSON},
	}
	// InjectsTable holds the schema information for the "injects" table.
	InjectsTable = &schema.Table{
		Name:       "injects",
		Columns:    InjectsColumns,
		PrimaryKey: []*schema.Column{InjectsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "inject_title",
				Unique:  false,
				Columns: []*schema.Column{InjectsColumns[3]},
			},
		},
	}
	// InjectSubmissionsColumns holds the columns for the "inject_submissions" table.
	InjectSubmissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "files", Type: field.TypeJSON},
		{Name: "notes", Type: field.TypeString},
		{Name: "rubric", Type: field.TypeJSON},
		{Name: "graded", Type: field.TypeBool},
		{Name: "inject_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// InjectSubmissionsTable holds the schema information for the "inject_submissions" table.
	InjectSubmissionsTable = &schema.Table{
		Name:       "inject_submissions",
		Columns:    InjectSubmissionsColumns,
		PrimaryKey: []*schema.Column{InjectSubmissionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "inject_submissions_injects_submissions",
				Columns:    []*schema.Column{InjectSubmissionsColumns[7]},
				RefColumns: []*schema.Column{InjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "inject_submissions_users_submissions",
				Columns:    []*schema.Column{InjectSubmissionsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
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
				Symbol:     "score_caches_rounds_scoreCaches",
				Columns:    []*schema.Column{ScoreCachesColumns[4]},
				RefColumns: []*schema.Column{RoundsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "score_caches_users_scoreCaches",
				Columns:    []*schema.Column{ScoreCachesColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
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
				Symbol:     "status_checks_statuses",
				Columns:    []*schema.Column{StatusColumns[6]},
				RefColumns: []*schema.Column{ChecksColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "status_rounds_statuses",
				Columns:    []*schema.Column{StatusColumns[7]},
				RefColumns: []*schema.Column{RoundsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "status_users_statuses",
				Columns:    []*schema.Column{StatusColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
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
		InjectsTable,
		InjectSubmissionsTable,
		RoundsTable,
		ScoreCachesTable,
		StatusTable,
		UsersTable,
	}
)

func init() {
	CheckConfigsTable.ForeignKeys[0].RefTable = ChecksTable
	CheckConfigsTable.ForeignKeys[1].RefTable = UsersTable
	InjectSubmissionsTable.ForeignKeys[0].RefTable = InjectsTable
	InjectSubmissionsTable.ForeignKeys[1].RefTable = UsersTable
	ScoreCachesTable.ForeignKeys[0].RefTable = RoundsTable
	ScoreCachesTable.ForeignKeys[1].RefTable = UsersTable
	StatusTable.ForeignKeys[0].RefTable = ChecksTable
	StatusTable.ForeignKeys[1].RefTable = RoundsTable
	StatusTable.ForeignKeys[2].RefTable = UsersTable
}
