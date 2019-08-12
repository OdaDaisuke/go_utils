package db

import "fmt"

type DataBaseSource struct {
	User string
	Password string
	Host string
	Port string
	DBName string
}

func (d DataBaseSource) genDataSource() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", d.User, d.Password, d.Host, d.Port, d.DBName)
}