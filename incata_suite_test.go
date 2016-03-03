package incata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestIncata(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Incata Suite")
}
