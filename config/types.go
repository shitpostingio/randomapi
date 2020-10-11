package config

type mongoConfig struct {
	DatabaseName   string
	Username       string
	Password       string
	AuthSource     string // memesapi
	CollectionName string // requests
	ReplicaSetName string // shitposting
	Hosts          []string
}

// Config is the base structure for all config values
type Config struct {
	MemeFolder string
	Endpoint   string
	Port       int
	MongoMemes mongoConfig
}
