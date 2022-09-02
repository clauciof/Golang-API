package main

import (
  //"github.com/gin-gonic/gin"
  "example.com/helloworld/routers"
)

func main() {
  router := routers.Routes()
  router.Run()
}