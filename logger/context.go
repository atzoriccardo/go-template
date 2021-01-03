package logger

import (
	"context"

	"github.com/sirupsen/logrus"
)

type loggerkey struct{}

var defaultLogger *logrus.Entry = logrus.NewEntry(logrus.StandardLogger())

// WithLogger retu
func WithLogger(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, loggerkey{}, logger)
}

// Get fun
func Get(ctx context.Context) *logrus.Entry {
	logger := ctx.Value(loggerkey{})

	if logger == nil {
		return defaultLogger
	}

	entry, ok := logger.(*logrus.Entry)

	if !ok {
		return defaultLogger
	}

	return entry
}
