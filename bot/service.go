package bot

import (
	"bytes"
	"fmt"
	"os/exec"
)

func AnalyzeSentiment(text string) (string, error) {
	cmd := exec.Command("model/.venv/bin/python3", "model/sentiment.py", text)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error running command: %v", err)
	}

	return out.String(), nil
}
