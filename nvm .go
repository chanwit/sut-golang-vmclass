package main

import (
	"fmt"
	"os"
	"encoding/binary"
)

const (
		CONSTANT_Class				=	7
		CONSTANT_Fieldref			=	9
		CONSTANT_Methodref			=	10
		CONSTANT_InterfaceMethodref	=	11
		CONSTANT_String				=	8
		CONSTANT_Integer			=	3
		CONSTANT_Float				=	4
		CONSTANT_Long				=	5
		CONSTANT_Double				=	6
		CONSTANT_NameAndType		=	12
		CONSTANT_Utf8				=	1
		CONSTANT_MethodHandle		=	15
		CONSTANT_MethodType			=	16
		CONSTANT_InvokeDynamic		=	18
)

type classFile struct {
	magic					uint32
	minor_version			uint16
	major_version			uint16
	constant_pool_count		uint16
	constant_pool			[]cp_info
	access_flags			uint16
	this_class				uint16
	super_class				uint16
	interfaces_count		uint16
	interfaces 				[]uint16
	fields_count			uint16
	fields 					[]field_info
	method_count 			uint16
	methods 				[]method_info
	attributes_count		uint16
	attributes 				[]attribute_info
}

const (
    ACC_PUBLIC      = 0x0001
    ACC_PRIVATE 	= 0x0002
    ACC_PROTECTED 	= 0x0004
    ACC_STATIC 		= 0x0008
    ACC_FINAL       = 0x0010
    ACC_SUPER       = 0x0020
    ACC_VOLATILE 	= 0x0040
    ACC_TRANSIENT 	= 0x0080
    ACC_INTERFACE   = 0x0200
    ACC_ABSTRACT    = 0x0400
    ACC_SYNTHETIC 	= 0x1000
    ACC_ENUM 		= 0x4000
)

const (
	nop					= iota
	aconst_null
	iconst_m1
	iconst_0
	iconst_1
	iconst_2
	iconst_3
	iconst_4
	iconst_5
	lconst_0
	lconst_1
	fconst_0
	fconst_1
	fconst_2
	dconst_0
	dconst_1
	bipush
	sipush
	ldc
	ldc_w
	ldc2_w
	iload
	lload
	fload
	dload
	aload
	iload_0
	iload_1
	iload_2
	iload_3
	lload_0
	lload_1
	lload_2
	lload_3
	fload_0
	fload_1
	fload_2
	fload_3
	dload_0
	dload_1
	dload_2
	dload_3
	aload_0
	aload_1
	aload_2
	aload_3
	iaload
	laload
	faload
	daload
	aaload
	baload
	caload
	saload
	istore
	lstore
	fstore
	dstore
	astore
	istore_0
	istore_1
	istore_2
	istore_3
	lstore_0
	lstore_1
	lstore_2
	lstore_3
	fstore_0
	fstore_1
	fstore_2
	fstore_3
	dstore_0
	dstore_1
	dstore_2
	dstore_3
	astore_0
	astore_1
	astore_2
	astore_3
	iastore
	lastore
	fastore
	dastore
	aastore
	bastore
	castore
	sastore
	pop
	pop2
	dup
	dup_x1
	dup_x2
	dup2
	dup2_x1
	dup2_x2
	swap
	iadd
	ladd
	fadd
	dadd
	isub
	lsub
	fsub
	dsub
	imul
	lmul
	fmul
	dmul
	idiv
	ldiv
	fdiv
	ddiv
	irem
	lrem
	frem
	drem
	ineg
	lneg
	fneg
	dneg
	ishl
	lshl
	ishr
	lshr
	iushr
	lushr
	iand
	land
	ior
	lor
	ixor
	lxor
	iinc
	i2l
	i2f
	i2d
	l2i
	l2f
	l2d
	f2i
	f2l
	f2d
	d2i
	d2l
	d2f
	i2b
	i2c
	i2s
	lcmp
	fcmpl
	fcmpg
	dcmpl
	dcmpg
	ifeq
	ifne
	iflt
	ifge
	ifgt
	ifle
	if_icmpeq
	if_icmpne
	if_icmplt
	if_icmpge
	if_icmpgt
	if_icmple
	if_acmpeq
	if_acmpne
	goto_x
	jsr
	ret
	tableswitch
	lookupswitch
	ireturn
	lreturn
	freturn
	dreturn
	areturn
	return_x
	getstatic
	putstatic
	getfield
	putfield
	invokevirtual
	invokespecial
	invokestatic
	invokeinterface
	invokedynamic
	new
	newarray
	anewarray
	arraylength
	athrow
	checkcast
	instanceof
	monitorenter
	monitorexit
	wide
	multianewarray
	ifnull
	ifnonnull
	goto_w
	jsr_w
	breakpoint
	//(no name)			= 0xcb-fd
	impdep1				= 254
	impdep2				= 255
)

