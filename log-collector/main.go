package main

import (
	"fmt"
	"log-collector/logging"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	errorLog         = "error.log"
	loggingDirectory = "performance"
	csvHeader        = "Real Time,Game Date"
	perfLogPrefix    = "perf_date: "
	timeStart        = "["
	timeEnd          = "]"
)

func main() {
	err := os.MkdirAll(loggingDirectory, 0755)
	if err != nil {
		logging.Fatalf("Failed to create application logging directory: %v", err)
		return
	}

	resultFileName := fmt.Sprintf("%s_performance_results.csv", time.Now().Format("20060102150405"))
	logging.Infof("Logging results into %s", resultFileName)

	resultPath := filepath.Join(loggingDirectory, resultFileName)

	resultFile, err := os.OpenFile(resultPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logging.Fatalf("Failed to create result file: %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logging.Fatalf("Failed to close result file: %v", err)
		}
	}(resultFile)

	if _, err := resultFile.WriteString(fmt.Sprint(csvHeader, "\n")); err != nil {
		logging.Fatalf("Failed to write csv header into result file: %v", err)
		return
	}

	results := make(map[string]string)
	for {
		data, err := os.ReadFile(errorLog)
		if err != nil {
			logging.Fatalf("Failed to read %s: %v", errorLog, err)
		}
		for line := range strings.Lines(string(data)) {
			if !strings.Contains(line, perfLogPrefix) {
				continue
			}

			elements := strings.Split(line, perfLogPrefix)
			if len(elements) != 2 {
				continue
			}

			if results[elements[1]] == "" {
				logging.Infof("Found %s", elements[1])

				gameTime := strings.TrimSpace(strings.ReplaceAll(elements[1], ",", "_"))
				realTime := strings.ReplaceAll(elements[0][strings.Index(elements[0], timeStart)+1:strings.Index(elements[0], timeEnd)], ",", "_")

				csv := fmt.Sprintf(
					"%s,%s",
					realTime,
					gameTime,
				)

				if _, err := resultFile.WriteString(fmt.Sprint(csv, "\n")); err != nil {
					logging.Warnf("Failed to write %s into csv, trying again in a second: %v", gameTime, err)
					return
				}
				results[elements[1]] = elements[0]
			}
		}

		time.Sleep(1 * time.Second)
	}

}
