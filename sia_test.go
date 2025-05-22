package go_sia__test

import (
	"math/big"
	"testing"

	sia "github.com/TimeleapLabs/go-sia/v2/pkg"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestSerialize(t *testing.T) {
	var sampleUint16 uint16 = 182
	sampleString := "hello world"

	rawByte := sia.New().
		AddUInt16(sampleUint16).
		AddString64(sampleString).
		Bytes()

	deserialized := sia.NewFromBytes(rawByte)
	gotSampleUint16 := deserialized.ReadUInt16()
	assert.Equal(t, sampleUint16, gotSampleUint16, "should be equal")
	gotSampleString := deserialized.ReadString64()
	assert.Equal(t, sampleString, gotSampleString, "should be equal")
}

type person struct {
	Name string
	Age  uint8
}

func (p *person) Sia() sia.Sia {
	return sia.New().
		AddString8(p.Name).
		AddUInt8(p.Age)
}

func TestArraySerialize(t *testing.T) {

	sampleArray := []person{
		{Name: "joe", Age: 12},
		{Name: "jane", Age: 13},
		{Name: "john", Age: 14},
	}

	s := sia.NewSiaArray[person]().
		AddArray8(sampleArray, func(s *sia.ArraySia[person], item person) {
			s.EmbedBytes(item.Sia().Bytes())
		})

	rawByte := s.Bytes()

	deserialized := sia.NewArrayFromBytes[person](rawByte).ReadArray8(func(s *sia.ArraySia[person]) person {
		return person{
			Name: s.ReadString8(),
			Age:  s.ReadUInt8(),
		}
	})

	if cmp.Diff(sampleArray, deserialized) != "" {
		t.Errorf("should be equal")
	}
}

func TestArraySerializeWithArrayReaders(t *testing.T) {

	sampleInt := 182

	s := sia.New()
	s.AddInt8(int8(sampleInt))

	writer := sia.NewArray[person](&s)

	sampleArray := []person{
		{Name: "joe", Age: 12},
		{Name: "jane", Age: 13},
		{Name: "john", Age: 14},
	}

	writer.
		AddArray8(sampleArray, func(s *sia.ArraySia[person], item person) {
			s.EmbedBytes(item.Sia().Bytes())
		})

	rawBytes := s.Bytes()

	desia := sia.NewFromBytes(rawBytes)
	gotSampleInt := desia.ReadInt8()

	assert.Equal(t, int8(sampleInt), gotSampleInt, "should be equal")

	reader := sia.NewArray[person](&desia)
	deserialized := reader.ReadArray8(func(s *sia.ArraySia[person]) person {
		return person{
			Name: s.ReadString8(),
			Age:  s.ReadUInt8(),
		}
	})

	if cmp.Diff(sampleArray, deserialized) != "" {
		t.Errorf("should be equal")
	}
}

func TestBoolSerialization(t *testing.T) {
	values := []bool{true, false, true}
	s := sia.New()
	for _, v := range values {
		s.AddBool(v)
	}

	desia := sia.NewFromBytes(s.Bytes())
	for i, v := range values {
		got := desia.ReadBool()
		assert.Equal(t, v, got, "bool mismatch at index %d", i)
	}
}

func TestBigIntSerialization(t *testing.T) {
	bigVals := []*big.Int{
		big.NewInt(123456789),
		big.NewInt(0),
	}

	// Add large number using SetString
	largeInt, ok := new(big.Int).SetString("9876543210987654321", 10)
	if !ok {
		t.Fatal("failed to parse large big.Int")
	}
	bigVals = append(bigVals, largeInt)

	s := sia.New()
	for _, v := range bigVals {
		s.AddBigInt(v)
	}

	desia := sia.NewFromBytes(s.Bytes())
	for i, v := range bigVals {
		got := desia.ReadBigInt()
		assert.Equal(t, 0, v.Cmp(got), "big.Int mismatch at index %d", i)
	}
}

func TestIntegerArraySerializationVariants(t *testing.T) {
	type arrayCase struct {
		label     string
		addArray  func(s *sia.ArraySia[int], arr []int)
		readArray func(s *sia.ArraySia[int]) []int
	}

	cases := []arrayCase{
		{
			label: "Array8",
			addArray: func(s *sia.ArraySia[int], arr []int) {
				s.AddArray8(arr, func(s *sia.ArraySia[int], val int) {
					s.AddInt32(int32(val)) // Use 32-bit representation
				})
			},
			readArray: func(s *sia.ArraySia[int]) []int {
				return s.ReadArray8(func(s *sia.ArraySia[int]) int {
					return int(s.ReadInt32())
				})
			},
		},
		{
			label: "Array16",
			addArray: func(s *sia.ArraySia[int], arr []int) {
				s.AddArray16(arr, func(s *sia.ArraySia[int], val int) {
					s.AddInt32(int32(val))
				})
			},
			readArray: func(s *sia.ArraySia[int]) []int {
				return s.ReadArray16(func(s *sia.ArraySia[int]) int {
					return int(s.ReadInt32())
				})
			},
		},
		{
			label: "Array32",
			addArray: func(s *sia.ArraySia[int], arr []int) {
				s.AddArray32(arr, func(s *sia.ArraySia[int], val int) {
					s.AddInt32(int32(val))
				})
			},
			readArray: func(s *sia.ArraySia[int]) []int {
				return s.ReadArray32(func(s *sia.ArraySia[int]) int {
					return int(s.ReadInt32())
				})
			},
		},
		{
			label: "Array64",
			addArray: func(s *sia.ArraySia[int], arr []int) {
				s.AddArray64(arr, func(s *sia.ArraySia[int], val int) {
					s.AddInt32(int32(val))
				})
			},
			readArray: func(s *sia.ArraySia[int]) []int {
				return s.ReadArray64(func(s *sia.ArraySia[int]) int {
					return int(s.ReadInt32())
				})
			},
		},
	}

	sample := []int{-100, 0, 42, 9999, -32768, 2147483647}

	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			s := sia.NewSiaArray[int]()
			c.addArray(s.(*sia.ArraySia[int]), sample)
			desia := sia.NewArrayFromBytes[int](s.Bytes())
			got := c.readArray(desia.(*sia.ArraySia[int]))
			assert.Equal(t, sample, got, "%s: array mismatch", c.label)
		})
	}
}

