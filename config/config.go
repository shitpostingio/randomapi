package config

import "go.mongodb.org/mongo-driver/mongo/options"

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
