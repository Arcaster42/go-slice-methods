package main

func FilterInt(arg []int, key int) []int {
	var filteredSlice []int
	for _, v := range arg {
		if v == key {
			filteredSlice = append(filteredSlice, v)
		}
	}
	return filteredSlice
}

func FilterString(arg []string, key string) []string {
	var filteredSlice []string
	for _, v := range arg {
		if v == key {
			filteredSlice = append(filteredSlice, v)
		}
	}
	return filteredSlice
}

func FilterFloat32(arg []float32, key float32) []float32 {
	var filteredSlice []float32
	for _, v := range arg {
		if v == key {
			filteredSlice = append(filteredSlice, v)
		}
	}
	return filteredSlice
}

func FilterFloat64(arg []float64, key float64) []float64 {
	var filteredSlice []float64
	for _, v := range arg {
		if v == key {
			filteredSlice = append(filteredSlice, v)
		}
	}
	return filteredSlice
}

func FilterBool(arg []bool, key bool) []bool {
	var filteredSlice []bool
	for _, v := range arg {
		if v == key {
			filteredSlice = append(filteredSlice, v)
		}
	}
	return filteredSlice
}

func FilterIntByFunc(arg []int, test func(int) bool) []int {
	var filteredSlice []int
	for _, v := range arg {
		if test(v) {
			filteredSlice = append(filteredSlice, v)
		}
	}
	return filteredSlice
}

func FilterFloat32ByFunc(arg []float32, test func(float32) bool) []float32 {
	var filteredSlice []float32
	for _, v := range arg {
		if test(v) {
			filteredSlice = append(filteredSlice, v)
		}
	}
	return filteredSlice
}

func FilterFloat64ByFunc(arg []float64, test func(float64) bool) []float64 {
	var filteredSlice []float64
	for _, v := range arg {
		if test(v) {
			filteredSlice = append(filteredSlice, v)
		}
	}
	return filteredSlice
}
