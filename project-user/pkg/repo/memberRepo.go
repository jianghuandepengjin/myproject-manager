package repo

import (
	"context"
	"test.com/project-user/internal/database"
	"test.com/project-user/internal/datatable"
)

type MemberDao interface {
	GetEmailFromMember(c context.Context, email string) (bool, error)
	GetPhoneFromMember(c context.Context, phone string) (bool, error)
	GetAccountFromMember(c context.Context, account string) (bool, error)
	InsertUserTOMember(conn database.DbConn, c context.Context, member datatable.Member) (bool, error)
}
