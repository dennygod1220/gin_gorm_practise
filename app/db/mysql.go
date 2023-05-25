package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	DB_Host     string
	DB_Port     string
	DB_Name     string
	DB_User     string
	DB_Password string
}

func (m *Mysql) Dsn() string {
	return m.DB_User + ":" + m.DB_Password + "@tcp(" + m.DB_Host + ":" + m.DB_Port + ")/" + m.DB_Name + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func (m *Mysql) Dialector() gorm.Dialector {
	return mysql.Open(m.Dsn())

}
