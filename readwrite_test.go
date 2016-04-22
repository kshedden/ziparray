package ziparray

import (
	"testing"
)

func Test_int64(t *testing.T) {

	m := 100
	vec := make([]int64, m)
	for k := 0; k < m; k++ {
		vec[k] = int64(k)
	}

	err := WriteInt64Array(vec, "test.bin.gz")
	if err != nil {
		panic(err)
	}

	x, err := ReadInt64Array("test.bin.gz")
	if err != nil {
		panic(err)
	}

	for k := 0; k < m; k++ {
		if vec[k] != x[k] {
			t.Fail()
		}
	}
}

func Test_int64sub(t *testing.T) {

	m := 100
	vec := make([]int64, m)
	for k := 0; k < m; k++ {
		vec[k] = int64(k)
	}

	err := WriteInt64Array(vec, "test.bin.gz")
	if err != nil {
		panic(err)
	}

	x, err := ReadInt64SubArray("test.bin.gz", 3, 6)
	if err != nil {
		panic(err)
	}

	svec := vec[3:6]
	for k := 0; k < 3; k++ {
		if svec[k] != x[k] {
			t.Fail()
		}
	}
}

func Test_float64(t *testing.T) {

	m := 100
	vec := make([]float64, m)
	for k := 0; k < m; k++ {
		vec[k] = float64(k)
	}

	err := WriteFloat64Array(vec, "test.bin.gz")
	if err != nil {
		panic(err)
	}

	x, err := ReadFloat64Array("test.bin.gz")
	if err != nil {
		panic(err)
	}

	for k := 0; k < m; k++ {
		if vec[k] != x[k] {
			t.Fail()
		}
	}
}

func Test_float64sub(t *testing.T) {

	m := 100
	vec := make([]float64, m)
	for k := 0; k < m; k++ {
		vec[k] = float64(k)
	}

	err := WriteFloat64Array(vec, "test.bin.gz")
	if err != nil {
		panic(err)
	}

	x, err := ReadFloat64SubArray("test.bin.gz", 2, 7)
	if err != nil {
		panic(err)
	}

	svec := vec[2:7]
	for k := 0; k < 5; k++ {
		if svec[k] != x[k] {
			t.Fail()
		}
	}
}
