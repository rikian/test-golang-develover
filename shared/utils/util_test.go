package utils

import (
	"go/service1/config"
	"strings"
	"testing"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUtil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Util")
}

var _ = BeforeSuite(func() {
	config.LoadEnvFile()
})

var _ = Describe("Util", func() {
	Context("Encrypt Session", func() {
		It("SUCCESS", func() {
			userid := uuid.New().String()
			session, err := EncryptSession(userid, 3600)

			Expect(err).To(BeNil())
			Expect(session).NotTo(Equal(""))
			Expect(len(strings.Split(session, "."))).To(Equal(3))
		})
	})
})
