package configs

type Server struct {
	Host string
	Port string
}

type KamipaDatabase struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

type SimipaDatabase struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

type JWT struct {
	Key    string
	Issuer string
}

type Redis struct {
	Addr     string
	Password string
}
