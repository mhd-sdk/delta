package scanner

type Observer interface {
	NotifyScan(results ScanResults)
}
