package main

import (
	artDb "ArtStore/artStore"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *artDb.DbArts

func main() {
	LoadConfig()
	client, err := Init()
	if err != nil {
		log.Fatal("Initialization error:", err)
	}
	defer client.Disconnect(context.Background())
	if err != nil {
		log.Println("Error creating customer:", err)
	}
	DB = &artDb.DbArts{
		ArtsCollection: client.Database("ArtsDB").Collection("Arts"),
	}
	startServer()
}

func Init() (*mongo.Client, error) {
	url := viper.GetString("MongoURL")
	log.Println("MongoDB URL:", url)

	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		return nil, fmt.Errorf("error creating MongoDB client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %w", err)
	}

	log.Println("Connected to MongoDB successfully")
	return client, nil
}

func LoadConfig() {
	viper.SetConfigFile("config.toml")
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("viper read configuration failed, %s", err)
	}

}

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Go backend!")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server is healthy!")
	})
	http.Handle("/artworks", corsMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			InsertArtworkHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))
	// List Artworks
	http.Handle("/listartworks", corsMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			ListArtWorks(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	port := ":8090" // Specify the port
	log.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow Angular origin
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5300")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight OPTIONS request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the actual handler
		next.ServeHTTP(w, r)
	})
}
func InsertArtworkHandler(w http.ResponseWriter, res *http.Request) {
	log.Println("InsertArtworkHandler")
	store := &artDb.ArtStore{}
	if err := json.NewDecoder(res.Body).Decode(store); err != nil {
		log.Println("while inserting art error:", err)
	}
	UUID, _ := uuid.NewRandom()
	store.ArtID = UUID.String()
	store.CreatedDate = time.Now()
	if _, err := DB.ArtsCollection.InsertOne(context.Background(), &store); err != nil {
		log.Println("error while inserting arts: ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store)
}

func ListArtWorks(w http.ResponseWriter, res *http.Request) {
	log.Println("InsertArtworkHandler")
	store := &artDb.ArtStore{}
	if err := json.NewDecoder(res.Body).Decode(store); err != nil {
		log.Println("while inserting art error:", err)
	}
	query := bson.M{}
	if store.ArtName != "" {
		query["artName"] = store.ArtName
	}
	cursor, err := DB.ArtsCollection.Find(context.Background(), store)
	if err != nil {
		log.Println("error while fetching arts: ", err)

	}
	defer cursor.Close(context.Background())
	var result []artDb.ArtStore
	if err := cursor.All(context.Background(), &result); err != nil {
		log.Println("error in decoding from cursor")
		return
	}
	w.Header().Set("Context-Type", "application/json")
	json.NewEncoder(w).Encode(&result)
}
