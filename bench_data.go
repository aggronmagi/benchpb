package benchpb

import (
	"math"

	gopb_msg "github.com/aggronmagi/benchpb/gopb/proto3"
)

var TestData = &gopb_msg.FieldTestMessage{
	OptionalBool:      true,
	OptionalEnum:      100,
	OptionalInt32:     990,
	OptionalSint32:    880,
	OptionalUint32:    550,
	OptionalInt64:     8880,
	OptionalSint64:    11110,
	OptionalUint64:    2220,
	OptionalSfixed32:  -3330,
	OptionalFixed32:   4440,
	OptionalFloat:     5.0,
	OptionalSfixed64:  -6.0,
	OptionalFixed64:   70,
	OptionalDouble:    888.80,
	OptionalString:    "sdf234",
	OptionalBytes:     []byte{1, 5, 2, 5, 6, 54, 5, 0},
	Optional_Message:  nil,
	Optional_Message2: &gopb_msg.FieldTestMessage_Message2{Bool: true, Enum: 100},
	RepeatedBool:      []bool{true, false, true, false},
	RepeatedEnum:      []gopb_msg.FieldTestMessage_Enum{1, 3},
	RepeatedInt32:     []int32{8, 899, 9999999, -4, math.MaxInt32, math.MinInt32},
	RepeatedSint32:    []int32{2, -5, 666, 55, math.MaxInt32, math.MinInt32},
	RepeatedUint32:    []uint32{99, 88, math.MaxUint32, 0},
	RepeatedInt64:     []int64{1, 2, -1, math.MaxInt64, math.MinInt64},
	RepeatedSint64:    []int64{1, -45, math.MaxInt64, math.MinInt64},
	RepeatedUint64:    []uint64{1, 33, 0, math.MaxUint64},
	RepeatedSfixed32:  []int32{0, 111, math.MaxInt32, math.MinInt32},
	RepeatedFixed32:   []uint32{111, math.MaxUint32},
	RepeatedFloat:     []float32{0.00001, 1.2, math.MaxFloat32},
	RepeatedSfixed64:  []int64{1, -500, math.MaxInt64, math.MinInt64},
	RepeatedFixed64:   []uint64{1, 0, math.MaxUint64},
	RepeatedDouble:    []float64{1.0, 222.0333, math.MaxFloat64},
	RepeatedString:    []string{"vskj", "中国", "v1ksksjdjid"},
	RepeatedBytes:     [][]byte{{1, 2, 3, 34, 45}, {123, 4, 4, 4, 4, 4, 22}},
	Repeated_Message:  []*gopb_msg.FieldTestMessage_Message{{}},
	Repeated_Message2: []*gopb_msg.FieldTestMessage_Message2{{Bool: true, Enum: 1}},
	NopackBool:        []bool{true, false, true, false},
	NopackEnum:        []gopb_msg.FieldTestMessage_Enum{1, 3},
	NopackInt32:       []int32{8, 899, 9999999, -4, math.MaxInt32, math.MinInt32},
	NopackSint32:      []int32{2, -5, 666, 55, math.MaxInt32, math.MinInt32},
	NopackUint32:      []uint32{99, 88, math.MaxUint32, 0},
	NopackInt64:       []int64{1, 2, -1, math.MaxInt64, math.MinInt64},
	NopackSint64:      []int64{1, -45, math.MaxInt64, math.MinInt64},
	NopackUint64:      []uint64{1, 33, 0, math.MaxUint64},
	NopackSfixed32:    []int32{0, 111, math.MaxInt32, math.MinInt32},
	NopackFixed32:     []uint32{111, math.MaxUint32},
	NopackFloat:       []float32{0.00001, 1.2, math.MaxFloat32},
	NopackSfixed64:    []int64{1, -500, math.MaxInt64, math.MinInt64},
	NopackFixed64:     []uint64{1, 0, math.MaxUint64},
	NopackDouble:      []float64{1.0, 222.0333, math.MaxFloat64},
	NopackString:      []string{"vskj", "中国", "v1ksksjdjid"},
	NopackBytes:       [][]byte{{1, 2, 3, 34, 45}, {123, 4, 4, 4, 4, 4, 22}},
	Nopack_Message:    []*gopb_msg.FieldTestMessage_Message{{}},
	Nopack_Message2:   []*gopb_msg.FieldTestMessage_Message2{{Bool: true, Enum: 1}},
	MapBool:           map[int32]bool{1: true, 100: false},
	MapEnum:           map[int32]gopb_msg.FieldTestMessage_Enum{1: 11},
	MapInt32:          map[int32]int32{1000: math.MaxInt32, -1: math.MinInt32},
	MapSint32:         map[int32]int32{1000: math.MaxInt32, -1: math.MinInt32},
	MapUint32:         map[uint32]uint32{1000: math.MaxInt32, math.MaxUint32: 0, 0: 11},
	MapInt64:          map[int64]int64{1000: math.MaxInt64, -1: math.MinInt64},
	MapSint64:         map[int64]int64{1000: math.MaxInt64, -1: math.MinInt64},
	MapUint64:         map[uint64]uint64{1000: math.MaxInt64, 0: math.MaxUint64},
	MapSfixed32:       map[int32]int32{1: 111, 32: 88888},
	MapFixed32:        map[uint32]uint32{math.MaxUint32: 0},
	MapFloat:          map[int32]float32{1: 1},
	MapSfixed64:       map[int32]int64{1: 2},
	MapFixed64:        map[int32]uint64{22: 222},
	MapDouble:         map[int32]float64{4: 444},
	MapString:         map[int32]string{1: "abskdf"},
	MapBytes:          map[int32][]byte{1: {1, 2, 3}},
	Map_Message:       map[int32]*gopb_msg.FieldTestMessage_Message{},
	Map_Message2:      map[int32]*gopb_msg.FieldTestMessage_Message2{1: {Bool: true, Enum: 8}},
	//MapKeyBool:        map[bool]int32{true: 11, false: 1112},
	MapKeyInt32:       map[int32]int32{0: 11, 11: 0, 1000: math.MaxInt32, -1: math.MinInt32},
	MapKeySint32:      map[int32]int32{0: 11, 1000: math.MaxInt32, -1: math.MinInt32},
	MapKeyUint32:      map[uint32]int32{0: 0, 1000: math.MaxInt32, math.MaxUint32: math.MinInt32},
	MapKeyInt64:       map[int64]int32{1000: math.MaxInt32, -1: math.MinInt32, math.MaxInt64: 0, 0: 0},
	MapKeySint64:      map[int64]int32{1000: math.MaxInt32, -1: math.MinInt32, math.MaxInt64: 0, 0: 0},
	MapKeyUint64:      map[uint64]int32{1000: math.MaxInt32, 0: math.MinInt32, math.MaxUint64: 0, 111: 0},
	MapKeySfixed32:    map[int32]int32{1000: math.MaxInt32, -1: math.MinInt32, math.MaxInt32: 0, 0: 0},
	MapKeyFixed32:     map[uint32]int32{1000: math.MaxInt32, 111: math.MinInt32, math.MaxUint32: 0, 0: 0},
	MapKeySfixed64:    map[int64]int32{1000: math.MaxInt32, -1: math.MinInt32, math.MaxInt64: 0, 0: 0},
	MapKeyFixed64:     map[uint64]int32{1000: math.MaxInt32, 111: math.MinInt32, math.MaxInt64: 0, 0: 0},
	MapKeyString:      map[string]int32{"": 1, "v": 33, "max": math.MaxInt32, "min": math.MinInt32},
	MapInt32Int64:     map[int32]int64{88: 2, 90: 5},
	MapStringMessage:  map[string]*gopb_msg.FieldTestMessage_Message{"svd": {}},
	MapStringMessage2: map[string]*gopb_msg.FieldTestMessage_Message2{"sss": {Bool: true}},
	MapFixed64Enum:    map[uint64]gopb_msg.FieldTestMessage_Enum{111: 1},
}

