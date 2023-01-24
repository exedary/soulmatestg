package main

import (
	"context"
	"log"

	docs "github.com/exedary/soulmates/docs"
	"github.com/exedary/soulmates/internal/pair"
	persistence "github.com/exedary/soulmates/internal/persistence/pair"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Run() {

}

func main() {
	app := gin.Default()

	v1 := app.Group("api/v1")
	docs.SwaggerInfo.BasePath = "/api/v1"

	ctx := context.Background()

	mongoClientPool, err := mongo.Connect(ctx,
		options.Client().ApplyURI("mongodb://localhost:27017"),
		/*options.Client().SetAuth(options.Credential{
			Username: "",
			Password: "",
		})*/)

	defer mongoClientPool.Disconnect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	pair.Register(v1, persistence.NewPairRepository(mongoClientPool))
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	app.Run(":8080")
}
