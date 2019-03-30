package main

import "fmt"

func binarySearch(array []int, key int) int {
	left := 0
	right := len(array) - 1
	for left <= right {
		mid := left + (right-left)/2
		if key == array[mid] {
			return mid
		} else if key > array[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
func binarySearch2(array []int, left int, right int, key int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	if key == array[mid] {
		return mid
	} else if key > array[mid] {
		return binarySearch2(array, mid+1, right, key)
	} else {
		return binarySearch2(array, left, mid-1, key)
	}
}

//数组之中的数据可能可以重复，要求返回匹配的数据的最小的下标
func binarySearch3(array []int, key int) int {
	left := 0
	right := len(array) - 1
	result := -1
	for left <= right {
		mid := left + (right-left)/2
		if key == array[mid] {
			result = mid
			right = mid - 1
		} else if key > array[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}

//数组之中的数据可能可以重复，要求返回匹配的数据的最大的下标
func binarySearch4(array []int, key int) int {
	left := 0
	right := len(array) - 1
	result := -1
	for left <= right {
		mid := left + (right-left)/2
		if key == array[mid] {
			result = mid
			left = mid + 1
		} else if key > array[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}
func main() {
	array := []int{-1, -2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	//	fmt.Println(binarySearch(array, 1))
	//	fmt.Println(binarySearch(array, 2))
	//	fmt.Println(binarySearch(array, 3))
	//	fmt.Println(binarySearch(array, 4))
	//	fmt.Println(binarySearch(array, 0))
	//	fmt.Println(binarySearch2(array, 0, len(array), 1))
	//	fmt.Println(binarySearch2(array, 0, len(array), 2))
	//	fmt.Println(binarySearch2(array, 0, len(array), 3))
	//	fmt.Println(binarySearch2(array, 0, len(array), 4))
	//	fmt.Println(binarySearch2(array, 0, len(array), 0))
	fmt.Println(binarySearch3(array, 10))
	fmt.Println(binarySearch4(array, 10))
}
