package config

import "go.mongodb.org/mongo-driver/mongo/options"

type mongoConfig struct {
	DatabaseName   string
	Username       string
	Password       string
	AuthSource     string // memesapi
	CollectionName string // requests
	ReplicaSetName string // shitposting
	Hosts          []string
}

type mariaDBConfig struct {
	Username     string
	Password     string
	Host         string
	DatabaseName string
}

// Config is the base structure for all config values
type Config struct {
	MemeFolder        string
	MemeSymlinkFolder string
	Endpoint          string
	StorageEndpoint   string
	Port              int
	MariaDB           mariaDBConfig
	Mongo             mongoConfig
}

//MongoDBConnectionOptions does the magic connection stuff
func (c *mongoConfig) MongoDBConnectionOptions() *options.ClientOptions {

	clientOptions := options.Client()
	clientOptions.SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		AuthSource:    c.AuthSource,
		Username:      c.Username,
		Password:      c.Password,
		PasswordSet:   true,
	})

	clientOptions.SetHosts(c.Hosts)
	clientOptions.SetReplicaSet(c.ReplicaSetName)
	return clientOptions
}
