package main

import (
	"fmt"

	"github.com/xterminator24/bootdev-blog-gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}

	if err := cfg.SetUser("xterminator24"); err != nil {
		panic(err)
	}

	cfg2, err := config.Read()
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg2)
}