type cp_info struct {
	tag						uint8
	info					[]uint8
}

type CONSTANT_Class_info struct {
	tag						uint8
	name_index				uint16
}

type CONSTANT_Fieldref_info struct {
    tag						uint8
    class_index				uint16
    name_and_type_index		uint16
}

type CONSTANT_Methodref_info struct {
    tag						uint8
    class_index				uint16
    name_and_type_index		uint16
}

type CONSTANT_InterfaceMethodref_info struct {
    tag						uint8
    class_index				uint16
    name_and_type_index		uint16
}

type CONSTANT_String_info struct {
    tag						uint8
    string_index			uint16
}

type CONSTANT_Integer_info struct {
    tag						uint8
    bytes					uint32
}

type CONSTANT_Float_info struct {
    tag						uint8
    bytes					uint32
}

type CONSTANT_Long_info struct {
    tag						uint8
    high_bytes				uint32
    low_bytes				uint32
}

type CONSTANT_Double_info struct {
    tag						uint8
    high_bytes				uint32
    low_bytes				uint32
}

type CONSTANT_NameAndType_info struct {
    tag						uint8
    name_index 				uint16
    descriptor_index 		uint16
}

type CONSTANT_Utf8_info struct {
    tag						uint8
    length					uint16
    bytes					[]uint8
}

type CONSTANT_MethodHandle_info struct {
    tag						uint8
    reference_kind			uint8
    reference_index			uint16
}

type CONSTANT_MethodType_info struct {
    tag						uint8
    descriptor_index		uint16
}

type CONSTANT_InvokeDynamic_info struct {
    tag						uint8
    bootstrap_method_attr_index	uint16
    name_and_type_index		uint16
}

type field_info struct {
    access_flags			uint16
    name_index 				uint16
    descriptor_index 		uint16
    attributes_count 		uint16
    attributes 				[]attribute_info
}

type method_info struct {
	access_flags			uint16
	name_index 				uint16
	descriptor_index 		uint16
	attributes_count 		uint16
	attributes 				[]attribute_info
}

type attribute_info struct {
    attribute_name_index	uint16
    attribute_length		uint32
    info 					[]uint8
}

type code_attribute struct {
	attribute_name_index	uint16
	attribute_length		uint32
	max_stack				uint16
	max_locals				uint16
	code_length				uint32
	code 					[]uint8
	exception_table_length	uint16
	exception 				[]exception_table
	attributes_count		uint16
	attributes				[]attribute_info
}

type LineNumberTable_attribute struct {
	attribute_name_index	uint16
    attribute_length		uint32
    line_number_table_length	uint16
    line_number_tables 		[]line_number_table
}

type line_number_table struct {
	start_pc				uint16
	line_number 			uint16
}

type exception_table struct {
	start_pc				uint16
	end_pc					uint16
	handler_pc				uint16
	catch_type				uint16
}

type decoder struct {
	file 		*os.File
	bo 			binary.ByteOrder
	cf 			*classFile
}

func (d *decoder) readMagic() {
	binary.Read(d.file, d.bo, &(d.cf.magic))
	fmt.Printf("magic : %x\n", d.cf.magic)
}

func (d *decoder) readVersion() {
	binary.Read(d.file, d.bo, &(d.cf.minor_version))
	binary.Read(d.file, d.bo, &(d.cf.major_version))
	fmt.Printf("version : %d.%d\n", d.cf.major_version, d.cf.minor_version)
}