func TestStringSerializationVariants(t *testing.T) {
	type stringCase struct {
		label  string
		add    func(s sia.Sia, v string) sia.Sia
		read   func(s sia.Sia) string
		values []string
	}

	cases := []stringCase{
		{
			label: "String8",
			add:   func(s sia.Sia, v string) sia.Sia { return s.AddString8(v) },
			read:  func(s sia.Sia) string { return s.ReadString8() },
			values: []string{
				"hello", "Sia", "", "a longer string of characters",
			},
		},
		{
			label: "String16",
			add:   func(s sia.Sia, v string) sia.Sia { return s.AddString16(v) },
			read:  func(s sia.Sia) string { return s.ReadString16() },
			values: []string{
				"short", "1234567890", "", "unicode ✓✓✓",
			},
		},
		{
			label: "String32",
			add:   func(s sia.Sia, v string) sia.Sia { return s.AddString32(v) },
			read:  func(s sia.Sia) string { return s.ReadString32() },
			values: []string{
				"one", "two", "", "🚀 rocket emoji",
			},
		},
		{
			label: "String64",
			add:   func(s sia.Sia, v string) sia.Sia { return s.AddString64(v) },
			read:  func(s sia.Sia) string { return s.ReadString64() },
			values: []string{
				"OK", "longer message with many characters", "",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			s := sia.New()
			for _, val := range c.values {
				c.add(s, val)
			}
			desia := sia.NewFromBytes(s.Bytes())
			for i, val := range c.values {
				got := c.read(desia)
				assert.Equal(t, val, got, "%s: value mismatch at index %d", c.label, i)
			}
		})
	}
}

func TestIntegerSerializationVariants(t *testing.T) {
	type intCase[T any] struct {
		label  string
		add    func(s sia.Sia, v T) sia.Sia
		read   func(s sia.Sia) T
		values []T
	}

	intCases := []intCase[any]{
		{
			label:  "Int8",
			add:    func(s sia.Sia, v any) sia.Sia { return s.AddInt8(v.(int8)) },
			read:   func(s sia.Sia) any { return s.ReadInt8() },
			values: []any{int8(0), int8(-1), int8(127), int8(-128)},
		},
		{
			label:  "UInt8",
			add:    func(s sia.Sia, v any) sia.Sia { return s.AddUInt8(v.(uint8)) },
			read:   func(s sia.Sia) any { return s.ReadUInt8() },
			values: []any{uint8(0), uint8(255), uint8(42)},
		},
		{
			label:  "Int16",
			add:    func(s sia.Sia, v any) sia.Sia { return s.AddInt16(v.(int16)) },
			read:   func(s sia.Sia) any { return s.ReadInt16() },
			values: []any{int16(-32768), int16(0), int16(32767)},
		},
		{
			label:  "UInt16",
			add:    func(s sia.Sia, v any) sia.Sia { return s.AddUInt16(v.(uint16)) },
			read:   func(s sia.Sia) any { return s.ReadUInt16() },
			values: []any{uint16(0), uint16(65535)},
		},
		{
			label:  "Int32",
			add:    func(s sia.Sia, v any) sia.Sia { return s.AddInt32(v.(int32)) },
			read:   func(s sia.Sia) any { return s.ReadInt32() },
			values: []any{int32(-2147483648), int32(0), int32(2147483647)},
		},
		{
			label:  "UInt32",
			add:    func(s sia.Sia, v any) sia.Sia { return s.AddUInt32(v.(uint32)) },
			read:   func(s sia.Sia) any { return s.ReadUInt32() },
			values: []any{uint32(0), uint32(4294967295)},
		},
		{
			label:  "Int64",
			add:    func(s sia.Sia, v any) sia.Sia { return s.AddInt64(v.(int64)) },
			read:   func(s sia.Sia) any { return s.ReadInt64() },
			values: []any{int64(-9223372036854775808), int64(0), int64(9223372036854775807)},
		},
		{
			label:  "UInt64",
			add:    func(s sia.Sia, v any) sia.Sia { return s.AddUInt64(v.(uint64)) },
			read:   func(s sia.Sia) any { return s.ReadUInt64() },
			values: []any{uint64(0), uint64(18446744073709551615)},
		},
	}

	for _, c := range intCases {
		t.Run(c.label, func(t *testing.T) {
			s := sia.New()
			for _, val := range c.values {
				c.add(s, val)
			}
			desia := sia.NewFromBytes(s.Bytes())
			for i, val := range c.values {
				got := c.read(desia)
				assert.Equal(t, val, got, "%s: mismatch at index %d", c.label, i)
			}
		})
	}
}
