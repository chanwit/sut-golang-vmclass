package gvm

import "os"
import "encoding/binary"
//import "fmt"

type Decoder struct {
    file *os.File
    bo   binary.ByteOrder
    cf   *ClassFile
}
//To create Decoder struct
func ClassDecoder(f *os.File, cf *ClassFile) *Decoder {
    return &Decoder{f, binary.BigEndian, cf}
}

func (d *Decoder) ReadMagic() {
    binary.Read(d.file, d.bo, &(d.cf.magic))
    //fmt.Printf("  magic : %x\n", d.cf.magic)
}

func (d *Decoder) ReadVersion() {
    binary.Read(d.file, d.bo, &(d.cf.minor_version))
    binary.Read(d.file, d.bo, &(d.cf.major_version))
    //fmt.Printf("  minor version: %d\n", d.cf.minor_version)
    //fmt.Printf("  major version: %d\n", d.cf.major_version)
}

func (d *Decoder) ReadConstantPool() {
    binary.Read(d.file, d.bo, &(d.cf.constant_pool_count))
    //fmt.Printf("Constant pool(%d):\n", d.cf.constant_pool_count)
    d.cf.constant_pool = make([]cp_info, d.cf.constant_pool_count)
    for i := uint16(1); i < d.cf.constant_pool_count; i++ {
        var tag uint8
        binary.Read(d.file, d.bo, &(tag))
        switch tag {
        case CONSTANT_Class:
            info := make([]byte, 2)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_Fieldref:
            info := make([]byte, 4)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_Methodref:
            info := make([]byte, 4)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_InterfaceMethodref:
            info := make([]byte, 4)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_String:
            info := make([]byte, 2)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_Integer:
            info := make([]byte, 4)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_Float:
            info := make([]byte, 4)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_Long:
            info := make([]byte, 8)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_Double:
            info := make([]byte, 8)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_NameAndType:
            info := make([]byte, 4)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_Utf8:
            length := uint16(0)
            binary.Read(d.file, d.bo, &(length))
            info := make([]byte, 2+length)
            binary.BigEndian.PutUint16(info[0:2], length)
            binary.Read(d.file, d.bo, info[2:])
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
            //fmt.Printf("  #%d = %s\n", i, info[2:])
        case CONSTANT_MethodHandle:
            info := make([]byte, 3)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_MethodType:
            info := make([]byte, 2)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        case CONSTANT_InvokeDynamic:
            info := make([]byte, 4)
            binary.Read(d.file, d.bo, info)
            d.cf.constant_pool[i] = cp_info{tag: tag, info: info}
        }
    }
}

func (d *Decoder) ReadFlag() {
    binary.Read(d.file, d.bo, &(d.cf.access_flags))
    //fmt.Println("  flags:")
    accessFlags := d.cf.access_flags
    if accessFlags & ACC_PUBLIC == ACC_PUBLIC {
        //fmt.Println(" ACC_PUBLIC,")
    }
    if accessFlags & ACC_PRIVATE == ACC_PRIVATE {
        //fmt.Println(" ACC_PRIVATE,")
    }
    if accessFlags & ACC_PROTECTED == ACC_PROTECTED {
        //fmt.Println(" ACC_PROTECTED,")
    }
    if accessFlags & ACC_STATIC == ACC_STATIC {
        //fmt.Println(" ACC_STATIC,")
    }
    if accessFlags & ACC_FINAL == ACC_FINAL {
        //fmt.Println(" ACC_FINAL,")
    }
    if accessFlags & ACC_SUPER == ACC_SUPER {
        //fmt.Println(" ACC_SUPER,")
    }
    if accessFlags & ACC_VOLATILE == ACC_VOLATILE {
        //fmt.Println(" ACC_VOLATILE,")
    }
    if accessFlags & ACC_TRANSIENT == ACC_TRANSIENT {
        //fmt.Println(" ACC_TRANSIENT,")
    }
    if accessFlags & ACC_INTERFACE == ACC_INTERFACE {
        //fmt.Println(" ACC_INTERFACE,")
    }
    if accessFlags & ACC_ABSTRACT == ACC_ABSTRACT {
        //fmt.Println(" ACC_ABSTRACT,")
    }
    if accessFlags & ACC_SYNTHETIC == ACC_SYNTHETIC {
        //fmt.Println(" ACC_SYNTHETIC,")
    }
    if accessFlags & ACC_ENUM == ACC_ENUM {
        //fmt.Println(" ACC_ENUM,")
    }
    //fmt.Println("\b \n")
}

