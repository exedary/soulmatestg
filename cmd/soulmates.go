package main

import (
	"context"
	"log"

	docs "github.com/exedary/soulmates/docs"
	"github.com/exedary/soulmates/internal/auth"
	"github.com/exedary/soulmates/internal/pair"
	pairPersistence "github.com/exedary/soulmates/internal/persistence/pair"
	personPersistence "github.com/exedary/soulmates/internal/persistence/person"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Run() {

}

func main() {
	app := gin.Default()

	v1 := app.Group("api/v1")
	docs.SwaggerInfo.BasePath = "/api/v1"

	store := cookie.NewStore([]byte("secret"))
	app.Use(sessions.Sessions("default", store))

	ctx := context.Background()

	mongoClientPool, err := mongo.Connect(ctx,
		options.Client().ApplyURI("mongodb://localhost:27017"),
		/*options.Client().SetAuth(options.Credential{
			Username: "",
			Password: "",
		})*/)

	if err != nil {
		log.Fatal(err)
	}

	if err := mongoClientPool.Ping(ctx, readpref.Nearest()); err != nil {
		log.Fatal(err)
	}

	defer mongoClientPool.Disconnect(ctx)

	auth.Register(&app.RouterGroup, personPersistence.NewPersonRepository(mongoClientPool))
	pair.Register(v1, pairPersistence.NewPairRepository(mongoClientPool))
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	app.Run(":8080")
}
