package array



// https://www.youtube.com/watch?v=quAS1iydq7U
// how permutation are build ?
// the first question is to answer this first !
// https://www.youtube.com/watch?v=quAS1iydq7U&t=3s

// https://leetcode.com/problems/next-permutation/solution/
// Am I the only ones that find the category marked as 'medium' strange?
// I think if you know the solution already, it's medium.
// But if you don't know this problem is much "harder" than some designated "hard" problem.

// hate such questions. first, the pattern is not easy to observe.
// second, the most important part is even if you observe the pattern,
// how to prove it is correct? usually if you cannot prove its correctness,
// the algorithm may not work for the corner cases.....

// 这道题不会做...
// 和下面的代码比较清晰，和视频里面的差不多。
func nextPermutation(nums []int) {
	if len(nums) == 0 {
		return
	}
	piv := -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			piv = i
			break
		}
	}
	// no next permutation
	if piv == -1 {
		reverse(&nums, 0, len(nums)-1)
		return
	}
	nextPiv := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > nums[piv] {
			nextPiv = i
			break
		}
	}
	// swap two items
	nums[piv], nums[nextPiv] = nums[nextPiv], nums[piv]

	reversev1(&nums, piv+1, len(nums)-1)
}

func reversev1(nums *[]int, start, end int) {
	for start < end {
		(*nums)[start], (*nums)[end] = (*nums)[end], (*nums)[start]
		start++
		end--
	}
}