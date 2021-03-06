// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// UserImageDao is the manager for logic model data accessing and custom defined data operations functions management.
type UserImageDao struct {
	gmvc.M                                      // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	C       userImageColumns // C is the short type for Columns, which contains all the column names of Table for convenient usage.	
	DB      gdb.DB                              // DB is the raw underlying database management object.
	Table   string                              // Table is the underlying table name of the DAO.
}

// UserImageColumns defines and stores column names for table user_image.
type userImageColumns struct {
	Id          string //                          
    ImgUrl      string //                          
    UserId      string //                          
    CreateTime  string //                          
    Status      string // 1-正常，2-删除,3-喜欢的
}

// NewUserImageDao creates and returns a new DAO object for table data access.
func NewUserImageDao() *UserImageDao {
    columns := userImageColumns{
		Id:         "id",           
            ImgUrl:     "img_url",      
            UserId:     "user_id",      
            CreateTime: "create_time",  
            Status:     "status",
	}
	return &UserImageDao{
		C: 	   columns,
		M:     g.DB("default").Model("user_image").Safe(),
		DB:    g.DB("default"),
		Table: "user_image",
	}
}