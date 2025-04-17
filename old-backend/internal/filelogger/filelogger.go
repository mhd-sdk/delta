package filelogger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/delta/internal/worker/scanner"
)

type FileLogger struct {
	file       *os.File
	logger     *log.Logger
	seenAssets map[string]bool
}

func NewLogger(logFile string) (*FileLogger, error) {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &FileLogger{
		file:       file,
		logger:     log.New(file, "", log.LstdFlags),
		seenAssets: make(map[string]bool),
	}, nil
}

// write in the log file the results of the scan
func (l *FileLogger) NotifyScan(results scanner.ScanResults) {
	for _, result := range results {
		// at least 2 times the average volume and a 5% change
		if isAssetValid(result.Asset.Name) && result.DayChange > 0.1 {
			if _, seen := l.seenAssets[result.Asset.Symbol]; !seen {
				logEntry := fmt.Sprintf(
					"[%s] Asset: %s (%s), AvgChange: %.2f, DayChange: %.2f, LastPrice: %.2f\n",
					time.Now().Format(time.RFC3339),
					result.Asset.Symbol, result.Asset.Name,
					result.AvgChange, result.DayChange, result.LastPrice,
				)
				l.logger.Println(logEntry)
				l.seenAssets[result.Asset.Symbol] = true
			}
		}
	}
}
func (l *FileLogger) Close() {
	l.file.Close()
}

func isAssetValid(name string) bool {
	if strings.Contains(name, "Warrant") {
		return false
	}
	if strings.Contains(name, "ETF") {
		return false
	}
	return true
}
