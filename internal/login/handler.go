package login

import (
	"context"
	"errors"
	"pandora/internal/login/message"
	"pandora/internal/login/repository"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

const (
	desiredVersion = "1.48.10"
)

type SessionHandler struct {
	accountRepo *repository.AccountRepository
	log         *zap.Logger
}

func NewSessionHandler(accountRepo *repository.AccountRepository, logger *zap.Logger) *SessionHandler {
	return &SessionHandler{
		accountRepo: accountRepo,
		log:         logger.Named("handler"),
	}
}

func (h *SessionHandler) HandleAskQueuePosition(ctx context.Context, c *Client, _ message.QueuePositionMessage) error {
	if c.Version != desiredVersion {
		return c.Send(message.NewInvalidVersionMessage(desiredVersion))
	}

	account, err := h.accountRepo.GetByToken(ctx, c.Token)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Send(message.NewInvalidPasswordMessage())
		}

		return err
	}

	if account.IsBanned {
		if account.BannedUntil.IsZero() {
			return c.Send(message.NewAccountBannedMessage())
		}

		return c.Send(message.NewAccountTempBannedMessage(account.BannedUntil))
	}

	if err := c.Send(message.NewQueuePositionMessage()); err != nil {
		return err
	}

	if len(account.Username) == 0 {
		c.state = StateWaitingNickname

		return c.Send(message.NewChooseNicknameMessage())
	}

	return nil
}
