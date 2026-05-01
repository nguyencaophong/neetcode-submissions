func topKFrequent(nums []int, k int) []int {

    freq := map[int]int{}
    for _, n := range nums {
        freq[n]++
    }

    buckets := make([][]int, len(nums)+1)

    for num, f := range freq {
        buckets[f] = append(buckets[f], num)
    }

    res := []int{}

    for i := len(buckets)-1; i >= 0 && len(res) < k; i-- {
        res = append(res, buckets[i]...)
    }

    return res[:k]
}