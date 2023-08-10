package api

import (
	"log"
	"rotten_world_server/api/routers"
)

func Run() {
	r := routers.UseGlobalRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
