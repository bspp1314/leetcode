package main

type NumArray2 struct {
	Nums []int
}


func Constructor(nums []int) NumArray {
	return NumArray{Nums: nums}
}

func (this *NumArray2) SumRange(i int, j int) int {
	res := 0
	for k :=i;k <=j;k++ {
		res += this.Nums[k]
	}

	return res
}



func (this *NumArray2) Update(i int, val int)  {

	this.Nums[i] = val
}

func main() {

}