func (d *decoder) readConstantPool() {
	binary.Read(d.file, d.bo, &(d.cf.constant_pool_count))
	fmt.Printf("constant pool count : %d\n", d.cf.constant_pool_count)
	d.cf.constant_pool = make([]cp_info, d.cf.constant_pool_count)
	for i := uint16(1); i < d.cf.constant_pool_count; i++ {
		var tag uint8
		binary.Read(d.file, d.bo, &(tag))
		switch tag {
			case CONSTANT_Class :
				info := make([]byte, 2)
				binary.Read(d.file, d.bo, info)
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }
			case CONSTANT_Fieldref :
				info := make([]byte, 4)
				binary.Read(d.file, d.bo, info)
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }
			case CONSTANT_Methodref :
				info := make([]byte, 4)
				binary.Read(d.file, d.bo, info)
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }
			case CONSTANT_InterfaceMethodref :
				info := make([]byte, 4)
				binary.Read(d.file, d.bo, info)
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }
			case CONSTANT_String :
				info := make([]byte, 2)
				binary.Read(d.file, d.bo, info)
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }
			case CONSTANT_Integer :
				info := make([]byte, 4)
				binary.Read(d.file, d.bo, info)
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }
			case CONSTANT_Float :
				info := make([]byte, 4)
				binary.Read(d.file, d.bo, info)
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }
			case CONSTANT_Long :
				info := make([]byte, 8)
				binary.Read(d.file, d.bo, info)
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }
			case CONSTANT_Double :
				info := make([]byte, 8)
				binary.Read(d.file, d.bo, info)
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }
			case CONSTANT_NameAndType :
				info := make([]byte, 4)
				binary.Read(d.file, d.bo, info)
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }
			case CONSTANT_Utf8 :
				var length uint16
				binary.Read(d.file, d.bo, &(length))
				info := make([]byte, 2 + length)
				binary.BigEndian.PutUint16(info[0:2], length)
				binary.Read(d.file, d.bo, info[2:])
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }
				fmt.Printf("Line %d\t %s\n",i,info[2:])
			case CONSTANT_MethodHandle :
				info := make([]byte, 3)
				binary.Read(d.file, d.bo, info)
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }
			case CONSTANT_MethodType :
				info := make([]byte, 2)
				binary.Read(d.file, d.bo, info)
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }
			case CONSTANT_InvokeDynamic :
				info := make([]byte, 4)
				binary.Read(d.file, d.bo, info)
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }
		}
	}
}

func (d *decoder) readFlag() {
	binary.Read(d.file, d.bo, &(d.cf.access_flags))
	fmt.Println("++++++++++++++++++++ Flag")
	if d.cf.access_flags & ACC_PUBLIC == ACC_PUBLIC {
		fmt.Println("ACC_PUBLIC")
	}
	if d.cf.access_flags & ACC_PRIVATE == ACC_PRIVATE {
		fmt.Println("ACC_PRIVATE")
	}
	if d.cf.access_flags & ACC_PROTECTED == ACC_PROTECTED {
		fmt.Println("ACC_PROTECTED")
	}
	if d.cf.access_flags & ACC_STATIC == ACC_STATIC {
		fmt.Println("ACC_STATIC")
	}
	if d.cf.access_flags & ACC_FINAL == ACC_FINAL {
		fmt.Println("ACC_FINAL")
	}
	if d.cf.access_flags & ACC_SUPER == ACC_SUPER {
		fmt.Println("ACC_SUPER")
	}
	if d.cf.access_flags & ACC_VOLATILE == ACC_VOLATILE {
		fmt.Println("ACC_VOLATILE")
	}
	if d.cf.access_flags & ACC_TRANSIENT == ACC_TRANSIENT {
		fmt.Println("ACC_TRANSIENT")
	}
	if d.cf.access_flags & ACC_INTERFACE == ACC_INTERFACE {
		fmt.Println("ACC_INTERFACE")
	}
	if d.cf.access_flags & ACC_ABSTRACT == ACC_ABSTRACT {
		fmt.Println("ACC_ABSTRACT")
	}
	if d.cf.access_flags & ACC_SYNTHETIC == ACC_SYNTHETIC {
		fmt.Println("ACC_SYNTHETIC")
	}
	if d.cf.access_flags & ACC_ENUM == ACC_ENUM {
		fmt.Println("ACC_ENUM")
	}
}

