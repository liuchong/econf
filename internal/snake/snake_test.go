package snake

import (
	"fmt"
	"testing"
)

func TestSnake(t *testing.T) {
	var s string
	var snakeS string
	var expectS string

	s = "ver1"
	snakeS = ToSnake(s, '_', true)
	fmt.Printf("%s -> %s\n", s, snakeS)

	s = "Ver1"
	snakeS = ToSnake(s, '_', true)
	fmt.Printf("%s -> %s\n", s, snakeS)

	s = "Ver 1"
	snakeS = ToSnake(s, '_', true)
	fmt.Printf("%s -> %s\n", s, snakeS)

	s = "Ver 1"
	snakeS = ToSnake(s, '_', false)
	fmt.Printf("%s -> %s\n", s, snakeS)

	s = "Ver 101"
	snakeS = ToSnake(s, '_', false)
	fmt.Printf("%s -> %s\n", s, snakeS)

	s = "ver101"
	snakeS = ToSnake(s, '_', false)
	fmt.Printf("%s -> %s\n", s, snakeS)

	s = "EConf JSONVer101beta.3"
	snakeS = ToSnake(s, '_', true)
	fmt.Printf("%s -> %s\n", s, snakeS)

	expectS = "E_CONF_JSON_VER_101_BETA_3"
	if snakeS != expectS {
		t.Errorf("snake test failed! original string: %s, snake string: %s, expected string: %s",
			s, snakeS, expectS)
	}
}
