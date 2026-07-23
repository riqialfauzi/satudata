package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/satudata/backend/internal/config"
)

// EventType mendefinisikan tipe event yang dipublikasikan.
type EventType string

const (
	EventReleasePublished EventType = "release.published"
	EventReleaseUpdated   EventType = "release.updated"
	EventReleaseDeleted   EventType = "release.deleted"
	EventDatasetProcessed EventType = "dataset.processed"
)

// Event adalah struktur untuk message yang dikirim via NATS.
type Event struct {
	Type      EventType   `json:"type"`
	Timestamp time.Time   `json:"timestamp"`
	Data      interface{} `json:"data"`
}

// NATSClient adalah wrapper untuk koneksi NATS.
type NATSClient struct {
	conn *nats.Conn
	cfg  config.NATSConfig
}

// NewNATSClient membuat koneksi ke NATS server.
func NewNATSClient(cfg config.NATSConfig) (*NATSClient, error) {
	opts := []nats.Option{
		nats.Name(cfg.ClientID),
		nats.Timeout(10 * time.Second),
		nats.ReconnectWait(2 * time.Second),
		nats.MaxReconnects(10),
	}

	conn, err := nats.Connect(cfg.URL, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	client := &NATSClient{
		conn: conn,
		cfg:  cfg,
	}

	log.Println("[NATS] Successfully connected to NATS server")
	return client, nil
}

// PublishEvent mempublikasikan event ke NATS.
func (n *NATSClient) PublishEvent(ctx context.Context, eventType EventType, data interface{}) error {
	event := Event{
		Type:      eventType,
		Timestamp: time.Now().UTC(),
		Data:      data,
	}

	payload, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	subject := string(eventType)
	if err := n.conn.Publish(subject, payload); err != nil {
		return fmt.Errorf("failed to publish event: %w", err)
	}

	// Flush to ensure delivery
	return n.conn.Flush()
}

// SubscribeEvent berlangganan ke event tertentu.
func (n *NATSClient) SubscribeEvent(eventType EventType, handler func(msg *nats.Msg)) (*nats.Subscription, error) {
	subject := string(eventType)
	sub, err := n.conn.Subscribe(subject, handler)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to %s: %w", subject, err)
	}

	_ = n.conn.Flush()
	log.Printf("[NATS] Subscribed to: %s", subject)
	return sub, nil
}

// QueueSubscribeEvent berlangganan ke event dengan queue group (load balancing).
func (n *NATSClient) QueueSubscribeEvent(eventType EventType, queue string, handler func(msg *nats.Msg)) (*nats.Subscription, error) {
	subject := string(eventType)
	sub, err := n.conn.QueueSubscribe(subject, queue, handler)
	if err != nil {
		return nil, fmt.Errorf("failed to queue subscribe to %s: %w", subject, err)
	}

	log.Printf("[NATS] Queue subscribed to %s (queue: %s)", subject, queue)
	return sub, nil
}

// Close menutup koneksi NATS.
func (n *NATSClient) Close() {
	if n.conn != nil && n.conn.IsConnected() {
		n.conn.Close()
		log.Println("[NATS] Connection closed")
	}
}

// IsConnected memeriksa apakah masih terhubung ke NATS.
func (n *NATSClient) IsConnected() bool {
	return n.conn != nil && n.conn.IsConnected()
}