func (d *decoder) readClass() {
	binary.Read(d.file, d.bo, &(d.cf.this_class))
	binary.Read(d.file, d.bo, &(d.cf.super_class))
	fmt.Println("++++++++++++++++++++ Class")
	thisc := d.cf.constant_pool[d.cf.this_class]
	superc := d.cf.constant_pool[d.cf.super_class]
	fmt.Println(string(d.cf.constant_pool[(d.bo.Uint16(thisc.info))].info[2:]))
	fmt.Println(string(d.cf.constant_pool[(d.bo.Uint16(superc.info))].info[2:]))
}

func (d *decoder) readInterface() {
	binary.Read(d.file, d.bo, &(d.cf.interfaces_count))
	fmt.Println("++++++++++++++++++++ Interface")
	fmt.Printf("count : %d\n", d.cf.interfaces_count)
	d.cf.interfaces = make([]uint16, d.cf.interfaces_count)
	for i := uint16(0); i < d.cf.interfaces_count; i++ {
		binary.Read(d.file, d.bo, &(d.cf.interfaces[i]))
		inter := d.cf.constant_pool[d.cf.interfaces[i]]
		fmt.Println(string(d.cf.constant_pool[(d.bo.Uint16(inter.info))].info[2:]))
	}
}

func (d *decoder) readField() {
	binary.Read(d.file, d.bo, &(d.cf.fields_count))
	fmt.Println("++++++++++++++++++++ Field")
	fmt.Printf("count : %d\n", d.cf.fields_count)
	d.cf.fields = make([]field_info, d.cf.fields_count)
	for i := uint16(0); i < d.cf.fields_count; i++ {
		var fi field_info
		binary.Read(d.file, d.bo, &fi.access_flags)
		binary.Read(d.file, d.bo, &fi.name_index)
		binary.Read(d.file, d.bo, &fi.descriptor_index)
		binary.Read(d.file, d.bo, &fi.attributes_count)
		fie := d.cf.constant_pool[fi.name_index]
		fmt.Println(string(fie.info[2:]))

		fi.attributes = make([]attribute_info, fi.attributes_count)
		for j := uint16(0); j < fi.attributes_count; j++ {
			var name_index 	uint16
			var length 		uint32
			binary.Read(d.file, d.bo, &name_index)
			binary.Read(d.file, d.bo, &length)
			info := make([]uint8, length)
			binary.Read(d.file, d.bo, &info)
		}
	}
}

