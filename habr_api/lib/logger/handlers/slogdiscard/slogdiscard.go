package slogdiscard

import (
	"context"
	"log/slog"
)

func NewDiscardLogger() *slog.Logger {
	return slog.New(NewDiscardHandler())
}

type DiscardHandler struct{}

func NewDiscardHandler() *DiscardHandler {
	return &DiscardHandler{}
}
func (h *DiscardHandler) Handle(_ context.Context, _ slog.Record) error {
	// игнорируем запись журнала
	return nil
}

func (h *DiscardHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	//возвращает тот же обработчик, так как нет аттрибутов для сохранения
	return h
}

func (h *DiscardHandler) WithGroup(_ string) slog.Handler {
	//возвращает тот же обработчик, так как нет группы для сохранения
	return h
}
func (h *DiscardHandler) Enabled(_ context.Context, _ slog.Level) bool {
	// всегда возвращает false, так как запись журнала игнорируется
	return false
}
