package config

type WithPortConfig struct {
	Port string `json:"port"`
}

type WithRedisConfig struct {
	RedisConfig *RedisConfig `json:"redis"`
}

type RedisConfig struct {
	Address       string `json:"address"`
	PubSubChannel string `json:"pub_sub_channel"`
}
