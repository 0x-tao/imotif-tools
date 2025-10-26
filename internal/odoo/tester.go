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
		return fmt.Errorf("addon name is required (e.g. odoo test my_module)")
	}
	addons := t.args[0]
	return t.runDockerTest(addons)
}

func (t *Tester) runDockerTest(addons string) error {
	const composeFile = "docker-compose.test.yaml"

	// Default fallback
	if addons == "" {
		return fmt.Errorf("addon name is required")
	}

	// Generate coverage pattern list
	addonList := strings.Split(addons, ",")
	var coveragePatterns []string
	for _, a := range addonList {
		a = strings.TrimSpace(a)
		if a != "" {
			coveragePatterns = append(coveragePatterns, fmt.Sprintf("*/%s/*", a))
		}
	}
	coverageInclude := strings.Join(coveragePatterns, ",")

	// Define test command (inline equivalent of bash TEST_COMMAND)
	testCommand := fmt.Sprintf(`coverage run --rcfile=.coveragerc -m odoo \
		-d test_db \
		--addons-path=./addons,./additional-addons \
		-i %s \
		--test-enable \
		--stop-after-init \
		--log-level=test ; coverage report -m --include='%s'`,
		addons, coverageInclude)

	// Step 1: Build image
	fmt.Println("Building test image...")
	if err := runCmd("docker-compose", "-f", composeFile, "build"); err != nil {
		return fmt.Errorf("build failed: %w", err)
	}

	// Step 2: Run test inside container
	fmt.Printf("Installing and testing addons: %s\n", addons)
	if err := runCmd("docker-compose", "-f", composeFile, "run", "--rm", "odoo", "/bin/bash", "-c", testCommand); err != nil {
		return fmt.Errorf("test failed: %w", err)
	}

	// Step 3: Clean up containers & volumes
	fmt.Println("Test run finished. Cleaning up...")
	if err := runCmd("docker-compose", "-f", composeFile, "down", "-v"); err != nil {
		return fmt.Errorf("cleanup failed: %w", err)
	}

	fmt.Println("Done.")
	return nil
}

// Helper to run external commands and stream output to current stdout/stderr
func runCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
