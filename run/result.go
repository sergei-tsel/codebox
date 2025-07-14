package run

import "time"

type Result struct {
	Id        int       // сам уникальный идентификатор
	RequestId int       // идентификатор запроса
	Code      string    // исполняемый код
	Language  string    // язык программирования
	Image     string    // использованный образ
	Output    string    // результат выполнения кода
	CreatedAt time.Time // дата создания результата
}
