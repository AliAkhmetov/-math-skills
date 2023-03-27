package tools

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadData(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	stat, err := file.Stat()
	lenOfText := stat.Size()
	arrOfBytes := make([]byte, lenOfText)
	file.Read(arrOfBytes)
	file.Close()
	arrOfStrs := strings.Split(string(arrOfBytes), "\n")
	PrintMetrics(arrOfStrs)
}

func PrintMetrics(arrOfStrs []string) {
	var arrOfInts []int
	for i, s := range arrOfStrs {
		if i == len(arrOfStrs)-1 {
			continue
		}
		nb, errNb := (strconv.Atoi(s))
		if errNb != nil {
			fmt.Printf(errNb.Error())
			return
		}
		arrOfInts = append(arrOfInts, nb)
	}
	average := FindAverage(arrOfInts)
	fmt.Println("Average:", average)
	median := FindMedian(arrOfInts)
	fmt.Println("Median:", median)
	variance := FindVariance(arrOfInts)
	fmt.Println("Variance:", variance)
	standardDeviation := FindStandardDeviation(variance)
	fmt.Println("Standard Deviation:", standardDeviation)
}
