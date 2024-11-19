package models

type StatusOfTask int

const (
	NeedToDo StatusOfTask = iota
	Blocked
	InProgress
	Review
	Done
)

func (s StatusOfTask) toString() string {
	switch s {
	case NeedToDo:
		return "Need to do"
	case Blocked:
		return "Blocked"
	case InProgress:
		return "In progress"
	case Review:
		return "Review"
	case Done:
		return "Done"
	default:
		return "Unknown"
	}
}
