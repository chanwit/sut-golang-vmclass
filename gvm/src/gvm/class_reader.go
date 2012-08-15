package gvm

import "os"
import "encoding/binary"
import "fmt"

type ClassReader struct {
    file *os.File
    bo   binary.ByteOrder
    cf   *ClassFile
}

func NewClassReader(f *os.File, cf *ClassFile) *ClassReader {
    return &ClassReader{f, binary.BigEndian, cf}
}

func (d *ClassReader) ReadMagic() {
    binary.Read(d.file, d.bo, &(d.cf.magic))
    _debugf("  magic : %x\n", d.cf.magic)
}

func (d *ClassReader) ReadVersion() {
    binary.Read(d.file, d.bo, &(d.cf.minor_version))
    binary.Read(d.file, d.bo, &(d.cf.major_version))
    _debugf("  minor version: %d\n", d.cf.minor_version)
    _debugf("  major version: %d\n", d.cf.major_version)
}

func (d *ClassReader) ReadConstantPool() {
    binary.Read(d.file, d.bo, &(d.cf.constant_pool_count))
    _debugf("Constant pool(%d):\n", d.cf.constant_pool_count)
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
            _debugf("  #%d = %s\n", i, info[2:])
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

func (d *ClassReader) ReadFlag() {
    binary.Read(d.file, d.bo, &(d.cf.access_flags))
    _debug("  flags:")
    accessFlags := d.cf.access_flags
    if accessFlags & ACC_PUBLIC == ACC_PUBLIC {
        _debug(" ACC_PUBLIC,")
    }
    if accessFlags & ACC_PRIVATE == ACC_PRIVATE {
        _debug(" ACC_PRIVATE,")
    }
    if accessFlags & ACC_PROTECTED == ACC_PROTECTED {
        _debug(" ACC_PROTECTED,")
    }
    if accessFlags & ACC_STATIC == ACC_STATIC {
        _debug(" ACC_STATIC,")
    }
    if accessFlags & ACC_FINAL == ACC_FINAL {
        _debug(" ACC_FINAL,")
    }
    if accessFlags & ACC_SUPER == ACC_SUPER {
        _debug(" ACC_SUPER,")
    }
    if accessFlags & ACC_VOLATILE == ACC_VOLATILE {
        _debug(" ACC_VOLATILE,")
    }
    if accessFlags & ACC_TRANSIENT == ACC_TRANSIENT {
        _debug(" ACC_TRANSIENT,")
    }
    if accessFlags & ACC_INTERFACE == ACC_INTERFACE {
        _debug(" ACC_INTERFACE,")
    }
    if accessFlags & ACC_ABSTRACT == ACC_ABSTRACT {
        _debug(" ACC_ABSTRACT,")
    }
    if accessFlags & ACC_SYNTHETIC == ACC_SYNTHETIC {
        _debug(" ACC_SYNTHETIC,")
    }
    if accessFlags & ACC_ENUM == ACC_ENUM {
        _debug(" ACC_ENUM,")
    }
    _debug("\b \n")
}

func (d *ClassReader) ReadClass() {
    binary.Read(d.file, d.bo, &(d.cf.this_class ))
    binary.Read(d.file, d.bo, &(d.cf.super_class))
    _debug("Class:")
    thisc  := d.cf.constant_pool[d.cf.this_class ]
    superc := d.cf.constant_pool[d.cf.super_class]
    _debug("  this class:", string(d.cf.constant_pool[(d.bo.Uint16(thisc.info))].info[2:]))
    _debug("  super class:", string(d.cf.constant_pool[(d.bo.Uint16(superc.info))].info[2:]))
}

func (d *ClassReader) ReadInterface() {
    binary.Read(d.file, d.bo, &(d.cf.interfaces_count))
    interfaceCount := d.cf.interfaces_count
    _debugf("Interface(%d):\n", interfaceCount)
    d.cf.interfaces = make([]uint16, interfaceCount)
    for i := uint16(0); i < interfaceCount; i++ {
        binary.Read(d.file, d.bo, &(d.cf.interfaces[i]))
        inter := d.cf.constant_pool[d.cf.interfaces[i]]
        _debug(" ", string(d.cf.constant_pool[(d.bo.Uint16(inter.info))].info[2:]))
    }
}

func (d *ClassReader) ReadField() {
    binary.Read(d.file, d.bo, &(d.cf.fields_count))
    _debugf("Field(%d):\n", d.cf.fields_count)
    d.cf.fields = make([]field_info, d.cf.fields_count)
    for i := uint16(0); i < d.cf.fields_count; i++ {
        var fi field_info
        binary.Read(d.file, d.bo, &fi.access_flags)
        binary.Read(d.file, d.bo, &fi.name_index)
        binary.Read(d.file, d.bo, &fi.descriptor_index)
        binary.Read(d.file, d.bo, &fi.attributes_count)
        d.cf.fields[i] = field_info{access_flags: fi.access_flags, name_index: fi.name_index, descriptor_index: fi.descriptor_index, attributes_count: fi.attributes_count}
        ni := d.cf.constant_pool[fi.name_index]
        _debug(" ", string(ni.info[2:]))

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

func (d *ClassReader) ReadMethod() {
    binary.Read(d.file, d.bo, &(d.cf.method_count))
    _debugf("Method(%d):\n", d.cf.method_count)
    d.cf.methods = make([]method_info, d.cf.method_count)
    for i := uint16(0); i < d.cf.method_count; i++ {
        var mi method_info
        binary.Read(d.file, d.bo, &mi.access_flags)
        binary.Read(d.file, d.bo, &mi.name_index)
        binary.Read(d.file, d.bo, &mi.descriptor_index)
        binary.Read(d.file, d.bo, &mi.attributes_count)
        d.cf.methods[i] = method_info{access_flags: mi.access_flags, name_index: mi.name_index, descriptor_index: mi.descriptor_index, attributes_count: mi.attributes_count}
        ni := d.cf.constant_pool[mi.name_index]
        _debug(" ", string(ni.info[2:]))

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
                    _debugf("      %d: ", k)
                    switch ca.code[k] {
                    case NOP:
                        _debug("nop")
                    case ACONST_NULL:
                        _debug("aconst_null")
                    case ICONST_M1:
                        _debug("iconst_m1")
                    case ICONST_0:
                        _debug("iconst_0")
                    case ICONST_1:
                        _debug("iconst_1")
                    case ICONST_2:
                        _debug("iconst_2")
                    case ICONST_3:
                        _debug("iconst_3")
                    case ICONST_4:
                        _debug("iconst_4")
                    case ICONST_5:
                        _debug("iconst_5")
                    case LCONST_0:
                        _debug("lconst_0")
                    case LCONST_1:
                        _debug("lconst_1")
                    case FCONST_0:
                        _debug("fconst_0")
                    case FCONST_1:
                        _debug("fconst_1")
                    case FCONST_2:
                        _debug("fconst_2")
                    case DCONST_0:
                        _debug("dconst_0")
                    case DCONST_1:
                        _debug("dconst_1")
                    case BIPUSH:
                        _debug("bipush")
                        k = k + 1
                    case SIPUSH:
                        _debug("sipush")
                        k = k + 2
                    case LDC:
                        _debug("ldc")
                        k = k + 1
                    case LDC_W:
                        _debug("ldc_w")
                        k = k + 2
                    case LDC2_W:
                        _debug("ldc2_w")
                        k = k + 2
                    case ILOAD:
                        _debug("iload")
                        k = k + 1
                    case LLOAD:
                        _debug("lload")
                        k = k + 1
                    case FLOAD:
                        _debug("fload")
                        k = k + 1
                    case DLOAD:
                        _debug("dload")
                        k = k + 1
                    case ALOAD:
                        _debug("aload")
                        k = k + 1
                    case ILOAD_0:
                        _debug("iload_0")
                    case ILOAD_1:
                        _debug("iload_1")
                    case ILOAD_2:
                        _debug("iload_2")
                    case ILOAD_3:
                        _debug("iload_3")
                    case LLOAD_0:
                        _debug("lload_0")
                    case LLOAD_1:
                        _debug("lload_1")
                    case LLOAD_2:
                        _debug("lload_2")
                    case LLOAD_3:
                        _debug("lload_3")
                    case FLOAD_0:
                        _debug("fload_0")
                    case FLOAD_1:
                        _debug("fload_1")
                    case FLOAD_2:
                        _debug("fload_2")
                    case FLOAD_3:
                        _debug("fload_3")
                    case DLOAD_0:
                        _debug("dload_0")
                    case DLOAD_1:
                        _debug("dload_1")
                    case DLOAD_2:
                        _debug("dload_2")
                    case DLOAD_3:
                        _debug("dload_3")
                    case ALOAD_0:
                        _debug("aload_0")
                    case ALOAD_1:
                        _debug("aload_1")
                    case ALOAD_2:
                        _debug("aload_2")
                    case ALOAD_3:
                        _debug("aload_3")
                    case IALOAD:
                        _debug("iaload")
                    case LALOAD:
                        _debug("laload")
                    case FALOAD:
                        _debug("faload")
                    case DALOAD:
                        _debug("daload")
                    case AALOAD:
                        _debug("aaload")
                    case BALOAD:
                        _debug("baload")
                    case CALOAD:
                        _debug("caload")
                    case SALOAD:
                        _debug("saload")
                    case ISTORE:
                        _debug("istore")
                        k = k + 1
                    case LSTORE:
                        _debug("lstore")
                        k = k + 1
                    case FSTORE:
                        _debug("fstore")
                        k = k + 1
                    case DSTORE:
                        _debug("dstore")
                        k = k + 1
                    case ASTORE:
                        _debug("astore")
                        k = k + 1
                    case ISTORE_0:
                        _debug("istore_0")
                    case ISTORE_1:
                        _debug("istore_1")
                    case ISTORE_2:
                        _debug("istore_2")
                    case ISTORE_3:
                        _debug("istore_3")
                    case LSTORE_0:
                        _debug("lstore_0")
                    case LSTORE_1:
                        _debug("lstore_1")
                    case LSTORE_2:
                        _debug("lstore_2")
                    case LSTORE_3:
                        _debug("lstore_3")
                    case FSTORE_0:
                        _debug("fstore_0")
                    case FSTORE_1:
                        _debug("fstore_1")
                    case FSTORE_2:
                        _debug("fstore_2")
                    case FSTORE_3:
                        _debug("fstore_3")
                    case DSTORE_0:
                        _debug("dstore_0")
                    case DSTORE_1:
                        _debug("dstore_1")
                    case DSTORE_2:
                        _debug("dstore_2")
                    case DSTORE_3:
                        _debug("dstore_3")
                    case ASTORE_0:
                        _debug("astore_0")
                    case ASTORE_1:
                        _debug("astore_1")
                    case ASTORE_2:
                        _debug("astore_2")
                    case ASTORE_3:
                        _debug("astore_3")
                    case IASTORE:
                        _debug("iastore")
                    case LASTORE:
                        _debug("lastore")
                    case FASTORE:
                        _debug("fastore")
                    case DASTORE:
                        _debug("dastore")
                    case AASTORE:
                        _debug("aastore")
                    case BASTORE:
                        _debug("bastore")
                    case CASTORE:
                        _debug("castore")
                    case SASTORE:
                        _debug("sastore")
                    case POP:
                        _debug("pop")
                    case POP2:
                        _debug("pop2")
                    case DUP:
                        _debug("dup")
                    case DUP_X1:
                        _debug("dup_x1")
                    case DUP_X2:
                        _debug("dup_x2")
                    case DUP2:
                        _debug("dup2")
                    case DUP2_X1:
                        _debug("dup2_x1")
                    case DUP2_X2:
                        _debug("dup2_x2")
                    case SWAP:
                        _debug("swap")
                    case IADD:
                        _debug("iadd")
                    case LADD:
                        _debug("ladd")
                    case FADD:
                        _debug("fadd")
                    case DADD:
                        _debug("dadd")
                    case ISUB:
                        _debug("isub")
                    case LSUB:
                        _debug("lsub")
                    case FSUB:
                        _debug("fsub")
                    case DSUB:
                        _debug("dsub")
                    case IMUL:
                        _debug("imul")
                    case LMUL:
                        _debug("lmul")
                    case FMUL:
                        _debug("fmul")
                    case DMUL:
                        _debug("dmul")
                    case IDIV:
                        _debug("idiv")
                    case LDIV:
                        _debug("ldiv")
                    case FDIV:
                        _debug("fdiv")
                    case DDIV:
                        _debug("ddiv")
                    case IREM:
                        _debug("irem")
                    case LREM:
                        _debug("lrem")
                    case FREM:
                        _debug("frem")
                    case DREM:
                        _debug("drem")
                    case INEG:
                        _debug("ineg")
                    case LNEG:
                        _debug("lneg")
                    case FNEG:
                        _debug("fneg")
                    case DNEG:
                        _debug("dneg")
                    case ISHL:
                        _debug("ishl")
                    case LSHL:
                        _debug("lshl")
                    case ISHR:
                        _debug("ishr")
                    case LSHR:
                        _debug("lshr")
                    case IUSHR:
                        _debug("iushr")
                    case LUSHR:
                        _debug("lushr")
                    case IAND:
                        _debug("iand")
                    case LAND:
                        _debug("land")
                    case IOR:
                        _debug("ior")
                    case LOR:
                        _debug("lor")
                    case IXOR:
                        _debug("ixor")
                    case LXOR:
                        _debug("lxor")
                    case IINC:
                        _debug("iinc")
                        k = k + 2
                    case I2L:
                        _debug("i2l")
                    case I2F:
                        _debug("i2f")
                    case I2D:
                        _debug("i2d")
                    case L2I:
                        _debug("l2i")
                    case L2F:
                        _debug("l2f")
                    case L2D:
                        _debug("l2d")
                    case F2I:
                        _debug("f2i")
                    case F2L:
                        _debug("f2l")
                    case F2D:
                        _debug("f2d")
                    case D2I:
                        _debug("d2i")
                    case D2L:
                        _debug("d2l")
                    case D2F:
                        _debug("d2f")
                    case I2B:
                        _debug("i2b")
                    case I2C:
                        _debug("i2c")
                    case I2S:
                        _debug("i2s")
                    case LCMP:
                        _debug("lcmp")
                    case FCMPL:
                        _debug("fcmpl")
                    case FCMPG:
                        _debug("fcmpg")
                    case DCMPL:
                        _debug("dcmpl")
                    case DCMPG:
                        _debug("dcmpg")
                    case IFEQ:
                        _debug("ifeq")
                        k = k + 2
                    case IFNE:
                        _debug("ifne")
                        k = k + 2
                    case IFLT:
                        _debug("iflt")
                        k = k + 2
                    case IFGE:
                        _debug("ifge")
                        k = k + 2
                    case IFGT:
                        _debug("ifgt")
                        k = k + 2
                    case IFLE:
                        _debug("ifle")
                        k = k + 2
                    case IF_ICMPEQ:
                        _debug("if_icmpeq")
                        k = k + 2
                    case IF_ICMPNE:
                        _debug("if_icmpne")
                        k = k + 2
                    case IF_ICMPLT:
                        _debug("if_icmplt")
                        k = k + 2
                    case IF_ICMPGE:
                        _debug("if_icmpge")
                        k = k + 2
                    case IF_ICMPGT:
                        _debug("if_icmpgt")
                        k = k + 2
                    case IF_ICMPLE:
                        _debug("if_icmple")
                        k = k + 2
                    case IF_ACMPEQ:
                        _debug("if_acmpeq")
                        k = k + 2
                    case IF_ACMPNE:
                        _debug("if_acmpne")
                        k = k + 2
                    case GOTO:
                        _debug("goto")
                        k = k + 2
                    case JSR:
                        _debug("jsr")
                        k = k + 2
                    case RET:
                        _debug("ret")
                        k = k + 1
                    case TABLESWITCH:
                        _debug("tableswitch")
                        //k = k+???
                    case LOOKUPSWITCH:
                        _debug("lookupswitch")
                        //k = k+???
                    case IRETURN:
                        _debug("ireturn")
                    case LRETURN:
                        _debug("lreturn")
                    case FRETURN:
                        _debug("freturn")
                    case DRETURN:
                        _debug("dreturn")
                    case ARETURN:
                        _debug("areturn")
                    case RETURN:
                        _debug("return")
                    case GETSTATIC:
                        _debug("getstatic")
                        k = k + 2
                    case PUTSTATIC:
                        _debug("putstatic")
                        k = k + 2
                    case GETFIELD:
                        _debug("getfield")
                        k = k + 2
                    case PUTFIELD:
                        _debug("putfield")
                        k = k + 2
                    case INVOKEVIRTUAL:
                        _debug("invokevirtual")
                        k = k + 2
                    case INVOKESPECIAL:
                        _debug("invokespecial")
                        k = k + 2
                    case INVOKESTATIC:
                        _debug("invokestatic")
                        k = k + 2
                    case INVOKEINTERFACE:
                        _debug("invokeinterface")
                        k = k + 4
                    case INVOKEDYNAMIC:
                        _debug("invokedynamic")
                        k = k + 4
                    case NEW:
                        _debug("new")
                        k = k + 2
                    case NEWARRAY:
                        _debug("newarray")
                        k = k + 1
                    case ANEWARRAY:
                        _debug("anewarray")
                        k = k + 2
                    case ARRAYLENGTH:
                        _debug("arraylength")
                    case ATHROW:
                        _debug("athrow")
                    case CHECKCAST:
                        _debug("checkcast")
                        k = k + 2
                    case INSTANCEOF:
                        _debug("instanceof")
                        k = k + 2
                    case MONITORENTER:
                        _debug("monitorenter")
                    case MONITOREXIT:
                        _debug("monitorexit")
                    case WIDE:
                        _debug("wide")
                        //k = k+???
                    case MULTIANEWARRAY:
                        _debug("multianewarray")
                        k = k + 3
                    case IFNULL:
                        _debug("ifnull")
                        k = k + 2
                    case IFNONNULL:
                        _debug("ifnonnull")
                        k = k + 2
                    case GOTO_W:
                        _debug("goto_w")
                        k = k + 4
                    case JSR_W:
                        _debug("jsr_w")
                        k = k + 4
                    case BREAKPOINT:
                        _debug("breakpoint")
                    case IMPDEP1:
                        _debug("impdep1")
                    case IMPDEP2:
                        _debug("impdep2")
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
                    _debug(start_pc, end_pc, handler_pc, catch_type)
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
                    _debug("   ", string(d.cf.constant_pool[name_index].info[2:]), ":")
                    d.cf.methods[i].attributes[j].line_number_table_att[m].line_number_tables = make([]line_number_table, lnt_a.line_number_table_length)
                    for o := uint16(0); o < lnt_a.line_number_table_length; o++ {
                        var start_pc uint16
                        var line_number uint16
                        start_pc = d.bo.Uint16(info[index+20+(o*4) : index+22+(o*4)])
                        line_number = d.bo.Uint16(info[index+22+(o*4) : index+24+(o*4)])

                        d.cf.methods[i].attributes[j].line_number_table_att[m].line_number_tables[o] = line_number_table{start_pc: start_pc, line_number: line_number}

                        _debug("    line", line_number, ":", start_pc)
                    }

                    d.cf.methods[i].attributes[j].line_number_table_att[m] = LineNumberTable_attribute{attribute_name_index: lnt_a.attribute_name_index, attribute_length: lnt_a.attribute_length, line_number_table_length: lnt_a.line_number_table_length, line_number_tables: d.cf.methods[i].attributes[j].line_number_table_att[m].line_number_tables}

                }

                d.cf.methods[i].attributes[j] = code_attribute{attribute_name_index: name_index, attribute_length: length, max_stack: ca.max_stack, max_locals: ca.max_locals, code_length: ca.code_length, code: d.cf.methods[i].attributes[j].code, exception_table_length: ca.exception_table_length, exception: d.cf.methods[i].attributes[j].exception, attributes_count: ca.attributes_count, line_number_table_att: d.cf.methods[i].attributes[j].line_number_table_att}

            }
        }
    }
}

func (cr *ClassReader) ReadAttribute() {
    binary.Read(cr.file, cr.bo, &(cr.cf.attributes_count))
    _debugf("attribute count : %d\n", cr.cf.attributes_count)
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

        att := cr.cf.constant_pool[name_index]
        fmt.Println(att, string(att.info[2:]))
    }
}
