package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterInt(t *testing.T) {
	testSlice := Filter([]int{1, 4, 19, 4, 2, 64, 43}, func(i int) bool { return i == 4 })
	assert.Equal(t, testSlice, []int{4, 4})
	assert.Equal(t, len(testSlice), 2)
}

func TestFilterString(t *testing.T) {
	testSlice := Filter([]string{"one", "two", "three", "three", "four", "five"}, func(s string) bool { return s == "three" })
	assert.Equal(t, testSlice, []string{"three", "three"})
	assert.Equal(t, len(testSlice), 2)
}

func TestFilterFloat32(t *testing.T) {
	testSlice := Filter([]float32{1.00243, 3.2059023, 8.2039502, 3.2059023, 12.402952}, func(f float32) bool { return f == 3.2059023 })
	assert.Equal(t, testSlice, []float32{3.2059023, 3.2059023})
	assert.Equal(t, len(testSlice), 2)
}

func TestFilterFloat64(t *testing.T) {
	testSlice := Filter([]float64{1.00243, 3.2059023, 8.2039502, 3.2059023, 12.402952}, func(f float64) bool { return f == 3.2059023 })
	assert.Equal(t, testSlice, []float64{3.2059023, 3.2059023})
	assert.Equal(t, len(testSlice), 2)
}

func TestFilterBool(t *testing.T) {
	testSlice := Filter([]bool{true, false, false, true, false, false}, func(b bool) bool { return b == true })
	assert.Equal(t, testSlice, []bool{true, true})
	assert.Equal(t, len(testSlice), 2)
}

func TestFilterIntByFunc(t *testing.T) {
	input := []int{-10, -6, 0, 2, 5, 8}
	testSlice := Filter(input, func(i int) bool { return i > 0 })
	testSlice2 := Filter(input, func(i int) bool { return i%2 == 0 })
	assert.Equal(t, testSlice, []int{2, 5, 8})
	assert.Equal(t, len(testSlice), 3)
	assert.Equal(t, testSlice2, []int{-10, -6, 0, 2, 8})
	assert.Equal(t, len(testSlice2), 5)
}

func TestFilterStringByFunc(t *testing.T) {
	input := []string{"one", "two", "three", "four", "five"}
	testSlice := Filter(input, func(s string) bool { return len(s) > 3 })
	assert.Equal(t, testSlice, []string{"three", "four", "five"})
	assert.Equal(t, len(testSlice), 3)
}

func TestFilterFloat32ByFunc(t *testing.T) {
	input := []float32{-1.253263, -4.2959302, 0.00002, 3.049234, 8.2085204}
	testSlice := Filter(input, func(f float32) bool { return f > 0.1 })
	assert.Equal(t, testSlice, []float32{3.049234, 8.2085204})
	assert.Equal(t, len(testSlice), 2)
}

func TestFilterFloat64ByFunc(t *testing.T) {
	input := []float64{-1.256463, -4.2959302, 0.00002, 3.049234, 8.2085204}
	testSlice := Filter(input, func(f float64) bool { return f > 0.1 })
	assert.Equal(t, testSlice, []float64{3.049234, 8.2085204})
	assert.Equal(t, len(testSlice), 2)
}

func TestMapInt(t *testing.T) {
	input := []int{-3, 0, 3, 6}
	testSlice := Map(input, func(i int) int { return i * 2 })
	assert.Equal(t, testSlice, []int{-6, 0, 6, 12})
}

func TestMapString(t *testing.T) {
	input := []string{"bird", "cow", "sheep"}
	testSlice := Map(input, func(s string) string {
		if s != "sheep" {
			return s + "s"
		} else {
			return s
		}
	})
	assert.Equal(t, testSlice, []string{"birds", "cows", "sheep"})
}

func TestSliceIncludes(t *testing.T) {
	intTest1 := SliceIncludes([]int{1, 4, 8, 9}, 8)
	intTest2 := SliceIncludes([]int{1, 4, 8, 9}, 3)
	stringTest1 := SliceIncludes([]string{"one", "two", "three"}, "two")
	stringTest2 := SliceIncludes([]string{"one", "two", "three"}, "four")
	float32Test1 := SliceIncludes([]float32{1.042, 4.2342, 8.412, 9.2342}, 8.412)
	float32Test2 := SliceIncludes([]float32{1.042, 4.2342, 8.412, 9.2342}, 8.412242)
	float64Test1 := SliceIncludes([]float64{1.042, 4.2342, 8.412, 9.2342}, 8.412)
	float64Test2 := SliceIncludes([]float64{1.042, 4.2342, 8.412, 9.2342}, 8.412242)
	assert.Equal(t, intTest1, true)
	assert.Equal(t, intTest2, false)
	assert.Equal(t, stringTest1, true)
	assert.Equal(t, stringTest2, false)
	assert.Equal(t, float32Test1, true)
	assert.Equal(t, float32Test2, false)
	assert.Equal(t, float32Test1, true)
	assert.Equal(t, float32Test2, false)
	assert.Equal(t, float64Test1, true)
	assert.Equal(t, float64Test2, false)
}
