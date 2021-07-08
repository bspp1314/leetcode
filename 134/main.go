package main

// 1 2 3 4 5
// 3 4 5 1 2

//车能开完全程需要满足两个条件
//所有站里的油总量要>=车子的总耗油量。
// 那么，假设从编号为0站开始，一直到k站都正常，在开往k+1站时车子没油了。这时，应该将起点设置为k+1站。(贪心法)
//问题1: 为什么应该将起始站点设为k+1？

//因为k->k+1站耗油太大，0->k站剩余油量都是不为负的，每减少一站，就少了一些剩余油量。所以如果从k前面的站点作为起始站，剩余油量不可能冲过k+1站。
func canCompleteCircuit(gas []int, cost []int) int {
	res := 0
	run := 0
	start := 0
	n := len(gas)

	for i := 0; i < n; i++ {
		run += gas[i] - cost[i]
		res +=  gas[i] - cost[i]
		if run < 0 {
			start = i + 1
			run = 0
		}
	}

	if res < 0 {
		return -1
	}else{
		return start
	}
}
