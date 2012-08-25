package gvm
//import "fmt"

func Interpret(ca code_attribute, cp []cp_info) {
     s := new(Stack)
     s.Init(ca.max_stack)
     locals := make([]interface{}, ca.max_locals)
    //s := NewFrame(ca.max_stack, ca.max_locals)
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
                    // case CONSTANT_Class:
                    // case CONSTANT_Float:
                    // case CONSTANT_Long:
                    // case CONSTANT_Double:
                    // case CONSTANT_String:
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
                index:=code[pc]
                locals[int(index)] = s.Pop()
                pc++
            case ISTORE_0:
                locals[0] = s.Pop()//s.Store(0)
            case ISTORE_1:
                locals[1] = s.Pop()//s.Store(1)
            case ISTORE_2:
                locals[2] = s.Pop()//s.Store(2)
            case ISTORE_3:
                locals[3] = s.Pop()//s.Store(3)

            case ILOAD:
                s.Push(int(code[pc]))
                pc++
            case ILOAD_0:
                s.Push(locals[0])
            case ILOAD_1:
                s.Push(locals[1])
            case ILOAD_2:
                s.Push(locals[2])
            case ILOAD_3:
                s.Push(locals[3])

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
                //fmt.Println("CONSTANT_Fieldref : ")
                //fmt.Println("fieldref=", cp[value].info)
                ownerIndex := u16(cp[value].info[:2])
                nameAndTypeIndex := u16(cp[value].info[2:])
                ownerClassIndex  := u16(cp[ownerIndex].info[:2])
                nameIndex := u16(cp[nameAndTypeIndex].info[:2])
                //typeIndex := u16(cp[nameAndTypeIndex].info[2:])

                ownerName := string(cp[ownerClassIndex].info[2:])
                fieldName := string(cp[nameIndex].info[2:])
                //fieldTypeName := string(cp[typeIndex].info[2:])
                obj := CT(ownerName).StaticFields[fieldName]
                s.Push(obj)

                //fmt.Println(fieldTypeName)

                pc = pc + 2
            case INVOKEVIRTUAL:
                //fmt.Println("INVOKEVIRTUAL")
                methodRefIndex := u16(code[pc:pc+2])
                ownerIndex := u16(cp[methodRefIndex].info[:2])
                nameAndTypeIndex := u16(cp[methodRefIndex].info[2:])

                ownerClassIndex := u16(cp[ownerIndex].info)

                //fmt.Println(cp[nameAndTypeIndex].info)

                nameIndex := u16(cp[nameAndTypeIndex].info[:2])
                typeIndex := u16(cp[nameAndTypeIndex].info[2:])
                owner := string(cp[ownerClassIndex].info[2:])

                //fmt.Println(owner)

                desc := string(cp[typeIndex].info[2:])
                signature := string(cp[nameIndex].info[2:]) + desc

                //fmt.Println(signature)
                method := CT(owner).Methods[signature]
                argCount := method.GetArgCount()
                args := make([]*Object, argCount)
                for i := 0; i < argCount; i++ {
                    a := s.Pop()
                    switch a.(type) {
                        case int:
                            args[i] = &Object{Native:a.(int)} 
                        case *Object:
                           args[i] = &Object{Native:(a.(*Object)).Native}
                    }
                }                    
                recv := s.Pop().(*Object)
                if void, ret := method.Invoke(recv, args); !void {
                    s.Push(ret)                    
                }
                pc = pc + 2
            case INVOKESPECIAL:
                methodRefIndex := u16(code[pc:pc+2])
                //ownerIndex := u16(cp[methodRefIndex].info[:2])
                nameAndTypeIndex := u16(cp[methodRefIndex].info[2:])

                //ownerClassIndex := u16(cp[ownerIndex].info)

                //fmt.Println(cp[nameAndTypeIndex].info)

                nameIndex := u16(cp[nameAndTypeIndex].info[:2])
                typeIndex := u16(cp[nameAndTypeIndex].info[2:])
                //owner := string(cp[ownerClassIndex].info[2:])

                //fmt.Println(owner)

                desc := string(cp[typeIndex].info[2:])
                signature := string(cp[nameIndex].info[2:]) + desc
                //fmt.Println(signature)
                reciver := s.Pop().(*Object)
                method := CT(reciver.ClassName).StaticFields[signature]
                reciver.Native = method.Native
                pc = pc + 2
            case NEW:
                value :=make([]byte,4)
                    value[0] = code[pc]
                    value[1] = code[pc+1]
                index := u16(value)
                classindex := u16(cp[index].info[:])
                obj:=&Object{ClassName:string(cp[classindex].info[2:])}
                s.Push(obj)
                pc = pc + 2
            case DUP:
                top := s.Pop()
                s.Push(top)
                s.Push(top)
            case RETURN:
                //fmt.Println(s.locals)
                return
        }
    }
}