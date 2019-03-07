package main

import (
	"context"

	"github.com/k81/log"
)

func foo(ctx context.Context) {
	logger := log.With("context", "foo")

	logger.Info(ctx, "foo called")
}

func bar(ctx context.Context) {
	logger := log.With("context", "bar")

	logger.Info(ctx, "bar called")
}

func helloworld(ctx context.Context) {
	logger := log.With("context", "hello")
	logger = logger.With("context", "world")
	logger.Info(ctx, "this is hello world")
}

func main() {
	ctx := context.TODO()

	log.Info(ctx, "program started")
	foo(ctx)
	bar(ctx)
	helloworld(ctx)
	log.Tag("__OK_TAG__").Info(ctx, "program exited")
}
