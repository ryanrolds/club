package main

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	It("should return config", func() {
		config, err := GetConfig("../../testdata/test_club_config.yaml")
		Expect(err).To(BeNil())
		Expect(config).ToNot(BeNil())
		Expect(config.Environment).To(Equal(EnvironmentTests))
		Expect(config.DefaultGroupLimit).To(Equal(42))
		Expect(config.ReaperInterval).To(Equal(time.Second * 60))
		Expect(config.Port).To(Equal(3002))
	})
})
