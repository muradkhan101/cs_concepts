package main

func merge(list *[]int, start, mid, end int) {
	lengthLeft := mid - start + 1
	leftCopy := make([]int, lengthLeft)
	for i := 0; i < lengthLeft; i++ {
		leftCopy[i] = (*list)[i+start]
	}

	lengthRight := end - mid
	rightCopy := make([]int, lengthRight)
	for i := 0; i < lengthRight; i++ {
		rightCopy[i] = (*list)[i+mid+1]
	}

	left := 0
	right := 0
	original := start
	for left < lengthLeft && right < lengthRight {
		if leftCopy[left] < rightCopy[right] {
			(*list)[original] = leftCopy[left]
			left++
		} else {
			(*list)[original] = rightCopy[right]
			right++
		}
		original++
	}
	for left < lengthLeft {
		(*list)[original] = leftCopy[left]
		left++
		original++
	}
	for right < lengthRight {
		(*list)[original] = rightCopy[right]
		right++
		original++
	}
}

func mergeSort(list *[]int, start, end int) {
	if start < end {
		mid := int((start + end) / 2)
		mergeSort(list, start, mid)
		mergeSort(list, mid+1, end)
		merge(list, start, mid, end)
	}
}
