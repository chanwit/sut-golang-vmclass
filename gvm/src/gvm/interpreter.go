package gvm

func Interpret(ca code_attribute, cp []cp_info) {
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
                index := u16(cp[code[pc]].info)
                switch cp[index].tag {
                    case CONSTANT_Utf8:
                        obj := &Object{ClassName:"java/lang/String",
                                       Native: string(cp[index].info[2:])}
                        s.Push(obj)
                    case CONSTANT_Integer:
                        value := i32(cp[index].info)
                        s.Push(value)
                    case CONSTANT_Class:
                        panic("LDC pushing CONSTANT_Class, NYI")
                    case CONSTANT_Float:
                        panic("LDC pushing CONSTANT_Float, NYI")
                    case CONSTANT_Long:
                        panic("LDC pushing CONSTANT_Long, NYI")
                    case CONSTANT_Double:
                        panic("LDC pushing CONSTANT_Double, NYI")
                    case CONSTANT_String:
                        panic("LDC pushing CONSTANT_String, NYI")
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

            case ALOAD_0:
                s.Load(0)

            case BIPUSH:
                s.Push(int(code[pc]))
                pc++
            case SIPUSH:
                value := int(u16(code[pc:pc+2]))
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
                value := u16(code[pc:pc+2])
                if cp[value].tag != CONSTANT_Fieldref {
                    panic("CONSTANT_Fieldref")
                }

                _debug("CONSTANT_Fieldref : ")
                _debug("fieldref=", cp[value].info)
                ownerIndex := u16(cp[value].info[:2])
                nameAndTypeIndex := u16(cp[value].info[2:])
                ownerClassIndex  := u16(cp[ownerIndex].info[:2])
                nameIndex := u16(cp[nameAndTypeIndex].info[:2])
                typeIndex := u16(cp[nameAndTypeIndex].info[2:])

                ownerName := string(cp[ownerClassIndex].info[2:])
                fieldName := string(cp[nameIndex].info[2:])
                fieldTypeName := string(cp[typeIndex].info[2:])

                obj := CT(ownerName).StaticFields[fieldName]
                s.Push(obj)
                _debug(fieldTypeName)

                pc = pc + 2

            case INVOKEVIRTUAL:
                _debug("INVOKEVIRTUAL")
                methodRefIndex := u16(code[pc:pc+2])
                ownerIndex := u16(cp[methodRefIndex].info[:2])
                nameAndTypeIndex := u16(cp[methodRefIndex].info[2:])

                ownerClassIndex := u16(cp[ownerIndex].info)

                _debug(cp[nameAndTypeIndex].info)

                nameIndex := u16(cp[nameAndTypeIndex].info[:2])
                typeIndex := u16(cp[nameAndTypeIndex].info[2:])
                owner := string(cp[ownerClassIndex].info[2:])

                _debug(owner)

                desc := string(cp[typeIndex].info[2:])
                signature := string(cp[nameIndex].info[2:]) + desc

                _debug(signature)

                method := CT(owner).Methods[signature]
                if method == nil {
                    _debug(CT(owner).Methods)
                    panic("Method not found in CT: " + signature)
                }
                argCount := method.GetArgCount()
                args := make([]interface{}, argCount)
                for i := 0; i < argCount; i++ {
                    args[i] = s.Pop()
                }
                recv := s.Pop().(*Object)
                if void, ret := method.Invoke(recv, args); !void {
                    s.Push(ret)
                }
                pc = pc + 2

            case RETURN:
                _debug(s.locals)
                return

            default:
                panic("NYI opcode: " + opCodes[op])
        }
    }
}
