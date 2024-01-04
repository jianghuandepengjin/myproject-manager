package dao

import (
	"context"
	"test.com/project-user/internal/database"
	"test.com/project-user/internal/database/gorms"
	"test.com/project-user/internal/datatable"
)

type OrganizationDao struct {
	Conn *gorms.GormConn
}

func NewOrganizationDao() *OrganizationDao {
	return &OrganizationDao{
		Conn: gorms.New(),
	}
}

func (o *OrganizationDao) InsertOrganization(conn database.DbConn, ctx context.Context, organization datatable.Organization) (bool, error) {
	o.Conn = conn.(*gorms.GormConn)
	err := o.Conn.Tx(ctx).Create(organization).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
