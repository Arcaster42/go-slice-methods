package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterInt(t *testing.T) {
	testSlice := FilterInt([]int{1, 4, 19, 4, 2, 64, 43}, 4)
	assert.Equal(t, testSlice, []int{4, 4})
	assert.Equal(t, len(testSlice), 2)
}

func TestFilterString(t *testing.T) {
	testSlice := FilterString([]string{"one", "two", "three", "three", "four", "five"}, "three")
	assert.Equal(t, testSlice, []string{"three", "three"})
	assert.Equal(t, len(testSlice), 2)
}

func TestFilterFloat32(t *testing.T) {
	testSlice := FilterFloat32([]float32{1.00243, 3.2059023, 8.2039502, 3.2059023, 12.402952}, 3.2059023)
	assert.Equal(t, testSlice, []float32{3.2059023, 3.2059023})
	assert.Equal(t, len(testSlice), 2)
}

func TestFilterFloat64(t *testing.T) {
	testSlice := FilterFloat64([]float64{1.00243, 3.2059023, 8.2039502, 3.2059023, 12.402952}, 3.2059023)
	assert.Equal(t, testSlice, []float64{3.2059023, 3.2059023})
	assert.Equal(t, len(testSlice), 2)
}

func TestFilterBool(t *testing.T) {
	testSlice := FilterBool([]bool{true, false, false, true, false, false}, true)
	assert.Equal(t, testSlice, []bool{true, true})
	assert.Equal(t, len(testSlice), 2)
}

func TestFilterIntByFunc(t *testing.T) {
	input := []int{-10, -6, 0, 2, 5, 8}
	testSlice := FilterIntByFunc(input, func(i int) bool { return i > 0 })
	testSlice2 := FilterIntByFunc(input, func(i int) bool { return i%2 == 0 })
	assert.Equal(t, testSlice, []int{2, 5, 8})
	assert.Equal(t, len(testSlice), 3)
	assert.Equal(t, testSlice2, []int{-10, -6, 0, 2, 8})
	assert.Equal(t, len(testSlice2), 5)
}

func TestFilterStringByFunc(t *testing.T) {
	input := []string{"one", "two", "three", "four", "five"}
	testSlice := FilterStringByFunc(input, func(s string) bool { return len(s) > 3 })
	assert.Equal(t, testSlice, []string{"three", "four", "five"})
	assert.Equal(t, len(testSlice), 3)
}

func TestFilterFloat32ByFunc(t *testing.T) {
	input := []float32{-1.253263, -4.2959302, 0.00002, 3.049234, 8.2085204}
	testSlice := FilterFloat32ByFunc(input, func(f float32) bool { return f > 0.1 })
	assert.Equal(t, testSlice, []float32{3.049234, 8.2085204})
	assert.Equal(t, len(testSlice), 2)
}

func TestFilterFloat64ByFunc(t *testing.T) {
	input := []float64{-1.256463, -4.2959302, 0.00002, 3.049234, 8.2085204}
	testSlice := FilterFloat64ByFunc(input, func(f float64) bool { return f > 0.1 })
	assert.Equal(t, testSlice, []float64{3.049234, 8.2085204})
	assert.Equal(t, len(testSlice), 2)
}
