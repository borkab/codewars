/*
Finish the solution so that it sorts the passed in array of numbers. If the function passes in an empty array or null/nil value then it should return an empty array.

For example:

solution(c(1, 2, 3, 10, 5)) # should return c(1, 2, 3, 5, 10)
solution(NULL)              # should return NULL
*/

package kata

import "sort"

func SortNumbers(numbers []int) []int {
	sort.Ints(numbers)
	return numbers
}

// other solutions:
/*
func SortNumbers(numbers []int) []int {

  for i := 0; i < len(numbers); i++{
    for j := 0; j < len(numbers); j++ {
      if numbers[j] > numbers[i] {
        temp := numbers[i]
        numbers[i] = numbers[j]
        numbers[j] = temp
      }
    }

  }

  return numbers // your code here
}

*/

/*
func SortNumbers(arr []int)  []int{
	if len(arr) == 0 {
		return arr
	}
	qSort(arr, 0, len(arr) - 1)
	return arr
}

func qSort(arr []int, low, high int) {
	if low < high {
		p := pivot(arr, low, high)
		qSort(arr, low, p - 1)
		qSort(arr, p + 1, high)
	}
}

func pivot(arr[]int, low, high int) int {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if pivot > arr[j] {
			swap(&arr[i], &arr[j])
			i++
		}
	}

	swap(&arr[i], &arr[high])
	return i
}

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}
*/

/*
func SortNumbers(numbers []int) []int {
	if numbers == nil {
		return nil
	}
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	return numbers
}

*/
