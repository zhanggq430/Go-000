package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
	"log"
)

var ErrUserNotFound = errors.New("user not found")

type Users struct {
	Name    string
	Level   int
	Address string
}

func (u *Users) GetById(id int) error {
	db, err := gorm.Open("mysql", "root:root@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return errors.Wrap(err, "连接数据库失败")
	}
	ctx := db.Where("id = ?", id).Find(u)
	if ctx.Error == gorm.ErrRecordNotFound {
		return errors.WithStack(sql.ErrNoRows)
	}

	// sql错误自己处理掉不往上抛
	if ctx.Error != nil {
		log.Println(ctx.Error)
	}
	return nil
}
