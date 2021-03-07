package dao

import (
	"MicroConf/ConfigService/config"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	client *gorm.DB
}

var defaultDao *Dao

func GetDao() *Dao {
	return defaultDao
}

// Init 初始化数据库连接
func Init(c *config.Config) (err error) {
	defaultDao, err = newDao(c)
	return
}

func newDao(c *config.Config) (*Dao, error) {
	var (
		d   Dao
		err error
	)
	if d.client, err = gorm.Open(c.DB.DriverName, c.DB.URL); err != nil {
		return nil, err
	}
	if d.client, err = gorm.Open(c.DB.DriverName, c.DB.URL); err != nil {
		return nil, err
	}
	d.client.SingularTable(true)       //表名采用单数形式
	d.client.DB().SetMaxOpenConns(100) //SetMaxOpenConns用于设置最大打开的连接数
	d.client.DB().SetMaxIdleConns(10)  //SetMaxIdleConns用于设置闲置的连接数
	//d.client.LogMode(true)

	return &d, nil
}
