package logger

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

type Logger interface {
	Info(format string, args ...interface{})
	Error(format string, args ...interface{})
	Success(format string, args ...interface{})
}

type ConsoleLogger struct{}

func NewConsoleLogger() Logger {
	return &ConsoleLogger{}
}

func (l *ConsoleLogger) Info(format string, args ...interface{}) {
	blue := color.New(color.FgBlue).SprintFunc()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s %s\n", timestamp, blue("[INFO]"), fmt.Sprintf(format, args...))
}

func (l *ConsoleLogger) Error(format string, args ...interface{}) {
	red := color.New(color.FgRed).SprintFunc()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s %s\n", timestamp, red("[ERROR]"), fmt.Sprintf(format, args...))
}

func (l *ConsoleLogger) Success(format string, args ...interface{}) {
	green := color.New(color.FgGreen).SprintFunc()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s %s\n", timestamp, green("[SUCCESS]"), fmt.Sprintf(format, args...))
}
