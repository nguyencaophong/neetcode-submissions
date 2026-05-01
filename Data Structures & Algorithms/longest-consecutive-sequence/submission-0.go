func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    sort.Ints(nums)

    longest := 1
    current := 1

    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            continue // skip duplicates
        }
        if nums[i]-nums[i-1] == 1 {
            current++
            if current > longest {
                longest = current // track running max
            }
        } else {
            current = 1 // reset on break
        }
    }
    return longest
}