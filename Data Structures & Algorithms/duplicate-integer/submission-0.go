func hasDuplicate(nums []int) bool {
    sort.Ints(nums) // sắp xếp mảng tăng dần

    for i := 0; i < len(nums)-1; i++ {
        if nums[i] == nums[i+1] {
            return true
        }
    }
    return false
}
