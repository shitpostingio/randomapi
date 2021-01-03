package config

type mongoConfig struct {
	UseAuthentication bool `type:"optional"`
	UseReplicaSet     bool `type:"optional"`
	DatabaseName      string
	AuthMechanism     string   `type:"optional"`
	Username          string   `type:"optional"`
	Password          string   `type:"optional"`
	AuthSource        string   `type:"optional"`
	ReplicaSetName    string   `type:"optional"`
	Hosts             []string `type:"optional"`
}

// Config is the base structure for all config values
type Config struct {
	PostFolder string
	Endpoint   string
	Port       int
	Mongo      mongoConfig
}
