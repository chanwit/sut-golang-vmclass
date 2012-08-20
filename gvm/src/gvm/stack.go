package gvm

type Stack struct {
    data    []interface{}
    tos     int
}

func (s *Stack) Init(size int) {
    s.data = make([]interface{}, size)
    s.tos = -1
}

func (s *Stack) Push(obj interface{}) {
    s.tos++
    s.data[s.tos] = obj
}

func (s *Stack) Pop() (obj interface{}) {
    obj = s.data[s.tos]
    s.tos--
    return
}