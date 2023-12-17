package main

import (
	"CalcInfixNotation/internal/calculator"
	"CalcInfixNotation/internal/stack"
	"bufio"
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := setupLogger()
	logger.Info("starting application")

	ctx, cancel := context.WithCancel(context.Background())

	expressionChan := make(chan string, 10)
	outputChan := make(chan float64, 10)
	go scanInput(ctx, logger, expressionChan)

	s := stack.New(50)
	go calculator.Calculate(ctx, logger, s, expressionChan, outputChan)

	go printResults(ctx, logger, outputChan)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sig := <-stop
	close(expressionChan)
	close(outputChan)
	cancel()
	logger.Info("stopping aplication", slog.String("signal", sig.String()))
	logger.Info("application stopped")
}

func printResults(ctx context.Context, logger *slog.Logger, outputChan chan float64) {
	var result float64
	select {
	case <-ctx.Done():
		break
	case result = <-outputChan:
		logger.With("app", "main").Info("calculated:", slog.Float64("result", result))
	default:
	}
}

func scanInput(logger *slog.Logger, ch chan string) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		logger.With("app", "main").Info("Введите арифметическое выражение: ")
		scanner.Scan()
		expression := scanner.Text()
		ch <- expression
	}
}

func setupLogger() *slog.Logger {
	var l *slog.Logger

	l = slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	return l
}
