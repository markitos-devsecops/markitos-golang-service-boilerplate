package api

import (
	"markitos-golang-service-boilerplate/internal/services"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Server) boilerCreateHandler(ctx *gin.Context) {
	var request services.BoilerCreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	var service services.BoilerCreateService = services.NewBoilerCreateService(s.repository)
	boiler, err := service.Execute(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	ctx.JSON(http.StatusCreated, boiler)
}

func (s *Server) boilerListHandler(ctx *gin.Context) {
	var service services.BoilerListService = services.NewBoilerListService(s.repository)
	boiler, err := service.Execute()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	ctx.JSON(http.StatusOK, boiler)
}

func (s *Server) boilerOneHandler(ctx *gin.Context) {
	var request services.BoilerOneRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	var service services.BoilerOneService = services.NewBoilerOneService(s.repository)
	boiler, err := service.Execute(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	ctx.JSON(http.StatusOK, boiler)
}

func (s *Server) boilerUpdateHandler(ctx *gin.Context) {
	request, shouldExitByError := createRequestOrExitWithError(ctx)
	if shouldExitByError {
		return
	}

	var service services.BoilerUpdateService = services.NewBoilerUpdateService(s.repository)
	boiler, err := service.Execute(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	ctx.JSON(http.StatusOK, boiler)
}

func (s *Server) boilerSearchHandler(ctx *gin.Context) {
	searchTerm := ctx.Query("search")
	pageNumberStr := ctx.DefaultQuery("page", "1")
	if pageNumberStr == "" {
		pageNumberStr = "1"
	}
	pageSizeStr := ctx.DefaultQuery("size", "10")
	if pageSizeStr == "" {
		pageSizeStr = "10"
	}

	pageNumber, err := strconv.Atoi(pageNumberStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	var service services.BoilerSearchService = services.NewBoilerSearchService(s.repository)
	var request services.BoilerSearchRequest = services.BoilerSearchRequest{
		SearchTerm: searchTerm,
		PageNumber: pageNumber,
		PageSize:   pageSize,
	}
	boiler, err := service.Execute(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	ctx.JSON(http.StatusOK, boiler)
}

func createRequestOrExitWithError(ctx *gin.Context) (services.BoilerUpdateRequest, bool) {
	var requestUri services.BoilerUpdateRequestUri
	if err := ctx.ShouldBindUri(&requestUri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return services.BoilerUpdateRequest{}, true
	}
	var requestBody services.BoilerUpdateRequestBody
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return services.BoilerUpdateRequest{}, true
	}

	var request services.BoilerUpdateRequest = services.BoilerUpdateRequest{
		Id:      requestUri.Id,
		Message: requestBody.Message,
	}

	return request, false
}

func (s *Server) boilerMotdHandler(ctx *gin.Context) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	ctx.JSON(http.StatusOK, "[BOILERPLATE] - Marco Antonio - markitos say, Hi all!! at "+time.Now().Format(time.RFC3339)+" on "+hostname)
}
