package day02

import (
	"bufio"
	"log/slog"
	"strconv"
	"strings"
)

type Draw map[string]int

type Game struct {
	Id    int
	Draws []Draw
}

func parseGameLine(line string) *Game {
	game := Game{}
	fields := strings.Split(line, ": ")
	gameIdFields := strings.Fields(fields[0])
	id, _ := strconv.Atoi(gameIdFields[1])
	game.Id = id
	drawsFields := strings.Split(fields[1], "; ")

	for _, drawStr := range drawsFields {
		draw := Draw{}
		cubeStrs := strings.Split(drawStr, ", ")
		for _, cubeStr := range cubeStrs {
			cubeFields := strings.Fields(cubeStr)
			count, _ := strconv.Atoi(cubeFields[0])
			draw[cubeFields[1]] = count
		}
		game.Draws = append(game.Draws, draw)
	}
	slog.Debug("parseGameLine", "line", line, "game", game)
	return &game
}

func parseGameInput(input []string) (games []*Game) {
	for _, line := range input {
		games = append(games, parseGameLine(line))
	}
	return
}

func checkGame(game *Game, red int, green int, blue int) bool {
	for _, draw := range game.Draws {
		if draw["red"] > red || draw["green"] > green || draw["blue"] > blue {
			return false
		}
	}
	return true
}

func minCubes(game *Game) (red int, green int, blue int) {
	for _, draw := range game.Draws {
		if draw["red"] > red {
			red = draw["red"]
		}
		if draw["green"] > green {
			green = draw["green"]
		}
		if draw["blue"] > blue {
			blue = draw["blue"]
		}
	}
	return
}

func part1Handler(games []*Game) (total int) {
	for _, game := range games {
		if checkGame(game, 12, 13, 14) {
			total += game.Id
		}
	}
	return
}

func part2Handler(games []*Game) (total int) {
	for _, game := range games {
		r, g, b := minCubes(game)
		slog.Debug("part2Handler", "game", game, "minRed", r, "minGreen", g, "minBlue", b, "power", r*g*b)
		total += r * g * b
	}
	return
}

func readInput(scanner *bufio.Scanner) (output []string) {
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	slog.Debug("Lines of input read", "lines", len(output))
	return
}

func Execute(scanner *bufio.Scanner) {
	input := readInput(scanner)
	games := parseGameInput(input)
	part1Total := part1Handler(games)
	part2Total := part2Handler(games)
	slog.Info("Part 1", "result", part1Total)
	slog.Info("Part 2", "result", part2Total)
}
