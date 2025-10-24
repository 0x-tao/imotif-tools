package cli

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/imotif-tools/pkg/text"
)

type CommitPrompter struct {
	args []string
}

func NewCommitPrompter(args []string) *CommitPrompter {
	return &CommitPrompter{
		args: args,
	}
}

func (c *CommitPrompter) RunCommit(msg ...string) error {
	var message string
	if len(msg) > 0 {
		fmt.Println("Using commit message:", msg[0])
		message = msg[0]
	} else {
		parser := text.NewParser(c.args)
		msg, err := parser.Parse("commit message is required")
		if err != nil {
			return err
		}
		message = msg
	}
	taskInput, err := c.promptTaskID()
	if err != nil {
		return err
	}
	formattedTask := c.formatTaskID(taskInput)

	commitType, err := c.promptCommitType()
	if err != nil {
		return err
	}

	finalMessage := fmt.Sprintf("%s[%s] %s", formattedTask, commitType, message)

	err = c.runGitCommit(finalMessage)
	if err != nil {
		return err
	}

	return nil
}

func (c *CommitPrompter) promptTaskID() (string, error) {
	var taskInput string
	err := survey.AskOne(&survey.Input{
		Message: "Enter Task NO. (e.g. Task-1 or Task-1,Task-2):",
	}, &taskInput)
	return strings.TrimSpace(taskInput), err
}

func (c *CommitPrompter) formatTaskID(input string) string {
	tasks := strings.Split(input, ",")
	for i := range tasks {
		tasks[i] = strings.TrimSpace(tasks[i])
	}
	return "[" + strings.Join(tasks, ",") + "]"
}

func (c *CommitPrompter) promptCommitType() (string, error) {
	typeMap := map[string]string{
		"FIX":   "Fix a bug or incorrect behavior",
		"ADD":   "Add new code, files, or features",
		"REF":   "Refactor code without changing behavior",
		"MIG":   "Migration related changes (e.g., DB schema)",
		"DOC":   "Documentation updates or changes",
		"STYLE": "Code style changes (no logic impact)",
		"TEST":  "Add or update tests",
		"CHORE": "Maintenance tasks (e.g., config, build)",
		"PERF":  "Performance improvements",
		"FEAT":  "Introduce a new feature",
		"ISS":   "Work related to a reported issue or bug task",
	}

	fmt.Println("Available Commit Types:")
	for k, v := range typeMap {
		fmt.Printf("  [%s] %s\n", k, v)
	}

	var inputType string
	err := survey.AskOne(&survey.Input{
		Message: "Enter commit type:",
	}, &inputType)

	return strings.ToUpper(strings.TrimSpace(inputType)), err
}

func (c *CommitPrompter) runGitCommit(msg string) error {
	cmd := exec.Command("git", "commit", "-m", msg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (c *CommitPrompter) ConfirmOrEditMessage(defaultMsg string) (string, error) {
	var useDefault bool
	confirmPrompt := &survey.Confirm{
		Message: fmt.Sprintf("Use this AI commit message?\n→ %s\n", defaultMsg),
		Default: true,
	}
	err := survey.AskOne(confirmPrompt, &useDefault)
	if err != nil {
		return "", err
	}

	if useDefault {
		return defaultMsg, nil
	}

	var newMsg string
	inputPrompt := &survey.Input{
		Message: "Enter your custom commit message:",
	}
	err = survey.AskOne(inputPrompt, &newMsg)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(newMsg), nil
}
