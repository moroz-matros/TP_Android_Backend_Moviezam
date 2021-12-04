package config

// Sudo su postgres
//psql -c "CREATE USER android WITH superuser login password 'moviezam’;"
//createdb -O android moviedb

const (
	Port      = ":8088"
	DBConnect = "user=android password=moviezam dbname=moviedb host=localhost port=5432 sslmode=disable pool_max_conns=10"
)
