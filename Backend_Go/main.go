package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

	//"firebase.google.com/go/auth"
	"github.com/d-mittal-21/GymkhanaCalendar/Backend_Go/handlers"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func main() {

	client, err := createClient()
	if err != nil {
		log.Printf("error creating client: %v", err)
		return
	}

	r := gin.Default()

	//check health
	r.GET("/api/health", handlers.HealthCheckHandler())

	//create event
	r.POST("/api/event", handlers.CreateEventHandler(client))

	//get all the events
	r.GET("/api/event", handlers.ListEventHandler(client))

	//update an event
	r.PATCH("/api/event/:title", handlers.UpdateEventHandler(client))

	//delete an event
	r.DELETE("/api/event/:title", handlers.DeleteEventHandler(client))

	r.Run("127.0.0.1:3000")

}

func createClient() (*firestore.Client, error) {

	//firebase init
	opt := option.WithCredentialsFile("./serviceKey.json")
	ctx := context.Background()
	// conf := &firebase.Config{ProjectID: "gymkhanacalendar"}
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error when creating client: %v\n", err)
	}

	// // testing function
	// ref := client.Collection("test").NewDoc()
	// result, err := ref.Set(ctx, map[string]interface{}{
	// 	"title": "test1",
	// 	"desc": "testing 123",
	// })
	// if err != nil {
	//     log.Fatalf("error creating test collection: %v\n", err)
	// }
	// log.Printf("Result is [%v]\n", result)

	return client, err

}
