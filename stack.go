package dsa

import (
    "sync"
    "errors"
)

type Stack[T comparable] struct {
     lock sync.Mutex // you don't have to do this if you don't want thread safety
     s []T
}

func NewStack[T comparable]() *Stack[T] {
    return &Stack[T]{sync.Mutex{}, make([]T,0), }
}

func (s *Stack[T]) Len() int {
    return len(s.s)
}

func (s *Stack[T]) IsEmpty() bool {
    return len(s.s) == 0
}

func (s *Stack[T]) Peek() (T, error) {
    s.lock.Lock()
    defer s.lock.Unlock()

    l := len(s.s)
    if l == 0 {
        var t T
        return t, errors.New("Empty Stack")
    }

    return s.s[l-1], nil
}

func (s *Stack[T]) Push(value T) {
    s.lock.Lock()
    defer s.lock.Unlock()

    s.s = append(s.s, value)
}

func (s *Stack[T]) Pop() (T, error) {
    s.lock.Lock()
    defer s.lock.Unlock()


    l := len(s.s)
    if l == 0 {
        var t T
        return t, errors.New("Empty Stack")
    }

    res := s.s[l-1]
    s.s = s.s[:l-1]
    return res, nil
}


