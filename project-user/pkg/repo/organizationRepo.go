package repo

import (
	"context"
	"test.com/project-user/internal/data"
)

type Organization interface {
	InsertOrganization(context.Context, data.Organization) (bool, error)
}
