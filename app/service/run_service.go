package service

import (
	"codebox/model"
	"codebox/repository"
	"time"
)

type RunRequest struct {
	Id       int    `json:"id"`
	Code     string `json:"code"`
	Language string `json:"language"`
	Image    string `json:"image"`
}

type RunService struct {
	Repo *repository.RunRepo
}

func (s RunService) Run(req RunRequest) error {
	output := "Test output"

	result := &model.Result{}
	result.RequestId = req.Id
	result.Code = req.Code
	result.Language = req.Language
	result.Image = req.Image
	result.Output = output
	result.CreatedAt = time.Now()

	err := s.Repo.CreateResult(result)

	if err != nil {
		return err
	}

	return nil
}
