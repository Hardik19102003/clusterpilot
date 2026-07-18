package events

type EventType string

const (
	Create EventType = "create"
	Update EventType = "update"
	Delete EventType = "delete"
)

type Event struct {
	Resource string
	Name     string
	Type     EventType
}
