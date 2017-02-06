package groot_test

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"

	"code.cloudfoundry.org/grootfs/integration/runner"
	"code.cloudfoundry.org/grootfs/testhelpers"
	"code.cloudfoundry.org/lager"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

var (
	GrootFSBin string
	DraxBin    string
	Driver     string
	MountPath  string
	Runner     runner.Runner
	StorePath  string
	StoreName  string

	CurrentUserID    string
	CurrentUserIDInt int
	RegistryUsername string
	RegistryPassword string
)

const btrfsMountPath = "/mnt/btrfs"
const xfsMountPath = "/mnt/xfs"

func TestGroot(t *testing.T) {
	RegisterFailHandler(Fail)

	SynchronizedBeforeSuite(func() []byte {
		grootFSBin, err := gexec.Build("code.cloudfoundry.org/grootfs")
		Expect(err).NotTo(HaveOccurred())

		return []byte(grootFSBin)
	}, func(data []byte) {
		CurrentUserIDInt = os.Getuid()
		CurrentUserID = strconv.Itoa(CurrentUserIDInt)
		GrootFSBin = string(data)
		Driver = os.Getenv("VOLUME_DRIVER")
		if Driver == "overlay-xfs" {
			MountPath = xfsMountPath
		} else {
			Driver = "btrfs"
			MountPath = btrfsMountPath
		}

		fmt.Fprintf(os.Stderr, "============> RUNNING %s TESTS <=============", Driver)
	})

	SynchronizedAfterSuite(func() {
	}, func() {
		gexec.CleanupBuildArtifacts()
	})

	BeforeEach(func() {
		StoreName = fmt.Sprintf("test-store-%d", GinkgoParallelNode())
		StorePath = path.Join(MountPath, StoreName)
		Expect(os.Mkdir(StorePath, 0755)).NotTo(HaveOccurred())

		var err error
		DraxBin, err = gexec.Build("code.cloudfoundry.org/grootfs/store/filesystems/btrfs/drax")
		Expect(err).NotTo(HaveOccurred())
		testhelpers.SuidDrax(DraxBin)

		RegistryUsername = os.Getenv("REGISTRY_USERNAME")
		RegistryPassword = os.Getenv("REGISTRY_PASSWORD")

		Runner = runner.Runner{
			GrootFSBin: GrootFSBin,
			StorePath:  StorePath,
			DraxBin:    DraxBin,
			Driver:     Driver,
		}.WithLogLevel(lager.DEBUG).WithStderr(GinkgoWriter)
	})

	AfterEach(func() {
		testhelpers.CleanUpBtrfsSubvolumes(btrfsMountPath, StoreName)
		testhelpers.CleanUpOverlayMounts(xfsMountPath, StoreName)
		Expect(os.RemoveAll(StorePath)).To(Succeed())
	})

	RunSpecs(t, "GrootFS Integration Suite - Running as groot")
}

func writeMegabytes(outputPath string, mb int) error {
	cmd := exec.Command("dd", "if=/dev/zero", fmt.Sprintf("of=%s", outputPath), "bs=1048576", fmt.Sprintf("count=%d", mb))
	sess, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	if err != nil {
		return err
	}
	Eventually(sess).Should(gexec.Exit())
	if sess.ExitCode() > 0 {
		return errors.New(string(sess.Err.Contents()))
	}
	return nil
}
