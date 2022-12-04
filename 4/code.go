package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Seq struct {
	start int
	end   int
}

func newSeq(input string) *Seq {
	parts := strings.Split(input, "-")
	start, _ := strconv.ParseInt(parts[0], 0, 64)
	end, _ := strconv.ParseInt(parts[1], 0, 64)
	return &Seq{start: int(start), end: int(end)}
}

func readFile(filename string) ([]string, error) {
	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func doesSeqContainAnother(seq1, seq2 *Seq) bool {
	if (seq1.start >= seq2.start && seq1.end <= seq2.end) ||
		(seq2.start >= seq1.start && seq2.end <= seq1.end) {
		return true
	}
	return false
}

func doSeqsOverlap(seq1, seq2 *Seq) bool {
	if seq1.start > seq2.end || seq2.start > seq1.end {
		return false
	}
	return true
}

func checkOverlap(input []string) int {
	overlapCount := 0

	for _, seqStrs := range input {
		seqs := strings.Split(seqStrs, ",")
		seq1 := newSeq(seqs[0])
		seq2 := newSeq(seqs[1])
		if doesSeqContainAnother(seq1, seq2) {
			overlapCount++
		}
	}
	return overlapCount
}

func checkOverlap2(input []string) int {
	overlapCount := 0

	for _, seqStrs := range input {
		seqs := strings.Split(seqStrs, ",")
		seq1 := newSeq(seqs[0])
		seq2 := newSeq(seqs[1])
		if doSeqsOverlap(seq1, seq2) {
			overlapCount++
		}
	}
	return overlapCount
}

func main() {
	test, _ := readFile("test.txt")
	input, _ := readFile("input.txt")
	fmt.Println("Part1")
	fmt.Println(checkOverlap(test))
	fmt.Println(checkOverlap(input))
	fmt.Println("Part2")
	fmt.Println(checkOverlap2(test))
	fmt.Println(checkOverlap2(input))
}
