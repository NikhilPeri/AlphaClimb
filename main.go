package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Climb struct {
	word        string
	biggestMove int
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func biggestMove(word string) int {
	letterPos := []byte(word)
	letterPos = append([]byte{'a'}, letterPos...)
	letterPos = append(letterPos, 'z')
	sort.Slice(letterPos, func(i int, j int) bool { return letterPos[i] < letterPos[j] })

	maxWordDistance := 0
	for i, a := range letterPos[:len(letterPos)-1] {
		maxWordDistance = max(maxWordDistance, int(letterPos[i+1]-a))
	}

	return maxWordDistance
}

func main() {
	f, err := os.Open("words_alpha.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var climbs []Climb
	for scanner.Scan() {
		word := scanner.Text()
		climbs = append(climbs, Climb{word: word, biggestMove: biggestMove(word)})
	}

	sort.Slice(climbs, func(i, j int) bool {
		return climbs[i].biggestMove < climbs[j].biggestMove
	})

	fmt.Println("Top 10 Easiest")
	for i := 1; i <= 10; i++ {
		fmt.Printf("Rank: %d, Score: %d, Word: %s\n", i, climbs[i].biggestMove, climbs[i].word)
	}

	csvFile, err := os.Create("words_scored.csv")
	defer csvFile.Close()
	csvwriter := csv.NewWriter(csvFile)
	for _, climb := range climbs {
		_ = csvwriter.Write([]string{strconv.Itoa(climb.biggestMove), climb.word})
	}
	csvwriter.Flush()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
