package main

/*
matrix = [
   [ 1,  5,  9],
   [10, 11, 13],
   [12, 13, 15]
],
k = 8,
return 13.
*/

func kthSmallest(matrix [][]int, k int) int {

	arr := []int{}
	for _, k := range matrix {
		arr = append(arr, k...)
	}

	for i := 0; i < k; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] < arr[j+1] {
				tmp := 0
				tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp

			}
		}
	}
	return arr[len(arr)-k]
}

func main() {

}
