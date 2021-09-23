package configs

import "os"

var Cfgs *SysConfig

type SysConfig struct {
	Port        string
	Env         string
	DBURL       string
	DBTest      string
	TokenSecret string
}

func Get() *SysConfig {
	if Cfgs == nil {
		return fetchConfigs()
	}
	return Cfgs
}

func fetchConfigs() *SysConfig {
	port := os.Getenv("PORT")
	env := os.Getenv("env")
	_db := os.Getenv("DB_URL")
	tokenSecret := os.Getenv("JWT_SECRET")

	if port == "" {
		port = "5000"
	}

	if env == "" {
		env = "local"
	}

	if _db == "" {
		_db = "sell-boot.db"
	}

	if tokenSecret == "" {
		tokenSecret = "secret"
	}

	return &SysConfig{
		Port:   port,
		Env:    env,
		DBURL:  _db,
		DBTest: "test.db",
	}
}
