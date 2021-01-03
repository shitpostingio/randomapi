package config

import "go.mongodb.org/mongo-driver/mongo/options"

//MongoDBConnectionOptions does the magic connection stuff
func (c *mongoConfig) MongoDBConnectionOptions() *options.ClientOptions {

	//TODO: CHECK
	//
	clientOptions := options.Client()
	clientOptions.SetHosts(c.Hosts)

	//
	if c.UseAuthentication {
		clientOptions.SetAuth(options.Credential{
			AuthMechanism: c.AuthMechanism,
			AuthSource:    c.AuthSource,
			Username:      c.Username,
			Password:      c.Password,
			PasswordSet:   true,
		})
	}

	//
	if c.UseReplicaSet {
		clientOptions.SetReplicaSet(c.ReplicaSetName)
	}

	return clientOptions

}
