package config

var GlobalConfig Config

type Config struct {
	Mongo MongoConfig
}

type MongoConfig struct {
	Host string
	Port string
}