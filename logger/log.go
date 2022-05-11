package logger

import (
	go_logger "github.com/phachon/go-logger"
)

var L *go_logger.Logger

func init() {
	L = go_logger.NewLogger()

	L.Detach("console")

	consoleConfig := &go_logger.ConsoleConfig{
		Color:      true,
		JsonFormat: false,
		Format:     "%timestamp_format% [%level_string%] %body%",
	}
	L.Attach("console", go_logger.LOGGER_LEVEL_DEBUG, consoleConfig)

}
