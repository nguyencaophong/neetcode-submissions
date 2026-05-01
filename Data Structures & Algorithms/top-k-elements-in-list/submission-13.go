func topKFrequent(nums []int, k int) []int {
	mapS := map[int]int{}
	for _,v := range(nums) {
		mapS[v]++
	}

	buckets := make([][]int, len(nums) +1)
	for k,v := range mapS {
		buckets[v] = append(buckets[v],k)
	}

	for k,v := range buckets {
		fmt.Print("key",k, "value",v)
	}
	
	res := []int{}
	for i := len(buckets) -1; i >= 0; i -- {
		res = append(res, buckets[i]...)
		if(len(res) >= k){
			break;
		}
	}
	return res[:k]
}
