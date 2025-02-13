package broker

import (
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Message struct {
	Key       string
	Exchange  ExchangeName
	Mandatory bool
	Immediate bool
	*amqp.Publishing
}

type MessageBuilder struct {
	message *Message
}

func NewMessageBuilder() *MessageBuilder {
	return &MessageBuilder{
		message: &Message{
			Publishing: &amqp.Publishing{},
		},
	}
}

func (b *MessageBuilder) WithExchange(exchange ExchangeName) *MessageBuilder {
	b.message.Exchange = exchange
	return b
}

func (b *MessageBuilder) WithKey(key string) *MessageBuilder {
	b.message.Key = key
	return b
}

func (b *MessageBuilder) WithMandatory(mandatory bool) *MessageBuilder {
	b.message.Mandatory = mandatory
	return b
}

func (b *MessageBuilder) WithImmediate(immediate bool) *MessageBuilder {
	b.message.Immediate = immediate
	return b
}

func (b *MessageBuilder) WithHeaders(headers amqp.Table) *MessageBuilder {
	b.message.Publishing.Headers = headers
	return b
}

func (b *MessageBuilder) WithContentType(contentType string) *MessageBuilder {
	b.message.Publishing.ContentType = contentType
	return b
}

func (b *MessageBuilder) WithContentEncoding(contentEncoding string) *MessageBuilder {
	b.message.Publishing.ContentEncoding = contentEncoding
	return b
}

func (b *MessageBuilder) WithDeliveryMode(deliveryMode uint8) *MessageBuilder {
	b.message.Publishing.DeliveryMode = deliveryMode
	return b
}

func (b *MessageBuilder) WithPriority(priority uint8) *MessageBuilder {
	b.message.Publishing.Priority = priority
	return b
}

func (b *MessageBuilder) WithCorrelationId(correlationId string) *MessageBuilder {
	b.message.Publishing.CorrelationId = correlationId
	return b
}

func (b *MessageBuilder) WithReplyTo(replyTo string) *MessageBuilder {
	b.message.Publishing.ReplyTo = replyTo
	return b
}

func (b *MessageBuilder) WithExpiration(expiration string) *MessageBuilder {
	b.message.Publishing.Expiration = expiration
	return b
}

func (b *MessageBuilder) WithMessageId(messageId string) *MessageBuilder {
	b.message.Publishing.MessageId = messageId
	return b
}

func (b *MessageBuilder) WithTimestamp(timestamp time.Time) *MessageBuilder {
	b.message.Publishing.Timestamp = timestamp
	return b
}

func (b *MessageBuilder) WithType(messageType string) *MessageBuilder {
	b.message.Publishing.Type = messageType
	return b
}

func (b *MessageBuilder) WithUserId(userId string) *MessageBuilder {
	b.message.Publishing.UserId = userId
	return b
}

func (b *MessageBuilder) WithAppId(appId string) *MessageBuilder {
	b.message.Publishing.AppId = appId
	return b
}

func (b *MessageBuilder) WithBody(body []byte) *MessageBuilder {
	b.message.Publishing.Body = body
	return b
}

func (b *MessageBuilder) Build() *Message {
	return b.message
}
