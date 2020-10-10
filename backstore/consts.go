package backstore

import "fmt"

const (
	mysqlConnectionStringTemplate string = "%s:%s@tcp(%s)/%s?charset=utf8mb4,utf8&parseTime=True&loc=Local"
)

func dbConnString(username, password, host, database string) string {
	return fmt.Sprintf(mysqlConnectionStringTemplate, username, password, host, database)
}
