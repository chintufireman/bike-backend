package main

import (
	"bike-website/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// // Initialize the database connection
	// client, err := db.InitDatabase()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // You can now use the 'client' to interact with the MongoDB database.

	// // Create a new context for the database operation
	// ctx := context.Background()
	// // options := make(map[string]interface{})
	// // List the databases
	// databases, err := client.ListDatabaseNames(ctx, map[string]interface{}{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Iterate through the database names and print them
	// for _, dbName := range databases {
	// 	fmt.Println(dbName)
	// }

	// // Don't forget to defer client.Disconnect() to close the connection when your application exits.
	// defer client.Disconnect(context.Background())

	// // Your application logic goes here.

	r := mux.NewRouter()
	r.HandleFunc("/new-bike", controller.Demo).Methods("POST")

	http.Handle("/", r)

	// Start the HTTP server
	http.ListenAndServe(":9090", nil)

}
