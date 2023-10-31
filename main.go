package main

import (
	"fmt"
	"math/rand"
	"time"
)

func monobitTest(sequence []int) bool {
	ones := 0
	zeros := 0

	for _, bit := range sequence {
		if bit == 1 {
			ones++
		} else {
			zeros++
		}
	}

	return ones >= 9654 && ones <= 10346
}

func maxSeriesLengthTest(sequence []int) bool {
	maxZeroSeries := 0
	maxOneSeries := 0
	currentZeroSeries := 0
	currentOneSeries := 0

	for _, bit := range sequence {
		if bit == 0 {
			currentZeroSeries++
			currentOneSeries = 0
			if currentZeroSeries > maxZeroSeries {
				maxZeroSeries = currentZeroSeries
			}
		} else {
			currentOneSeries++
			currentZeroSeries = 0
			if currentOneSeries > maxOneSeries {
				maxOneSeries = currentOneSeries
			}
		}
	}

	return maxZeroSeries <= 36 && maxOneSeries <= 36
}

func pokerTest(sequence []int) bool {
	m := 4
	k := len(sequence) / m
	blockCount := make(map[int]int)

	for i := 0; i < k; i++ {
		block := sequence[i*m : (i+1)*m]
		blockValue := 0
		for j, bit := range block {
			blockValue += bit << uint(j)
		}
		blockCount[blockValue]++
	}

	x3 := 0.0
	for _, count := range blockCount {
		x3 += float64(count * count)
	}
	x3 = x3*16/float64(k) - float64(k)

	return x3 >= 1.03 && x3 <= 57.4
}

func seriesLengthTest(sequence []int) bool {
	seriesCounts := make(map[int]int)

	currentSeries := 0
	for _, bit := range sequence {
		if bit == 0 {
			if currentSeries > 0 {
				seriesCounts[currentSeries]++
				currentSeries = 0
			}
		} else {
			currentSeries++
		}
	}

	fmt.Println(seriesCounts)

	for seriesLen, count := range seriesCounts {
		switch seriesLen {
		case 1:
			if count < 2267 || count > 2733 {
				return false
			}
		case 2:
			if count < 1079 || count > 1421 {
				return false
			}
		case 3:
			if count < 502 || count > 748 {
				return false
			}
		case 4:
			if count < 223 || count > 402 {
				return false
			}
		case 5:
			if count < 90 || count > 223 {
				return false
			}
		}
	}

	return true
}

func generateRandomSequence(length int) []int {
	rand.Seed(time.Now().UnixNano())

	sequence := make([]int, length)
	for i := 0; i < length; i++ {
		sequence[i] = rand.Intn(2)
	}

	return sequence
}

func main() {
	sequenceLength := 20000
	randomSequence := generateRandomSequence(sequenceLength)

	fmt.Printf("Monobit: %v\n", monobitTest(randomSequence))
	fmt.Printf("maxSeriesLengthTest: %v\n", maxSeriesLengthTest(randomSequence))
	fmt.Printf("pokerTest: %v\n", pokerTest(randomSequence))
	fmt.Printf("seriesLengthTest: %v\n", seriesLengthTest(randomSequence))
}
