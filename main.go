package main

type InputType interface {
	int | float32 | float64 | string | bool
}

func Filter[i InputType](arg []i, fn func(i) bool) []i {
	var filteredSlice []i
	for _, v := range arg {
		if fn(v) {
			filteredSlice = append(filteredSlice, v)
		}
	}
	return filteredSlice
}

func Map[i InputType](arg []i, fn func(i) i) []i {
	var mappedSlice []i
	for _, v := range arg {
		mappedSlice = append(mappedSlice, fn(v))
	}
	return mappedSlice
}

func SliceIncludes[i InputType](arg []i, key i) bool {
	for _, v := range arg {
		if v == key {
			return true
		}
	}
	return false
}
