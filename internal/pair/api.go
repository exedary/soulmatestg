package pair

import "github.com/gin-gonic/gin"

const resourceName = "/pairs"

func RegisterHandlers(router *gin.RouterGroup, service *Service) {
	pairHandler := &pairHandler{service: *service}

	//router.GET(resourceName)
	router.GET(resourceName+"/:id", pairHandler.get)
	router.POST(resourceName, pairHandler.create)
}

type pairHandler struct {
	service Service
}

// @BasePath /api/v1
// Create Pair godoc
//
//	@Summary		Creates a pair without acceptance from another user
//	@Description	add by json account
//	@Tags			pairs
//	@Accept			json
//	@Produce		json
//	@Param			account	body	createPairDto	true "Create pair"
//	@Success		200		{object} string
//	@Failure		400		{object} string
//	@Failure		422		{object} string
//	@Router			/pairs [post]
func (handler *pairHandler) create(context *gin.Context) {
	dto := &createPairDto{}

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

//	@BasePath	/api/v1
//	@Summary		Get pair
//	@Description	get pair
//	@Tags			pairs
//	@Accept			json
//	@Produce		json
//	@Param			account	path	string	true	"Get pair"
//	@Success		200		{object} string
//	@Failure		400		{object} string
//	@Failure		422		{object} string
//	@Router			/pairs [get]
func (handler *pairHandler) get(context *gin.Context) {
	value, success := context.Params.Get("id")
	if !success {
		context.AbortWithStatus(422)
	}

	pair, err := handler.service.GetById(context.Request.Context(), value)

	if err != nil {
		context.AbortWithError(400, err)
	}

	context.JSON(200, pair)
}
