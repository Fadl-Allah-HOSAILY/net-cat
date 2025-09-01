package functions

import (
	"sync"

	gb "netCat/global"
)

func SendHistory(c *gb.Client, history *[]string, historyMu *sync.Mutex) {
	historyMu.Lock()
	for _, h := range *history {
		c.Ch <- h
	}
	historyMu.Unlock()
}
