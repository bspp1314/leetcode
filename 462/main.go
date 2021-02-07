package main

import "fmt"

func midNum(a, b, c int) int {
	if (a-b)*(a-c) <= 0 {
		return a
	}

	if (b-a)*(b-c) <= 0 {
		return b
	}

	return c
}



func quickSort(arr []int,start,end int)  {
	if start >= end {
		return
	}
	pivotIndex := partition(arr,start,end)
	quickSort(arr,start,pivotIndex -1 )
	quickSort(arr,pivotIndex+1,end)
}


func partition(arr []int, startIndex, endIndex int) int {
	dealPivot(arr,startIndex,endIndex)
	var (
		pivot = arr[startIndex]
		left  = startIndex
		right = endIndex
	)
	for left != right {
		for left < right && pivot <= arr[right] {
			right--
		}

		for left < right && pivot >= arr[left] {
			left++
		}

		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}

	arr[startIndex], arr[left] = arr[left], arr[startIndex]


	return left
}

func dealPivot(arr []int,left,right int)  {
	mid := left + (right -left) / 2

	pivotIndex := mid

	if (arr[left] - arr[right]) * (arr[left]- arr[mid]) <= 0 {
		pivotIndex = left
	}

	if (arr[right] - arr[left]) * (arr[right] - arr[mid]) <= 0 {
		pivotIndex = right
	}

	arr[left],arr[pivotIndex] = arr[pivotIndex],arr[left]
}

// a1 a2 a3 ...an-1 an an+1.... a2n+1
// (N + M)

func partition2(arr []int, startIndex, endIndex int) int {
	if startIndex >= endIndex {
		return arr[endIndex]
	}

	dealPivot(arr,startIndex,endIndex)
	var (
		pivot = arr[startIndex]
		left  = startIndex
		right = endIndex
	)
	for left != right {
		for left < right && pivot <= arr[right] {
			right--
		}

		for left < right && pivot >= arr[left] {
			left++
		}

		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}

	arr[startIndex], arr[left] = arr[left], arr[startIndex]

	
	if left == len(arr) / 2  {
		return arr[left]
	}else if left > len(arr) / 2 {
		return partition2(arr,startIndex,left-1)
	}else{
		return partition2(arr,left+1,endIndex)
	}

}

func abs(a,b int) int   {
	if a >= b {
		return a -b
	}else{
		return b -a
	}
}
func minMoves2(nums []int) int {
	if len(nums) <=1  {
		return  0
	}

	mid := partition2(nums,0,len(nums)-1)

	var res int
	for i :=0;i<len(nums);i++ {
		res  += abs(nums[i],mid)
	}

	return res
}

func main() {
	k := []int{}
	//partition(k, 0, len(k)-1)
	//k2 := []int{6,3,4,5,11,33,33,2,4,5}
	//partition2(k2, 0, len(k2)-1)
	//quickSort(k,0,len(k)-1)
	//fmt.Println(k)
	fmt.Println(minMoves2(k))
	//fmt.Println(k2)
}
