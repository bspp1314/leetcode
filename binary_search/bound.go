package main


//flag = true,在从小到大的排序数组中， 从数组的begin位置到end-1位置二分查找第一个大于或等于num的数字，找到返回该数字的index，不存在则返回end
//flag = false,在从大到小的排序数组中 从数组的begin位置到end-1位置二分查找第一个小于或等于num的数字，找到返回该数字的index，不存在则返回end。
func LowerBound(begin int,end int,key int,nums []int,flag bool ) int   {
	left := begin
	right := end - 1
	if flag {
		mid := left + ((right-left) >> 1)
		if nums[mid] >= key {
			right = mid //如果是把mid - 1就会把自己给错过
		}else{
			left = mid + 1
		}
	}else{
		mid := left + ((right-left) >> 1)
		if nums[mid] <= key {
			right = mid //如果是把mid - 1就会把自己给错过
		}else{
			left = mid + 1
		}
	}

	return right
}

//flag = true,在从小到大的排序数组中，从数组的begin位置到end-1位置二分查找第一个大于num的数字，找到返回该数字的index，不存在则返回end
//flag = false,在从大到小的排序数组中 从数组的begin位置到end-1位置二分查找第一个小于num的数字，找到返回该数字的index，不存在则返回end。
func UpperBound(begin int,end int,key int,nums []int,flag bool ) int   {
	left := begin
	right := end - 1
	if flag {
		mid := left + ((right-left) >> 1)
		if nums[mid] > key {
			right = mid //如果是把mid - 1就会把自己给错过
		}else{
			left = mid + 1
		}
	}else{
		mid := left + ((right-left) >> 1)
		if nums[mid] < key {
			right = mid //如果是把mid - 1就会把自己给错过
		}else{
			left = mid + 1
		}
	}

	return right
}