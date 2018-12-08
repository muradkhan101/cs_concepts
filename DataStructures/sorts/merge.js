function merge(list, start, mid, end) {
    const lenArray1 = mid - start + 1;
    const array1 = new Array(lenArray1)
    for (let i = 0; i < lenArray1; i++) {
        array1[i] = list[i + start]
    }
    const lenArray2 = end - mid;
    const array2 = new Array(lenArray2)
    for (let i = 0; i < lenArray2; i++) {
        array2[i] = list[i + mid + 1]
    }

    let leftPointer = 0,
        rightPointer = 0,
        originalPointer = start;
    while (leftPointer < lenArray1 && rightPointer < lenArray2) {
        if (array1[leftPointer] < array2[rightPointer]) {
            list[originalPointer] = array1[leftPointer]
            leftPointer++
        } else {
            list[originalPointer] = array2[rightPointer]
            rightPointer++
        }
        originalPointer++
    }
    while (leftPointer < lenArray1) {
        list[originalPointer] = array1[leftPointer]
        leftPointer++
        originalPointer++
    }
    while (rightPointer < lenArray2) {
        list[originalPointer] = array2[rightPointer]
        rightPointer++
        originalPointer++
    }
}

function mergeSort(arr, start = 0, end = arr.length - 1) {
    if (start < end) {
        const mid = Math.floor((start + end) / 2)
        mergeSort(arr, start, mid)
        mergeSort(arr, mid + 1, end)
        merge(arr, start, mid, end)
    }
    return arr

}