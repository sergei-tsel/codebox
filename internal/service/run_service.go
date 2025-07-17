package service

import (
	"codebox/internal/model"
	"codebox/internal/repository"
	"time"
)

type RunRequest struct {
	Id       int    `json:"id"`
	Code     string `json:"code"`
	Language string `json:"language"`
	Image    string `json:"image"`
}

func Run(req RunRequest) error {
	output := "Test output"

	result := &model.Result{}
	result.RequestId = req.Id
	result.Code = req.Code
	result.Language = req.Language
	result.Image = req.Image
	result.Output = output
	result.CreatedAt = time.Now()

	err := repository.CreateResult(result)

	if err != nil {
		return err
	}

	return nil
}
