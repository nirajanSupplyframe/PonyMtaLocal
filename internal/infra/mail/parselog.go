package mail

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

type Parsing struct {
}

type PostfixLog struct {
	Timestamp time.Time
	ProcessID string
	To        string
	Relay     string
	Delay     string
	Delays    string
	DSN       string
	Status    string
	Message   string
	Raw       string // store original line (optional)
}

func NewParsedMessage() *PostfixLog {
	cmd := exec.Command("log", "stream", "--predicate", "process CONTAINS \"smtp\"", "--info")
	stdout, _ := cmd.StdoutPipe()
	err := cmd.Start()
	if err != nil {
		return nil
	}
	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() {
		line := scanner.Text()
		postfixData := ParsePostfixLog(line)

		if strings.Contains(strings.ToLower(line), "dsn") {
			if entry := ParsePostfixLog(line); entry != nil {
				fmt.Printf("Parsed: %+v\n", entry)
				return postfixData
			}
		}
	}
	return nil
}

var postfixRegex = regexp.MustCompile(
	`^(\d{4}-\d{2}-\d{2} [\d:.+-]+)\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+smtp:\s+(\w+): to=<([^>]+)>, relay=([^,]+), delay=([^,]+), delays=([^,]+), dsn=([^,]+), status=(\w+) \((.+)\)$`,
)

func ParsePostfixLog(line string) *PostfixLog {

	matches := postfixRegex.FindStringSubmatch(line)
	if len(matches) == 0 {
		return nil
	}

	timestamp, _ := time.Parse("2006-01-02 15:04:05.000000-0700", matches[1])

	return &PostfixLog{
		Timestamp: timestamp,
		ProcessID: matches[2],
		To:        matches[3],
		Relay:     matches[4],
		Delay:     matches[5],
		Delays:    matches[6],
		DSN:       matches[7],
		Status:    matches[8],
		Message:   matches[9],
		Raw:       line,
	}
}
