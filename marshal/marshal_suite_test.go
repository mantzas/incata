package marshal_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMarshal(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Marshal Suite")
}
