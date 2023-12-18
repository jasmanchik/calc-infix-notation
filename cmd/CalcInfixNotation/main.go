package main

import (
	"CalcInfixNotation/internal/calculator"
	"bufio"
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	logger := setupLogger()
	logger.Info("starting application")

	var wg sync.WaitGroup
	wg.Add(3)
	ctx, cancel := context.WithCancel(context.Background())
	expressionChan := make(chan string, 10)
	outputChan := make(chan float64, 10)

	go scanInput(ctx, &wg, logger, expressionChan)

	calc := calculator.NewCalculator(ctx, &wg, logger)
	go calc.Calculate(expressionChan, outputChan)

	go printResults(ctx, &wg, logger, outputChan)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	select {
	case sig := <-stop:
		logger.Info("stopping aplication", slog.String("signal", sig.String()))
		close(expressionChan)
		close(outputChan)
		cancel()
		wg.Wait()
		logger.Info("application stopped")
	}
}

func printResults(ctx context.Context, wg *sync.WaitGroup, logger *slog.Logger, outputChan chan float64) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case result := <-outputChan:
			logger.With("app", "main").Info("calculated:", slog.Float64("result", result))
		default:
		}
	}
}

func scanInput(ctx context.Context, wg *sync.WaitGroup, logger *slog.Logger, ch chan string) {
	defer wg.Done()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		logger.With("app", "main").Info("Введите арифметическое выражение: ")

		select {
		case <-ctx.Done():
			return
		default:
			newInput := make(chan string)
			go func() {
				if scanner.Scan() {
					newInput <- scanner.Text()
				} else {
					close(newInput)
				}
			}()

			select {
			case <-ctx.Done():
				return
			case expression, ok := <-newInput:
				if !ok {
					return
				}
				ch <- expression
			}
		}
	}
}

func setupLogger() *slog.Logger {
	var l *slog.Logger

	l = slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	return l
}
