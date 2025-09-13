package main

import (
	"fmt"
	store "tasks_manager/dataStore"
	handlers "tasks_manager/httpUtils"
	server "tasks_manager/server"
)

func main() {
	dataStore := store.NewDataStore()
	httpHandlers := handlers.NewHTTPHandlers(dataStore)
	serv := server.NewServer(httpHandlers)

	if err := serv.StartServer("9091"); err != nil {
		fmt.Println("failed to start http server:", err)
	}
}
