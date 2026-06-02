package utils

type Log struct {
	Id     int
	Action string
}

var Logs = []Log{}

func AppendLog(action string) {
	Logs = append(Logs, Log{
		Id:     len(Logs) + 1,
		Action: action,
	})
}
