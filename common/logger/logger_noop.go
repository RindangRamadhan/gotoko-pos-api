package logger

import "context"

type NoopContextLogger struct{}

func (n *NoopContextLogger) Debug(context.Context, string, ...Field) {}

func (n *NoopContextLogger) Info(context.Context, string, ...Field) {}

func (n *NoopContextLogger) Warn(context.Context, string, ...Field) {}

func (n *NoopContextLogger) Error(context.Context, string, ...Field) {}

func (n *NoopContextLogger) Fatal(context.Context, string, ...Field) {}

func (n *NoopContextLogger) Panic(context.Context, string, ...Field) {}

func (n *NoopContextLogger) TDR(LogTDRModel) {}

func (n *NoopContextLogger) Close() {}

func NewNoopLogger() Logger { return &NoopContextLogger{} }
