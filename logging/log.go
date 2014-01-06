package logging

import "log"
import z "github.com/zaiuz/zaiuz"

const ContextKey = "zaius.LoggingModule"

type LoggingModule struct{
	logger *log.Logger
}

var _ z.Module = new(LoggingModule)

func Stdout() *LoggingModule {
	return nil
}

func Log(logger *log.Logger) *LoggingModule {
	return nil
}

func Print(c *z.Context) {
	// TODO: this and other logging methods
}

func (l *LoggingModule) Attach(c *z.Context) error {
	// TODO: Log incoming request.
	return nil
}

func (l *LoggingModule) Detach(c *z.Context) error {
	// TODO: Time finished responses
	return nil
}

