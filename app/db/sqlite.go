package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Sqlite struct {
	DBName string
}

func (s *Sqlite) Dsn() string {
	return s.DBName
}

func (s *Sqlite) Dialector() gorm.Dialector {
	return sqlite.Open(s.Dsn())
}
