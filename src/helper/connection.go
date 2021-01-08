package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	c "github.com/coolblknerd/cookie-api/src/config"
	"github.com/spf13/viper"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(col string) *mongo.Collection {
	configs := SetUpConfigs()
	client, err := mongo.NewClient(options.Client().ApplyURI(dbURI(configs)))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Mongo Successfully connected.")
	collection := client.Database(configs.Database.Name).Collection(col)
	return collection
}

type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

func GetError(err error, w http.ResponseWriter) {
	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	_, err = w.Write(message)
	if err != nil {
		log.Println("There was an issue with writing the response.")
	}
}

func SetUpConfigs() c.Configurations {
	viper.SetConfigName("config")
	viper.AddConfigPath("/Users/madblkman/go/src/github.com/coolblknerd/cookie-api")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	var configuration c.Configurations
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return configuration
}

func dbURI(configs c.Configurations) string {
	return fmt.Sprintf("mongodb://%s:%s@%s/%s", configs.Database.User, configs.Database.Password, configs.Database.Host, configs.Database.Name)
}
