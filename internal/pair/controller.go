package pair

import (
	"net/http"

	"github.com/exedary/soulmates/internal/domain/pair"
	"github.com/gin-gonic/gin"
)

const resourceName = "/pairs"

type Controller struct {
	repository pair.Repository
}

func Register(router *gin.RouterGroup, repository pair.Repository) {
	pairController := &Controller{repository: repository}

	//router.GET(resourceName)
	router.GET(resourceName+"/:id", pairController.get)
	router.POST(resourceName, pairController.create)
}

// @BasePath /api/v1
// Create Pair godoc
//
//	@Summary		Creates a pair without acceptance from another user
//	@Description	add by json account
//	@Tags			pairs
//	@Accept			json
//	@Produce		json
//	@Param			account	body	createDto	true "Create pair"
//	@Success		200		{object} string
//	@Failure		400		{object} string
//	@Failure		422		{object} string
//	@Router			/pairs [post]
func (controller *Controller) create(context *gin.Context) {

	response, err := Create(context.Request.Context(), controller.repository)

	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}

	context.JSON(http.StatusOK, response)
}

// @BasePath	/api/v1
// @Summary		Get pair
// @Description	get pair
// @Tags			pairs
// @Accept			json
// @Produce		json
// @Param			account	path	string	true	"Get pair"
// @Success		200		{object} string
// @Failure		400		{object} string
// @Failure		422		{object} string
// @Router			/pairs [get]
func (controller *Controller) get(context *gin.Context) {
	id := context.Param("id")

	response, err := GetById(context.Request.Context(), controller.repository, id)

	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}

	context.JSON(http.StatusOK, response)
}
