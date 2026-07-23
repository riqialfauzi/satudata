package messaging

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/satudata/backend/pkg/storage"
)

// MessageHandler adalah handler untuk memproses event dari NATS.
type MessageHandler struct {
	storageClient *storage.StorageClient
}

// NewMessageHandler membuat instance baru MessageHandler.
func NewMessageHandler(storageClient *storage.StorageClient) *MessageHandler {
	return &MessageHandler{
		storageClient: storageClient,
	}
}

// ProcessDataset memproses dataset (contoh: validasi, transformasi).
func (h *MessageHandler) ProcessDataset(msg *nats.Msg) {
	var event Event
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		log.Printf("[NATS Handler] Failed to unmarshal event: %v", err)
		return
	}

	log.Printf("[NATS Handler] Processing dataset event: %s", event.Type)
	// TODO: Implement actual dataset processing logic
	// - Validate file format
	// - Parse CSV/JSON/Excel
	// - Extract column information
	// - Count rows
	// - Update dataset_metadata
	log.Printf("[NATS Handler] Dataset processing completed for event: %s", event.Type)
}

// SendNotification mengirim notifikasi (contoh: email, SMS).
func (h *MessageHandler) SendNotification(msg *nats.Msg) {
	var event Event
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		log.Printf("[NATS Handler] Failed to unmarshal event: %v", err)
		return
	}

	log.Printf("[NATS Handler] Sending notification for event: %s", event.Type)
	// TODO: Implement actual notification logic
	// - Send email to subscribers
	// - Send webhook notifications
}

// GenerateThumbnail membuat thumbnail untuk gambar artikel.
func (h *MessageHandler) GenerateThumbnail(msg *nats.Msg) {
	var event Event
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		log.Printf("[NATS Handler] Failed to unmarshal event: %v", err)
		return
	}

	log.Printf("[NATS Handler] Generating thumbnail for event: %s", event.Type)
	// TODO: Implement actual thumbnail generation logic
}
