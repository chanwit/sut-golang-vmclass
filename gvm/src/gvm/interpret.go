package gvm

import "fmt"
import "encoding/binary"

func Interpret(ca code_attribute, cf *ClassFile) {
    s := new(Stack)
    s.Init(int(ca.max_stack))
    locals := make([]interface{}, ca.max_locals)
    code := ca.code
    pc := 0

    for {
        op := code[pc]
        switch op {
            case LDC:
                //fmt.Println(code[pc], code[pc+1])
                if len(cf.constant_pool[code[pc+1]].info) == 2 {
                    s.Push(int(binary.BigEndian.Uint16(cf.constant_pool[code[pc+1]].info)))
                }else if len(cf.constant_pool[code[pc+1]].info) == 4 {
                    s.Push(int(binary.BigEndian.Uint32(cf.constant_pool[code[pc+1]].info)))
                }else if len(cf.constant_pool[code[pc+1]].info) > 4 {
                    //s.Push(cf.constant_pool[code[pc+1]].info)
                }
                pc = pc + 2
            case ICONST_M1:
                s.Push(-1)
                pc++
            case ICONST_0:
                s.Push(0)
                pc++
            case ICONST_1:
                s.Push(1)
                pc++
            case ICONST_2:
                s.Push(2)
                pc++
            case ICONST_3:
                s.Push(3)
                pc++
            case ICONST_4:
                s.Push(4)
                pc++
            case ICONST_5:
                s.Push(5)
                pc++
            ////case ASTORE_1:
            ////    locals[1] =
            case ISTORE:
                locals[code[pc+1]] = s.Pop()
                pc = pc + 2
            case ISTORE_1:
                locals[1] = s.Pop()
                pc++
            case ISTORE_2:
                locals[2] = s.Pop()
                pc++
            case ISTORE_3:
                locals[3] = s.Pop()
                pc++
            case ILOAD:
                s.Push(locals[code[pc+1]])
                pc = pc + 2
            case ILOAD_1:
                s.Push(locals[1])
                pc++
            case ILOAD_2:
                s.Push(locals[2])
                pc++
            case ILOAD_3:
                s.Push(locals[3])
                pc++
            case BIPUSH:
                s.Push(int(code[pc+1]))
                pc = pc + 2
            case SIPUSH:
                getb := []byte{code[pc+1], code[pc+2]}
                value := int(binary.BigEndian.Uint16(getb))
                s.Push(value)
                pc = pc + 3

            case IADD:
                o1 := s.Pop()
                o2 := s.Pop()
                s.Push(o2.(int) + o1.(int))
                pc++
            case ISUB:
                o1 := s.Pop()
                o2 := s.Pop()
                s.Push(o2.(int) - o1.(int))
                pc++
            case IMUL:
                o1 := s.Pop()
                o2 := s.Pop()
                s.Push(o2.(int) * o1.(int))
                pc++
            case IDIV:
                o1 := s.Pop()
                o2 := s.Pop()
                s.Push(o2.(int) / o1.(int))
                pc++

            case GETSTATIC:
                getb := []byte{code[pc+1], code[pc+2]}
                value := binary.BigEndian.Uint16(getb)
                if cf.constant_pool[value].tag == CONSTANT_Fieldref {
                    fmt.Print("CONSTANT_Fieldref : ")
                    //fmt.Println("fieldref=", cf.constant_pool[value].info)
                    //fmt.Println("class_index=", binary.BigEndian.Uint16(cf.constant_pool[value].info[:2]))
                    //fmt.Println("name_and_type_index=", binary.BigEndian.Uint16(cf.constant_pool[value].info[2:]))
                }
                pc = pc + 3

            case INVOKEVIRTUAL:
                getb := []byte{code[pc+1], code[pc+2]}
                value := binary.BigEndian.Uint16(getb)
                methodRef := cf.constant_pool[value].info
                //class := binary.BigEndian.Uint16(methodRef[:2])
                nameAndType := binary.BigEndian.Uint16(methodRef[2:])

                method := string(cf.constant_pool[binary.BigEndian.Uint16(cf.constant_pool[nameAndType].info[:2])].info[2:])
                signature := string(cf.constant_pool[binary.BigEndian.Uint16(cf.constant_pool[nameAndType].info[2:])].info[2:])

                str := s.Pop().(int)
                varType := cf.constant_pool[str].tag

                if method == "println" && signature == "(Ljava/lang/String;)V" && varType == CONSTANT_Utf8 {
                    fmt.Println(string(cf.constant_pool[str].info[2:]))
                }

                pc = pc + 3

            case RETURN:
                fmt.Println(locals)
                pc++
                return
        }
    }
}