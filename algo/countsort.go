package algo

func CountSortUint(arr []uint) {
	maxVal := arr[0]

	// max=8  seq: 4, 2, 2, 8, 3, 3, 1
	for _, v := range arr {
		if v > maxVal {
			maxVal = v
		}
	}

	// len=8, seq: 0, 0, 0, 0, 0, 0, 0, 0
	countArr := make([]uint, maxVal+1)
	outputArr := make([]uint, maxVal+1)

	// idx: 0  1  2  3  4  5  6  7  8
	// seq: 0  1  2  2  1  0  0  0  1
	for _, v := range arr {
		countArr[v] += 1
	}

	// idx: 0  1  2  3  4  5  6  7  8
	// seq: 0  1  3  5  6  6  6  6  7
	var i uint = 1
	for i = 1; i <= maxVal; i++ {
		countArr[i] += countArr[i-1]
	}

	for _, v := range arr {
		outputArr[countArr[v]-1] = v
		countArr[v] -= 1
	}
}
