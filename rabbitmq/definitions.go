package broker

import amqp "github.com/rabbitmq/amqp091-go"

// The common types are "direct", "fanout", "topic" and "headers".
type ExchangeKind string

const (
	DirectExchangeKind  ExchangeKind = "direct"
	FanoutExchangeKind  ExchangeKind = "fanout"
	TopicExchangeKind   ExchangeKind = "topic"
	HeadersExchangeKind ExchangeKind = "headers"
)

type ExchangeName string

const (
	// Packager events
	PackagerGeneratePlaylistExchangeName ExchangeName = "packager.generate.playlist"
	// Transcoder events
	TranscoderTranscodeMediaExchangeName ExchangeName = "transcoder.transcode.media"
)

type QueueName string

const MessagePersistent = amqp.Persistent
