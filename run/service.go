package run

import "time"

type Service struct {
	Repo *Repo
}

func (s Service) Run(req Request) error {
	output := "Test output"

	result := &Result{}
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
