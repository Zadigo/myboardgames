package observer

type BaseObserver struct {
	BaseSubject
	subjectsList []SubjectInterface `json:"-"`
	chErrors     chan error         `json:"-"`
}

func (g *BaseObserver) SubjectCount() int {
	return len(g.subjectsList)
}

func (g *BaseObserver) Register(observer SubjectInterface) {
	g.subjectsList = append(g.subjectsList, observer)
}

func (g *BaseObserver) Deregister(observer SubjectInterface) {
	for i, o := range g.subjectsList {
		if o == observer {
			g.subjectsList[i] = nil
			g.subjectsList = append(g.subjectsList[:i], g.subjectsList[i+1:]...)
			break
		}
	}
}

func (g *BaseObserver) NotifyAll(message ObserverMessage) {
	for _, o := range g.subjectsList {
		go func(o SubjectInterface) {
			if message.Sender == nil {
				message.Sender = o
			}
			o.Update(message)
		}(o)
	}
}

func NewBaseObserver() ObserverInterface {
	return &BaseObserver{
		subjectsList: []SubjectInterface{},
		chErrors:     make(chan error),
	}
}
