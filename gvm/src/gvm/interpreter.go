package gvm

import "encoding/binary"

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
                index := binary.BigEndian.Uint16(cp[code[pc]].info)
                tag := cp[index].tag
                switch tag {
                    case CONSTANT_Utf8:
                        s.Push(string(cp[index].info[2:]))
                    case CONSTANT_Integer:
                        value := binary.BigEndian.Uint32(cp[index].info)
                        s.Push(int(value))
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
                if cp[value].tag == CONSTANT_Fieldref {
                    _debug("CONSTANT_Fieldref : ")
                    _debug("fieldref=", cp[value].info)
                    ownerIndex := binary.BigEndian.Uint16(cp[value].info[:2])
                    nameAndTypeIndex := binary.BigEndian.Uint16(cp[value].info[2:])
                    ownerClassIndex  := binary.BigEndian.Uint16(cp[ownerIndex].info[:2])
                    nameIndex := binary.BigEndian.Uint16(cp[nameAndTypeIndex].info[:2])
                    typeIndex := binary.BigEndian.Uint16(cp[nameAndTypeIndex].info[2:])

                    ownerName := string(cp[ownerClassIndex].info[2:])
                    fieldName := string(cp[nameIndex].info[2:])
                    fieldTypeName := string(cp[typeIndex].info[2:])

                    obj := CT(ownerName).staticFields[fieldName]
                    s.Push(obj)

                    // _info(ownerName)
                    // _info(fieldName)
                    _debug(fieldTypeName)
                    // _info(cp[nameAndTypeIndex].info)
                }
                pc = pc + 2

            case INVOKEVIRTUAL:
                _info("INVOKEVIRTUAL")
                methodRefIndex := binary.BigEndian.Uint16(code[pc:pc+2])
                ownerIndex := binary.BigEndian.Uint16(cp[methodRefIndex].info[:2])
                nameAndTypeIndex := binary.BigEndian.Uint16(cp[methodRefIndex].info[2:])

                ownerClassIndex := binary.BigEndian.Uint16(cp[ownerIndex].info)

                _info(cp[nameAndTypeIndex].info)

                nameIndex := binary.BigEndian.Uint16(cp[nameAndTypeIndex].info[:2])
                typeIndex := binary.BigEndian.Uint16(cp[nameAndTypeIndex].info[2:])

                owner := string(cp[ownerClassIndex].info[2:])
                _info(owner)
                _info(string(cp[nameIndex].info[2:]))
                _info(string(cp[typeIndex].info[2:]))
                switch value := s.Pop().(type) {
                    case string:
                        _info(value)
                }
                pc = pc + 2

            case RETURN:
                _info(s.locals)
                return
        }
    }
}