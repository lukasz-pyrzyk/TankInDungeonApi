package main

type Config struct {
	Mongo MongoConfig
}

type MongoConfig struct {
	Host     string
	Database string
	Table    string
}
