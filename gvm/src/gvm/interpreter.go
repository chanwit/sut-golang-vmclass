package gvm

import "encoding/binary"
import "fmt"

func Interpret(ca code_attribute, cf *ClassFile) {
    // s := new(Stack)
    // s.Init(int(ca.max_stack))
    // locals := make([]interface{}, ca.max_locals)
    s := NewFrame(ca.max_stack, ca.max_locals)
    code := ca.code
    pc := 0

    for {

        op := code[pc]
        pc++

        switch op {
            case LDC:
                switch len(cf.constant_pool[code[pc]].info) {
                    case 2:
                        s.Push(int(binary.BigEndian.Uint16(cf.constant_pool[code[pc]].info)))
                    case 4:
                        s.Push(int(binary.BigEndian.Uint32(cf.constant_pool[code[pc]].info)))
                }
                pc++

            case ICONST_M1:
                s.Push(-1)
            case ICONST_0:
                s.Push(0)
            case ICONST_1:
                s.Push(1)
            case ICONST_2:
                s.Push(2)
            case ICONST_3:
                s.Push(3)
            case ICONST_4:
                s.Push(4)
            case ICONST_5:
                s.Push(5)

            case ISTORE:
                s.Store(int(code[pc]))
                pc++
            case ISTORE_0:
                s.Store(0)
            case ISTORE_1:
                s.Store(1)
            case ISTORE_2:
                s.Store(2)
            case ISTORE_3:
                s.Store(3)

            case ILOAD:
                s.Load(int(code[pc]))
                pc++
            case ILOAD_0:
                s.Load(0)
            case ILOAD_1:
                s.Load(1)
            case ILOAD_2:
                s.Load(2)
            case ILOAD_3:
                s.Load(3)

            case BIPUSH:
                s.Push(int(code[pc]))
                pc++
            case SIPUSH:
                // bytes := []byte{code[pc+1], code[pc+2]}
                value := int(binary.BigEndian.Uint16(code[pc:pc+2]))
                s.Push(value)
                pc = pc + 2
            case IADD:
                o1 := s.Pop()
                o2 := s.Pop()
                s.Push(o2.(int) + o1.(int))
            case ISUB:
                o1 := s.Pop()
                o2 := s.Pop()
                s.Push(o2.(int) - o1.(int))
            case IMUL:
                o1 := s.Pop()
                o2 := s.Pop()
                s.Push(o2.(int) * o1.(int))
            case IDIV:
                o1 := s.Pop()
                o2 := s.Pop()
                s.Push(o2.(int) / o1.(int))
            case GETSTATIC:
                value := binary.BigEndian.Uint16(code[pc:pc+2])
                if cf.constant_pool[value].tag == CONSTANT_Fieldref {
                    fmt.Print("CONSTANT_Fieldref : ")
                    //fmt.Println("fieldref=", cf.constant_pool[value].info)
                    //fmt.Println("class_index=", binary.BigEndian.Uint16(cf.constant_pool[value].info[:2]))
                    //fmt.Println("name_and_type_index=", binary.BigEndian.Uint16(cf.constant_pool[value].info[2:]))
                }
                pc = pc + 2

            case INVOKEVIRTUAL:
                strIndex := s.Pop().(int)
                fmt.Println(string(cf.constant_pool[strIndex].info[2:]))
                pc = pc + 2

            case RETURN:
                fmt.Println(s.locals)
                return
        }
    }
}