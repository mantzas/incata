package model_test

import (
	. "github.com/mantzas/incata/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/satori/go.uuid"
	"time"
)

var _ = Describe("Event", func() {

	It("create a new event when calling initializer function", func() {

		username := "user name"
		event := NewEvent(uuid.NewV4(), username, "test type", 1)

		Expect(event.Version).To(Equal(1))
		Expect(event.EventType).To(Equal("test type"))
		Expect(event.Payload).To(Equal(username))
		Expect(event.Created.After(time.Now().UTC())).To(BeFalse())
	})
})
