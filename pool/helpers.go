package pool

import "go.uber.org/zap"

type poolOption = func(*Pool)

func WithPoolLogger(logger *zap.Logger) func(*Pool) {
	return func(p *Pool) {
		p.logger = logger
	}
}

func WithPoolConcurrency(concurrency int) func(*Pool) {
	return func(p *Pool) {
		p.concurrency = concurrency
	}
}

func WithPoolName(name string) func(*Pool) {
	return func(p *Pool) {
		p.name = name
	}
}
