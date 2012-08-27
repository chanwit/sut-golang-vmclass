package gvm

import "fmt"

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
                index := u16(cf.constant_pool[code[pc+1]].info)
                switch cf.constant_pool[index].tag {
                    case CONSTANT_Utf8:
                        obj := &Object{ClassName: "java/lang/String", Native: string(cf.constant_pool[index].info[2:])}
                        s.Push(obj)
                    case CONSTANT_Integer:
                        value := i32(cf.constant_pool[index].info)
                        s.Push(value)
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

            case ASTORE_0:
                locals[0] = s.Pop()
                pc++
            case ASTORE_1:
                locals[1] = s.Pop()
                pc++
            case ASTORE_2:
                locals[2] = s.Pop()
                pc++
            case ASTORE_3:
                locals[3] = s.Pop()
                pc++

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
                bytes := []byte{code[pc+1], code[pc+2]}
                value := int(u16(bytes))
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
                bytes := []byte{code[pc+1], code[pc+2]}
                value := u16(bytes)
                if cf.constant_pool[value].tag == CONSTANT_Fieldref {
                    fmt.Print("CONSTANT_Fieldref : ")
                    ownerIndex := u16(cf.constant_pool[value].info[:2])
                    nameAndTypeIndex := u16(cf.constant_pool[value].info[2:])
                    ownerClassIndex := u16(cf.constant_pool[ownerIndex].info[:2])
                    nameIndex := u16(cf.constant_pool[nameAndTypeIndex].info[:2])
                    typeIndex := u16(cf.constant_pool[nameAndTypeIndex].info[2:])

                    ownerName := string(cf.constant_pool[ownerClassIndex].info[2:])
                    fieldName := string(cf.constant_pool[nameIndex].info[2:])
                    fieldTypeName := string(cf.constant_pool[typeIndex].info[2:])

                    obj := CT(ownerName).StaticFields[fieldName]
                    s.Push(obj)

                    fmt.Println(fieldTypeName)
                }
                pc = pc + 3

            case INVOKEVIRTUAL:
                fmt.Println("INVOKEVIRTUAL")
                methodRefIndex := u16(code[pc+1:pc+3])
                ownerIndex := u16(cf.constant_pool[methodRefIndex].info[:2])
                nameAndTypeIndex := u16(cf.constant_pool[methodRefIndex].info[2:])

                ownerClassIndex := u16(cf.constant_pool[ownerIndex].info)

                fmt.Println(cf.constant_pool[nameAndTypeIndex].info)

                nameIndex := u16(cf.constant_pool[nameAndTypeIndex].info[:2])
                typeIndex := u16(cf.constant_pool[nameAndTypeIndex].info[2:])
                owner := string(cf.constant_pool[ownerClassIndex].info[2:])

                fmt.Println(owner)

                desc := string(cf.constant_pool[typeIndex].info[2:])
                signature := string(cf.constant_pool[nameIndex].info[2:]) + desc

                fmt.Println(signature)

                method := CT(owner).Methods[signature]
                argCount := method.GetArgCount()
                args := make([]*Object, argCount)
                for i := 0; i < argCount; i++ {
                    a := s.Pop()
                    switch a.(type) {
                        case int:
                            args[i] = &Object{Native: a.(int)}
                        case *Object:
                            args[i] = a.(*Object)
                    }
                }
                recv := s.Pop().(*Object)
                if void, ret := method.Invoke(recv, args); !void {
                    s.Push(ret)
                }
                pc = pc + 3

            case INVOKESPECIAL:
                fmt.Println("INVOKESPECIAL")
                methodRefIndex := u16(code[pc+1:pc+3])
                ownerIndex := u16(cf.constant_pool[methodRefIndex].info[:2])
                nameAndTypeIndex := u16(cf.constant_pool[methodRefIndex].info[2:])

                ownerClassIndex := u16(cf.constant_pool[ownerIndex].info)

                fmt.Println(cf.constant_pool[nameAndTypeIndex].info)

                nameIndex := u16(cf.constant_pool[nameAndTypeIndex].info[:2])
                typeIndex := u16(cf.constant_pool[nameAndTypeIndex].info[2:])
                owner := string(cf.constant_pool[ownerClassIndex].info[2:])

                fmt.Println(owner)

                desc := string(cf.constant_pool[typeIndex].info[2:])
                signature := string(cf.constant_pool[nameIndex].info[2:]) + desc

                fmt.Println(signature)

                method := CT(owner).Methods[signature]
                argCount := method.GetArgCount()
                args := make([]*Object, argCount)
                for i := 0; i < argCount; i++ {
                    a := s.Pop()
                    switch a.(type) {
                        case int:
                            args[i] = &Object{Native: a.(int)}
                        case *Object:
                            args[i] = a.(*Object)
                    }
                }
                recv := s.Pop().(*Object)
                if void, ret := method.Invoke(recv, args); !void {
                    s.Push(ret)
                }

                pc = pc + 3

            case NEW:
                fmt.Println("NEW")
                methodRefIndex := u16(code[pc+1:pc+3])
                ownerIndex := u16(cf.constant_pool[methodRefIndex].info[:2])

                ownerClassIndex := u16(cf.constant_pool[ownerIndex].info)
                owner := string(cf.constant_pool[ownerClassIndex].info[2:])

                obj := &Object{ClassName: owner}
                s.Push(obj)
                pc = pc + 3

            case DUP:
                obj := s.Pop()
                s.Push(obj)
                s.Push(obj)
                pc++

            case RETURN:
                fmt.Println(locals)
                pc++
                return
        }
    }
}