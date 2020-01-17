package main

import (
	"os"
	"requiem-example/api"
	"strconv"

	"github.com/borderstech/requiem"
)

func main() {
	s := requiem.NewServer(
		api.HelloController{})
	s.Port, _ = strconv.Atoi(os.Getenv("PORT"))
	s.Start()
}
