package observer

import "sync"

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

func (g *BaseObserver) Notify(nameOrId string, message ObserverMessage) {
	for _, o := range g.subjectsList {
		if o.GetID() == nameOrId || o.GetName() == nameOrId {
			g.chErrors <- o.Update(message)
		}
	}
}

func (g *BaseObserver) NotifyAll(message ObserverMessage) {
	var wg sync.WaitGroup

	for _, o := range g.subjectsList {
		wg.Add(1)
		go func(o SubjectInterface) {
			if message.Sender == nil {
				message.Sender = o
			}
			g.chErrors <- o.Update(message)
			wg.Done()
		}(o)
	}

	wg.Wait()

	<-g.chErrors

	close(g.chErrors)
}

func NewBaseObserver() ObserverInterface {
	return &BaseObserver{
		subjectsList: []SubjectInterface{},
		chErrors:     make(chan error),
	}
}
