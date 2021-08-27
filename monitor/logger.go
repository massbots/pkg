package monitor

import (
	"encoding/json"
	"time"

	tele "gopkg.in/tucnak/telebot.v3"
)

func (m *Monitor) Info(c tele.Context, msg string, payload ...M) {
	m.log(c, "info", msg, payload...)
}

func (m *Monitor) Debug(c tele.Context, msg string, payload ...M) {
	m.log(c, "debug", msg, payload...)
}

func (m *Monitor) Warn(c tele.Context, msg string, payload ...M) {
	m.log(c, "warn", msg, payload...)
}

func (m *Monitor) Error(c tele.Context, msg string, payload ...M) {
	m.log(c, "error", msg, payload...)
}

func (m *Monitor) log(c tele.Context, level, msg string, payload ...M) {
	record := Record{
		Time:    time.Now(),
		Level:   level,
		Message: msg,
	}
	if c != nil {
		record.UpdateID = uint(c.Update().ID)
	}
	if len(payload) > 0 {
		data, _ := json.Marshal(payload[0])
		record.Payload = string(data)
	}

	m.bus <- record
}
