package main

import (
	"blog-gin/internal/routers"
	//"github.com/gin-gonic/gin"
)


func main() {
	r := routers.NewRouter()
	r.Run(":8899")
}