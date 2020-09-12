package main
import "sort"

func ThreeNumberSum(array []int, target int) [][]int {
	// Write your code here.
	var res [][]int 
	sort.Ints(array)
	mapArr :=make(map[int]bool)
	for _, val := range(array){
		mapArr[val] = true
	}
	for i:=0;i < len(array);i++{
		for j:=i+1; j<len(array);j++{
			temp:= target - array[i]-array[j]
			if _, ok := mapArr[temp];ok{
				if temp > array[j]{
					res = append(res,[]int{array[i],array[j],temp})
				}
			}
		}
	}
	return res
}

