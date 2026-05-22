package login

import (
	"context"
	"pandora/core"
	"pandora/internal/login/message"
)

type MessageDispatcher struct {
	tree *core.MessageTree
}

func NewMessageDispatcher(handler *SessionHandler) *MessageDispatcher {
	dispatcher := &MessageDispatcher{
		tree: core.NewMessageTree(),
	}

	dispatcher.registerHandlers(handler)

	return dispatcher
}

func (d *MessageDispatcher) registerHandlers(h *SessionHandler) {
	d.tree.Register("Af", func(payload string) (any, error) {
		msg := message.NewQueuePositionMessage()

		return func(ctx context.Context, c *Client) error {
			return h.HandleAskQueuePosition(ctx, c, msg)
		}, nil
	})
}

func (d *MessageDispatcher) Dispatch(ctx context.Context, c *Client, rawMessage string) error {
	result, err := d.tree.Parse(rawMessage)
	if err != nil {
		return err
	}

	// Cast the result to the expected handler function type and execute it
	if handler, ok := result.(func(context.Context, *Client) error); ok {
		return handler(ctx, c)
	}

	return nil
}
