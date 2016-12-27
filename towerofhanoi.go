package tower_of_hanoi

func solve(n int, source, spare, destination *[]int) {
	if n > 0 {
		solve(n-1, source, destination, spare)
		var popIndex int
		if len(*source) > 0 {
			popIndex = len(*source) - 1
		}

		*destination = append(*destination, (*source)[popIndex])
		*source = (*source)[:popIndex]
		solve(n-1, spare, source, destination)
	}
}
