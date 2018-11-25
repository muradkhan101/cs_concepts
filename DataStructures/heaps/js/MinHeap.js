class MinHeap {
    constructor() {
        this.heap = [];
        this.size = 0;
    }

    getParentIndex(i) { return  Math.floor((i - 1) / 2); }
    getLeftIndex(i) { return i * 2 + 1; }
    getRightIndex(i) { return i * 2 + 2; }

    getParent(i) { return this.heap[this.getParentIndex(i)]; }
    getLeftChild(i) { return this.heap[this.getLeftIndex(i)]; }
    getRightChild(i) { return this.heap[this.getRightIndex(i)]; }

    hasParent(i) { return this.getParent(i) !== undefined; }
    hasLeftChild(i) { return this.getLeftChild(i) !== undefined; }
    hasRightChild(i) { return this.getRightChild(i) !== undefined; }

    swap(i, j) {
        const temp = this.heap[i];
        this.heap[i] = this.heap[j];
        this.heap[j] = temp;
    }

    add(val) {
        this.heap[this.size - 1] = val;
        this.size++;
        this.heapifyUp();
    }
    heapifyUp() {
        let index = this.size - 1;
        while (this.hasParent(index) && this.heap[index] < this.getParent(index)) {
            this.swap(index, this.getParentIndex(index));
            index = this.getParentIndex(index);
        }
    }

    peek() { return this.heap[0]; }
    pop() {
        const val = this.heap[0];
        this.heap[0] = this.heap[this.size - 1];
        this.heap[this.size - 1] = undefined;
        this.size--;

        this.heapifyDown();
        return val;
    }

    heapifyDown() {
        let index = 0;
        while (this.hasLeftChild(index)) {
            let minIndex = this.getLeftIndex(index);
            if (this.hasRightChild(index) && this.getRightChild(index) < this.getLeftChild(index)) {
                minIndex = this.getRightIndex(index);
            }

            if (this.heap[minIndex] > this.heap[index]) break;

            this.swap(index, minIndex);
            index = minIndex;
        }
    }
}