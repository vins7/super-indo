package db

type DatabaseList struct {
	SuperIndo Database
	Redis     DB
}

// Database :
type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

// Database :
type DB struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
