package domain

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewPerson() *Person {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	phones := []*PhoneNumber{
		{
			Number: fmt.Sprintf("%d-%d-%d", r.Intn(100), r.Intn(100), r.Intn(100)),
			Type:   PhoneType_MOBILE,
		},
		{
			Number: fmt.Sprintf("%d-%d-%d", r.Intn(100), r.Intn(100), r.Intn(100)),
			Type:   PhoneType_HOME,
		},
		{
			Number: fmt.Sprintf("%d-%d-%d", r.Intn(100), r.Intn(100), r.Intn(100)),
			Type:   PhoneType_WORK,
		},
	}

	return &Person{
		Id:          uuid.New().String(),
		Name:        fmt.Sprintf("name %d", r.Intn(1000)),
		Email:       fmt.Sprintf("email %d", r.Intn(1000)),
		Phones:      phones,
		LastUpdated: timestamppb.New(time.Now().UTC()),
	}
}
