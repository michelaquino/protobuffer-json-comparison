package main

import (
	"testing"

	marshalunmarshal "github.com/michelaquino/protobuffer-json-comparison/marshalUnmarshal"
)

func BenchmarkSerializeJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		marshalunmarshal.MarshalJSON()
	}
}

func BenchmarkSerializeProtocolBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		marshalunmarshal.MarshalPROTO()
	}
}

func BenchmarkSerializeProtocolBufferAsJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		marshalunmarshal.MarshalPROTOAsJSON()
	}
}

func BenchmarkDeserializeJSON(b *testing.B) {
	msgJson := marshalunmarshal.MarshalJSON()
	for i := 0; i < b.N; i++ {
		marshalunmarshal.UnmarshalJSON(msgJson)
	}
}

func BenchmarkDeserializeProtocolBuffer(b *testing.B) {
	msgProto := marshalunmarshal.MarshalPROTO()
	for i := 0; i < b.N; i++ {
		marshalunmarshal.UnmarshalPROTO(msgProto)
	}
}

func BenchmarkDeserializeProtocolBufferAsJSON(b *testing.B) {
	msgProtoAsJson := marshalunmarshal.MarshalPROTOAsJSON()
	for i := 0; i < b.N; i++ {
		marshalunmarshal.UnmarshalPROTOAsJSON(msgProtoAsJson)
	}
}
