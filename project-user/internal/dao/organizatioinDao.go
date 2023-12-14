package dao

import (
	"context"
	"test.com/project-user/internal/data"
	"test.com/project-user/internal/database/gorms"
)

type OrganizationDao struct {
	Conn *gorms.GormConn
}

func NewOrganizationDao() *OrganizationDao {
	return &OrganizationDao{
		Conn: gorms.New(),
	}
}

func (o *OrganizationDao) InsertOrganization(ctx context.Context, organization data.Organization) (bool, error) {
	err := o.Conn.DB.Create(organization).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
