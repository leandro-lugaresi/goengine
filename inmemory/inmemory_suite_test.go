package inmemory_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoengine(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "In Memory Suite")
}
