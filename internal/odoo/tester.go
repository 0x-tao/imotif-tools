package odoo

import (
	"os"
	"os/exec"

	"github.com/imotif-tools/pkg/text"
)

type Tester struct {
	args []string
}

func NewTester(args []string) *Tester {
	return &Tester{
		args: args,
	}
}

func (t *Tester) RunTest() error {
	parser := text.NewParser(t.args)
	addon, err := parser.Parse("addon name is required")
	if err != nil {
		return err
	}
	if err := t.runDockerTest(addon); err != nil {
		return err
	}
	return nil
}

func (t *Tester) runDockerTest(addons string) error {
	cmd := exec.Command("./odoo-test.sh", addons)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
