package main

import (
	"backend/router"
	"backend/setup"
	"fmt"
)

func main() {
	r := router.NewRouter()
	setup.SetupEndpoints(r)
	err := r.Run(":3000")
	if err != nil {
		fmt.Println(err)
	}
}
