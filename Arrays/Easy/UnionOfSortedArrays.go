package main

import (
	"fmt"
)

// clear comments if you dont want to repeat the same element
// check bottom comments for ultra better solution
func union(a1 []int, a2 []int, s1 int, s2 int) []int {
	var uni []int
	var i, j int
	for i < s1 && j < s2 {
		if a1[i] <= a2[j] {
			// if len(uni) == 0 || uni[len(uni)-1] != a1[i] {
			// 	uni = append(uni, a1[i])
			// }
			uni = append(uni, a1[i])
			i++
		} else if a1[i] >= a2[j] {
			// if len(uni) == 0 || uni[len(uni)-1] != a2[j] {
			// 	uni = append(uni, a2[j])
			// }
			uni = append(uni, a2[j])
			j++
		}
	}
	for i < s1 {
		if len(uni) == 0 || uni[len(uni)-1] != a1[i] {
			uni = append(uni, a1[i])
		}
		i++
	}
	for j < s2 {
		if len(uni) == 0 || uni[len(uni)-1] != a2[j] {
			uni = append(uni, a2[j])
		}
		j++
	}
	return uni
	// for i := range uni {
	// 	a1[i] = uni[i]
	// }
	// a1 = a1[:len(uni)]
	// if you want to return from a1 itself
}
func main() {
	a1 := []int{1, 3, 5, 7, 9}
	a2 := []int{2, 4, 6, 8}

	fmt.Print(union(a1, a2, len(a1), len(a2)))

}

// func merge(nums1 []int, m int, nums2 []int, n int)  {
//     for n != 0 {
//         if m != 0 && nums1[m-1] > nums2[n-1] {
//             nums1[m+n-1] = nums1[m-1]
//             m--
//         } else {
//             nums1[m+n-1] = nums2[n-1]
//             n--
//         }
//     }
// }
