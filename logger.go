package mx

import (
	"bytes"
	"fmt"

	"github.com/mdigger/log3"
)

// LogINOUT задает символы, используемые для вывода направления
// (true - входящие, false - исходящие)
var LogINOUT = map[bool]string{true: "→", false: "←"}

// csta форматирует вывод лога с командами CSTA.
func (c *Conn) csta(inFlag bool, id uint16, data []byte) {
	c.mul.RLock()
	if c.logger == nil {
		c.mul.RUnlock()
		return
	}
	var name = data
	if indx := bytes.IndexAny(data, " />"); indx > 1 {
		name = data[1:indx]
	}
	var msg = fmt.Sprintf("%s %s", LogINOUT[inFlag], name)
	if id > 0 && id < 9999 {
		c.logger.Debug(msg, "id", fmt.Sprintf("%04d", id), "xml", string(data))
	} else {
		c.logger.Debug(msg, "xml", string(data))
	}
	c.mul.RUnlock()
}

// SetLogger устанавливает лог.
func (c *Conn) SetLogger(l log.Logger) {
	c.mul.Lock()
	c.logger = l
	c.mul.Unlock()
}
