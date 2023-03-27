package tools

import "math"

func FindAverage(arr []int) int {
	sum := float64(0)

	for _, n := range arr {
		sum += float64(n)
	}
	return int(math.Round(sum / float64(len(arr))))
}

func SortIntegerTable(table []int) {
	len := len(table)
	for i := 0; i < len-1; i++ {
		if table[i] > table[i+1] {
			temp := table[i]
			table[i] = table[i+1]
			table[i+1] = temp
			i = -1
		}
	}
}

func FindMedian(arr []int) int {
	arrSorted := append(arr)
	median := 0
	SortIntegerTable(arrSorted)
	if len(arrSorted)%2 == 0 {
		midle := float64(arrSorted[len(arr)/2-1] + arrSorted[len(arr)/2])
		median = int(math.Round(midle / 2))

	} else {
		median = arrSorted[len(arr)/2]
	}

	return median
}

func FindVariance(arr []int) int {
	var arrVariance []float64
	sumVariance := float64(0)
	sumArr := float64(0)

	for _, n := range arr {
		sumArr += float64(n)
	}
	average := sumArr / float64(len(arr))

	for _, n := range arr {
		current := math.Pow(float64(n)-float64(average), 2)
		arrVariance = append(arrVariance, math.Round(current))
		sumVariance += math.Round(current)

	}

	variance := math.Round(sumVariance / float64(len(arrVariance)))
	return int(variance)
}

func FindStandardDeviation(variance int) int {
	standardDeviation := math.Round(math.Sqrt(float64(variance)))
	return int(standardDeviation)
}
