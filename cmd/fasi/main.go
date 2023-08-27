package main

import (
	"fmt"

	"github.com/wefasi/fasi/server"
	"github.com/wefasi/fasi/server/infraestructure"
)

func main() {
	infraestructure.InitEnv()
	infraestructure.InitCache()
	infraestructure.InitS3()
	app := server.NewApp()

	fmt.Println("🚀 Listening http://localhost:3210")
	err := app.Listen("localhost:3210")
	if err != nil {
		panic(err)
	}
}
