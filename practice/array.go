package main

import (
	"fmt"
	"slices"
)

func removeItem(nums []int, targetIndex int) []int {
	if targetIndex < 0 || targetIndex >= len(nums) {
		return nums
	}
	newItems := append([]int{}, nums[0:targetIndex]...)
	newItems = append(newItems, nums[targetIndex+1:]...)
	return newItems
}
func contains(nums []int, target int) bool {
	//Using loop
	// targetExist := false
	// for i := range nums {
	// 	if nums[i] == target {
	// 		targetExist = true
	// 		break
	// 	}
	// }
	// return targetExist

	//Using in-built method
	return slices.Contains(nums, target)
}

func filterOdds(nums []int) []int {
	oddItems := []int{}
	for _, num := range nums {
		if num%2 != 0 {
			oddItems = append(oddItems, num)
		}
	}
	return oddItems
}

func mergeSlices(s1 []int, s2 []int) []int {
	merged := make([]int, 0, len(s1)+len(s2))

	merged = append(merged, s1...)
	merged = append(merged, s2...)

	return merged
}

func removeDuplicate(nums []int) []int {
	seen := map[int]bool{}
	uniqueItems := []int{}

	for _, num := range nums {
		if seen[num] {
			continue
		}
		seen[num] = true
		uniqueItems = append(uniqueItems, num)
	}
	return uniqueItems
}

func main() {
	// QUE: 1. Create and print an array
	// var arr [5]int = [5]int{1, 2, 3, 4, 5}
	// fmt.Println(arr)

	// QUE: 2. Access array elements
	// nums := [4]int{5, 10, 15, 20}
	// numsLength := len(nums)
	// fmt.Println(nums[0], nums[numsLength-1])

	// QUE: 3. Update an array element
	// var marks [3]int32 = [3]int32{70, 80, 90}
	// marks[1] = 85
	// fmt.Println(marks)

	// QUE: 4. Loop through an array
	// names := [3]string{2: "Amit", 1: "Rohit", 0: "Neha"}

	// Using standard for loop
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// Using range in for loop
	// for index, value := range names {
	// 	fmt.Printf("Item %v at %v index\n", value, index)
	// }

	// Using for loop as while loop
	// iterationCount := 0
	// for iterationCount < len(names) {
	// 	fmt.Printf("Item %v at %v index\n", names[iterationCount], iterationCount)
	// 	iterationCount++
	// }

	//Using infinite loop
	// iterationCount := 0
	// for {
	// 	if iterationCount >= len(names) {
	// 		break
	// 	}
	// 	name := names[iterationCount]
	// 	fmt.Printf("Item %v at %v index\n", name, iterationCount)
	// 	iterationCount++
	// }

	// QUE: 5. Find sum of array elements
	// nums := [5]int{1, 2, 3, 4, 5}

	// sum := 0
	// for i := range nums {
	// 	sum += nums[i]
	// }
	// fmt.Printf("Total sum is: %d\n", sum)

	// QUE: 6. Convert an array to a Slice, append a new value and print the new slice
	// arr := [3]int{10, 20, 30}
	// slices := arr[:] //arr[:] means arr[0:len(arr)] -> arr[0:3] -> 3 because last number is not included
	// slices = append(slices, 10)

	// fmt.Println(slices)

	// QUE: 9. Find length and capacity
	// nums := []int{10, 20, 30, 40, 50}
	// part := nums[1:4]

	// fmt.Println(len(part)) // OUTPUT: 3
	// fmt.Println(cap(part)) // OUTPUT: 4 -> capacity will be same as original slice starting point till the end. Even if we take 1 element, cap will be till the last element.

	// QUE: 10. Remove 2nd index item from a slice
	// nums := []int{10, 20, 30, 40, 50}

	// fmt.Println(removeItem(nums, 1))

	// QUE: 12. Count even numbers
	// nums := []int{1, 2, 3, 4, 5, 6}
	// evenNumsCount := 0
	// for i := range nums {
	// 	num := nums[i]
	// 	if num%2 == 0 {
	// 		evenNumsCount += 1
	// 	}
	// }
	// fmt.Println(evenNumsCount)

	// QUE: 13. Reverse a slice
	// nums := []int{1, 2, 3, 4, 5, 6}

	//Using in-built Reverse method
	// slices.Reverse(nums)
	// fmt.Println(nums)

	//Using loops
	// reversedNums := []int{}
	// lastindex := len(nums) - 1
	// for i := range nums {
	// 	reversedNums = append(reversedNums, nums[lastindex-i])
	// }
	// fmt.Println(reversedNums)

	// QUE: 14. Create a function to check if a value exists
	// 	nums := []int{1, 2, 3, 4, 5, 6}
	// 	fmt.Println(contains(nums, -1))

	// QUE: 15. Filter odd numbers
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(filterOdds(nums))

	// QUE: 16. Merge two slices
	// a := []int{1, 2, 3}
	// b := []int{4, 5, 6}

	// fmt.Println(mergeSlices(a, b))

	// QUE: 20. Remove duplicates from a slice
	fmt.Println(removeDuplicate([]int{1, 2, 2, 3, 1}))
}
