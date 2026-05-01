func topKFrequent(nums []int, k int) []int {

    freq := map[int]int{}
    for _, n := range nums {
        freq[n]++
    }

    buckets := make([][]int, len(nums)+1)

    for key, value := range freq {
        buckets[value] = append(buckets[value], key)
    }
 
    res := []int{}

    for i := len(buckets)-1; i >= 0 && len(res) <= k; i-- {
        res = append(res, buckets[i]...)
    }

    return res[:k]
}