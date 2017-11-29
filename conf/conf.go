package conf

/*
	DB
*/
const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "min0917"
	DbName   = "ev"
)

/*
	EV Charger API
*/
const (
	URL        = "http://open.ev.or.kr:8080/openapi/services/rest/EvChargerService?ServiceKey="                     // 공공포털 서비스 주소
	ServiceKey = "traoC%2BJDroxln0JBNb6JmeXPWdGXzk%2FzhIUTO5QRx5HCZnAMUsXT7NtoMz9gLSRYVEfO8svYbdPQW3gYCisraA%3D%3D" // 인증키 값
)
