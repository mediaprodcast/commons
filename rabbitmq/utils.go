package broker

import (
	"context"

	"go.opentelemetry.io/otel"
)

type AmqpHeaderCarrier map[string]interface{}

func (a AmqpHeaderCarrier) Get(k string) string {
	value, ok := a[k]
	if !ok {
		return ""
	}

	return value.(string)
}

func (a AmqpHeaderCarrier) Set(k string, v string) {
	a[k] = v
}

func (a AmqpHeaderCarrier) Keys() []string {
	keys := make([]string, len(a))
	i := 0

	for k := range a {
		keys[i] = k
		i++
	}

	return keys
}

func InjectAMQPHeaders(ctx context.Context) map[string]interface{} {
	carrier := make(AmqpHeaderCarrier)
	otel.GetTextMapPropagator().Inject(ctx, carrier)
	return carrier
}

func ExtractAMQPHeader(ctx context.Context, headers map[string]interface{}) context.Context {
	return otel.GetTextMapPropagator().Extract(ctx, AmqpHeaderCarrier(headers))
}
