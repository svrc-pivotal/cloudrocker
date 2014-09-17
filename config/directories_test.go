package config_test

import (
	"github.com/cloudcredo/cloudfocker/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Directories", func() {
	Describe("Provide a structure for directories", func() {
		var (
			cloudFockerHomeDir string
			testDirectories    *config.Directories
		)

		BeforeEach(func() {
			cloudFockerHomeDir = "/path/to/buildpacks"
			testDirectories = config.NewDirectories(cloudFockerHomeDir)
		})

		It("should return the buildpacks directory", func() {
			Expect(testDirectories.Buildpacks()).To(Equal(cloudFockerHomeDir + "/buildpacks"))
		})

		It("should return the droplet directory", func() {
			Expect(testDirectories.Droplet()).To(Equal(cloudFockerHomeDir + "/droplet"))
		})

		It("should return the result directory", func() {
			Expect(testDirectories.Result()).To(Equal(cloudFockerHomeDir + "/result"))
		})

		It("should return the cache directory", func() {
			Expect(testDirectories.Cache()).To(Equal(cloudFockerHomeDir + "/cache"))
		})

		It("should return the focker directory", func() {
			Expect(testDirectories.Focker()).To(Equal(cloudFockerHomeDir + "/focker"))
		})
	})
})
