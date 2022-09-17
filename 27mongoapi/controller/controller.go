package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/itzganesh03/Golang-Practical-/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://ganesh:<password>@cluster0.hgnj086.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// MOST IMPORTANT COLLECTION

var collection *mongo.Collection

// connect with mongoDB

func init() {
	//client syntx
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		// panic(err)
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Success")

	collection = client.Database(dbName).Collection(colName)

	// Collection instence ready for me

	fmt.Println("Collection instence is Ready")
}

// Mongo helpers -file

// insert 1 record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in DB with id", inserted.InsertedID)
}

// Update One record

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	// Bsno.M is Used For shorter and Clearnce Result

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("Modified Conntent:", result.ModifiedCount)
}

// Delete One Record

func deleteOneMoive(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie got deleted with delete count", deleteCount)
}

// delete all record from Mongo DB

func deleteAllMobie() int64 {

	// filter := bson.D{{}}
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number Of Movie is Deleted", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// Get All Movie From DataBase

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	// for-loop but style of while-loop

	for cur.Next(context.Background()) {
		var movie bson.M

		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}

//main Controller in this file

func GetMyAllMobies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMoive(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	// params := mux.Vars(r)
	// deleteOneMoive(params["id"])
	count := deleteAllMobie()
	json.NewEncoder(w).Encode(count)
}
