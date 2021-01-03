package logger

import (
	"github.com/sirupsen/logrus"
)

// InitOptions is the struct of options to configure logger
type InitOptions struct {
	Level string
}

// Init is a function to init json logger
func Init(option InitOptions) (*logrus.Logger, error) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	if option.Level == "" {
		return logger, nil
	}

	level, err := logrus.ParseLevel(option.Level)

	if err != nil {
		return nil, err
	}

	logger.SetLevel(level)

	return logger, nil
}