var BenchData = []struct {
	Name   string
	Source *gopb_msg.FieldTestMessage
}{
	{
		Name:   "complex-",
		Source: TestData,
	},
	{
		Name: "basic_type-",
		Source: &gopb_msg.FieldTestMessage{
			OptionalBool:      true,
			OptionalEnum:      100,
			OptionalInt32:     990,
			OptionalSint32:    880,
			OptionalUint32:    550,
			OptionalInt64:     8880,
			OptionalSint64:    11110,
			OptionalUint64:    2220,
			OptionalSfixed32:  -3330,
			OptionalFixed32:   4440,
			OptionalFloat:     5.0,
			OptionalSfixed64:  -6.0,
			OptionalFixed64:   70,
			OptionalDouble:    888.80,
			OptionalString:    "sdf234",
			OptionalBytes:     []byte{1, 5, 2, 5, 6, 54, 5, 0},
			Optional_Message:  nil,
			Optional_Message2: &gopb_msg.FieldTestMessage_Message2{Bool: true, Enum: 100},
		},
	},
	{
		Name: "repeated-",
		Source: &gopb_msg.FieldTestMessage{
			RepeatedBool:      []bool{true, false, true, false},
			RepeatedEnum:      []gopb_msg.FieldTestMessage_Enum{1, 3},
			RepeatedInt32:     []int32{8, 899, 9999999, -4, math.MaxInt32, math.MinInt32},
			RepeatedSint32:    []int32{2, -5, 666, 55, math.MaxInt32, math.MinInt32},
			RepeatedUint32:    []uint32{99, 88, math.MaxUint32, 0},
			RepeatedInt64:     []int64{1, 2, -1, math.MaxInt64, math.MinInt64},
			RepeatedSint64:    []int64{1, -45, math.MaxInt64, math.MinInt64},
			RepeatedUint64:    []uint64{1, 33, 0, math.MaxUint64},
			RepeatedSfixed32:  []int32{0, 111, math.MaxInt32, math.MinInt32},
			RepeatedFixed32:   []uint32{111, math.MaxUint32},
			RepeatedFloat:     []float32{0.00001, 1.2, math.MaxFloat32},
			RepeatedSfixed64:  []int64{1, -500, math.MaxInt64, math.MinInt64},
			RepeatedFixed64:   []uint64{1, 0, math.MaxUint64},
			RepeatedDouble:    []float64{1.0, 222.0333, math.MaxFloat64},
			RepeatedString:    []string{"vskj", "中国", "v1ksksjdjid"},
			RepeatedBytes:     [][]byte{{1, 2, 3, 34, 45}, {123, 4, 4, 4, 4, 4, 22}},
			Repeated_Message:  []*gopb_msg.FieldTestMessage_Message{{}},
			Repeated_Message2: []*gopb_msg.FieldTestMessage_Message2{{Bool: true, Enum: 1}},
		},
	},
	{
		Name: "nopack-",
		Source: &gopb_msg.FieldTestMessage{
			NopackBool:      []bool{true, false, true, false},
			NopackEnum:      []gopb_msg.FieldTestMessage_Enum{1, 3},
			NopackInt32:     []int32{8, 899, 9999999, -4, math.MaxInt32, math.MinInt32},
			NopackSint32:    []int32{2, -5, 666, 55, math.MaxInt32, math.MinInt32},
			NopackUint32:    []uint32{99, 88, math.MaxUint32, 0},
			NopackInt64:     []int64{1, 2, -1, math.MaxInt64, math.MinInt64},
			NopackSint64:    []int64{1, -45, math.MaxInt64, math.MinInt64},
			NopackUint64:    []uint64{1, 33, 0, math.MaxUint64},
			NopackSfixed32:  []int32{0, 111, math.MaxInt32, math.MinInt32},
			NopackFixed32:   []uint32{111, math.MaxUint32},
			NopackFloat:     []float32{0.00001, 1.2, math.MaxFloat32},
			NopackSfixed64:  []int64{1, -500, math.MaxInt64, math.MinInt64},
			NopackFixed64:   []uint64{1, 0, math.MaxUint64},
			NopackDouble:    []float64{1.0, 222.0333, math.MaxFloat64},
			NopackString:    []string{"vskj", "中国", "v1ksksjdjid"},
			NopackBytes:     [][]byte{{1, 2, 3, 34, 45}, {123, 4, 4, 4, 4, 4, 22}},
			Nopack_Message:  []*gopb_msg.FieldTestMessage_Message{{}},
			Nopack_Message2: []*gopb_msg.FieldTestMessage_Message2{{Bool: true, Enum: 1}},
		},
	},
	{
		Name: "map-int-",
		Source: &gopb_msg.FieldTestMessage{
			MapBool:      map[int32]bool{1: true, 100: false},
			MapEnum:      map[int32]gopb_msg.FieldTestMessage_Enum{1: 11},
			MapInt32:     map[int32]int32{1000: math.MaxInt32, -1: math.MinInt32},
			MapSint32:    map[int32]int32{1000: math.MaxInt32, -1: math.MinInt32},
			MapUint32:    map[uint32]uint32{1000: math.MaxInt32, math.MaxUint32: 0, 0: 11},
			MapInt64:     map[int64]int64{1000: math.MaxInt64, -1: math.MinInt64},
			MapSint64:    map[int64]int64{1000: math.MaxInt64, -1: math.MinInt64},
			MapUint64:    map[uint64]uint64{1000: math.MaxInt64, 0: math.MaxUint64},
			MapSfixed32:  map[int32]int32{1: 111, 32: 88888},
			MapFixed32:   map[uint32]uint32{math.MaxUint32: 0},
			MapFloat:     map[int32]float32{1: 1},
			MapSfixed64:  map[int32]int64{1: 2},
			MapFixed64:   map[int32]uint64{22: 222},
			MapDouble:    map[int32]float64{4: 444},
			MapString:    map[int32]string{1: "abskdf"},
			MapBytes:     map[int32][]byte{1: {1, 2, 3}},
			Map_Message:  map[int32]*gopb_msg.FieldTestMessage_Message{},
			Map_Message2: map[int32]*gopb_msg.FieldTestMessage_Message2{1: {Bool: true, Enum: 8}},
		},
	},
	{
		Name: "map-all-",
		Source: &gopb_msg.FieldTestMessage{
			MapKeyInt32:       map[int32]int32{0: 11, 11: 0, 1000: math.MaxInt32, -1: math.MinInt32},
			MapKeySint32:      map[int32]int32{0: 11, 1000: math.MaxInt32, -1: math.MinInt32},
			MapKeyUint32:      map[uint32]int32{0: 0, 1000: math.MaxInt32, math.MaxUint32: math.MinInt32},
			MapKeyInt64:       map[int64]int32{1000: math.MaxInt32, -1: math.MinInt32, math.MaxInt64: 0, 0: 0},
			MapKeySint64:      map[int64]int32{1000: math.MaxInt32, -1: math.MinInt32, math.MaxInt64: 0, 0: 0},
			MapKeyUint64:      map[uint64]int32{1000: math.MaxInt32, 0: math.MinInt32, math.MaxUint64: 0, 111: 0},
			MapKeySfixed32:    map[int32]int32{1000: math.MaxInt32, -1: math.MinInt32, math.MaxInt32: 0, 0: 0},
			MapKeyFixed32:     map[uint32]int32{1000: math.MaxInt32, 111: math.MinInt32, math.MaxUint32: 0, 0: 0},
			MapKeySfixed64:    map[int64]int32{1000: math.MaxInt32, -1: math.MinInt32, math.MaxInt64: 0, 0: 0},
			MapKeyFixed64:     map[uint64]int32{1000: math.MaxInt32, 111: math.MinInt32, math.MaxInt64: 0, 0: 0},
			MapKeyString:      map[string]int32{"": 1, "v": 33, "max": math.MaxInt32, "min": math.MinInt32},
			MapInt32Int64:     map[int32]int64{88: 2, 90: 5},
			MapStringMessage:  map[string]*gopb_msg.FieldTestMessage_Message{"svd": {}},
			MapStringMessage2: map[string]*gopb_msg.FieldTestMessage_Message2{"sss": {Bool: true}},
			MapFixed64Enum:    map[uint64]gopb_msg.FieldTestMessage_Enum{111: 1},
		},
	},
}
