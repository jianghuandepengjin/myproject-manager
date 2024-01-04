package repo

import (
	"context"
	"test.com/project-user/internal/database"
	"test.com/project-user/internal/datatable"
)

type Organization interface {
	InsertOrganization(database.DbConn, context.Context, datatable.Organization) (bool, error)
}
