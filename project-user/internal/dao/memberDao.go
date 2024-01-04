package dao

import (
	"context"
	"test.com/project-user/internal/database"
	"test.com/project-user/internal/database/gorms"
	"test.com/project-user/internal/datatable"
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
	err := m.Conn.DB.Model(&datatable.Member{}).Where("email=?", email).Count(&count).Error
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
	//err := m.Conn.DB.Model(&datatable.Member{}).Where("email=?", phone).Count(&count).Error()
	err := m.Conn.DB.Raw("select count(*) as count from ms_member where mobile = ?", phone).Scan(&count).Error
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
	err := m.Conn.DB.Model(&datatable.Member{}).Where("email=?", account).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (m MemberDao) InsertUserTOMember(conn database.DbConn, c context.Context, member datatable.Member) (bool, error) {
	m.Conn = conn.(*gorms.GormConn)
	result := m.Conn.Tx(c).Create(&member)
	err := result.Error
	if err != nil {
		return false, err
	}
	return true, nil
}
