package service

import (
	"codebox/internal/model"
	"codebox/internal/repository"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type RunRequest struct {
	Id       int    `json:"id"`
	Code     string `json:"code"`
	Language string `json:"language"`
	Image    string `json:"image"`
}

func Run(req RunRequest) error {
	// Создание временной директории с кодом
	tempDir, _ := os.MkdirTemp("", "code-")
	defer os.RemoveAll(tempDir)

	codeFilePath := ""

	if req.Language == "golang" {
		codeFilePath = filepath.Join(tempDir, "code.go")
	} else {
		codeFilePath = filepath.Join(tempDir, "code")
	}

	codeFile, _ := os.OpenFile(codeFilePath, os.O_CREATE|os.O_WRONLY, 0644)
	defer codeFile.Close()
	codeFile.WriteString(req.Code)

	if req.Language == "golang" {
		// Сборка исполняемого файла для Go
		buildCmd := exec.Command("go", "build", "-o", "main", codeFilePath)
		buildCmd.Dir = tempDir
		buildCmd.Run()
		codeFilePath = filepath.Join(tempDir, "main")
	}

	// Запуск контейнера с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	args := []string{"run", "--rm", "-v", fmt.Sprintf("%s:/app", tempDir), req.Image}

	if req.Language == "golang" {
		args = append(args, "/app/"+filepath.Base(codeFilePath))
	}

	dockerCmd := exec.CommandContext(ctx, "docker", args[0:]...)
	command := getExecutionCommand(req.Language)
	dockerCmd.Args = append(dockerCmd.Args, command...)

	// Запуск кода
	byteResult, err := dockerCmd.CombinedOutput()

	output := ""

	switch {
	case errors.Is(ctx.Err(), context.DeadlineExceeded):
		output = "timeout exceeded"
	case err != nil:
		output = err.Error()
		return fmt.Errorf("%s", string(byteResult))
	default:
		output = string(byteResult)
	}

	// Сохранение результата
	result := &model.Result{}
	result.RequestId = req.Id
	result.Code = req.Code
	result.Language = req.Language
	result.Image = req.Image
	result.Output = output
	result.CreatedAt = time.Now()

	err = repository.CreateResult(result)

	if err != nil {
		return err
	}

	return nil
}

func getExecutionCommand(language string) []string {
	switch language {
	case "golang":
		return []string{"go", "run", "/app/code.go"}
	case "python":
		return []string{"python", "/app/code"}
	case "php":
		return []string{"php", "/app/code"}
	default:
		return []string{"/app/code"}
	}
}
