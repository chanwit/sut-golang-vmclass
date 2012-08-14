package gvm

type Stack struct {
    data    []int
    tos     int
}

func (s *Stack) Init(size int) {
    s.data = make([]int, size)
    s.tos = -1
}

func (s *Stack) Push(num int) {
    s.tos++
    s.data[s.tos] = num
}

func (s *Stack) Pop() (num int) {
    num = s.data[s.tos]
    s.tos--
    return
}