package config_test

import (
	"os"

	"github.com/cloudcredo/cloudrocker/config"

	. "github.com/cloudcredo/cloudrocker/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/cloudcredo/cloudrocker/Godeps/_workspace/src/github.com/onsi/gomega"
)

var _ = Describe("Directories", func() {
	var (
		cloudRockerHomeDir string
		testDirectories    *config.Directories
	)

	BeforeEach(func() {
		cloudRockerHomeDir = "/path/to"
		testDirectories = config.NewDirectories(cloudRockerHomeDir)
	})

	Describe("Provide a structure for directories", func() {
		It("should return the cloudrocker home directory", func() {
			Expect(testDirectories.Home()).To(Equal(cloudRockerHomeDir))
		})

		It("should return the buildpacks directory", func() {
			Expect(testDirectories.Buildpacks()).To(Equal(cloudRockerHomeDir + "/buildpacks"))
		})

		It("should return the container's buildpacks directory", func() {
			Expect(testDirectories.ContainerBuildpacks()).To(Equal("/cloudrockerbuildpacks"))
		})

		It("should return the runtime droplet directory", func() {
			Expect(testDirectories.Droplet()).To(Equal(cloudRockerHomeDir + "/droplet"))
		})

		It("should return the rocker directory", func() {
			Expect(testDirectories.Rocker()).To(Equal(cloudRockerHomeDir + "/rocker"))
		})

		It("should return the staging directory", func() {
			Expect(testDirectories.Staging()).To(Equal(cloudRockerHomeDir + "/staging"))
		})

		It("should return the host cloudrocker tmp directory", func() {
			Expect(testDirectories.Tmp()).To(Equal(cloudRockerHomeDir + "/tmp"))
		})

		It("should return the host directory for holding the base container configuration", func() {
			Expect(testDirectories.BaseConfig()).To(Equal(cloudRockerHomeDir + "/baseConfig"))
		})

		It("should return the application directory", func() {
			pwd, _ := os.Getwd()
			Expect(testDirectories.App()).To(Equal(pwd))
		})
	})

	Describe("Providing the directories to be mounted in the container", func() {
		It("should return a mapping of host to container directories", func() {
			Expect(testDirectories.Mounts()).To(Equal(map[string]string{ // host dir: container dir
				"/path/to/tmp":        "/tmp",
				"/path/to/rocker":     "/rocker",
				"/path/to/buildpacks": "/cloudrockerbuildpacks",
				"/path/to/staging":    "/tmp/app",
			}))
		})
	})

	Describe("Providing the directories to be created before staging", func() {
		It("should return a set of directories to be created", func() {
			Expect(testDirectories.HostDirectories()).To(ConsistOf(
				"/path/to",
				"/path/to/buildpacks",
				"/path/to/droplet",
				"/path/to/rocker",
				"/path/to/staging",
				"/path/to/tmp",
				"/path/to/baseConfig",
			))
		})
	})

	Describe("Providing the directories to be cleaned before staging", func() {
		It("should return a set of directories to be cleaned", func() {
			Expect(testDirectories.HostDirectoriesToClean()).To(ConsistOf(
				"/path/to/droplet",
				"/path/to/staging",
			))
		})
	})
})
