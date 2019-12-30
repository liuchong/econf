package conf

import (
	"fmt"
)

type redis struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func (r *redis) Addr() string {
	return fmt.Sprintf("%s:%s", r.Host, r.Port)
}

// Redis is configuration of redis connection
var Redis = redis{
	Host:     "redis",
	Port:     "6379",
	Password: "",
	DB:       0,
}
