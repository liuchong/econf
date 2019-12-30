package main

import (
	"fmt"

	"fromfile/conf"
)

func main() {
	fmt.Printf("Redis: %+v\n", conf.Redis)
	fmt.Printf("Postgres: %+v\n", conf.Postgres)
}
