package signaling_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSignaling(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Signaling Suite")
}
