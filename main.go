package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/ilyassyaf/yeyebackend/config"
	"github.com/ilyassyaf/yeyebackend/controllers"
	"github.com/ilyassyaf/yeyebackend/repository"
	"github.com/ilyassyaf/yeyebackend/routes"
	"github.com/ilyassyaf/yeyebackend/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	ctx         context.Context
	mongoclient *mongo.Client
	redisclient *redis.Client

	userService         services.UserService
	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	counterCollection      *mongo.Collection
	counterService         services.CounterService

	authCollection      *mongo.Collection
	authService         services.AuthService
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	tokenController      controllers.TokenCotroller
	tokenService         services.TokenService
	TokenRouteController routes.TokenRouteController
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load config", err)
	}

	defer mongoclient.Disconnect(ctx)

	value, err := redisclient.Get(ctx, "test").Result()
	if err == redis.Nil {
		fmt.Println("key: test does not exist")
	} else if err != nil {
		panic(err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", "http://localhost:3000"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": value})
	})

	AuthRouteController.AuthRoute(router, userService)
	UserRouteController.UserRoute(router, userService)
	TokenRouteController.TokenRoute(router, userService)
	log.Fatal(server.Run(":" + config.Port))
}

func init() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	ctx = context.TODO()

	// connect mongo
	mongoconn := options.Client().ApplyURI(config.DBUri)
	mongoclient, err := mongo.Connect(ctx, mongoconn)
	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	// connect redis
	redisclient = redis.NewClient(&redis.Options{
		Addr: config.RedisUri,
	})

	if _, err := redisclient.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	err = redisclient.Set(ctx, "test", "Welcome to Golang with Redis and MongoDB", 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis client connected successfully...")

	// collections
	authCollection = mongoclient.Database("golang_mongodb").Collection("users")
	userService = repository.NewUserServiceImpl(authCollection, ctx)
	authService = repository.NewAuthService(authCollection, ctx)
	AuthController = controllers.NewAuthController(authService, userService)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(userService)
	UserRouteController = routes.NewUserRouteController(UserController)

	counterCollection = mongoclient.Database("golang_mongodb").Collection("sequence")
	counterService = repository.NewCounterServiceImpl(counterCollection, ctx)

	tokenService = repository.NewTokenServiceImpl(mongoclient.Database("golang_mongodb"), ctx)
	tokenController = controllers.NewTokenController(tokenService, counterService)
	TokenRouteController = routes.NewTokenRouteController(tokenController)

	server = gin.Default()
}
