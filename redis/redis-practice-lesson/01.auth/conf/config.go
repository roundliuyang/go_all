package conf

import "sync"

type Config struct {
	Language string
	Token    string
	Super    string
	RedisPre string
	Host     string
	OpenJwt  bool
	Routes   []string
}

var (
	Cfg   Config
	mutex sync.Mutex
	once  sync.Once
)

const (
	AccessToken  = "access_token"
	RefreshToken = "refresh_token"
)

func Set(cfg Config) {
	mutex.Lock()
	defer mutex.Unlock()
	Cfg.RedisPre = setDefault(cfg.RedisPre, "", "auth.redis")
	Cfg.Language = setDefault(cfg.Language, "", "cn")
	Cfg.Token = setDefault(cfg.Token, "", "token")
	Cfg.Super = setDefault(cfg.Super, "", "admin")
	Cfg.Host = setDefault(cfg.Host, "", "http://localhost:9100")
	Cfg.Routes = cfg.Routes
	Cfg.OpenJwt = cfg.OpenJwt
}

func setDefault(value, def, defValue string) string {
	if value == def {
		return defValue
	}
	return value
}

type MyJWT struct {
	Uid          int64
	AccessToken  string
	RefreshToken string
}
