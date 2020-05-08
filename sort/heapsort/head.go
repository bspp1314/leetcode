package main

import (
	"fmt"
)

//索引为i的左孩子的索引是 (2*i+1);
//索引为i的右孩子的索引是 (2*i+2)
//父节点的索引为 （i-1）/2

func getLeastNumbers(array []int, k int) []int {
	//1 构建小顶堆
	beginIndex := (len(array) >> 1) - 1 //第一个非叶子节点。

	//1 构建小顶堆
	for i := beginIndex; i >= 0; i-- {
		heap(array, i, len(array)-1)
	}

	//2.调整堆结构+交换堆顶元素与末尾元素
	res := make([]int, k)

	m := 0
	for i := len(array) - 1; i >= 0; i-- {
		if m >  k {
			break
		}

		array[i], array[0] = array[0], array[i]
		res[m] = array[i]
		heap(array, 0, i-1)
		m = m + 1
	}

	return res
}

func heap(array []int, i, end int) {
	parent := i
	child := 2*parent + 1

	for child <= end {
		//比较左右两个节点
		if child+1 <= end && array[child] > array[child+1] {
			child = child + 1
		}

		if array[parent] < array[child] {
			return
		} else {
			//交换父节点和字节点
			array[parent], array[child] = array[child], array[parent]
			//由于父节点和子堆得交换将有可能导致子堆的子堆被破坏
			parent = child //向下转移
			child = 2*parent + 1
		}
	}
}


func HeapSort(array []int) {
	//1 构建大顶堆
	beginIndex := (len(array) >> 1) - 1 //第一个非叶子节点。

	//1 构建大顶堆
	for i := beginIndex; i >= 0; i-- {
		heap(array, i, len(array)-1)
	}

	//2.调整堆结构+交换堆顶元素与末尾元素
	for i := len(array) - 1; i > 0; i-- {
		array[i], array[0] = array[0], array[i]
		heap(array, 0, i-1)
	}
}

func main() {
	array := []int{0,0,2,3,2,1,1,2,0,4}
	fmt.Println(getLeastNumbers(array, 10))
	array = []int{0,0,2,3,2,1,1,2,0,4}
	HeapSort(array)
	fmt.Println(array)
}
