function swap(list, i, j) {
    let temp = list[i]
    list[i] = list[j]
    list[j] = temp
}

function partition(list, start, end) {
    const pivot = list[end];
    let leftPointer = start - 1;
    for (let i = start; i < end; ++i) {
        if (list[i] < pivot) {
            leftPointer++;
            swap(list, leftPointer, i)
        }
    }
    swap(list, leftPointer + 1, end);
    return leftPointer + 1;
}

function quickSort(list, start = 0, end = list.length - 1) {
    if (start < end) {
        const pivot = partition(list, start, end)
        quickSort(list, start, pivot - 1)
        quickSort(list, pivot + 1, end)
    }
    return list
}

