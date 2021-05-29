package marshalunmarshal

import (
	"github.com/michelaquino/protobuffer-json-comparison/domain"
)

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
