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
	file   *os.File
	logger *log.Logger
}

func NewLogger(logFile string) (*FileLogger, error) {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &FileLogger{
		file:   file,
		logger: log.New(file, "", log.LstdFlags),
	}, nil
}

// write in the log file the results of the scan
func (l *FileLogger) NotifyScan(results scanner.ScanResults) {
	for _, result := range results {
		// at least 2 times the average volume and a 5% change
		if isAssetValid(result.Asset.Name) && result.RelativeVolume > 2 && result.DayChange > 0.05 {
			logEntry := fmt.Sprintf(
				"[%s] Asset: %s (%s), AvgVolume: %d, CurrentVolume: %d, RelativeVolume: %d, AvgChange: %.2f, DayChange: %.2f, LastPrice: %.2f\n",
				time.Now().Format(time.RFC3339),
				result.Asset.Symbol, result.Asset.Name,
				result.AvgVolume, result.CurrentVolume, result.RelativeVolume,
				result.AvgChange, result.DayChange, result.LastPrice,
			)
			l.logger.Println(logEntry)
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
