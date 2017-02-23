package cmd_test

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"runtime"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FruStart", func() {
	var binLocation string

	BeforeEach(func() {
		binLocation = fmt.Sprintf("../bin/%s/workflow-cli", runtime.GOOS)
	})

	Context("When command is called", func() {
		It("UNIT should run the 'fru start' command successfully", func() {
			cmd := exec.Command(binLocation, "fru", "start")

			stdin, err := cmd.StdinPipe()
			Expect(err).To(BeNil())
			defer stdin.Close()

			stdout, err := cmd.StdoutPipe()
			Expect(err).To(BeNil())
			defer stdout.Close()

			stderr, err := cmd.StderrPipe()
			Expect(err).To(BeNil())
			defer stderr.Close()

			cmd.Start()

			io.WriteString(stdin, "a\n")
			time.Sleep(500 * time.Millisecond)

			io.WriteString(stdin, "b\n")
			time.Sleep(500 * time.Millisecond)

			io.WriteString(stdin, "c\n")
			time.Sleep(500 * time.Millisecond)

			io.WriteString(stdin, "1\n")
			time.Sleep(500 * time.Millisecond)

			io.WriteString(stdin, "2\n")
			time.Sleep(500 * time.Millisecond)

			io.WriteString(stdin, "3\n")
			time.Sleep(500 * time.Millisecond)

			buf := new(bytes.Buffer)
			buf.ReadFrom(stderr)
			Expect(buf.String()).To(ContainSubstring("Workflow complete"))

			cmd.Wait()

		})
	})
})
