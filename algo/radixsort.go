package algo

type RadixSort struct {
	arr []int
}

func (rs *RadixSort) maxItem() int {
	max := 0
	for _, v := range rs.arr {
		if max < v {
			max = v
		}
	}

	return max
}

func (rs *RadixSort) countSort(place int) {

	// 个位开始求解对应位最大的那个数字
	max := (rs.arr[0] / place) % 10

	for i := 1; i < len(rs.arr); i++ {
		if ((rs.arr[i] / place) % 10) > max {
			max = rs.arr[i] / place % 10
		}
	}

	countArr := make([]int, max+1)

	// 以元素对应位数的值作为索引，
	// 出现的次数作为值存储到数组
	for _, v := range rs.arr {
		placeIdx := (v / place) % 10
		countArr[placeIdx] += 1
	}

	// 求计数数组中前面的次数总和
	for i := 1; i <= max; i++ {
		countArr[i] += countArr[i-1]
	}

	output := make([]int, len(rs.arr))
	for i := len(rs.arr) - 1; i >= 0; i-- {
		v := rs.arr[i]
		vCountIdx := (v / place) % 10
		outIdx := countArr[vCountIdx] - 1

		output[outIdx] = v
		countArr[vCountIdx] -= 1
	}

	for i, v := range output {
		rs.arr[i] = v
	}
}

func (rs *RadixSort) radixSort() {
	max := rs.maxItem()

	for place := 1; max/place > 0; place *= 10 {
		rs.countSort(place)
	}
}
