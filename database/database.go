package database

var connection string

func init() {
	connection = "MySQL"
}

func GetDatabases() string {
	return connection
}
