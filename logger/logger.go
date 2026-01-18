package logger

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type LogEntry struct {
	Time     string `json:"time"`
	IP       string `json:"ip"`
	Method   string `json:"method"`
	Path     string `json:"path"`
	Score    int    `json:"score"`
	Decision string `json:"decision"`
	Reason   string `json:"reason"`
}

var output = log.New(os.Stdout, "", 0)

func Write(entry LogEntry) {
	entry.Time = time.Now().Format(time.RFC3339)

	data, _ := json.Marshal(entry)
	output.Println(string(data))
}
