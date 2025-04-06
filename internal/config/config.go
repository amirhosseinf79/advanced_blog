package config

type Conf struct {
	DbString        string
	TokenExpireTime int
}

var Config Conf = Conf{
	DbString:        "host=localhost user=postgres password=Amir2001 dbname=blog_advanced port=5432 sslmode=disable TimeZone=Asia/Tehran",
	TokenExpireTime: 8,
}
