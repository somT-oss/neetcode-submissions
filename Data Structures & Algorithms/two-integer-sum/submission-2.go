func twoSum(nums []int, target int) []int {
    
    var res []int

    hashmap := make(map[int]int)

    for i := 0; i < len(nums); i ++{
        val, ok := hashmap[target - nums[i]]
        if !(ok){
            hashmap[nums[i]] = i
        } else {
            res = append(res, val)
            res = append(res, i)
        }
    }

    return res
}
