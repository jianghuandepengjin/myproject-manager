package dao

import (
	"context"
	"test.com/project-user/internal/data"
	"test.com/project-user/internal/database/gorms"
)

type MemberDao struct {
	Conn *gorms.GormConn
}

func NewMeberDao() *MemberDao {
	return &MemberDao{
		Conn: gorms.New(),
	}
}

func (m MemberDao) GetEmailFromMember(c context.Context, email string) (bool, error) {
	var count int64
	err := m.Conn.DB.Model(&data.Member{}).Where("email=?", email).Count(&count).Error()
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m MemberDao) GetPhoneFromMember(c context.Context, phone string) (bool, error) {
	var count int64
	err := m.Conn.DB.Model(&data.Member{}).Where("email=?", phone).Count(&count).Error()
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m MemberDao) GetAccountFromMember(c context.Context, account string) (bool, error) {
	var count int64
	err := m.Conn.DB.Model(&data.Member{}).Where("email=?", account).Count(&count).Error()
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
