package pair

import "github.com/gin-gonic/gin"

const resourceName = "/pairs"

func RegisterHandlers(router *gin.Engine, service *IService) {
	pairHandler := &pairHandler{service: *service}

	router.GET(resourceName)
	router.GET(resourceName+":id", pairHandler.get)
	router.POST(resourceName, pairHandler.create)
}

type pairHandler struct {
	service IService
}

func (handler *pairHandler) create(context *gin.Context) {
	dto := &CreatePairDto{}

	if err := context.ShouldBindJSON(dto); err != nil {
		context.AbortWithError(422, err)
	}

	id, err := handler.service.Create(context.Request.Context(), dto)

	if err != nil {
		context.AbortWithError(400, err)
	}

	context.JSON(200, gin.H{
		"id": id,
	})
}

func (handler *pairHandler) get(context *gin.Context) {
	dto := &GetByIdQueryDto{}

	if err := context.ShouldBindUri(dto); err != nil {
		context.AbortWithError(422, err)
	}

	pair, err := handler.service.GetById(context.Request.Context())

	if err != nil {
		context.AbortWithError(400, err)
	}

	context.JSON(200, pair)
}
