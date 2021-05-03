package main

import "sort"

func findItinerary(tickets [][]string) []string {
	m := make(map[string][]string)
	res := make([]string,0)

	for _,trick := range tickets {
		src,dst := trick[0],trick[1]
		subTricks,ok := m[src]
		if !ok {
			subTricks = make([]string,0)
		}

		m[src] =append(subTricks,dst)
	}

	for key := range m {
		sort.Strings(m[key])
	}

	var dfs func(current string)

	dfs = func(current string) {
		for {
			v,_ := m[current]
			if len(v) == 0 {
				break
			}

			tem := m[current][0]
			m[current] = m[current][1:]
			dfs(tem)
		}

		res = append(res,current)
	}

	dfs("JFK")
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res) - 1 - i] = res[len(res) - 1 - i], res[i]
	}
	return res

}