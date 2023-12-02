package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const MaxRed = 12
const MaxGreen = 13
const MaxBlue = 14

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to open input.txt, %v", err)
	}
	defer f.Close()
	// fmt.Println(sumValidGames(f))
	fmt.Println(sumMinimumBlocks(f))
}
func sumMinimumBlocks(r io.Reader) int {
	var sum int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		game, err := NewGame(scanner.Text())
		if err != nil {
			log.Fatalf("unecpected invalida game, %v", err)
		}
		sum += game.MinimumBlocks()

	}
	return sum
}

func sumValidGames(r io.Reader) int {
	var sum int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		game, err := NewGame(scanner.Text())
		if err != nil {
			log.Fatalf("unecpected invalida game, %v", err)
		}
		if game.IsValid() {
			sum += game.ID
		}
	}
	return sum
}

type Game struct {
	ID      int
	Reveals []Reveal
}

func NewGame(str string) (*Game, error) {

	// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	parts := strings.Split(str, ":")
	gameNumber, err := parseGameNumber(parts[0])
	if err != nil {
		return nil, err
	}
	reveals, err := parseReveals(parts[1])
	if err != nil {
		return nil, err
	}
	return &Game{
		ID:      gameNumber,
		Reveals: reveals,
	}, nil
}

func (g Game) MinimumBlocks() int {
	// set to 1 to avoid 0-multiplications
	res := Reveal{Blue: 1, Red: 1, Green: 1}
	for _, r := range g.Reveals {
		if r.Blue > res.Blue {
			res.Blue = r.Blue
		}
		if r.Green > res.Green {
			res.Green = r.Green
		}
		if r.Red > res.Red {
			res.Red = r.Red
		}
	}
	return res.Blue * res.Green * res.Red
}

func (g Game) IsValid() bool {
	for _, reveal := range g.Reveals {
		if !reveal.IsValid() {
			return false
		}
	}
	return true
}

func parseGameNumber(s string) (int, error) {
	// Game 2
	if len(s) < 6 {
		return 0, fmt.Errorf("invalid game number: %s", s)
	}
	gameNumber, err := strconv.ParseInt(s[5:], 10, 64)
	if err != nil {
		return 0, err
	}
	return int(gameNumber), nil
}

func parseReveals(s string) ([]Reveal, error) {
	//  1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	revealsParts := strings.Split(s, ";")
	reveals := make([]Reveal, 0, len(revealsParts))
	for _, revealPart := range revealsParts {
		reveal := Reveal{}
		// 1 blue, 2 green
		colorParts := strings.Split(revealPart, ",")
		for _, colorPart := range colorParts {
			// 1 blue
			finalParts := strings.Split(strings.Trim(colorPart, " "), " ")
			if len(finalParts) != 2 {
				return []Reveal{}, fmt.Errorf("invalid reveal: %s", revealPart)
			}
			num, err := strconv.ParseInt(finalParts[0], 10, 64)
			if err != nil {
				return []Reveal{}, fmt.Errorf("invalid reveal: %s", revealPart)
			}
			switch finalParts[1] {
			case "blue":
				reveal.Blue = int(num)
			case "red":
				reveal.Red = int(num)
			case "green":
				reveal.Green = int(num)
			default:
				return []Reveal{}, fmt.Errorf("invalid reveal: %s", revealPart)
			}
		}
		reveals = append(reveals, reveal)
	}
	return reveals, nil
}

type Reveal struct {
	Blue  int
	Red   int
	Green int
}

func (r Reveal) IsValid() bool {
	if r.Blue > MaxBlue || r.Red > MaxRed || r.Green > MaxGreen {
		return false
	}
	return true
}
