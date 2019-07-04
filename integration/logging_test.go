package integration_test

import (
	"io/ioutil"
	"os"

	"code.cloudfoundry.org/grootfs/groot"
	"code.cloudfoundry.org/grootfs/integration"
	"code.cloudfoundry.org/grootfs/integration/runner"
	"code.cloudfoundry.org/lager"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Logging", func() {
	var spec groot.CreateSpec

	BeforeEach(func() {
		spec = groot.CreateSpec{
			ID:           "my-image",
			BaseImageURL: integration.String2URL("/non/existent/rootfs"),
			Mount:        mountByDefault(),
		}
	})

	It("forwards human ouput to stdout", func() {
		buffer := gbytes.NewBuffer()

		_, err := Runner.WithStdout(buffer).Create(spec)
		Expect(err).To(HaveOccurred())

		Eventually(buffer).Should(gbytes.Say("no such file or directory"))
	})

	Describe("--log-level and --log-file flags", func() {
		Context("when the --log-file is not set", func() {
			Context("and --log-level is set", func() {
				It("writes logs to stderr", func() {
					buffer := gbytes.NewBuffer()

					_, err := Runner.WithStderr(buffer).WithLogLevel(lager.DEBUG).Create(spec)
					Expect(err).To(HaveOccurred())

					Expect(buffer).To(gbytes.Say(`"error":".*no such file or directory"`))
				})
			})

			Context("and --log-level is not set", func() {
				It("does not write anything to stderr", func() {
					buffer := gbytes.NewBuffer()

					_, err := Runner.WithStderr(buffer).WithoutLogLevel().Create(spec)
					Expect(err).To(HaveOccurred())

					Expect(buffer.Contents()).To(BeEmpty())
				})
			})
		})

		Context("when the --log-file is set", func() {
			var (
				logFilePath string
			)

			getAllTheLogs := func() (string, error) {
				allTheLogs, err := ioutil.ReadFile(logFilePath)
				if err != nil {
					return "", err
				}

				return string(allTheLogs), nil
			}

			BeforeEach(func() {
				logFile, err := ioutil.TempFile("", "log")
				Expect(err).NotTo(HaveOccurred())
				logFilePath = logFile.Name()
				Expect(os.Chmod(logFilePath, 0777)).To(Succeed())
			})

			AfterEach(func() {
				Expect(os.RemoveAll(logFilePath)).To(Succeed())
			})

			Context("and --log-level is set", func() {
				It("forwards logs to the given file", func() {
					_, err := Runner.WithLogFile(logFilePath).WithLogLevel(lager.DEBUG).Create(spec)
					Expect(err).To(HaveOccurred())

					allTheLogs, err := getAllTheLogs()
					Expect(err).NotTo(HaveOccurred())
					Expect(string(allTheLogs)).To(ContainSubstring("\"log_level\":0"))
				})
			})

			Context("and --log-level is not set", func() {
				It("forwards logs to the given file with the log level set to INFO", func() {
					_, err := Runner.WithLogFile(logFilePath).WithoutLogLevel().Create(spec)
					Expect(err).To(HaveOccurred())

					allTheLogs, err := getAllTheLogs()
					Expect(err).NotTo(HaveOccurred())
					Expect(string(allTheLogs)).NotTo(ContainSubstring("\"log_level\":0"))
					Expect(string(allTheLogs)).To(ContainSubstring("\"log_level\":1"))
				})
			})

			Context("and the log file cannot be created", func() {
				It("returns an error to stdout", func() {
					buffer := gbytes.NewBuffer()

					_, err := Runner.WithStdout(buffer).WithLogFile("/path/to/log_file.log").Create(spec)
					Expect(err).To(HaveOccurred())

					Expect(buffer).To(gbytes.Say("no such file or directory"))
				})
			})
		})
	})

	Describe("--clean-log-file", func() {
		var (
			baseImagePath   string
			sourceImagePath string
			cleanLogFile    string
			stderr          *gbytes.Buffer
			r               runner.Runner
		)

		BeforeEach(func() {
			f, err := ioutil.TempFile("", "cleanlog")
			Expect(err).NotTo(HaveOccurred())
			defer f.Close()
			Expect(os.Chown(f.Name(), GrootfsTestUid, GrootfsTestGid)).To(Succeed())

			cleanLogFile = f.Name()
			sourceImagePath = integration.CreateBaseImage(rootUID, rootGID, GrootUID, GrootGID)
			baseImageFile := integration.CreateBaseImageTar(sourceImagePath)
			baseImagePath = baseImageFile.Name()
			spec.BaseImageURL = integration.String2URL(baseImagePath)
			spec.CleanOnCreate = true

			stderr = gbytes.NewBuffer()
			r = Runner.WithStderr(stderr).WithLogLevel(lager.DEBUG).WithClean().WithCleanLogFile(cleanLogFile)
		})

		JustBeforeEach(func() {
			_, err := r.Create(spec)
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			Expect(os.Remove(cleanLogFile)).To(Succeed())
			Expect(os.RemoveAll(sourceImagePath)).To(Succeed())
			Expect(os.RemoveAll(baseImagePath)).To(Succeed())
		})

		readCleanLog := func() (string, error) {
			contents, err := ioutil.ReadFile(cleanLogFile)
			return string(contents), err
		}

		It("writes to the clean log file", func() {
			Eventually(readCleanLog, "10s").Should(ContainSubstring("groot-cleaning"))
		})

		It("does not log to stderr", func() {
			Consistently(stderr).ShouldNot(gbytes.Say("groot-cleaning"))
		})

		Context("when --clean-log-file is not set", func() {
			BeforeEach(func() {
				r = Runner.WithStderr(stderr).WithLogLevel(lager.DEBUG)
			})

			It("writes clean after create entries to stderr", func() {
				Expect(stderr).To(gbytes.Say("groot-cleaning"))
			})
		})
	})
})
