package config

type Config struct {
	DB DataBase
}

type DataBase struct {
	DriverName string
	URL        string
}
