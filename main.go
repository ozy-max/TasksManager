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

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start http server:", err)
	}
	server.NewServer(httpHandlers).StartServer(":9091")
}
