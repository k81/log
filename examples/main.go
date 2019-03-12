package main

import (
	"context"

	"github.com/k81/log"
)

func foo(ctx context.Context) {
	fctx := log.WithContext(ctx, "context", "foo")
	log.Info(fctx, "foo called")
}

func bar(ctx context.Context) {
	bctx := log.WithContext(ctx, "context", "bar")
	log.Info(bctx, "bar called")
}

func helloworld(ctx context.Context) {
	hctx := log.WithContext(ctx, "hello", "world")
	log.Info(hctx, "this is hello world")
	log.Tag("__supernova").Info(hctx, "unknown world")
}

func main() {
	mctx := context.Background()

	log.Info(mctx, "program started")
	foo(mctx)
	bar(mctx)
	helloworld(mctx)
	log.Tag("__OK_TAG__").Info(mctx, "program exited")
}
