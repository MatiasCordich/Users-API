package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/MatiasCordich/Golang-api/controllers"
	"github.com/MatiasCordich/Golang-api/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Vamos a crear alguans variables que van a ser la base para nuestro servidor

var (
	server         *gin.Engine
	userservice    services.UserService
	usercontroller controllers.UserController
	ctx            context.Context
	usercollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func loadEnv() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env")
	}
}

func init() {

	loadEnv()

	MONGO_URI := os.Getenv("MONGO_URL")

	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI(MONGO_URI)

	mongoclient, err = mongo.Connect(ctx, mongoconn)

	if err != nil {
		log.Fatal(err)
	}

	err = mongoclient.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the DB")

	usercollection = mongoclient.Database("userdb").Collection("users")

	userservice = services.NewUserService(usercollection, ctx)

	usercontroller = controllers.New(userservice)

	server = gin.Default()
}

func main() {

	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")

	usercontroller.RegisterUserRoutes(basepath)

	log.Fatal(server.Run("https://usersapi.vercel.app"))

}
