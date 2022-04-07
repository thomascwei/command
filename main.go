package main

import (
	"command/api"
	"command/internal"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

var Trace = log.New(os.Stdout, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	Trace.Println("Starting...")
	defer internal.CtxCancel()
	defer internal.Client.Disconnect(internal.Ctx)


	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/command/V1/get_all_command_template", api.GetAllCommandTemplateRoute)
	r.Run(":9218")

	/*

		   Connect to DB

		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://thomas:123456@localhost:27017"))
		if err != nil {
			log.Fatal(err)
		}
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Disconnect(ctx)


		   //List databases

		databases, err := client.ListDatabaseNames(ctx, bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		Trace.Println(databases)

		// get all commands
		collection := client.Database("command").Collection("command_template")
		cur, err := collection.Find(ctx, bson.D{})
		if err != nil {
			log.Fatal(err)
		}
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var result internal.ReadCommandTemplate
			err := cur.Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
			// do something with result....
			Trace.Println(result.ID.Hex())

			if err := cur.Err(); err != nil {
				log.Fatal(err)
			}
		}
		// find command template by id
		objectId, err := primitive.ObjectIDFromHex("62328f796e40faa6d40df49d")
		doc := collection.FindOne(context.Background(), bson.M{"_id": objectId})
		var result internal.ReadCommandTemplate
		doc.Decode(&result)
		Trace.Println(result)
		Trace.Println(result.Protocol.Headers.Account)

		// create single Template return id
		temp0 := internal.WriteCommandTemplate{
			Name:        "test000",
			Description: "manual insert",
			Protocol: internal.Protocol{
				Protocol: "HTTP",
				Method:   "GET",
				URL:      "https://gorest.co.in/public/v2/users",
				Headers: internal.Header{
					Method:   "basic",
					Account:  "thomas",
					Password: "123456",
				},
			},
			Insert: internal.Insert{
				Check: false,
			},
		}
		res, err := collection.InsertOne(ctx, temp0)
		if err != nil {
			Trace.Println(err)
		}
		Trace.Println(res.InsertedID)

		// delete single template by id
		_, err = collection.DeleteOne(ctx, bson.D{{"_id", res.InsertedID}})
		if err != nil {
			Trace.Println(err)
		}

		// update single template by _id
		zzz, err := collection.UpdateByID(
			ctx,
			objectId,
			bson.D{
				{"$set", temp0},
			},
		)
		if err != nil {
			log.Fatal(err)
		}
		Trace.Printf("Updated %v Documents!\n", zzz.MatchedCount)
	*/
}
