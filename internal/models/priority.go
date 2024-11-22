package models

type Priority int

const (
	Low Priority = iota
	Middle
	High
)

func (p Priority) toString() string {
	switch p {
	case Low:
		return "Low"
	case Middle:
		return "Middle"
	case High:
		return "High"
	default:
		return "Unknown"
	}
}