func (d *Decoder) ReadClass() {
    binary.Read(d.file, d.bo, &(d.cf.this_class ))
    binary.Read(d.file, d.bo, &(d.cf.super_class))
    //fmt.Println("Class:")
    //thisc  := d.cf.constant_pool[d.cf.this_class ]
    //superc := d.cf.constant_pool[d.cf.super_class]
    //fmt.Println("  this class:", string(d.cf.constant_pool[(d.bo.Uint16(thisc.info))].info[2:]))
    //fmt.Println("  super class:", string(d.cf.constant_pool[(d.bo.Uint16(superc.info))].info[2:]))
}

func (d *Decoder) ReadInterface() {
    binary.Read(d.file, d.bo, &(d.cf.interfaces_count))
    interfaceCount := d.cf.interfaces_count
    //fmt.Printf("Interface(%d):\n", interfaceCount)
    d.cf.interfaces = make([]uint16, interfaceCount)
    for i := uint16(0); i < interfaceCount; i++ {
        binary.Read(d.file, d.bo, &(d.cf.interfaces[i]))
        //inter := d.cf.constant_pool[d.cf.interfaces[i]]
        //fmt.Println(" ", string(d.cf.constant_pool[(d.bo.Uint16(inter.info))].info[2:]))
    }
}

func (d *Decoder) ReadField() {
    binary.Read(d.file, d.bo, &(d.cf.fields_count))
    //fmt.Printf("Field(%d):\n", d.cf.fields_count)
    d.cf.fields = make([]field_info, d.cf.fields_count)
    for i := uint16(0); i < d.cf.fields_count; i++ {
        var fi field_info
        binary.Read(d.file, d.bo, &fi.access_flags)
        binary.Read(d.file, d.bo, &fi.name_index)
        binary.Read(d.file, d.bo, &fi.descriptor_index)
        binary.Read(d.file, d.bo, &fi.attributes_count)
        d.cf.fields[i] = field_info{access_flags: fi.access_flags, name_index: fi.name_index, descriptor_index: fi.descriptor_index, attributes_count: fi.attributes_count}
        //ni := d.cf.constant_pool[fi.name_index]
        //fmt.Println(" ", string(ni.info[2:]))

        fi.attributes = make([]attribute_info, fi.attributes_count)
        for j := uint16(0); j < fi.attributes_count; j++ {
            var name_index uint16
            var length uint32
            binary.Read(d.file, d.bo, &name_index)
            binary.Read(d.file, d.bo, &length)
            info := make([]uint8, length)
            binary.Read(d.file, d.bo, &info)
            d.cf.fields[i].attributes[j] = attribute_info{attribute_name_index: name_index, attribute_length: length}
        }
    }
}

