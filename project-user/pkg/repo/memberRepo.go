package repo

import (
	"context"
	"test.com/project-user/internal/data"
)

type MemberDao interface {
	GetEmailFromMember(c context.Context, email string) (bool, error)
	GetPhoneFromMember(c context.Context, phone string) (bool, error)
	GetAccountFromMember(c context.Context, account string) (bool, error)
	InsertUserTOMember(c context.Context, member data.Member) (bool, error)
}
