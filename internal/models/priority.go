package models

type Priority int

const (
	High Priority = iota
	Middle
	Low
)

func (p Priority) toString() string {
	switch p {
	case High:
		return "High"
	case Middle:
		return "Middle"
	case Low:
		return "Low"
	default:
		return "Unknown"
	}
}
