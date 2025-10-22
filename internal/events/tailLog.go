package events

import (
	"bufio"
	"os/exec"
	"regexp"
	"strings"
)

var (
	reMessageID = regexp.MustCompile(`(?i)([0-9A-F]+): .*message-id=<([^>]+)>`)
	reStatus    = regexp.MustCompile(`(?i)([0-9A-F]+): .*status=(\w+)\s*\((.*)\)`)
)

func TailLog(sm *StateManager) {
	cmd := exec.Command("log", "stream", "--predicate", "process contains \"smtp\" OR process contains \"cleanup\"", "--info")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()

		if !strings.Contains(line, ":") {
			continue
		}

		//message-id link
		if m := reMessageID.FindStringSubmatch(line); len(m) >= 3 {
			qid := m[1]
			msgid := m[2]

			mesID := extractMessageIDFromLogs(msgid)
			if mesID != "" {
				sm.Publish(Event{
					Type:      EventLinked,
					RequestID: mesID,
					QueueID:   qid,
					Raw:       line,
				})
			}
			println("email LINKED with (tail log) request id :", mesID)
			continue
		}

		if m := reStatus.FindStringSubmatch(line); len(m) >= 4 {
			qid := m[1]
			status := strings.ToUpper(m[2])
			reason := m[3]
			sm.Publish(Event{
				Type:    EventStatus,
				QueueID: qid,
				Status:  status,
				Raw:     reason,
			})
			println("email STATUS check  with qid :", qid)
			continue
		}

		//if m:=reMessageID.FindStringSubmatch(line); len(m) > 0 {
		//	sm.Publish(Event{
		//		Type: EmailDelivered,
		//		QueueID: qid
		//	})
		//}

	}
}

func extractMessageIDFromLogs(msgid string) string {
	msgid = strings.TrimSpace(msgid)
	msgid = strings.TrimPrefix(msgid, "<")
	msgid = strings.TrimSuffix(msgid, ">")
	parts := strings.SplitN(msgid, "@", 2)
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}
