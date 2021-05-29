package marshalunmarshal

import (
	"github.com/michelaquino/protobuffer-json-comparison/domain"
	"google.golang.org/protobuf/encoding/protojson"
)

func MarshalPROTOAsJSON() []byte {
	msg, err := protojson.Marshal(getAddreddBook())
	if err != nil {
		panic(err)
	}

	return msg
}

func UnmarshalPROTOAsJSON(msg []byte) *domain.AddressBook {
	addressBook := &domain.AddressBook{}
	protojson.Unmarshal(msg, addressBook)

	return addressBook
}
