package handler_candidate

import (
	"github.com/gin-gonic/gin"
	"helicopter-hr/config"
	"helicopter-hr/internal/app_rest/middleware/jwtx"
	"helicopter-hr/internal/app_rest/service/service_candidate"
)

type Handler struct {
	cfg              *config.ConfigApp
	jwtAuth          jwtx.AuthenticationInterface
	candidateService service_candidate.CandidateServiceInterface
}

func NewHandlerUser(
	cfg *config.ConfigApp,
	router *gin.Engine,
	jwtAuth jwtx.AuthenticationInterface,
	candidateService service_candidate.CandidateServiceInterface) {
	handler := &Handler{
		cfg:              cfg,
		jwtAuth:          jwtAuth,
		candidateService: candidateService,
	}

	v1 := router.Group("/api/v1/candidate")
	{
		v1.POST("", handler.HandlerStoreCandidate)

	}
}
