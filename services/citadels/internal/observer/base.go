package observer

import (
	"fmt"
	"sync"
)

type BaseObserver struct {
	BaseSubject
	subjectsList []SubjectInterface `json:"-"`
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
	// for _, o := range g.subjectsList {
	// 	if o.GetID() == nameOrId || o.GetName() == nameOrId {
	// 		chErrors <- o.Update(message)
	// 	}
	// }
}

func (g *BaseObserver) NotifyAll(message ObserverMessage) {
	var wg sync.WaitGroup
	chErrors := make(chan error, len(g.subjectsList))

	for _, o := range g.subjectsList {
		wg.Add(1)
		go func(o SubjectInterface) {
			defer wg.Done()
			_message := message
			if _message.Sender == nil {
				_message.Sender = o
			}
			chErrors <- o.Update(_message)
		}(o)
	}

	wg.Wait()
	close(chErrors)

	for err := range chErrors {
		if err != nil {
			fmt.Println(err)
		}
	}
}

func NewBaseObserver() ObserverInterface {
	return &BaseObserver{
		subjectsList: []SubjectInterface{},
		// chErrors:     make(chan error),
	}
}
