package pair

import (
	"net/http"

	"github.com/exedary/soulmates/internal/domain/invitation"
	"github.com/exedary/soulmates/internal/domain/pair"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

const resourceName = "/pairs"

var Module = fx.Invoke(Register)

type controller struct {
	pairRepository        pair.Repository
	invitationReposirtory invitation.Repository
}

func Register(router *gin.RouterGroup, pairRepository pair.Repository, invitationRepository invitation.Repository) {
	pairController := &controller{pairRepository: pairRepository, invitationReposirtory: invitationRepository}

	router.GET(resourceName+"/:id", pairController.get)
	//router.POST(resourceName, pairController.create)
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
func (controller *controller) create(context *gin.Context) {
	ctx := context.Request.Context()
	personId := ctx.Value("personId").(string)

	response, err := Create(ctx, controller.pairRepository, controller.invitationReposirtory, personId)

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
// @Produce			json
// @Param			account	path	string	true	"Get pair"
// @Success		200		{object} string
// @Failure		400		{object} string
// @Failure		422		{object} string
// @Router			/pairs [get]
func (controller *controller) get(context *gin.Context) {
	id := context.Param("id")

	response, err := GetById(context.Request.Context(), controller.pairRepository, id)

	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}

	context.JSON(http.StatusOK, response)
}