func (d *decoder) readMethod() {
	binary.Read(d.file, d.bo, &(d.cf.method_count))
	fmt.Println("++++++++++++++++++++ Method")
	fmt.Printf("count : %d\n", d.cf.method_count)
	d.cf.methods = make([]method_info, d.cf.method_count)
	for i := uint16(0); i < d.cf.method_count; i++ {
		var mi method_info
		binary.Read(d.file, d.bo, &mi.access_flags)
		binary.Read(d.file, d.bo, &mi.name_index)
		binary.Read(d.file, d.bo, &mi.descriptor_index)
		binary.Read(d.file, d.bo, &mi.attributes_count)
		met := d.cf.constant_pool[mi.name_index]
		fmt.Println(string(met.info[2:]))

		mi.attributes = make([]attribute_info, mi.attributes_count)
		for j := uint16(0); j < mi.attributes_count; j++ {
			var name_index 	uint16
			var length 		uint32
			binary.Read(d.file, d.bo, &name_index)
			lookup := string(d.cf.constant_pool[name_index].info[2:])
			binary.Read(d.file, d.bo, &length)
			info := make([]uint8, length)
			binary.Read(d.file, d.bo, &info)
			if lookup == "Code" {
				var ca code_attribute
				ca.attribute_name_index = name_index
				ca.attribute_length = length
				ca.max_stack = d.bo.Uint16(info[0:2])
				ca.max_locals = d.bo.Uint16(info[2:4])
				ca.code_length = d.bo.Uint32(info[4:8])
				ca.code = info[8:8+ca.code_length]

				mi.attributes[j] = attribute_info {attribute_name_index:name_index, attribute_length:length, info:info}

				for k := uint32(0); k < ca.code_length; k++ {
					switch ca.code[k] {
						case nop :
							fmt.Println("++nop")
						case aconst_null :
							fmt.Println("++aconst_null")
						case iconst_m1 :
							fmt.Println("++iconst_m1")
						case iconst_0 :
							fmt.Println("++iconst_0")
						case iconst_1 :
							fmt.Println("++iconst_1")
						case iconst_2 :
							fmt.Println("++iconst_2")
						case iconst_3 :
							fmt.Println("++iconst_3")
						case iconst_4 :
							fmt.Println("++iconst_4")
						case iconst_5 :
							fmt.Println("++iconst_5")
						case lconst_0 :
							fmt.Println("++lconst_0")
						case lconst_1 :
							fmt.Println("++lconst_1")
						case fconst_0 :
							fmt.Println("++fconst_0")
						case fconst_1 :
							fmt.Println("++fconst_1")
						case fconst_2 :
							fmt.Println("++fconst_2")
						case dconst_0 :
							fmt.Println("++dconst_0")
						case dconst_1 :
							fmt.Println("++dconst_1")
						case bipush :
							fmt.Println("++bipush")
						case sipush :
							fmt.Println("++sipush")
						case ldc :
							fmt.Println("++ldc")
						case ldc_w :
							fmt.Println("++ldc_w")
						case ldc2_w :
							fmt.Println("++ldc2_w")
						case iload :
							fmt.Println("++iload")
						case lload :
							fmt.Println("++lload")
						case fload :
							fmt.Println("++fload")
						case dload :
							fmt.Println("++dload")
						case aload :
							fmt.Println("++aload")
						case iload_0 :
							fmt.Println("++iload_0")
						case iload_1 :
							fmt.Println("++iload_1")
						case iload_2 :
							fmt.Println("++iload_2")
						case iload_3 :
							fmt.Println("++iload_3")
						case lload_0 :
							fmt.Println("++lload_0")
						case lload_1 :
							fmt.Println("++lload_1")
						case lload_2 :
							fmt.Println("++lload_2")
						case lload_3 :
							fmt.Println("++lload_3")
						case fload_0 :
							fmt.Println("++fload_0")
						case fload_1 :
							fmt.Println("++fload_1")
						case fload_2 :
							fmt.Println("++fload_2")
						case fload_3 :
							fmt.Println("++fload_3")
						case dload_0 :
							fmt.Println("++dload_0")
						case dload_1 :
							fmt.Println("++dload_1")
						case dload_2 :
							fmt.Println("++dload_2")
						case dload_3 :
							fmt.Println("++dload_3")
						case aload_0 :
							fmt.Println("++aload_0")
						case aload_1 :
							fmt.Println("++aload_1")
						case aload_2 :
							fmt.Println("++aload_2")
						case aload_3 :
							fmt.Println("++aload_3")
						case iaload :
							fmt.Println("++iaload")
						case laload :
							fmt.Println("++laload")
						case faload :
							fmt.Println("++faload")
						case daload :
							fmt.Println("++daload")
						case aaload :
							fmt.Println("++aaload")
						case baload :
							fmt.Println("++baload")
						case caload :
							fmt.Println("++caload")
						case saload :
							fmt.Println("++saload")
						case istore :
							fmt.Println("++istore")
						case lstore :
							fmt.Println("++lstore")
						case fstore :
							fmt.Println("++fstore")
						case dstore :
							fmt.Println("++dstore")
						case astore :
							fmt.Println("++astore")
						case istore_0 :
							fmt.Println("++istore_0")
						case istore_1 :
							fmt.Println("++istore_1")
						case istore_2 :
							fmt.Println("++istore_2")
						case istore_3 :
							fmt.Println("++istore_3")
						case lstore_0 :
							fmt.Println("++lstore_0")
						case lstore_1 :
							fmt.Println("++lstore_1")
						case lstore_2 :
							fmt.Println("++lstore_2")
						case lstore_3 :
							fmt.Println("++lstore_3")
						case fstore_0 :
							fmt.Println("++fstore_0")
						case fstore_1 :
							fmt.Println("++fstore_1")
						case fstore_2 :
							fmt.Println("++fstore_2")
						case fstore_3 :
							fmt.Println("++fstore_3")
						case dstore_0 :
							fmt.Println("++dstore_0")
						case dstore_1 :
							fmt.Println("++dstore_1")
						case dstore_2 :
							fmt.Println("++dstore_2")
						case dstore_3 :
							fmt.Println("++dstore_3")
						case astore_0 :
							fmt.Println("++astore_0")
						case astore_1 :
							fmt.Println("++astore_1")
						case astore_2 :
							fmt.Println("++astore_2")
						case astore_3 :
							fmt.Println("++astore_3")
						case iastore :
							fmt.Println("++iastore")
						case lastore :
							fmt.Println("++lastore")
						case fastore :
							fmt.Println("++fastore")
						case dastore :
							fmt.Println("++dastore")
						case aastore :
							fmt.Println("++aastore")
						case bastore :
							fmt.Println("++bastore")
						case castore :
							fmt.Println("++castore")
						case sastore :
							fmt.Println("++sastore")
						case pop :
							fmt.Println("++pop")
						case pop2 :
							fmt.Println("++pop2")
						case dup :
							fmt.Println("++dup")
						case dup_x1 :
							fmt.Println("++dup_x1")
						case dup_x2 :
							fmt.Println("++dup_x2")
						case dup2 :
							fmt.Println("++dup2")
						case dup2_x1 :
							fmt.Println("++dup2_x1")
						case dup2_x2 :
							fmt.Println("++dup2_x2")
						case swap :
							fmt.Println("++swap")
						case iadd :
							fmt.Println("++iadd")
						case ladd :
							fmt.Println("++ladd")
						case fadd :
							fmt.Println("++fadd")
						case dadd :
							fmt.Println("++dadd")
						case isub :
							fmt.Println("++isub")
						case lsub :
							fmt.Println("++lsub")
						case fsub :
							fmt.Println("++fsub")
						case dsub :
							fmt.Println("++dsub")
						case imul :
							fmt.Println("++imul")
						case lmul :
							fmt.Println("++lmul")
						case fmul :
							fmt.Println("++fmul")
						case dmul :
							fmt.Println("++dmul")
						case idiv :
							fmt.Println("++idiv")
						case ldiv :
							fmt.Println("++ldiv")
						case fdiv :
							fmt.Println("++fdiv")
						case ddiv :
							fmt.Println("++ddiv")
						case irem :
							fmt.Println("++irem")
						case lrem :
							fmt.Println("++lrem")
						case frem :
							fmt.Println("++frem")
						case drem :
							fmt.Println("++drem")
						case ineg :
							fmt.Println("++ineg")
						case lneg :
							fmt.Println("++lneg")
						case fneg :
							fmt.Println("++fneg")
						case dneg :
							fmt.Println("++dneg")
						case ishl :
							fmt.Println("++ishl")
						case lshl :
							fmt.Println("++lshl")
						case ishr :
							fmt.Println("++ishr")
						case lshr :
							fmt.Println("++lshr")
						case iushr :
							fmt.Println("++iushr")
						case lushr :
							fmt.Println("++lushr")
						case iand :
							fmt.Println("++iand")
						case land :
							fmt.Println("++land")
						case ior :
							fmt.Println("++ior")
						case lor :
							fmt.Println("++lor")
						case ixor :
							fmt.Println("++ixor")
						case lxor :
							fmt.Println("++lxor")
						case iinc :
							fmt.Println("++iinc")
						case i2l :
							fmt.Println("++i2l")
						case i2f :
							fmt.Println("++i2f")
						case i2d :
							fmt.Println("++i2d")
						case l2i :
							fmt.Println("++l2i")
						case l2f :
							fmt.Println("++l2f")
						case l2d :
							fmt.Println("++l2d")
						case f2i :
							fmt.Println("++f2i")
						case f2l :
							fmt.Println("++f2l")
						case f2d :
							fmt.Println("++f2d")
						case d2i :
							fmt.Println("++d2i")
						case d2l :
							fmt.Println("++d2l")
						case d2f :
							fmt.Println("++d2f")
						case i2b :
							fmt.Println("++i2b")
						case i2c :
							fmt.Println("++i2c")
						case i2s :
							fmt.Println("++i2s")
						case lcmp :
							fmt.Println("++lcmp")
						case fcmpl :
							fmt.Println("++fcmpl")
						case fcmpg :
							fmt.Println("++fcmpg")
						case dcmpl :
							fmt.Println("++dcmpl")
						case dcmpg :
							fmt.Println("++dcmpg")
						case ifeq :
							fmt.Println("++ifeq")
						case ifne :
							fmt.Println("++ifne")
						case iflt :
							fmt.Println("++iflt")
						case ifge :
							fmt.Println("++ifge")
						case ifgt :
							fmt.Println("++ifgt")
						case ifle :
							fmt.Println("++ifle")
						case if_icmpeq :
							fmt.Println("++if_icmpeq")
						case if_icmpne :
							fmt.Println("++if_icmpne")
						case if_icmplt :
							fmt.Println("++if_icmplt")
						case if_icmpge :
							fmt.Println("++if_icmpge")
						case if_icmpgt :
							fmt.Println("++if_icmpgt")
						case if_icmple :
							fmt.Println("++if_icmple")
						case if_acmpeq :
							fmt.Println("++if_acmpeq")
						case if_acmpne :
							fmt.Println("++if_acmpne")
						case goto_x :
							fmt.Println("++goto")
						case jsr :
							fmt.Println("++jsr")
						case ret :
							fmt.Println("++ret")
						case tableswitch :
							fmt.Println("++tableswitch")
						case lookupswitch :
							fmt.Println("++lookupswitch")
						case ireturn :
							fmt.Println("++ireturn")
						case lreturn :
							fmt.Println("++lreturn")
						case freturn :
							fmt.Println("++freturn")
						case dreturn :
							fmt.Println("++dreturn")
						case areturn :
							fmt.Println("++areturn")
						case return_x :
							fmt.Println("++return")
						case getstatic :
							fmt.Println("++getstatic")
						case putstatic :
							fmt.Println("++putstatic")
						case getfield :
							fmt.Println("++getfield")
						case putfield :
							fmt.Println("++putfield")
						case invokevirtual :
							fmt.Println("++invokevirtual")
						case invokespecial :
							fmt.Println("++invokespecial")
						case invokestatic :
							fmt.Println("++invokestatic")
						case invokeinterface :
							fmt.Println("++invokeinterface")
						case invokedynamic :
							fmt.Println("++invokedynamic")
						case new :
							fmt.Println("++new")
						case newarray :
							fmt.Println("++newarray")
						case anewarray :
							fmt.Println("++anewarray")
						case arraylength :
							fmt.Println("++arraylength")
						case athrow :
							fmt.Println("++athrow")
						case checkcast :
							fmt.Println("++checkcast")
						case instanceof :
							fmt.Println("++instanceof")
						case monitorenter :
							fmt.Println("++monitorenter")
						case monitorexit :
							fmt.Println("++monitorexit")
						case wide :
							fmt.Println("++wide")
						case multianewarray :
							fmt.Println("++multianewarray")
						case ifnull :
							fmt.Println("++ifnull")
						case ifnonnull :
							fmt.Println("++ifnonnull")
						case goto_w :
							fmt.Println("++goto_w")
						case jsr_w :
							fmt.Println("++jsr_w")
						case breakpoint :
							fmt.Println("++breakpoint")
						case impdep1 :
							fmt.Println("++impdep1")
						case impdep2 :
							fmt.Println("++impdep2")
					}
				}
				ca.exception_table_length = d.bo.Uint16(info[8+ca.code_length:10+ca.code_length])
				ca.exception = make([]exception_table, ca.exception_table_length)
				for l := uint16(0); l < ca.exception_table_length; l++ {

				}
				index := uint16(ca.code_length) + ca.exception_table_length
				ca.attributes_count = d.bo.Uint16(info[index+10:index+12])
				ca.attributes = make([]attribute_info, ca.attributes_count)
				var lnt LineNumberTable_attribute
				for m := uint16(0); m < ca.attributes_count; m++ {
					var name_index 	uint16
					var length 		uint32
					name_index = d.bo.Uint16(info[index+12:index+14])
					length = d.bo.Uint32(info[index+14:index+18])
					lnt.attribute_name_index = name_index
					lnt.attribute_length = length
					lnt.line_number_table_length = d.bo.Uint16(info[index+18:index+20])
					lnt.line_number_tables = make([]line_number_table, lnt.line_number_table_length)
					fmt.Println("(a)", string(d.cf.constant_pool[name_index].info[2:]))
					for o := uint16(0); o < lnt.line_number_table_length; o++ {
						var startpc uint16
						var lineNumber uint16
						startpc = d.bo.Uint16(info[index+20+(o*4):index+22+(o*4)])
						lineNumber = d.bo.Uint16(info[index+22+(o*4):index+24+(o*4)])
						lnt.line_number_tables[o] = line_number_table { start_pc:startpc, line_number:lineNumber }
						fmt.Println("line",lineNumber,":", startpc)
					}
					//fmt.Println(name_index, length, lnt.line_number_table_length)
				}
				//fmt.Println(ca.max_stack, ca.max_locals, ca.code_length, ca.exception_table_length)
				//fmt.Println(ca.attributes_count)
			}
		}
		d.cf.methods[i] = method_info {access_flags:mi.access_flags, name_index:mi.name_index, descriptor_index:mi.descriptor_index, attributes_count:mi.attributes_count}
	}
}

