package tests

import (
	"fmt"
	"testing"

	"github.com/Zadigo/citadels/internal/observer"
	"github.com/stretchr/testify/assert"
)

type SimpleSubject struct {
	observer.BaseSubject
}

func (s *SimpleSubject) Update(msg observer.ObserverMessage) error {
	return nil
}

type ErroringSubject struct {
	observer.BaseSubject
}

func (s *ErroringSubject) Update(msg observer.ObserverMessage) error {
	return fmt.Errorf("Some error")
}

func TestBaseObserver(t *testing.T) {
	t.Skip()

	t.Run("simple subject", func(t *testing.T) {
		t.Parallel()

		o := observer.NewBaseObserver()
		assert.IsType(t, o, &observer.BaseObserver{})

		s := &SimpleSubject{}
		o.Register(s)
		o.NotifyAll(observer.ObserverMessage{StrMessage: "Some message"})

		assert.Equal(t, 1, o.SubjectCount())
	})

	t.Run("erroring subject", func(t *testing.T) {
		t.Parallel()

		o := observer.NewBaseObserver()
		assert.IsType(t, o, &observer.BaseObserver{})

		s := &ErroringSubject{}
		o.Register(s)
		o.NotifyAll(observer.ObserverMessage{StrMessage: "Some message"})

		assert.Equal(t, 1, o.SubjectCount())
	})
}
