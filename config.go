package main

type Config struct {
	MongoURL    string `config:",static" default:"mongodb://localhost:27017"`
	MongoDBName string `config:",static" default:"ArtStore"`
}
