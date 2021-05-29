package marshalunmarshal

import (
	"github.com/michelaquino/protobuffer-json-comparison/domain"
	"google.golang.org/protobuf/proto"
)

func MarshalPROTO() []byte {
	msg, err := proto.Marshal(getAddreddBook())
	if err != nil {
		panic(err)
	}

	return msg
}

func UnmarshalPROTO(msg []byte) *domain.AddressBook {
	addressBookProto := &domain.AddressBook{}
	err := proto.Unmarshal(msg, addressBookProto)
	if err != nil {
		panic(err)
	}

	return addressBookProto
}
