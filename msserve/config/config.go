package config

type Config struct {
	HostAddr string
	ConnStr  string
	TmpDir   string
}

func Init() *Config {
	return &Config{
		HostAddr: "127.0.0.1:6789",
		ConnStr:  "port=9876 user=postgres password=asd dbname=asd sslmode=disable host=localhost",
		TmpDir:   "./templates/*",
	}
}
