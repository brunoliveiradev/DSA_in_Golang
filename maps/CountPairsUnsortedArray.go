package maps

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// CountPairsUnsortedArray calculates the number of pairs of socks with matching colors in an unsorted array.
// Parameters:
//   - arr: slice of int32 representing the colors of the socks.
//   - n: number of socks in the array (should match len(arr)).
//
// Returns:
//   - int32: number of pairs of socks with matching colors.
//
// Example:
//
//	arr := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
//	n := int32(len(arr))
//	pairs := CountPairsUnsortedArray(arr, n) // pairs == 3
func CountPairsUnsortedArray(arr []int32, n int32) int32 {

	if n != int32(len(arr)) {
		panic("The length of the array does not match the provided size n.")
	}

	frequencyMap := make(map[int32]int32)

	for _, num := range arr {
		frequencyMap[num]++
	}

	count := int32(0)
	// Count pairs for each color
	for _, freq := range frequencyMap {
		count += freq / 2
	}

	return count
}

// CountPairs reads from standard input, processes the data, and prints the number of pairs of socks.
// Expects the first line of input to be the number of socks (n)
// and the second line to contain n integers representing the colors of the socks.
func CountPairs() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var ar []int32

	for i := 0; i < int(n); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := CountPairsUnsortedArray(ar, n)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
