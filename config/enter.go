package config

// Config 包含所有配置信息
type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
	Jwt    Jwt    `yaml:"jwt"`
	Redis  Redis  `yaml:"redis"`
}
