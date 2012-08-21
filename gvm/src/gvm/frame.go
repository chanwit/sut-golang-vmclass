package gvm

type stack struct {
    data    []interface{}
    tos     int
}

func (s *stack) Init(size uint16) {
    s.data = make([]interface{}, size)
    s.tos = -1
}

func (s *stack) Push(obj interface{}) {
    s.tos++
    s.data[s.tos] = obj
}

func (s *stack) Pop() (obj interface{}) {
    obj = s.data[s.tos]
    s.tos--
    return
}

type Frame struct {
    stack   *stack
    locals  []interface{}
}

func NewFrame(maxStack uint16, maxLocals uint16) *Frame {
    stack := new(stack)
    stack.Init(maxStack)
    frame := &Frame{stack, make([]interface{}, maxLocals)}
    return frame
}

func (s *Frame) Push(obj interface{}) {
    s.stack.Push(obj)
}

func (s *Frame) Pop() interface{} {
    return s.stack.Pop()
}

func (s *Frame) Store(index int) {
    s.locals[index] = s.Pop()
}

func (s *Frame) Load(index int) {
    s.Push(s.locals[index])
}