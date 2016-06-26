package model_test

import (
	. "github.com/mantzas/incata/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"

	"github.com/satori/go.uuid"
)

var _ = Describe("Event", func() {

	It("create a new event when calling initializer function", func() {

		username := "user name"
		created := time.Now()
		event := NewEvent(uuid.NewV4(), created, username, "test type", 1)

		Expect(event.Version).To(Equal(1))
		Expect(event.EventType).To(Equal("test type"))
		Expect(event.Payload).To(Equal(username))
		Expect(event.Created).To(Equal(created.UTC()))
	})
})
