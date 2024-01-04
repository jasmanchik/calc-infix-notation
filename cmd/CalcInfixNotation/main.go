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

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(3)

	inputChan := make(chan string, 10)
	outputChan := make(chan float64, 10)

	go scanInput(ctx, &wg, logger, inputChan)

	calc := calculator.NewCalculator(logger)
	go calc.Calculate(&wg, inputChan, outputChan)

	go printResults(ctx, &wg, logger, outputChan)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	select {
	case sig := <-stop:
		logger.Info("stopping application", slog.String("signal", sig.String()))
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
			close(outputChan)
			for result := range outputChan {
				logger.With("app", "main").Info("calculated:", slog.Float64("result", result))
			}
			return
		case result := <-outputChan:
			logger.With("app", "main").Info("calculated:", slog.Float64("result", result))
		default:
		}
	}
}

func scanInput(ctx context.Context, wg *sync.WaitGroup, logger *slog.Logger, masterChan chan string) {
	defer wg.Done()
	scanner := bufio.NewScanner(os.Stdin)
	inputChan := make(chan string)
	go func() {
		logger.With("app", "main").Info("Введите арифметическое выражение: ")
		for scanner.Scan() {
			logger.With("app", "main").Info("Введите арифметическое выражение: ")
			inputChan <- scanner.Text()
		}
	}()

	for {
		select {
		case <-ctx.Done():
			close(inputChan)
			close(masterChan)
			return
		case input := <-inputChan:
			masterChan <- input
		default:
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
