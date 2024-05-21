package config

type System struct {
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	Addr         int    `mapstructure:"addr" json:"addr" yaml:"addr"` // 端口值
	LimitCountIP int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP  int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
	UseRedis     bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"` // 使用redis
}
