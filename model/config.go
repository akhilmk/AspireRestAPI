package model

// Struct to hold application configurations.
type Config struct {
	DbInt  bool   `json:"db_init"`
	DbHost string `json:"db_host"`
	DbUser string `json:"db_user"`
	DbName string `json:"db_name"`
}
