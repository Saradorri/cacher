package utils

type ServiceConfig struct {
	System system `mapstructure:"system"`
	Nodes  []node `mapstructure:"nodes"`
}

type system struct {
	LogPath     string `mapstructure:"log_path"`
	LogName     string `mapstructure:"log_name"`
	DevelopMode bool   `mapstructure:"develop_mode"`
	HTTPPort    int    `mapstructure:"http_port"`
	VNodeCount  int    `mapstructure:"vnode_count"`
}

type node struct {
	Name    string `mapstructure:"name"`
	Address string `mapstructure:"address"`
}
