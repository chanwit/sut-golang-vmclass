package main

import (
	"fmt"
	"os"
	//"io"
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
}

const (
    ACC_PUBLIC      = 0x0001
    ACC_FINAL       = 0x0010
    ACC_SUPER       = 0x0020
    ACC_INTERFACE   = 0x0200
    ACC_ABSTRACT    = 0x0400
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

type attribute_info struct {
    attribute_name_index	uint16
    attribute_length		uint32
    info 					[]uint8
}

type method_info struct {
	access_flags			uint16
	name_index 				uint16
	descriptor_index 		uint16
	attributes_count 		uint16
	attributes 				[]attribute_info
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
	fmt.Println("flag")
	if d.cf.access_flags & ACC_PUBLIC == ACC_PUBLIC {
		fmt.Println("ACC_PUBLIC")
	}
	if d.cf.access_flags & ACC_FINAL == ACC_FINAL {
		fmt.Println("ACC_FINAL")
	}
	if d.cf.access_flags & ACC_SUPER == ACC_SUPER {
		fmt.Println("ACC_SUPER")
	}
	if d.cf.access_flags & ACC_INTERFACE == ACC_INTERFACE {
		fmt.Println("ACC_INTERFACE")
	}
	if d.cf.access_flags & ACC_ABSTRACT == ACC_ABSTRACT {
		fmt.Println("ACC_ABSTRACT")
	}
}

func (d *decoder) readClass() {
	binary.Read(d.file, d.bo, &(d.cf.this_class))
	binary.Read(d.file, d.bo, &(d.cf.super_class))
	fmt.Println("class")
	fmt.Println(d.cf.constant_pool[d.cf.this_class])
	fmt.Println(d.cf.constant_pool[d.cf.super_class])
	thisc := d.cf.constant_pool[d.cf.this_class]
	superc := d.cf.constant_pool[d.cf.super_class]
	fmt.Println(string(d.cf.constant_pool[(d.bo.Uint16(thisc.info))].info[2:]))
	fmt.Println(string(d.cf.constant_pool[(d.bo.Uint16(superc.info))].info[2:]))
}

func (d *decoder) readInterface() {
	binary.Read(d.file, d.bo, &(d.cf.interfaces_count))
	fmt.Printf("interface count : %d\n", d.cf.interfaces_count)
	d.cf.interfaces = make([]uint16, d.cf.interfaces_count)
	for i := uint16(0); i < d.cf.interfaces_count; i++ {
		binary.Read(d.file, d.bo, &(d.cf.interfaces[i]))
		fmt.Print(d.cf.constant_pool[d.cf.interfaces[i]], " ")
		inter := d.cf.constant_pool[d.cf.interfaces[i]]
		fmt.Println(string(d.cf.constant_pool[(d.bo.Uint16(inter.info))].info[2:]))
	}
}

func (d *decoder) readField() {
	binary.Read(d.file, d.bo, &(d.cf.fields_count))
	fmt.Printf("field count : %d\n", d.cf.fields_count)
	d.cf.fields = make([]field_info, d.cf.fields_count)
	for i := uint16(0); i < d.cf.fields_count; i++ {
		var fi field_info
		binary.Read(d.file, d.bo, &fi.access_flags)
		binary.Read(d.file, d.bo, &fi.name_index)
		binary.Read(d.file, d.bo, &fi.descriptor_index)
		binary.Read(d.file, d.bo, &fi.attributes_count)
		fmt.Println(fi.access_flags, fi.name_index, fi.descriptor_index, fi.attributes_count)

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
	fmt.Printf("method count : %d\n", d.cf.method_count)
	d.cf.methods = make([]method_info, d.cf.method_count)
	for i := uint16(0); i < d.cf.method_count; i++ {
		var mi method_info
		binary.Read(d.file, d.bo, &mi.access_flags)
		binary.Read(d.file, d.bo, &mi.name_index)
		binary.Read(d.file, d.bo, &mi.descriptor_index)
		binary.Read(d.file, d.bo, &mi.attributes_count)
		fmt.Println(mi.access_flags, mi.name_index, mi.descriptor_index, mi.attributes_count)

		mi.attributes = make([]attribute_info, mi.attributes_count)
		for j := uint16(0); j < mi.attributes_count; j++ {
			var name_index 	uint16
			var length 		uint32
			binary.Read(d.file, d.bo, &name_index)
			binary.Read(d.file, d.bo, &length)
			info := make([]uint8, length)
			binary.Read(d.file, d.bo, &info)
		}
	}
}

func readFile(fileClass string, cf classFile) {
	f, err := os.Open(fileClass)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	d := decoder{file:f, bo:binary.BigEndian, cf:&cf}
	d.readMagic()
	d.readVersion()
	d.readConstantPool()
	d.readFlag()
	d.readClass()
	d.readInterface()
	d.readField()
	d.readMethod()
}

func main() {

	var cf classFile

	if len(os.Args) == 1 {
		fmt.Println("please input fileName !!!")
	}else{
		fileName := os.Args[1]
		fileClass := fileName + ".class"
		fmt.Printf("%s\n\n", fileClass)
		readFile(fileClass, cf)
	}

}