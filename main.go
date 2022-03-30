package main

import (
	"fmt"
	"math/rand"
)

// Função que retorna uma array de números aleatórios.
func randomNumbers(n int) []int {
	numbers := make([]int, n)
	for i := range numbers {
		numbers[i] = i
	}
	for i := range numbers {
		j := rand.Intn(i + 1)
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

// Bubble Sort
func bubbleSort(numbers []int) ([]int, int) {
	iterations := 0
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			iterations++
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	return numbers, iterations
}

//Selection Sort
func selectionSort(numbers []int) ([]int, int) {
	iterations := 0
	for i := 0; i < len(numbers)-1; i++ {
		min := i
		for j := i + 1; j < len(numbers); j++ {
			iterations++
			if numbers[j] < numbers[min] {
				min = j
			}
		}
		numbers[i], numbers[min] = numbers[min], numbers[i]
	}
	return numbers, iterations
}

//Insertion Sort
func insertionSort(numbers []int) ([]int, int) {
	iterations := 0
	for i := 1; i < len(numbers); i++ {
		j := i
		for j > 0 && numbers[j] < numbers[j-1] {
			iterations++
			numbers[j], numbers[j-1] = numbers[j-1], numbers[j]
			j--
		}
	}
	return numbers, iterations
}

//Shell Sort
func shellSort(numbers []int) ([]int, int) {
	iterations := 0

	for increment := len(numbers) / 2; increment > 0; increment /= 2 {
		for i := increment; i < len(numbers); i++ {
			temp := numbers[i]
			j := i
			for j >= increment && numbers[j-increment] > temp {
				iterations++
				numbers[j] = numbers[j-increment]
				j -= increment
			}
			numbers[j] = temp
		}
	}
	return numbers, iterations
}

//Merge Sort
func mergeSort(numbers []int, iterations int) ([]int, int) {
	iterations++
	if len(numbers) > 1 {
		mid := len(numbers) / 2
		left := numbers[:mid]
		right := numbers[mid:]
		_, iterations = mergeSort(left, iterations)
		_, iterations = mergeSort(right, iterations)
		i := 0
		j := 0
		k := 0
		for i < len(left) && j < len(right) {
			iterations++
			if left[i] < right[j] {
				numbers[k] = left[i]
				i++
			} else {
				numbers[k] = right[j]
				j++
			}
			k++
		}
		for i < len(left) {
			iterations++
			numbers[k] = left[i]
			i++
			k++
		}
		for j < len(right) {
			iterations++
			numbers[k] = right[j]
			j++
			k++
		}
	}
	return numbers, iterations
}

//Quick Sort
func quickSort(numbers []int, p, r int, iterations int) ([]int, int) {
	iterations++
	var q int
	if p < r {
		q, iterations = partition(numbers, p, r, iterations)
		_, iterations = quickSort(numbers, p, q-1, iterations)
		_, iterations = quickSort(numbers, q+1, r, iterations)
	}
	return numbers, iterations
}

//Partition
func partition(numbers []int, p, r int, iterations int) (int, int) {
	x := numbers[r]
	i := p - 1
	for j := p; j < r; j++ {
		iterations++
		if numbers[j] <= x {
			i++
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}
	numbers[i+1], numbers[r] = numbers[r], numbers[i+1]
	return i + 1, iterations
}

// Função main
func main() {
	fmt.Println("| ======= Menu Principal ======== |")
	fmt.Println("| [1] - Gerar um Vetor Aleatório  |")
	fmt.Println("| [2] - Bubble Sort               |")
	fmt.Println("| [3] - Selection Sort            |")
	fmt.Println("| [4] - Insertion Sort            |")
	fmt.Println("| [5] - Shell Sort                |")
	fmt.Println("| [6] - Merge Sort                |")
	fmt.Println("| [7] - Quick Sort                |")
	fmt.Println("|                                 |")
	fmt.Println("| [0] - Sair                      |")
	fmt.Println("| =============================== |")
	fmt.Print("  Escolha uma opcao -> ")

	var randomList []int

	var option int

	for ok := true; ok; ok = option != 0 {
		fmt.Scanf("%d", &option)
		switch option {
		case 1:
			println("Digite o tamanho da lista")
			var listSize int
			fmt.Scanf("%d", &listSize)
			randomList = randomNumbers(listSize)
			fmt.Print("Vetor Aleatório Gerado: ")
			fmt.Println(randomList)
		case 2:
			sortedList, iterations := bubbleSort(randomList)
			fmt.Print("Bubble Sort Executado: ")
			fmt.Println(sortedList)
			fmt.Println("Iterações: ", iterations)
		case 3:
			sortedList, iterations := selectionSort(randomList)
			fmt.Println("Selection Sort Executado")
			fmt.Println(sortedList)
			fmt.Println("Iterações: ", iterations)
		case 4:
			sortedList, iterations := insertionSort(randomList)
			fmt.Println("Insertion Sort Executado")
			fmt.Println(sortedList)
			fmt.Println("Iterações: ", iterations)
		case 5:
			sortedList, iterations := shellSort(randomList)
			fmt.Println("Shell Sort Executado")
			fmt.Println(sortedList)
			fmt.Println("Iterações: ", iterations)
		case 6:
			sortedList, iterations := mergeSort(randomList, 0)
			fmt.Println("Merge Sort Executado")
			fmt.Println(sortedList)
			fmt.Println("Iterações: ", iterations)
		case 7:
			sortedList, iterations := quickSort(randomList, 0, len(randomList)-1, 0)
			fmt.Println("Quick Sort Executado")
			fmt.Println(sortedList)
			fmt.Println("Iterações: ", iterations)
		case 0:
			fmt.Println("Saindo...")
			ok = false
		default:
			fmt.Println("Opcao Invalida")
		}

	}

}
