package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var numbers = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Println(calibrateBlock(f))
}

func stream(r io.Reader) <-chan string {
	ch := make(chan string)
	scanner := bufio.NewScanner(r)
	go func(in chan string) {
		defer close(ch)
		for scanner.Scan() {
			in <- scanner.Text()

		}
	}(ch)
	return ch
}

func calibrateBlock(r io.Reader) int {
	var sum int

	wg := sync.WaitGroup{}
	g := runtime.NumCPU()
	result := make(chan int, g)
	wg.Add(g)

	// file as channel to be iterated
	in := stream(r)

	// Fan in
	for i := 0; i < g; i++ {
		go func(out chan int) {
			var localSum int
			defer wg.Done()
			for line := range in {
				localSum += calibrateLine(line)
			}
			out <- localSum
		}(result)
	}
	wg.Wait()

	// Fan out
	for i := 0; i < g; i++ {
		sum += <-result
	}
	return sum
}

func isNumberContained(str string, isLeft bool) (bool, string) {
	var idx int

	// 1. check digit at the right index
	if isLeft {
		idx = len(str) - 1
	}
	char := string(str[idx])
	if _, err := strconv.ParseInt(char, 10, 16); err == nil {
		return true, char
	}

	// 2. check text number
	for k, v := range numbers {
		if strings.Contains(str, k) {
			return true, v
		}
	}
	return false, ""
}

func calibrateLine(str string) int {
	var final, char string
	var leftFound, rightFound bool
	n := len(str)
	// Two pointers, left and right
	leftP, rightP := 0, n-1
	for {
		// Left branch
		if !leftFound {
			leftFound, char = isNumberContained(str[:leftP+1], true)
			if leftFound {
				final = char + final
			} else {
				leftP++
			}
		}
		// Right branch
		if !rightFound {
			rightFound, char = isNumberContained(str[rightP:], false)
			if rightFound {
				final = final + char
			} else {
				rightP--
			}
		}

		if leftFound && rightFound {
			break
		}

		// Apparently all lines have numbers but just in case
		if leftP >= n || rightP < 0 {
			break
		}
	}
	finalNum, err := strconv.ParseInt(final, 10, 16)
	if err != nil {
		return -1
	}
	return int(finalNum)
}
