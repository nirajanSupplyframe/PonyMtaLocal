package mail

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

type Parsing struct {
}

func NewParsedMessage() string {
	cmd := exec.Command("log", "stream", "--predicate", "process CONTAINS \"smtp\"", "--info")
	stdout, _ := cmd.StdoutPipe()
	err := cmd.Start()
	if err != nil {
		return ""
	}
	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Process CONTAINS \"dsn\"") {
			return line
		}
		fmt.Println(line)
	}
	return "Parse mail"
}
