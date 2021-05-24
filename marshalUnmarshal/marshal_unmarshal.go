package marshalunmarshal

import (
	"encoding/json"

	"github.com/michelaquino/protobuffer-json-comparison/domain"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func MarshalJSON() []byte {
	msg, err := json.Marshal(getAddreddBook())
	if err != nil {
		panic(err)
	}

	return msg
}

func UnmarshalJSON(msg []byte) *domain.AddressBook {
	addressBook := &domain.AddressBook{}
	err := json.Unmarshal(msg, &addressBook)
	if err != nil {
		panic(err)
	}

	return addressBook
}

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

func getAddreddBook() *domain.AddressBook {
	people := []*domain.Person{
		domain.NewPerson(),
		domain.NewPerson(),
	}

	addressBook := &domain.AddressBook{
		People: people,
	}

	return addressBook
}
