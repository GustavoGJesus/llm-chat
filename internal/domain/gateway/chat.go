package gateway

import (
	"context"

	"github.com/GustavoGJesus/llm-chat/internal/domain/entity"
)

type ChatGateway interface {
	CreateChat(ctx context.Context, chat *entity.Chat) error
	FindChatById(ctx context.Context, chatID string) (*entity.Chat, error)
	SaveChat(ctx context.Context, chat *entity.Chat) error
}
