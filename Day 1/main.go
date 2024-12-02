package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	left, right := read()

	quickSort(left, 0, len(left)-1)
	quickSort(right, 0, len(right)-1)

	distance := totalDistance(left, right)
	fmt.Println(distance)
}

func read() (left, right []int) {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	leftList := []int{}
	rightList := []int{}

	for scanner.Scan() {
		text := scanner.Text()
		text = strings.ReplaceAll(text, " ", "")

		leftVal, err := strconv.Atoi(text[:5])
		if err != nil {
			panic(err)
		}
		rightVal, err := strconv.Atoi(text[5:])
		if err != nil {
			panic(err)
		}

		leftList = append(leftList, leftVal)
		rightList = append(rightList, rightVal)
	}
	return leftList, rightList
}

func totalDistance(left, right []int) int {
	if len(left) != len(right) {
		panic("Lists are not same size")
	}

	total := 0
	for i := 0; i < len(left); i++ {
		distance := left[i] - right[i]
		if distance < 0 {
			distance *= -1
		}
		total += distance
	}

	return total
}

/* Quick Sort */
func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {

			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
