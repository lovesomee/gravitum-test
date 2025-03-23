package config

type Settings struct {
	Port     int      `json:"port"`
	Database Database `json:"database"`
}

type Database struct {
	PostgresConnection string `json:"postgresConnection"`
}
