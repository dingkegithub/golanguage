package algo

func LinearSearch(arr []int, key int) int {

	for i := 0; i < len(arr); i++ {
		if arr[i] == key {
			return i
		}
	}
	return -1
}

func binarySearch(arr []int, key, i, j int) int {
	if i < j {
		mid := i + ((j - i) >> 1)
		if arr[mid] == key {
			return mid
		}

		if arr[mid] > key {
			return binarySearch(arr, key, i, mid-1)
		} else {
			return binarySearch(arr, key, mid+1, j)
		}
	}

	return -1
}

func BinarySearch(arr []int, key int) int {
	return binarySearch(arr, key, 0, len(arr)-1)
}

func interpolation(arr []int, lo, hi, key int) int {
	a := key - arr[lo]
	b := arr[hi] - arr[lo]
	c := hi - lo

	probe := lo + (a/b)*c

	return probe
}

func InterpolationSearch(arr []int, key int) int {

	lo := 0
	hi := len(arr) - 1

	for lo <= hi && key <= arr[hi] && key >= arr[lo] {
		if lo == hi {
			if arr[lo] == key {
				return lo
			}
			break
		}

		probe := interpolation(arr, lo, hi, key)
		if arr[probe] == key {
			return probe
		} else if arr[probe] > arr[lo] {
			lo += 1
			continue
		} else {
			hi -= 1
		}
	}

	return -1
}

func minInt(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func ExponentailSearch(arr []int, key int) int {
	n := len(arr)

	if key == arr[0] {
		return key
	}

	r := 1

	for r < n && arr[r] <= key {
		r = (r << 1)
	}

	l := (r >> 1)

	if arr[l] == key {
		return l
	} else {
		r = minInt(r, n)
		return binarySearch(arr, key, l, r)
	}
}
