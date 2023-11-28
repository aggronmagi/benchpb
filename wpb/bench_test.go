package benchpb

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/aggronmagi/benchpb"
	gogofast_msg "github.com/aggronmagi/benchpb/gogofast/proto3"
	gopb_msg "github.com/aggronmagi/benchpb/gopb/proto3"

	"github.com/aggronmagi/benchpb/wpb/wproto"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func generateSameValue(value *gopb_msg.FieldTestMessage) (
	wpb *wproto.FieldTestMessage,
	//gitpb *github_msg.FieldTestMessage,
	gogofastpb *gogofast_msg.FieldTestMessage,
	// gogopb *gogo_msg.FieldTestMessage,
	// gofastpb *gogofast_msg.FieldTestMessage,
	//	gofasterpb *gogofaster_msg.FieldTestMessage,
) {
	//gitpb = &github_msg.FieldTestMessage{}
	gogofastpb = &gogofast_msg.FieldTestMessage{}
	// gogopb = &gogo_msg.FieldTestMessage{}
	// gofastpb = &gogofast_msg.FieldTestMessage{}
	//gofasterpb = &gogofaster_msg.FieldTestMessage{}
	wpb = &wproto.FieldTestMessage{}

	data, err := json.Marshal(value)
	if err != nil {
		panic(err)

	}
	err = json.Unmarshal(data, gogofastpb)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, wpb)
	if err != nil {
		panic(err)
	}

	return
}

func equal(t *testing.T, v1 *wproto.FieldTestMessage, v2 *gogofast_msg.FieldTestMessage) {
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
	// v1 := benchpb.TestData
	// v2, v3, v4, v5, v6 := generateSameValue(v1)
	// t.Log(v1, v2, v3, v4, v5, v6)
	v1, v2 := generateSameValue(benchpb.TestData)
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
	v1 := benchpb.TestData
	buf, err := v1.MarshalObject()
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, v1.MarshalSize(), len(buf))
}

func BenchmarkVS(b *testing.B) {
	for _, v := range benchpb.BenchData {
		//v1 := v.Source
		// v2, v3, v4, v5, v6 := generateSameValue(v1)
		// t.Log(v1, v2, v3, v4, v5, v6)
		v1, v2 := generateSameValue(v.Source)

		d1, err := v1.MarshalObject()
		if err != nil {
			b.Fatal(err)
		}
		d2, err := proto.Marshal(v2)
		if err != nil {
			b.Fatal(err)
		}
		b.Run(v.Name+"gogofast-marshal", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				proto.Marshal(v2)
			}
		})
		b.Run(v.Name+"gopb-marshal", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				v1.MarshalObject()
			}
		})

		b.Run(v.Name+"gogofast-unmarshal", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				v2 = &gogofast_msg.FieldTestMessage{}
				proto.Unmarshal(d2, v2)
			}
		})
		// b.Run(v.name+"gogofastpb-unmarshal-cross", func(b *testing.B) {
		// 	for i := 0; i < b.N; i++ {
		// 		v2 = &gogofast_msg.FieldTestMessage{}
		// 		proto.Unmarshal(d1, v2)
		// 	}
		// })
		b.Run(v.Name+"gopb-unmarshal", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				v1 = &wproto.FieldTestMessage{}
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
