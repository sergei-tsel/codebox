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
	tempDir, err := os.MkdirTemp("", "code-")
	defer os.RemoveAll(tempDir)

	codeFilePath := filepath.Join(tempDir, "code")
	codeFile, _ := os.OpenFile(codeFilePath, os.O_CREATE|os.O_WRONLY, 0644)
	codeFile.WriteString(req.Code)
	defer codeFile.Close()

	// Создание временного файла докера
	dockerfilePath := filepath.Join(tempDir, "Dockerfile")
	dockerfile, _ := os.OpenFile(dockerfilePath, os.O_CREATE|os.O_WRONLY, 0644)
	defer dockerfile.Close()

	dockerfileContent := generateDockerfile(req.Language, req.Image)
	dockerfile.WriteString(dockerfileContent)

	// Сборка образа докера
	imageName := fmt.Sprintf("temp-%s-%d", req.Language, req.Id)
	buildCmd := exec.Command(
		"docker",
		"build",
		"-t",
		imageName,
		"-f",
		dockerfilePath,
		tempDir,
	)
	buildCmd.CombinedOutput()

	defer func() {
		removeCmd := exec.Command("docker", "rmi", "-f", imageName)
		removeCmd.Run()
	}()

	// Запуск контейнера с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dockerCmd := exec.CommandContext(
		ctx,
		"docker",
		"run",
		"--rm",
		imageName,
	)

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

func generateDockerfile(language, image string) string {
	switch language {
	case "golang":
		return fmt.Sprintf(`
            FROM %s AS builder
            WORKDIR /app
            COPY code code.go
            RUN CGO_ENABLED=0 GOOS=linux go build -o app code.go
			RUN chmod +x app

            FROM scratch
            WORKDIR /app
            COPY --from=builder /app/app .
            ENTRYPOINT ["/app/app"]
        `, image)
	case "python":
		return fmt.Sprintf(`
            FROM %s
            WORKDIR /app
            COPY code code.py
            ENTRYPOINT ["python", "code.py"]
        `, image)
	case "javascript":
		return fmt.Sprintf(`
            FROM %s
            WORKDIR /app
            COPY code code.js
			RUN chmod +x code.js
            ENTRYPOINT ["node", "code.js"]
        `, image)
	case "ruby":
		return fmt.Sprintf(`
            FROM %s
            WORKDIR /app
            COPY code code.rb
			RUN chmod +x code.rb
            ENTRYPOINT ["ruby", "code.rb"]
        `, image)
	case "php":
		return fmt.Sprintf(`
            FROM %s
            WORKDIR /app
            COPY code code.php
			RUN chmod +x code.php
            ENTRYPOINT ["php", "code.php"]
        `, image)
	case "bash":
		return fmt.Sprintf(`
            FROM %s
            WORKDIR /app
            COPY code code.sh
            RUN chmod +x code.sh
            ENTRYPOINT ["/app/code.sh"]
        `, image)
	default:
		return fmt.Sprintf(`
            FROM %s
            WORKDIR /app
            COPY code code
			RUN chmod +x code
            ENTRYPOINT ["/app/code"]
        `, image)
	}
}
