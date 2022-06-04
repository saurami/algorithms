package stack

import (
	"testing"
	"reflect"
)

func TestStackWithData(t *testing.T) {

	var s stack

	s.push(10)
	s.push(18)
	s.push(7)
	s.push(9)

	t.Run("view top of stack", func(t *testing.T) {
		val, _ := s.peek()
		want := 9

		if val != want {
			t.Errorf("Got %v; wanted %v", val, want)
		}
	})

	t.Run("check unempty stack", func(t *testing.T) {
		got := s.isEmpty()
		want := false

		if got != want {
			t.Errorf("Stack contains data. Got %v; wanted %v", got, want)
		}
	})

	t.Run("remove element from stack top", func(t *testing.T) {
		got := s.pop(); s.pop()
		want := true

		if got != want {
			t.Errorf("Got %v; wanted %v", got, want)
		}
	})

	t.Run("view all elements", func(t *testing.T) {
		var got stack  = s.show()
		var want stack = []int{10, 18}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v; wanted %v", got, want)
		}
	})

}

func TestStackWithoutData(t *testing.T) {

	var s stack

	t.Run("is empty function", func(t *testing.T) {
		got := s.isEmpty()
		want := true

		if got != want {
			t.Errorf("Got %v; wanted empty stack status %v", got, want)
		}
	})

	t.Run("peek function", func(t *testing.T) {
		val, msg := s.peek()
		want, err := 0, "Stack is empty..."

		if val != want && msg.Error() != err {
			t.Errorf("Got value %v and message %v; wanted value %v and message %v", val, msg, want, err)
		}
	})

	t.Run("pop function", func(t *testing.T) {
		got := s.pop()
		want := false

		if got != want {
			t.Errorf("Wanted pop operation to be unsuccessfull, i.e., %v; but got %v", want, got)
		}
	})
}

