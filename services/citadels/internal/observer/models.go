package observer

type ObserverMessage struct {
	Sender      SubjectInterface  `json:"sender"`
	StrMessage  string            `json:"str_message"`
	JsonMessage map[string]string `json:"json_message"`
}

type SubjectInterface interface {
	GetName() string
	GetID() string
	Update(ObserverMessage)
}

type ObserverInterface interface {
	SubjectCount() int
	Register(observer SubjectInterface)
	Deregister(observer SubjectInterface)
	NotifyAll(message ObserverMessage)
}

type BaseSubject struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (b *BaseSubject) GetID() string {
	return b.ID
}

func (b *BaseSubject) GetName() string {
	return b.Name
}
