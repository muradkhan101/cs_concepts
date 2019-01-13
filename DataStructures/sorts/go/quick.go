package main

func swap(list *[]int, i, j int) {
	temp := (*list)[i]
	(*list)[i] = (*list)[j]
	(*list)[i] = temp
}

func partition(list *[]int, start, end int) int {
	left := start - 1
	pivot := (*list)[end]

	for i := start; i < end; i++ {
		if (*list)[i] < pivot {
			left++
			swap(list, i, left)
		}
	}
	swap(list, left+1, end)
	return left + 1
}

func quickSort(list *[]int, start, end int) {
	if start < end {
		mid := partition(list, start, end)
		quickSort(list, start, mid-1)
		quickSort(list, mid+1, end)
	}
}