func (d *Decoder) ReadMethod() {
    binary.Read(d.file, d.bo, &(d.cf.method_count))
    //fmt.Printf("Method(%d):\n", d.cf.method_count)
    d.cf.methods = make([]method_info, d.cf.method_count)
    for i := uint16(0); i < d.cf.method_count; i++ {
        var mi method_info
        binary.Read(d.file, d.bo, &mi.access_flags)
        binary.Read(d.file, d.bo, &mi.name_index)
        binary.Read(d.file, d.bo, &mi.descriptor_index)
        binary.Read(d.file, d.bo, &mi.attributes_count)
        d.cf.methods[i] = method_info{access_flags: mi.access_flags, name_index: mi.name_index, descriptor_index: mi.descriptor_index, attributes_count: mi.attributes_count}
        //ni := d.cf.constant_pool[mi.name_index]
        //fmt.Println(" ", string(ni.info[2:]))

        d.cf.methods[i].attributes = make([]code_attribute, mi.attributes_count)
        for j := uint16(0); j < mi.attributes_count; j++ {
            var name_index uint16
            var length uint32
            binary.Read(d.file, d.bo, &name_index)
            binary.Read(d.file, d.bo, &length)
            info := make([]uint8, length)
            binary.Read(d.file, d.bo, &info)

            lookup := string(d.cf.constant_pool[name_index].info[2:])
            if lookup == "Code" {
                var ca code_attribute
                ca.attribute_name_index = name_index
                ca.attribute_length = length
                ca.max_stack = d.bo.Uint16(info[0:2])
                ca.max_locals = d.bo.Uint16(info[2:4])
                ca.code_length = d.bo.Uint32(info[4:8])
                ca.code = info[8 : 8+ca.code_length]
                d.cf.methods[i].attributes[j].code = make([]uint8, ca.code_length)
                d.cf.methods[i].attributes[j].code = ca.code
                for k := uint32(0); k < ca.code_length; k++ {
                    //fmt.Printf("      %d: ", k)
                    switch ca.code[k] {
                    case NOP:
                        //fmt.Println("nop")
                    case ACONST_NULL:
                        //fmt.Println("aconst_null")
                    case ICONST_M1:
                        //fmt.Println("iconst_m1")
                    case ICONST_0:
                        //fmt.Println("iconst_0")
                    case ICONST_1:
                        //fmt.Println("iconst_1")
                    case ICONST_2:
                        //fmt.Println("iconst_2")
                    case ICONST_3:
                        //fmt.Println("iconst_3")
                    case ICONST_4:
                        //fmt.Println("iconst_4")
                    case ICONST_5:
                        //fmt.Println("iconst_5")
                    case LCONST_0:
                        //fmt.Println("lconst_0")
                    case LCONST_1:
                        //fmt.Println("lconst_1")
                    case FCONST_0:
                        //fmt.Println("fconst_0")
                    case FCONST_1:
                        //fmt.Println("fconst_1")
                    case FCONST_2:
                        //fmt.Println("fconst_2")
                    case DCONST_0:
                        //fmt.Println("dconst_0")
                    case DCONST_1:
                        //fmt.Println("dconst_1")
                    case BIPUSH:
                        //fmt.Println("bipush")
                        k = k + 1
                    case SIPUSH:
                        //fmt.Println("sipush")
                        k = k + 2
                    case LDC:
                        //fmt.Println("ldc")
                        k = k + 1
                    case LDC_W:
                        //fmt.Println("ldc_w")
                        k = k + 2
                    case LDC2_W:
                        //fmt.Println("ldc2_w")
                        k = k + 2
                    case ILOAD:
                        //fmt.Println("iload")
                        k = k + 1
                    case LLOAD:
                        //fmt.Println("lload")
                        k = k + 1
                    case FLOAD:
                        //fmt.Println("fload")
                        k = k + 1
                    case DLOAD:
                        //fmt.Println("dload")
                        k = k + 1
                    case ALOAD:
                        //fmt.Println("aload")
                        k = k + 1
                    case ILOAD_0:
                        //fmt.Println("iload_0")
                    case ILOAD_1:
                        //fmt.Println("iload_1")
                    case ILOAD_2:
                        //fmt.Println("iload_2")
                    case ILOAD_3:
                        //fmt.Println("iload_3")
                    case LLOAD_0:
                        //fmt.Println("lload_0")
                    case LLOAD_1:
                        //fmt.Println("lload_1")
                    case LLOAD_2:
                        //fmt.Println("lload_2")
                    case LLOAD_3:
                        //fmt.Println("lload_3")
                    case FLOAD_0:
                        //fmt.Println("fload_0")
                    case FLOAD_1:
                        //fmt.Println("fload_1")
                    case FLOAD_2:
                        //fmt.Println("fload_2")
                    case FLOAD_3:
                        //fmt.Println("fload_3")
                    case DLOAD_0:
                        //fmt.Println("dload_0")
                    case DLOAD_1:
                        //fmt.Println("dload_1")
                    case DLOAD_2:
                        //fmt.Println("dload_2")
                    case DLOAD_3:
                        //fmt.Println("dload_3")
                    case ALOAD_0:
                        //fmt.Println("aload_0")
                    case ALOAD_1:
                        //fmt.Println("aload_1")
                    case ALOAD_2:
                        //fmt.Println("aload_2")
                    case ALOAD_3:
                        //fmt.Println("aload_3")
                    case IALOAD:
                        //fmt.Println("iaload")
                    case LALOAD:
                        //fmt.Println("laload")
                    case FALOAD:
                        //fmt.Println("faload")
                    case DALOAD:
                        //fmt.Println("daload")
                    case AALOAD:
                        //fmt.Println("aaload")
                    case BALOAD:
                        //fmt.Println("baload")
                    case CALOAD:
                        //fmt.Println("caload")
                    case SALOAD:
                        //fmt.Println("saload")
                    case ISTORE:
                        //fmt.Println("istore")
                        k = k + 1
                    case LSTORE:
                        //fmt.Println("lstore")
                        k = k + 1
                    case FSTORE:
                        //fmt.Println("fstore")
                        k = k + 1
                    case DSTORE:
                        //fmt.Println("dstore")
                        k = k + 1
                    case ASTORE:
                        //fmt.Println("astore")
                        k = k + 1
                    case ISTORE_0:
                        //fmt.Println("istore_0")
                    case ISTORE_1:
                        //fmt.Println("istore_1")
                    case ISTORE_2:
                        //fmt.Println("istore_2")
                    case ISTORE_3:
                        //fmt.Println("istore_3")
                    case LSTORE_0:
                        //fmt.Println("lstore_0")
                    case LSTORE_1:
                        //fmt.Println("lstore_1")
                    case LSTORE_2:
                        //fmt.Println("lstore_2")
                    case LSTORE_3:
                        //fmt.Println("lstore_3")
                    case FSTORE_0:
                        //fmt.Println("fstore_0")
                    case FSTORE_1:
                        //fmt.Println("fstore_1")
                    case FSTORE_2:
                        //fmt.Println("fstore_2")
                    case FSTORE_3:
                        //fmt.Println("fstore_3")
                    case DSTORE_0:
                        //fmt.Println("dstore_0")
                    case DSTORE_1:
                        //fmt.Println("dstore_1")
                    case DSTORE_2:
                        //fmt.Println("dstore_2")
                    case DSTORE_3:
                        //fmt.Println("dstore_3")
                    case ASTORE_0:
                        //fmt.Println("astore_0")
                    case ASTORE_1:
                        //fmt.Println("astore_1")
                    case ASTORE_2:
                        //fmt.Println("astore_2")
                    case ASTORE_3:
                        //fmt.Println("astore_3")
                    case IASTORE:
                        //fmt.Println("iastore")
                    case LASTORE:
                        //fmt.Println("lastore")
                    case FASTORE:
                        //fmt.Println("fastore")
                    case DASTORE:
                        //fmt.Println("dastore")
                    case AASTORE:
                        //fmt.Println("aastore")
                    case BASTORE:
                        //fmt.Println("bastore")
                    case CASTORE:
                        //fmt.Println("castore")
                    case SASTORE:
                        //fmt.Println("sastore")
                    case POP:
                        //fmt.Println("pop")
                    case POP2:
                        //fmt.Println("pop2")
                    case DUP:
                        //fmt.Println("dup")
                    case DUP_X1:
                        //fmt.Println("dup_x1")
                    case DUP_X2:
                        //fmt.Println("dup_x2")
                    case DUP2:
                        //fmt.Println("dup2")
                    case DUP2_X1:
                        //fmt.Println("dup2_x1")
                    case DUP2_X2:
                        //fmt.Println("dup2_x2")
                    case SWAP:
                        //fmt.Println("swap")
                    case IADD:
                        //fmt.Println("iadd")
                    case LADD:
                        //fmt.Println("ladd")
                    case FADD:
                        //fmt.Println("fadd")
                    case DADD:
                        //fmt.Println("dadd")
                    case ISUB:
                        //fmt.Println("isub")
                    case LSUB:
                        //fmt.Println("lsub")
                    case FSUB:
                        //fmt.Println("fsub")
                    case DSUB:
                        //fmt.Println("dsub")
                    case IMUL:
                        //fmt.Println("imul")
                    case LMUL:
                        //fmt.Println("lmul")
                    case FMUL:
                        //fmt.Println("fmul")
                    case DMUL:
                        //fmt.Println("dmul")
                    case IDIV:
                        //fmt.Println("idiv")
                    case LDIV:
                        //fmt.Println("ldiv")
                    case FDIV:
                        //fmt.Println("fdiv")
                    case DDIV:
                        //fmt.Println("ddiv")
                    case IREM:
                        //fmt.Println("irem")
                    case LREM:
                        //fmt.Println("lrem")
                    case FREM:
                        //fmt.Println("frem")
                    case DREM:
                        //fmt.Println("drem")
                    case INEG:
                        //fmt.Println("ineg")
                    case LNEG:
                        //fmt.Println("lneg")
                    case FNEG:
                        //fmt.Println("fneg")
                    case DNEG:
                        //fmt.Println("dneg")
                    case ISHL:
                        //fmt.Println("ishl")
                    case LSHL:
                        //fmt.Println("lshl")
                    case ISHR:
                        //fmt.Println("ishr")
                    case LSHR:
                        //fmt.Println("lshr")
                    case IUSHR:
                        //fmt.Println("iushr")
                    case LUSHR:
                        //fmt.Println("lushr")
                    case IAND:
                        //fmt.Println("iand")
                    case LAND:
                        //fmt.Println("land")
                    case IOR:
                        //fmt.Println("ior")
                    case LOR:
                        //fmt.Println("lor")
                    case IXOR:
                        //fmt.Println("ixor")
                    case LXOR:
                        //fmt.Println("lxor")
                    case IINC:
                        //fmt.Println("iinc")
                        k = k + 2
                    case I2L:
                        //fmt.Println("i2l")
                    case I2F:
                        //fmt.Println("i2f")
                    case I2D:
                        //fmt.Println("i2d")
                    case L2I:
                        //fmt.Println("l2i")
                    case L2F:
                        //fmt.Println("l2f")
                    case L2D:
                        //fmt.Println("l2d")
                    case F2I:
                        //fmt.Println("f2i")
                    case F2L:
                        //fmt.Println("f2l")
                    case F2D:
                        //fmt.Println("f2d")
                    case D2I:
                        //fmt.Println("d2i")
                    case D2L:
                        //fmt.Println("d2l")
                    case D2F:
                        //fmt.Println("d2f")
                    case I2B:
                        //fmt.Println("i2b")
                    case I2C:
                        //fmt.Println("i2c")
                    case I2S:
                        //fmt.Println("i2s")
                    case LCMP:
                        //fmt.Println("lcmp")
                    case FCMPL:
                        //fmt.Println("fcmpl")
                    case FCMPG:
                        //fmt.Println("fcmpg")
                    case DCMPL:
                        //fmt.Println("dcmpl")
                    case DCMPG:
                        //fmt.Println("dcmpg")
                    case IFEQ:
                        //fmt.Println("ifeq")
                        k = k + 2
                    case IFNE:
                        //fmt.Println("ifne")
                        k = k + 2
                    case IFLT:
                        //fmt.Println("iflt")
                        k = k + 2
                    case IFGE:
                        //fmt.Println("ifge")
                        k = k + 2
                    case IFGT:
                        //fmt.Println("ifgt")
                        k = k + 2
                    case IFLE:
                        //fmt.Println("ifle")
                        k = k + 2
                    case IF_ICMPEQ:
                        //fmt.Println("if_icmpeq")
                        k = k + 2
                    case IF_ICMPNE:
                        //fmt.Println("if_icmpne")
                        k = k + 2
                    case IF_ICMPLT:
                        //fmt.Println("if_icmplt")
                        k = k + 2
                    case IF_ICMPGE:
                        //fmt.Println("if_icmpge")
                        k = k + 2
                    case IF_ICMPGT:
                        //fmt.Println("if_icmpgt")
                        k = k + 2
                    case IF_ICMPLE:
                        //fmt.Println("if_icmple")
                        k = k + 2
                    case IF_ACMPEQ:
                        //fmt.Println("if_acmpeq")
                        k = k + 2
                    case IF_ACMPNE:
                        //fmt.Println("if_acmpne")
                        k = k + 2
                    case GOTO:
                        //fmt.Println("goto")
                        k = k + 2
                    case JSR:
                        //fmt.Println("jsr")
                        k = k + 2
                    case RET:
                        //fmt.Println("ret")
                        k = k + 1
                    case TABLESWITCH:
                        //fmt.Println("tableswitch")
                        //k = k+???
                    case LOOKUPSWITCH:
                        //fmt.Println("lookupswitch")
                        //k = k+???
                    case IRETURN:
                        //fmt.Println("ireturn")
                    case LRETURN:
                        //fmt.Println("lreturn")
                    case FRETURN:
                        //fmt.Println("freturn")
                    case DRETURN:
                        //fmt.Println("dreturn")
                    case ARETURN:
                        //fmt.Println("areturn")
                    case RETURN:
                        //fmt.Println("return")
                    case GETSTATIC:
                        //fmt.Println("getstatic")
                        k = k + 2
                    case PUTSTATIC:
                        //fmt.Println("putstatic")
                        k = k + 2
                    case GETFIELD:
                        //fmt.Println("getfield")
                        k = k + 2
                    case PUTFIELD:
                        //fmt.Println("putfield")
                        k = k + 2
                    case INVOKEVIRTUAL:
                        //fmt.Println("invokevirtual")
                        k = k + 2
                    case INVOKESPECIAL:
                        //fmt.Println("invokespecial")
                        k = k + 2
                    case INVOKESTATIC:
                        //fmt.Println("invokestatic")
                        k = k + 2
                    case INVOKEINTERFACE:
                        //fmt.Println("invokeinterface")
                        k = k + 4
                    case INVOKEDYNAMIC:
                        //fmt.Println("invokedynamic")
                        k = k + 4
                    case NEW:
                        //fmt.Println("new")
                        k = k + 2
                    case NEWARRAY:
                        //fmt.Println("newarray")
                        k = k + 1
                    case ANEWARRAY:
                        //fmt.Println("anewarray")
                        k = k + 2
                    case ARRAYLENGTH:
                        //fmt.Println("arraylength")
                    case ATHROW:
                        //fmt.Println("athrow")
                    case CHECKCAST:
                        //fmt.Println("checkcast")
                        k = k + 2
                    case INSTANCEOF:
                        //fmt.Println("instanceof")
                        k = k + 2
                    case MONITORENTER:
                        //fmt.Println("monitorenter")
                    case MONITOREXIT:
                        //fmt.Println("monitorexit")
                    case WIDE:
                        //fmt.Println("wide")
                        //k = k+???
                    case MULTIANEWARRAY:
                        //fmt.Println("multianewarray")
                        k = k + 3
                    case IFNULL:
                        //fmt.Println("ifnull")
                        k = k + 2
                    case IFNONNULL:
                        //fmt.Println("ifnonnull")
                        k = k + 2
                    case GOTO_W:
                        //fmt.Println("goto_w")
                        k = k + 4
                    case JSR_W:
                        //fmt.Println("jsr_w")
                        k = k + 4
                    case BREAKPOINT:
                        //fmt.Println("breakpoint")
                    case IMPDEP1:
                        //fmt.Println("impdep1")
                    case IMPDEP2:
                        //fmt.Println("impdep2")
                    }
                }
                ca.exception_table_length = d.bo.Uint16(info[8+ca.code_length : 10+ca.code_length])
                d.cf.methods[i].attributes[j].exception = make([]exception_table, ca.exception_table_length)
                for l := uint16(0); l < ca.exception_table_length; l++ {
                    var start_pc uint16
                    var end_pc uint16
                    var handler_pc uint16
                    var catch_type uint16
                    start_pc = d.bo.Uint16(info[10+ca.code_length : 12+ca.code_length])
                    end_pc = d.bo.Uint16(info[12+ca.code_length : 14+ca.code_length])
                    handler_pc = d.bo.Uint16(info[14+ca.code_length : 16+ca.code_length])
                    catch_type = d.bo.Uint16(info[16+ca.code_length : 18+ca.code_length])
                    d.cf.methods[i].attributes[j].exception[l] = exception_table{start_pc: start_pc, end_pc: end_pc, handler_pc: handler_pc, catch_type: catch_type}
                    //fmt.Println(start_pc, end_pc, handler_pc, catch_type)
                }
                index := uint16(ca.code_length) + ca.exception_table_length
                ca.attributes_count = d.bo.Uint16(info[index+10 : index+12])
                d.cf.methods[i].attributes[j].line_number_table_att = make([]LineNumberTable_attribute, ca.attributes_count)
                var lnt_a LineNumberTable_attribute
                for m := uint16(0); m < ca.attributes_count; m++ {
                    var name_index uint16
                    var length uint32
                    name_index = d.bo.Uint16(info[index+12 : index+14])
                    length = d.bo.Uint32(info[index+14 : index+18])
                    lnt_a.attribute_name_index = name_index
                    lnt_a.attribute_length = length
                    lnt_a.line_number_table_length = d.bo.Uint16(info[index+18 : index+20])
                    lnt_a.line_number_tables = make([]line_number_table, lnt_a.line_number_table_length)
                    //fmt.Println("   ", string(d.cf.constant_pool[name_index].info[2:]), ":")
                    d.cf.methods[i].attributes[j].line_number_table_att[m].line_number_tables = make([]line_number_table, lnt_a.line_number_table_length)
                    for o := uint16(0); o < lnt_a.line_number_table_length; o++ {
                        var start_pc uint16
                        var line_number uint16
                        start_pc = d.bo.Uint16(info[index+20+(o*4) : index+22+(o*4)])
                        line_number = d.bo.Uint16(info[index+22+(o*4) : index+24+(o*4)])

                        d.cf.methods[i].attributes[j].line_number_table_att[m].line_number_tables[o] = line_number_table{start_pc: start_pc, line_number: line_number}

                        //fmt.Println("    line", line_number, ":", start_pc)
                    }

                    d.cf.methods[i].attributes[j].line_number_table_att[m] = LineNumberTable_attribute{attribute_name_index: lnt_a.attribute_name_index, attribute_length: lnt_a.attribute_length, line_number_table_length: lnt_a.line_number_table_length, line_number_tables: d.cf.methods[i].attributes[j].line_number_table_att[m].line_number_tables}

                }

                d.cf.methods[i].attributes[j] = code_attribute{attribute_name_index: name_index, attribute_length: length, max_stack: ca.max_stack, max_locals: ca.max_locals, code_length: ca.code_length, code: d.cf.methods[i].attributes[j].code, exception_table_length: ca.exception_table_length, exception: d.cf.methods[i].attributes[j].exception, attributes_count: ca.attributes_count, line_number_table_att: d.cf.methods[i].attributes[j].line_number_table_att}

            }
        }
    }
}

func (cr *Decoder) ReadAttribute() {
    binary.Read(cr.file, cr.bo, &(cr.cf.attributes_count))
    //fmt.Printf("attribute count : %d\n", cr.cf.attributes_count)
    cr.cf.attributes = make([]attribute_info, cr.cf.attributes_count)
    for i := uint16(0); i < cr.cf.attributes_count; i++ {
        var name_index uint16
        var length uint32
        binary.Read(cr.file, cr.bo, &name_index)
        binary.Read(cr.file, cr.bo, &length)
        info := make([]uint8, length)
        binary.Read(cr.file, cr.bo, &info)
        cr.cf.attributes[i] = attribute_info{attribute_name_index: name_index,
                                            attribute_length: length,
                                            info: info}

        //att := cr.cf.constant_pool[name_index]
        //fmt.Println(att, string(att.info[2:]))
    }
}
