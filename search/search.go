package search

func BSearch(nums []int, target int) int {
	numLen := len(nums)

	var low = 0
	var high = numLen - 1

	for {

		if low > high {
			return -1
		}
		midKey := low + (high-low)/2

		midValue := nums[midKey]
		if midValue == target {
			return midKey
		}

		if midValue > target {
			high = midKey - 1
		}

		if midValue < target {
			low = midKey + 1
		}
	}

}
