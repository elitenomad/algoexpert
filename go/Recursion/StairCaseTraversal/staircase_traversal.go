package main

package main

/*
	Worst case solution
*/
func StaircaseTraversalInEfficient(height int, maxSteps int) int {	// O(steps ** h)
	return StaircaseTraversalHelper(height, maxSteps)
}

func StaircaseTraversalHelperInEfficient(height int, maxSteps int) int {
	if height <= 1 {
		return 1
	} 
	
	count := 0
	for step := 1; step < min(maxSteps, height) + 1; step++ {
		count += StaircaseTraversalHelper(height - step, maxSteps)
	}
	
	return count
}


/*
	Better approach than the above one O(k * n)
*/
func StaircaseTraversal(height int, maxSteps int) int {
	return StaircaseTraversalHelper(height, maxSteps, map[int]int{0:1,1:1})
}

func StaircaseTraversalHelper(height int, maxSteps int, h map[int]int) int {
	if v, ok := h[height]; ok {
		return v
	}
	
	count := 0
	for step := 1; step < min(maxSteps, height) + 1; step++ {
		count += StaircaseTraversalHelper(height - step, maxSteps, h)
	}
	h[height] = count
	
	return count
}

/*
	StairCase Traversal Linear - Non Recursive
*/
func StaircaseTraversal(height int, maxSteps int) int {
	result := make([]int, height + 1)
	result[0] = 1
	result[1] = 1

	for currentHeight := 2; currentHeight < height + 1; currentHeight++ {
		step := 1

		for step <= maxSteps && step <= currentHeight {
			result[currentHeight] = result[currentHeight] + result[currentHeight - step]
			step += 1
		}
	}

	return result[height]
}


// Helper method to find min of steps and height
func min(a, b int) int {
	if a < b {
		return a
	}
	
	return b
}

func main() {

}
