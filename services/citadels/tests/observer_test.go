package tests

import (
	"testing"

	"github.com/Zadigo/citadels/internal/observer"
	"github.com/stretchr/testify/assert"
)

type SimpleSubject struct {
	observer.BaseSubject
}

func (s *SimpleSubject) Update(msg observer.ObserverMessage) {}

func TestBaseObserver(t *testing.T) {
	o := observer.NewBaseObserver()
	assert.IsType(t, o, &observer.BaseObserver{})

	s := &SimpleSubject{}
	o.Register(s)
	o.NotifyAll(observer.ObserverMessage{StrMessage: "Some message"})

	assert.Equal(t, 1, o.SubjectCount())
}