func (d *decoder) readAttribute() {
	binary.Read(d.file, d.bo, &(d.cf.attributes_count))
	fmt.Printf("attribute count : %d\n", d.cf.attributes_count)
	d.cf.attributes = make([]attribute_info, d.cf.attributes_count)
	for i := uint16(0); i < d.cf.attributes_count; i++ {
		var name_index 	uint16
		var length 		uint32
		binary.Read(d.file, d.bo, &name_index)
		binary.Read(d.file, d.bo, &length)
		info := make([]uint8, length)
		binary.Read(d.file, d.bo, &info)
		d.cf.attributes[i] = attribute_info{ attribute_name_index:name_index, attribute_length:length, info:info }

		att := d.cf.constant_pool[name_index]
		fmt.Println(att, string(att.info[2:]))
	}
}

func readFile(fileClass string, cf *classFile) {
	f, err := os.Open(fileClass)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	d := decoder{file:f, bo:binary.BigEndian, cf:cf}
	d.readMagic()
	d.readVersion()
	d.readConstantPool()
	d.readFlag()
	d.readClass()
	d.readInterface()
	d.readField()
	d.readMethod()
	//d.readAttribute()
}

func findMethod(name string, cf *classFile) (ca code_attribute) {
	fmt.Println("Hey")
	for i := uint16(0); i < cf.method_count; i++ {
		met := cf.constant_pool[cf.methods[i].name_index]
		if string(met.info[2:]) == name {
			fmt.Println("Now found main")
			for j := uint16(0); j < cf.methods[i].attributes_count; j++ {
				metMain := cf.constant_pool[cf.methods[i].attributes[j].attribute_name_index]
				if string(metMain.info[2:]) == "Code" {
					fmt.Println("find code")
					//return
				}
			}
		}
	}
	return
}

func main() {

	cf := &classFile{}

	if len(os.Args) == 1 {
		fmt.Println("please input fileName !!!")
	}else{
		fileName := os.Args[1]
		fileClass := fileName + ".class"
		fmt.Printf("%s\n\n", fileClass)
		readFile(fileClass, cf)
		ca := findMethod("main",cf)
		fmt.Println(ca)
		//find method main
		//execute code of method main
		//execute (ca.code)
	}

}