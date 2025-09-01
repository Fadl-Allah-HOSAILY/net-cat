package functions

import (
	"fmt"

	gb "netCat/global"
)

func FormatMessage(m *gb.Message) string {
	return fmt.Sprintf("[%s][%s]:%s", m.Timestamp.Format("2006-01-02 15:04:05"), m.From, m.Text)
}
