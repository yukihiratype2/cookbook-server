package model

// Config struct
type Config struct {
	Server ConfigServer
	DB     ConfigDB
	OBS    ConfigOBS
}

// ConfigServer struct
type ConfigServer struct {
	ServerAddress string
	ServerPort    string
	Secret        string
}

// ConfigDB struct
type ConfigDB struct {
	DBAddress  string
	DBUser     string
	DBPassword string
	DBName     string
}

type ConfigOBS struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	Secure          bool
	BucketName      string
	Region          string
}
