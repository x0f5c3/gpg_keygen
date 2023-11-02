package logging

import (
	"github.com/pterm/pterm"
	slogmulti "github.com/samber/slog-multi"
	"log/slog"
)

var rootLogger *slog.Logger = initLog()

func initLog() *slog.Logger {
	return slog.New(slogmulti.Fanout(pterm.NewSlogHandler(pterm.DefaultLogger.WithTime(true).WithFormatter(pterm.LogFormatterColorful))))
}

func DefaultLogger() {

}
