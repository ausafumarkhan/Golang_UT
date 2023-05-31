package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongoapi/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://LearnMongoDB:4Ka4XpMZ3Fq8ogFC@cluster0.1plvek0.mongodb.net/?retryWrites=true&w=majority" // Location of database
const dbName = "netflix"
const colName = "watchlist"

//Most Important

var collection *mongo.Collection // mongo.Collection will import the mangoDB driver and collection contains all the data from MongoDB

// Connect with mongoDB using init method which runs at the very first time in entire application

func init() {
	// clinet option -- to bring up the client connection
	clientOption := options.Client().ApplyURI(connectionString) // Need to provide connection string in the ApplyURI function as an argument

	// Connect to mongoDB

	client, err := mongo.Connect(context.TODO(), clientOption) // refer: https://pkg.go.dev/context#TODO for more details
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection established")

	// Lets jump inside the database to get all references using var collection declared at the beginning of code
	collection = client.Database(dbName).Collection(colName)

	// collection instance
	fmt.Println("Connection instance is Ready")
}

// MONGODB helpers -- file

// Insert 1 movie
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in dB with ID:", inserted.InsertedID)

}

// Update 1 record
func updateOneMovie(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified Count:", result.ModifiedCount)
}

// Delete 1 record

func deleteOneMovie(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie git deleted with delete count:", deleteCount)
}

// Delete all records from MongoDB
func deleteAllMovie() int64 {
	// bson.D{{}} --- {} will select all the data
	deleteAllCount, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movies deleted with delete count:", deleteAllCount.DeletedCount)
	return deleteAllCount.DeletedCount
}

// get all movies from MongoDB
func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
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

// Actual Controller -- files

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Method", "POST")
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Method", "PUT")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// Delete one movie
func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Method", "DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// Delete All movies
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Method", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
