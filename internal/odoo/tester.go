package odoo

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Tester struct {
	args []string
}

func NewTester(args []string) *Tester {
	return &Tester{args: args}
}

func (t *Tester) RunTest() error {
	if len(t.args) == 0 {
		return fmt.Errorf("addon name is required (e.g. odoo test tifshop_product_sync)")
	}
	addons := strings.Join(t.args, ",")
	return t.runDockerTest(addons)
}

func (t *Tester) runDockerTest(addons string) error {
	const composeFile = "docker-compose.test.yml"

	if addons == "" {
		return fmt.Errorf("addon name is required (e.g. odoo test tifshop_product_sync)")
	}

	if err := runCmd("docker", "compose", "-f", composeFile, "build"); err != nil {
		return fmt.Errorf("build failed: %w", err)
	}

	cmd := exec.Command("docker", "compose",
		"-f", composeFile,
		"up", "--build", "--abort-on-container-exit",
		"--exit-code-from", "odoo_test",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(),
		fmt.Sprintf("ADDONS=%s", addons),
	)

	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			code := exitErr.ExitCode()
			fmt.Printf("Test failed (exit code %d)\n", code)
			os.Exit(code)
		}
		return fmt.Errorf("test command failed: %w", err)
	}

	os.Exit(0)
	return nil
}

func runCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
