package functions

import "sync"

func AppendHistory(line string, history *[]string, historyMu *sync.Mutex, maxHistory int) {
	historyMu.Lock()
	*history = append(*history, line)
	if len(*history) > maxHistory {
		*history = (*history)[len(*history)-maxHistory:]
	}
	historyMu.Unlock()
}
