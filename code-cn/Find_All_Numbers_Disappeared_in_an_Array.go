package main

func findDisappearedNumbers(nums []int) []int {

	le := len(nums)
	res := make([]int, le+1, le+1)

	for _, v := range nums {
		res[v] = -1
	}

	for i := 1; i < le+1; i++ {
		if res[i] != -1 {
			res = append(res, i)
		}
	}

	return res[le+1:]
}
