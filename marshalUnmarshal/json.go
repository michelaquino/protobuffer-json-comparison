package marshalunmarshal

import (
	"encoding/json"

	"github.com/michelaquino/protobuffer-json-comparison/domain"
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
