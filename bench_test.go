package benchpb

import (
	//github_msg "benchpb/githubgo/proto3"
	// gogo_msg "benchpb/gogo/proto3"
	// gogofast_msg "benchpb/gogofast/proto3"
	// gogofaster_msg "benchpb/gogofaster/proto3"
	google_msg "benchpb/googlego/proto3"
	gopb_msg "benchpb/gopb/proto3"
	"encoding/json"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

var testData = &gopb_msg.FieldTestMessage{
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

func generateSameValue(value *gopb_msg.FieldTestMessage) (
	//gitpb *github_msg.FieldTestMessage,
	googlepb *google_msg.FieldTestMessage,
	// gogopb *gogo_msg.FieldTestMessage,
	// gofastpb *gogofast_msg.FieldTestMessage,
	//	gofasterpb *gogofaster_msg.FieldTestMessage,
) {
	//gitpb = &github_msg.FieldTestMessage{}
	googlepb = &google_msg.FieldTestMessage{}
	// gogopb = &gogo_msg.FieldTestMessage{}
	// gofastpb = &gogofast_msg.FieldTestMessage{}
	//gofasterpb = &gogofaster_msg.FieldTestMessage{}

	data, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, googlepb)
	if err != nil {
		panic(err)
	}
	return
}

func equal(t *testing.T, v1 *gopb_msg.FieldTestMessage, v2 *google_msg.FieldTestMessage) {
	d1, err := json.Marshal(v1)
	if err != nil {
		t.Fatal(err)
	}
	d2, err := json.Marshal(v2)
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, string(d1), string(d2))
}

func TestBasic(t *testing.T) {
	v1 := testData
	// v2, v3, v4, v5, v6 := generateSameValue(v1)
	// t.Log(v1, v2, v3, v4, v5, v6)
	v2 := generateSameValue(v1)
	// t.Log(v1, v2)

	// 确保两个结构里面的数据是一致性的.
	equal(t, v1, v2)

	// 互相编码解码
	d1, err := v1.MarshalObject()
	if err != nil {
		t.Fatal(err)
	}
	d2, err := proto.Marshal(v2)
	if err != nil {
		t.Fatal(err)
	}
	v1.Reset()
	v2.Reset()

	err = v1.UnmarshalObject(d2)
	if err != nil {
		t.Fatal(err)
	}
	err = proto.Unmarshal(d1, v2)
	if err != nil {
		t.Fatal(time.Now(), err)
	}

	equal(t, v1, v2)

	v1.Reset()
	v2.Reset()

	err = v1.UnmarshalObject(d1)
	if err != nil {
		t.Fatal(err)
	}
	err = proto.Unmarshal(d2, v2)
	if err != nil {
		t.Fatal(time.Now(), err)
	}

	equal(t, v1, v2)

	assert.EqualValues(t, len(d2), len(d1), "compare marshal size")
}

func TestSize(t *testing.T) {
	v1 := &gopb_msg.FieldTestMessage{
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
	buf, err := v1.MarshalObject()
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, v1.MarshalSize(), len(buf))
}

func BenchmarkVS(b *testing.B) {
	bd := []struct {
		name   string
		source *gopb_msg.FieldTestMessage
	}{
		{
			name:   "complex-",
			source: testData,
		},
		{
			name: "basic_type-",
			source: &gopb_msg.FieldTestMessage{
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
			name: "repeated-",
			source: &gopb_msg.FieldTestMessage{
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
			name: "map-int-",
			source: &gopb_msg.FieldTestMessage{
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
			name: "map-all-",
			source: &gopb_msg.FieldTestMessage{
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
	for _, v := range bd {
		v1 := v.source
		// v2, v3, v4, v5, v6 := generateSameValue(v1)
		// t.Log(v1, v2, v3, v4, v5, v6)
		v2 := generateSameValue(v1)

		d1, err := v1.MarshalObject()
		if err != nil {
			b.Fatal(err)
		}
		d2, err := proto.Marshal(v2)
		if err != nil {
			b.Fatal(err)
		}
		b.Run(v.name+"google-marshal", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				proto.Marshal(v2)
			}
		})
		b.Run(v.name+"gopb-marshal", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				v1.MarshalObject()
			}
		})

		b.Run(v.name+"google-unmarshal", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				v2 = &google_msg.FieldTestMessage{}
				proto.Unmarshal(d2, v2)
			}
		})
		// b.Run(v.name+"googlepb-unmarshal-cross", func(b *testing.B) {
		// 	for i := 0; i < b.N; i++ {
		// 		v2 = &google_msg.FieldTestMessage{}
		// 		proto.Unmarshal(d1, v2)
		// 	}
		// })
		b.Run(v.name+"gopb-unmarshal", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				v1 = &gopb_msg.FieldTestMessage{}
				v1.UnmarshalObject(d1)
			}
		})
		// b.Run(v.name+"gopb-unmarshal-cross", func(b *testing.B) {
		// 	for i := 0; i < b.N; i++ {
		// 		v1 = &gopb_msg.FieldTestMessage{}
		// 		v1.UnmarshalObject(d2)
		// 	}
		// })
	}
}
