package service_candidate

import (
	"context"
	"helicopter-hr/config"
	"helicopter-hr/internal/app_rest/repositories/repo_candidate"
)

type CandidateServiceInterface interface {
	StoreCandidate(ctx context.Context, payload StoreCandidatePayload) error
}

type candidateService struct {
	config        *config.ConfigApp
	repoCandidate repo_candidate.CandidateRepositoryInterface
}

func NewCandidateService(config *config.ConfigApp, repoCandidate repo_candidate.CandidateRepositoryInterface) CandidateServiceInterface {
	return &candidateService{
		config:        config,
		repoCandidate: repoCandidate,
	}
}
