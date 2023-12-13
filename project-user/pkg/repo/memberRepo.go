package repo

import "context"

type MemberDao interface {
	GetEmailFromMember(c context.Context, email string) (bool, error)
	GetPhoneFromMember(c context.Context, phone string) (bool, error)
	GetAccountFromMember(c context.Context, account string) (bool, error)
}
