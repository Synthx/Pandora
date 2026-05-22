package core

import (
	"fmt"
)

type Handler func(message string) (any, error)

type node struct {
	children map[byte]*node
	handler  Handler
}

type MessageTree struct {
	root *node
}

func NewMessageTree() *MessageTree {
	return &MessageTree{
		root: &node{
			children: make(map[byte]*node),
		},
	}
}

func (t *MessageTree) Register(header string, handler Handler) {
	current := t.root
	for i := 0; i < len(header); i++ {
		char := header[i]
		if _, ok := current.children[char]; !ok {
			current.children[char] = &node{
				children: make(map[byte]*node),
			}
		}
		current = current.children[char]
	}
	current.handler = handler
}

func (t *MessageTree) Parse(message string) (any, error) {
	current := t.root
	var lastFoundHandler Handler
	var headerLength int

	for i := 0; i < len(message); i++ {
		char := message[i]
		if next, ok := current.children[char]; ok {
			current = next
			if current.handler != nil {
				lastFoundHandler = current.handler
				headerLength = i + 1
			}
		} else {
			break
		}
	}

	if lastFoundHandler == nil {
		return nil, fmt.Errorf("no handler found for message: %s", message)
	}

	// Pass the payload (everything after the header) to the handler
	return lastFoundHandler(message[headerLength:])
}
