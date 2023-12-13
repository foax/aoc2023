package cmd

import (
	"bufio"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/foax/aoc2023/internal/day01"
	"github.com/foax/aoc2023/internal/day02"
	"github.com/foax/aoc2023/internal/day09"
	"github.com/foax/aoc2023/internal/day10"
	"github.com/foax/aoc2023/internal/day11"
	"github.com/foax/aoc2023/internal/day12"
	"github.com/foax/aoc2023/internal/day13"
)

func Execute() {
	var logLevel slog.Level
	logLevelFlag := flag.String("loglevel", "INFO", "Log level to use for output")
	flag.Parse()
	switch strings.ToLower(*logLevelFlag) {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		panic(fmt.Sprintf("Invalid log level provided: %v", *logLevelFlag))
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))
	slog.SetDefault(logger)

	slog.Info("Start", "loglevel", logLevel)
	defer slog.Info("End")
	slog.Debug("flag args", "args", flag.Args())

	scanner := bufio.NewScanner(os.Stdin)

	args := flag.Args()
	if len(args) == 0 {
		logger.Error("No day argument passed on the command line")
		os.Exit(1)
	}
	switch args[0] {
	case "day01":
		day01.Execute(scanner)
	case "day02":
		day02.Execute(scanner)
	case "day09":
		day09.Execute(scanner)
	case "day10":
		day10.Execute(scanner)
	case "day11":
		day11.Execute(scanner)
	case "day12":
		day12.Execute(scanner)
	case "day13":
		day13.Execute(scanner)
	default:
		slog.Error("Invalid argument", "arg", args[0])

	}

}
