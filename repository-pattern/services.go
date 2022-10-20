package main

import (
	"context"
	"design-pattern-in-go/repository"
	"log"
)

type Service struct {
	repo repository.Repository
}

func (s *Service) GetHour(ctx context.Context) {
	if err := s.repo.GetHour(ctx); err != nil {
		log.Fatal(err.Error())
	}
}
