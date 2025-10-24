package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

type AliasInitializer struct {
	Alias    string
	ShellRC  string
	IsFish   bool
}

func NewAliasInitializer() *AliasInitializer {
	return &AliasInitializer{
		Alias: "alias itcm='imotif-tools commit'",
	}
}

func (a *AliasInitializer) Run() (string, error) {
	// Step 1: Prompt for shell config path
	var shellPath string
	err := survey.AskOne(&survey.Input{
		Message: "Enter shell config path (e.g. ~/.zshrc, ~/.bashrc, ~/.config/fish/config.fish):",
	}, &shellPath)
	if err != nil {
		return "", fmt.Errorf("failed to read shell input: %w", err)
	}

	// Expand ~
	shellPath = strings.TrimSpace(shellPath)
	if strings.HasPrefix(shellPath, "~") {
		home, _ := os.UserHomeDir()
		shellPath = filepath.Join(home, shellPath[1:])
	}

	a.ShellRC = shellPath
	a.IsFish = strings.Contains(shellPath, "fish")

	// Step 2: Prepare alias line
	if a.IsFish {
		a.Alias = "alias itcm 'imotif-tools commit'"
	}

	// Step 3: Check if alias already exists
	content, err := os.ReadFile(a.ShellRC)
	if err != nil {
		return "", fmt.Errorf("failed to read shell config: %w", err)
	}
	if strings.Contains(string(content), a.Alias) {
		return "Alias already exists", nil
	}

	// Step 4: Append alias
	f, err := os.OpenFile(a.ShellRC, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to open shell config: %w", err)
	}
	defer f.Close()

	if _, err := f.WriteString("\n" + a.Alias + "\n"); err != nil {
		return "", fmt.Errorf("failed to write alias: %w", err)
	}

	msg := fmt.Sprintf("Alias added to %s\nPlease run: source %s", a.ShellRC, a.ShellRC)
	return msg, nil
}
