package handler

import (
	"context"
	"fmt"
	"github.com/muhadif/wa-bot-wrapper/internal/biz"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types/events"
)

type WAHandler interface {
	RunSubscriberMessage()
	ListenMessage(event interface{})
}

type waHandler struct {
	waClient *whatsmeow.Client
	waUC     biz.WAUseCase
}

func NewWAHandler(waClient *whatsmeow.Client, waUC biz.WAUseCase) WAHandler {
	return &waHandler{
		waClient: waClient,
		waUC:     waUC,
	}
}

func (w waHandler) RunSubscriberMessage() {
	w.waClient.AddEventHandler(w.ListenMessage)
}

func (w waHandler) ListenMessage(event interface{}) {
	switch v := event.(type) {
	case *events.Message:
		ctx := context.Background()
		reply, err := w.waUC.ParseMessageAndDoCallback(ctx, v.Message.GetConversation())
		if err != nil {
			return
		}
		sentTo := v.Info.Sender
		if v.Info.IsGroup {
			sentTo = v.Info.Chat
		}
		_, err = w.waClient.SendMessage(ctx, sentTo, &waProto.Message{
			Conversation: &reply,
		})
		if err != nil {
			fmt.Println("err", err)
		}
	}
}
