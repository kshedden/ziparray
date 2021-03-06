package ziparray

import (
	"testing"
)

func TestInt64(t *testing.T) {

	m := 100
	vec := make([]int64, m)
	for k := 0; k < m; k++ {
		vec[k] = int64(k)
	}

	err := WriteInt64(vec, "test.bin.gz")
	if err != nil {
		panic(err)
	}

	x, err := ReadInt64("test.bin.gz")
	if err != nil {
		panic(err)
	}

	for k := 0; k < m; k++ {
		if vec[k] != x[k] {
			t.Fail()
		}
	}
}

func TestInt64Part(t *testing.T) {

	m := 100
	vec := make([]int64, m)
	for k := 0; k < m; k++ {
		vec[k] = int64(k)
	}

	err := WriteInt64(vec, "test.bin.gz")
	if err != nil {
		panic(err)
	}

	x, err := ReadInt64Part("test.bin.gz", 3, 6)
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

func TestFloat64(t *testing.T) {

	m := 100
	vec := make([]float64, m)
	for k := 0; k < m; k++ {
		vec[k] = float64(k)
	}

	err := WriteFloat64(vec, "test.bin.gz")
	if err != nil {
		panic(err)
	}

	x, err := ReadFloat64("test.bin.gz")
	if err != nil {
		panic(err)
	}

	for k := 0; k < m; k++ {
		if vec[k] != x[k] {
			t.Fail()
		}
	}
}

func TestFloat64Part(t *testing.T) {

	m := 100
	vec := make([]float64, m)
	for k := 0; k < m; k++ {
		vec[k] = float64(k)
	}

	err := WriteFloat64(vec, "test.bin.gz")
	if err != nil {
		panic(err)
	}

	x, err := ReadFloat64Part("test.bin.gz", 2, 7)
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

func TestString(t *testing.T) {

	avec := []string{"cat", "dog", "apple"}
	err := WriteString(avec, "test.bin.gz")
	if err != nil {
		panic(err)
	}

	bvec, err := ReadString("test.bin.gz")
	if err != nil {
		panic(err)
	}

	for k, v := range avec {
		if bvec[k] != v {
			t.Fail()
		}
	}
}

func TestStringPart(t *testing.T) {

	avec := []string{"cat", "dog", "apple", "banana", "peach", "melon"}
	err := WriteString(avec, "test.bin.gz")
	if err != nil {
		panic(err)
	}

	bvec, err := ReadStringPart("test.bin.gz", 1, 4)
	if err != nil {
		panic(err)
	}

	for k, v := range avec[1:4] {
		if bvec[k] != v {
			t.Fail()
		}
	}
}
