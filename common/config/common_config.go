package config

type (
	WithPortConfig struct {
		Port string `json:"port"`
	}

	WithRedisConfig struct {
		RedisConfig *RedisConfig `json:"redis"`
	}

	RedisConfig struct {
		Address       string `json:"address"`
		PubSubChannel string `json:"pub_sub_channel"`
	}
)
