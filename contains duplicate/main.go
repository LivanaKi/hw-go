package main

import(
	"fmt"
)

func main(){
	nums := []int{1,1,1,2,2,3}
	k := 2
	fmt.Println(containsDuplicate(nums, k))

}
func containsDuplicate(nums []int, k int) []int {
	mp := make(map[int]int)
	var result []int 
	for i := 0; i < len(nums); i++{
		_, ok := mp[nums[i]]
		if ok{
			mp[nums[i]]++
		}else{
			mp[nums[i]] = 0
		}
	}
	var sm []int
	for i := 0; i < len(mp); i++{
		for j := i + 1; j < len(mp); j++{
			if mp[nums[i]] > mp[nums[j]]{
				sm = append(sm, )
			}
		}
	}

	return result
}