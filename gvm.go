
package main

import ( "fmt"
		"os"
		"io"	
		"encoding/binary"
		//"bytes"
)

type classfile struct{
	magic 									uint32
	minor_version 							uint16
	major_version 							uint16
	constant_pool_count						uint16
	constant_pool							[]cp_info
	access_flags							uint16
	this_class								uint16
	super_class								uint16
	interfaces_count						uint16
	interfaces								[]uint16
	fields_count							uint16
	fields									[]field_info
	method_count							uint16
	method									[]method_info
	attributes_count						uint16
	attributes								[]attribute_info
}	
const (
nop 			= iota
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
//(no name) = 0xcb-fd
impdep1 = 254
impdep2 = 255
)

const (
    ACC_PUBLIC      = 0x0001
    ACC_FINAL       = 0x0010
    ACC_SUPER       = 0x0020
    ACC_INTERFACE   = 0x0200
    ACC_ABSTRACT    = 0x0400
)
const(		//ACC_PUBLIC 	= 0x0001 	
		ACC_PRIVATE 	= 0x0002 	
		ACC_PROTECTED 	= 0x0004 	
		ACC_STATIC 	= 0x0008 
		//ACC_FINAL 	= 0x0010 	
		ACC_VOLATILE 	= 0x0040 	 
		ACC_TRANSIENT 	= 0x0080 	
		ACC_SYNTHETIC 	= 0x1000 	
		ACC_ENUM 	= 0x4000
)


type cp_info struct{
	tag		uint8
	info 	[]uint8
}
type field_info struct{
    access_flags 		uint16
    name_index      	uint16
    descriptor_index	uint16
    attributes_count	uint16
    attributes			[]attribute_info 	//[attributes_count]
}
type constantValue_attribute struct{
    attribute_name_index	uint16
    attribute_length		uint32
    constantvalue_index		uint16
}
type method_info struct {
	access_flags		uint16
	name_index 			uint16
	descriptor_index 	uint16
	attributes_count 	uint16
	attributes 			[]code_attribute
}
type code_attribute struct{
    attribute_name_index	uint16
    attribute_length		uint32
    max_stack				uint16
    max_locals				uint16
    code_length				uint32
    code					[]uint8
 	exception_table_length 	uint16
    exception				[]exception_table
    attributes_count 		uint16
    line_number_table_att   []LineNumberTable_attribute
}
type exception_table struct{
	start_pc	uint16
    end_pc		uint16
    handler_pc	uint16
    catch_type	uint16
}
type attribute_info struct{
    attribute_name_index	uint16
    attribute_length		uint32
    info					[]uint8	//[attribute_length]
}
type LineNumberTable_attribute struct {
    attribute_name_index		uint16
    attribute_length			uint32
    line_number_table_length 	uint16
    line_number_tables			[]line_number_table
}
type  line_number_table struct{
	start_pc			uint16
    line_number 		uint16	
}

type CONSTANT_Class_info struct{
	tag		uint8
	name_index	uint16
}

type CONSTANT_Fieldref_info struct{
    tag 			uint8
    class_index			uint16
    name_and_type_index		uint16
}
type CONSTANT_Methodref_info struct{
    tag				uint8
    class_index			uint16
    name_and_type_index	uint16
}

type CONSTANT_InterfaceMethodref_info struct{
    tag				uint8
    class_index			uint16
    name_and_type_index		uint16
}

type CONSTANT_String_info struct{
    tag			uint8
    string_index	uint16
}

type CONSTANT_Integer_info struct{
    tag		uint8
    bytes	uint32
}

type CONSTANT_Float_info struct {
    tag		uint8
   	bytes	uint32
}

type CONSTANT_Long_info struct{
    tag			uint8
    high_bytes 	uint32
    low_bytes	uint32
}

type CONSTANT_Double_info struct{
    tag		uint8
    high_bytes	uint32
    low_bytes	uint32
}

type CONSTANT_NameAndType_info struct{
    tag					uint8
    name_index 			uint16
    descriptor_index	uint16
}

type CONSTANT_Utf8_info struct{
    tag			uint8
    length		uint16
    bytes 		[]uint8
}

type CONSTANT_MethodHandle_info struct{
    tag				uint8
    reference_kind		uint8
    reference_index		uint16
}
type CONSTANT_MethodType_info struct{
    tag			uint8
    descriptor_index 	uint16
}
type CONSTANT_InvokeDynamic_info struct{
    tag					uint8
    bootstrap_method_attr_index		uint16
    name_and_type_index			uint16
}


const (	CONSTANT_Class				=	7
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
type decoder struct{
	file 	io.Reader
	bo		binary.ByteOrder
	cf 		*classfile
}

func (d *decoder) readMagic() {

    binary.Read(d.file, d.bo, &(d.cf.magic))
    fmt.Printf("Magic code = %x\n", d.cf.magic)
}

func (d *decoder) readVersion() {
    binary.Read(d.file, d.bo, &(d.cf.minor_version))
    binary.Read(d.file, d.bo, &(d.cf.major_version))
    fmt.Printf("Version = %d.%d\n", d.cf.major_version,d.cf.minor_version)
}
func (d *decoder) readFlag() {
	binary.Read(d.file, d.bo, &(d.cf.access_flags))
	if d.cf.access_flags & ACC_PUBLIC == ACC_PUBLIC {
		fmt.Print("ACC_PUBLIC ")
	}
	if d.cf.access_flags & ACC_ABSTRACT == ACC_ABSTRACT{
		fmt.Print("ACC_ABSTRACT ")
	}
	if d.cf.access_flags & ACC_SUPER == ACC_SUPER{
		fmt.Print("ACC_SUPER ")
	}

}
func (d *decoder) readConstantPool() {
	binary.Read(d.file,d.bo,&(d.cf.constant_pool_count))
	fmt.Printf("Constant pool:\n")

	d.cf.constant_pool = make([]cp_info, d.cf.constant_pool_count)
	for i:= uint16(1); i<d.cf.constant_pool_count ;i++{
		
		var tag uint8
		binary.Read(d.file,d.bo,&(tag))
			switch tag { 
			case CONSTANT_Class :
				info := make([]byte, 2)
				binary.Read(d.file, d.bo, info)
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }
			case CONSTANT_String :fallthrough
			case CONSTANT_MethodType :
				info := make([]byte, 2)
				binary.Read(d.file, d.bo, info)
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }
			case CONSTANT_MethodHandle :
				info := make([]byte, 3)
				binary.Read(d.file, d.bo, info)
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }
			case CONSTANT_Fieldref :fallthrough
			case CONSTANT_Integer :fallthrough
			case CONSTANT_Float :fallthrough
			case CONSTANT_NameAndType :fallthrough
			case CONSTANT_InterfaceMethodref :fallthrough
			case CONSTANT_Methodref :fallthrough
			case CONSTANT_InvokeDynamic :
				info := make([]byte, 4)
				binary.Read(d.file, d.bo, info)
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }
			case CONSTANT_Long :fallthrough
			case CONSTANT_Double :
				info := make([]byte, 8)
				binary.Read(d.file, d.bo, info)
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }
		
			case CONSTANT_Utf8 :
				var length uint16
				binary.Read(d.file, d.bo, &(length))
				info := make([]byte, 2 + length)
				binary.BigEndian.PutUint16(info[0:2], length)
				binary.Read(d.file, d.bo, info[2:])
				d.cf.constant_pool[i] = cp_info{ tag:tag, info:info }
			//	fmt.Printf("#%d\t%s\n",i,info[2:])
		}	
	}
	// --------------------To out put all constant_pool ---------------------
	for j:= uint16(1); j<d.cf.constant_pool_count ;j++{
		check := d.cf.constant_pool[j]

		switch d.cf.constant_pool[j].tag  {
		case CONSTANT_Class :
			insec:=binary.BigEndian.Uint16(check.info[:])
			fmt.Printf("#%d = Class\t\t#%d\t// ",j,insec)
			fmt.Printf("%s\n",string(d.cf.constant_pool[insec].info[2:]))
		case CONSTANT_Methodref :
			classindex:=binary.BigEndian.Uint16(check.info[:2])
				nametypeindex:=binary.BigEndian.Uint16(check.info[2:4])
					fmt.Printf("#%d = Methodref\t\t#%d.#%d\t// ",j,classindex,nametypeindex)				//#3.#20
					nextclassindex:=binary.BigEndian.Uint16(d.cf.constant_pool[classindex].info[:2])
						fmt.Printf("%s",string(d.cf.constant_pool[nextclassindex].info[2:]))
					nextnatindex:=binary.BigEndian.Uint16(d.cf.constant_pool[nametypeindex].info[:2])
						fmt.Printf(".%q",string(d.cf.constant_pool[nextnatindex].info[2:]))
					overnatindex:=binary.BigEndian.Uint16(d.cf.constant_pool[nametypeindex].info[2:4])
						fmt.Printf(":%s\n",string(d.cf.constant_pool[overnatindex].info[2:]))
		case CONSTANT_Fieldref :
			classindex:=binary.BigEndian.Uint16(check.info[:2])
				nametypeindex:=binary.BigEndian.Uint16(check.info[2:4])
					fmt.Printf("#%d = Fieldref\t\t#%d.#%d\t// ",j,classindex,nametypeindex)				//#3.#20
					nextclassindex:=binary.BigEndian.Uint16(d.cf.constant_pool[classindex].info[:2])
						fmt.Printf("%s",string(d.cf.constant_pool[nextclassindex].info[2:]))
					nextnatindex:=binary.BigEndian.Uint16(d.cf.constant_pool[nametypeindex].info[:2])
						fmt.Printf(".%s",string(d.cf.constant_pool[nextnatindex].info[2:]))
					overnatindex:=binary.BigEndian.Uint16(d.cf.constant_pool[nametypeindex].info[2:4])
						fmt.Printf(":%s\n",string(d.cf.constant_pool[overnatindex].info[2:]))
		case CONSTANT_NameAndType :
			nameindex:=binary.BigEndian.Uint16(check.info[:2])
			desindex:=binary.BigEndian.Uint16(check.info[2:4])
			fmt.Printf("#%d = NameAndType\t#%d:#%d\t// ",j,nameindex,desindex)
			fmt.Printf("%q:%s\n",string(d.cf.constant_pool[nameindex].info[2:]),string(d.cf.constant_pool[desindex].info[2:]))
		case CONSTANT_Utf8 :
			fmt.Printf("#%d = Utf8\t\t%s\n",j,check.info[2:])

		}	
	}
}

func (d *decoder) readThis() {
	binary.Read(d.file,d.bo,&(d.cf.this_class))

	thisc := d.cf.constant_pool[d.cf.this_class]
	fmt.Println("\nthis_class is "+string(d.cf.constant_pool[(d.bo.Uint16(thisc.info))].info[2:]))

	binary.Read(d.file,d.bo,&(d.cf.super_class))
	thiss := d.cf.constant_pool[d.cf.super_class]
	fmt.Println("super_class is "+string(d.cf.constant_pool[(d.bo.Uint16(thiss.info))].info[2:]))

}
func (d *decoder) readInterface() {
	binary.Read(d.file,d.bo,&(d.cf.interfaces_count))
	d.cf.interfaces = make([]uint16 , d.cf.interfaces_count)
	

	for i:=uint16(0); i<d.cf.interfaces_count ;i++{
		binary.Read(d.file,d.bo,&(d.cf.interfaces[i]))
		inter := d.cf.constant_pool[d.cf.interfaces[i]]
		fmt.Printf("interface = ")
		fmt.Println(string(d.cf.constant_pool[(d.bo.Uint16(inter.info))].info[2:]))
	}
}

func (d *decoder) readField() {
		binary.Read(d.file, d.bo, &(d.cf.fields_count))
		//fmt.Printf("field count : %d\n", d.cf.fields_count)
		d.cf.fields = make([]field_info, d.cf.fields_count)
		for i := uint16(0); i < d.cf.fields_count; i++ {
			var fi field_info
				binary.Read(d.file, d.bo, &fi.access_flags)
					
				binary.Read(d.file, d.bo, &fi.name_index)
				binary.Read(d.file, d.bo, &fi.descriptor_index)
				binary.Read(d.file, d.bo, &fi.attributes_count)
			d.cf.fields[i]=field_info{access_flags:fi.access_flags,name_index:fi.name_index,descriptor_index:fi.descriptor_index,attributes_count:fi.attributes_count}
				//Access_flags and FieldName Print!
				fieldname := d.cf.constant_pool[fi.name_index]
				fielddes := d.cf.constant_pool[fi.descriptor_index]
					if fi.access_flags & ACC_PUBLIC == ACC_PUBLIC {
						fmt.Print("public ")
					}
					if fi.access_flags & ACC_PRIVATE == ACC_PRIVATE{
						fmt.Print("private ")
					}
					if fi.access_flags & ACC_PROTECTED == ACC_PROTECTED{
						fmt.Print("protected ")
					}

				fmt.Println(string(fieldname.info[2:]),string(fielddes.info[2:])) 
				//
			
			fi.attributes = make([]attribute_info, fi.attributes_count)
					for j := uint16(0); j < fi.attributes_count; j++ {
						var name_index uint16
						var length uint32
						var cv constantValue_attribute
						binary.Read(d.file, d.bo, &name_index)
						binary.Read(d.file, d.bo, &length)
						info := make([]uint8, length)
						binary.Read(d.file, d.bo, &info)
					
							nameCheck:=d.cf.constant_pool[name_index]
						//	fmt.Println(string(nameCheck.info[2:]))
								if string(nameCheck.info[2:]) == "ConstantValue"{
									cv.attribute_name_index = name_index
									cv.attribute_length = length
									fmt.Println(string(cv.attribute_name_index))
									//cv.constantvalue_index =  
/*
			fi.attributes = make([]constantValue_attribute, fi.attributes_count)
			for k := uint16(0); k < fi.attributes_count; k++ {
				var name_index uint16
				var length uint32
				var constantvalue_index uint16
				binary.Read(d.file, d.bo, &name_index)
				binary.Read(d.file, d.bo, &length)
				binary.Read(d.file, d.bo, &constantvalue_index)
				fi.attributes[k]=constantValue_attribute{attribute_name_index:name_index, attribute_length:length,constantvalue_index:constantvalue_index}
			}
		}
		
		cp := d.cf.constant_pool
    	for i := uint16(0); i < d.cf.fields_count; i++ {
       		fi := d.cf.fields[i]
       		cp1 := cp[fi.name_index]
       		cp2 := cp[fi.descriptor_index]
       				if fi.access_flags & ACC_PUBLIC == ACC_PUBLIC {
						fmt.Print("public ")
					}
					if fi.access_flags & ACC_PRIVATE == ACC_PRIVATE{
						fmt.Print("private ")
					}
					if fi.access_flags & ACC_PROTECTED == ACC_PROTECTED{
						fmt.Print("protected ")
					}
       		fmt.Println( string(cp2.info[2:]), string(cp1.info[2:]))
       		
       	}*/
}
}
}
}
func (d *decoder) readMethod() {
	binary.Read(d.file, d.bo, &(d.cf.method_count))
	fmt.Printf("method count : %d\n", d.cf.method_count)
	d.cf.method = make([]method_info, d.cf.method_count)
	for i := uint16(0); i < d.cf.method_count; i++ {
		var mi method_info
		binary.Read(d.file, d.bo, &mi.access_flags)
		binary.Read(d.file, d.bo, &mi.name_index)
		binary.Read(d.file, d.bo, &mi.descriptor_index)
		binary.Read(d.file, d.bo, &mi.attributes_count)
	//	fmt.Println("method detail: ",mi.access_flags, mi.name_index, mi.descriptor_index, mi.attributes_count)
		d.cf.method[i]=method_info{
			access_flags:mi.access_flags,
			name_index:mi.name_index,
			descriptor_index:mi.descriptor_index,
			attributes_count:mi.attributes_count }	

			metname := d.cf.constant_pool[mi.name_index]
			metdes := d.cf.constant_pool[mi.descriptor_index]
					if mi.access_flags & ACC_PUBLIC == ACC_PUBLIC {
						fmt.Print("public ")
					}
					if mi.access_flags & ACC_PRIVATE == ACC_PRIVATE{
						fmt.Print("private ")
					}
					if mi.access_flags & ACC_PROTECTED == ACC_PROTECTED{
						fmt.Print("protected ")
					}
			fmt.Println(string(metname.info[2:]),string(metdes.info[2:])) 
		
			d.cf.method[i].attributes = make([]code_attribute, mi.attributes_count)
					for j := uint16(0); j < mi.attributes_count; j++ {
						var name_index uint16
						var length uint32
						var ca code_attribute
						//var excep_length uint16
						binary.Read(d.file, d.bo, &name_index)
						binary.Read(d.file, d.bo, &length)
							info := make([]uint8, length)
						binary.Read(d.file, d.bo, info)
							nameCheck:=d.cf.constant_pool[name_index]
							fmt.Println(string(nameCheck.info[2:]))
								if string(nameCheck.info[2:]) == "Code"{
									ca.attribute_name_index=name_index
									ca.attribute_length=length
									ca.max_stack=d.bo.Uint16(info[0:2])
									ca.max_locals=d.bo.Uint16(info[2:4])
									ca.code_length=d.bo.Uint32(info[4:8])
									ca.code=info[8:8+ca.code_length]
									lookupcode(ca.code,ca.code_length)

									ca.exception_table_length = d.bo.Uint16(info[8+ca.code_length:10+ca.code_length])
									d.cf.method[i].attributes[j].exception = make([]exception_table, ca.exception_table_length)
        								for l := uint16(0); l < ca.exception_table_length; l++ {
        									var start_pc uint16
        									var end_pc  uint16
        									var handler_pc  uint16
        									var catch_type  uint16
        									start_pc = d.bo.Uint16(info[10+ca.code_length:12+ca.code_length])
        									end_pc = d.bo.Uint16(info[12+ca.code_length:14+ca.code_length])
        									handler_pc = d.bo.Uint16(info[14+ca.code_length:16+ca.code_length])
        									catch_type = d.bo.Uint16(info[16+ca.code_length:18+ca.code_length])
        								d.cf.method[i].attributes[j].exception[l] = exception_table {
        									start_pc:start_pc,
        									end_pc:end_pc,
        									handler_pc:handler_pc,
        									catch_type:catch_type}
        									fmt.Println(start_pc, end_pc, handler_pc, catch_type)
        								}
        							index := uint16(ca.code_length) + ca.exception_table_length
        							ca.attributes_count = d.bo.Uint16(info[index+10:index+12])
        							d.cf.method[i].attributes[j].line_number_table_att = make([]LineNumberTable_attribute, ca.attributes_count)
									var lnt_a LineNumberTable_attribute
										for m := uint16(0); m < ca.attributes_count; m++ {
											var name_index uint16
											var length uint32
											name_index = d.bo.Uint16(info[index+12:index+14])
											length = d.bo.Uint32(info[index+14:index+18])
											lnt_a.attribute_name_index = name_index
											lnt_a.attribute_length = length
											lnt_a.line_number_table_length = d.bo.Uint16(info[index+18:index+20])
											lnt_a.line_number_tables = make([]line_number_table, lnt_a.line_number_table_length)
											fmt.Println(" ", string(d.cf.constant_pool[name_index].info[2:]), ":")
										d.cf.method[i].attributes[j].line_number_table_att[m].line_number_tables = make([]line_number_table, lnt_a.line_number_table_length)
											for o := uint16(0); o < lnt_a.line_number_table_length; o++ {
												var start_pc uint16
												var line_number uint16
												start_pc = d.bo.Uint16(info[index+20+(o*4):index+22+(o*4)])
												line_number = d.bo.Uint16(info[index+22+(o*4):index+24+(o*4)])
							
											d.cf.method[i].attributes[j].line_number_table_att[m].line_number_tables[o] = line_number_table {start_pc:start_pc, line_number:line_number}
											fmt.Println(" line",line_number,":", start_pc)
											}
							
										d.cf.method[i].attributes[j].line_number_table_att[m] = LineNumberTable_attribute {
											attribute_name_index:lnt_a.attribute_name_index,
											attribute_length:lnt_a.attribute_length,
											line_number_table_length:lnt_a.line_number_table_length,
											line_number_tables:d.cf.method[i].attributes[j].line_number_table_att[m].line_number_tables}
							
										}

									d.cf.method[i].attributes[j] = code_attribute {attribute_name_index:name_index,
										attribute_length:length,
										max_stack:ca.max_stack,
										max_locals:ca.max_locals,
										code_length:ca.code_length,
										code:ca.code,//d.cf.method[i].attributes[j].code, // There is no []code inside.
										exception_table_length:ca.exception_table_length,
										exception:d.cf.method[i].attributes[j].exception,
										attributes_count:ca.attributes_count,
										line_number_table_att:d.cf.method[i].attributes[j].line_number_table_att}
								}
					}
	}
}
								/*
									ca.exception_table_length=d.bo.Uint16(info[8+ca.code_length:10+ca.code_length])
									ca.exception_tables=make([]exception_table,ca.exception_table_length)
										// loopfor exception
										for k:=uint16(0) ;k < ca.exception_table_length ;k++{}

										index:=uint16(ca.code_length)+ca.exception_table_length // Starting Point of Attributes count.
									ca.attributes_count=d.bo.Uint16(info[index+10:index+12])
									ca.attributes=make([]LineNumberTable_attribute,ca.attributes_count)
										var lnt LineNumberTable_attribute
										for l:=uint16(0) ;l < ca.attributes_count ; l++{
											var nameindex 			uint16
											var attributelength 	uint32
											nameindex = d.bo.Uint16(info[index+12:index+14])
											attributelength = d.bo.Uint32(info[index+14:index+18])
											lnt.attribute_name_index = nameindex
											lnt.attribute_length = attributelength
											lnt.line_number_table_length = d.bo.Uint16(info[index+18:index+20])
											lnt.line_number_tables = make([]line_number_table, lnt.line_number_table_length)
											fmt.Println("(a)", string(d.cf.constant_pool[nameindex].info[2:]))
											for o := uint16(0); o < lnt.line_number_table_length; o++ {
												var startpc uint16
												var lineNumber uint16
												startpc = d.bo.Uint16(info[index+20+(o*4):index+22+(o*4)])
												lineNumber = d.bo.Uint16(info[index+22+(o*4):index+24+(o*4)])
												lnt.line_number_tables[o] = line_number_table { start_pc:startpc, line_number:lineNumber }
												fmt.Println("line",lineNumber,":", startpc)
											}
											d.cf.method[i].attributes[j].line_number_tables[l] = LineNumberTable_attribute {
												attribute_name_index:lnt.attribute_name_index, 
												attribute_length:lnt.attribute_length,
												line_number_table_length:lnt.line_number_table_length,
												line_number_tables:d.cf.method[i].attributes[j].line_number_table_att[l].line_number_tables}
										}
								

								}
								d.cf.method[i].attributes[j]=code_attribute{
									attribute_name_index:ca.attribute_name_index,
									attribute_length:ca.attribute_length,
									max_stack:ca.max_stack,
									max_locals:ca.max_locals,
									code_length:ca.code_length,
									code:ca.code,
									exception_table_length:ca.exception_table_length,
									exception_tables:ca.exception_tables,
									attributes_count:ca.attributes_count,
									attributes:ca.attributes}		
						
								
					}	
	}	
} */
func lookupcode(ca []uint8,length uint32) {
   	//fmt.Printf("\n%x\n",ca[:1])
   	var m = make(map[int]string)
   	m[0]="nop"				
	m[1]="aconst_null"
	m[2]="iconst_m1"				
	m[3]="iconst_0"					
	m[4]="iconst_1"					
	m[5]="iconst_2"					
	m[6]="iconst_3"					
	m[7]="iconst_4"					
	m[8]="iconst_5"					
	m[9]="lconst_0"					
	m[10]="lconst_1"					
	m[11]="fconst_0"				
	m[12]="fconst_1"				
	m[13]="fconst_2"				
	m[14]="dconst_0"					
	m[15]="dconst_1"					
	m[16]="bipush"						
	m[17]="sipush"						
	m[18]="ldc"						
	m[19]="ldc_w"						
	m[20]="ldc2_w"						
	m[21]="iload"						
	m[22]="lload"						
	m[23]="fload"						
	m[24]="dload"						
	m[25]="aload"						
	m[26]="iload_0"						
	m[27]="iload_1"						
	m[28]="iload_2"						
	m[29]="iload_3"						
	m[30]="lload_0"						
	m[31]="lload_1"						
	m[32]="lload_2"						
	m[33]="lload_3"						
	m[34]="fload_0"						
	m[35]="fload_1"						
	m[36]="fload_2"						
	m[37]="fload_3"						
	m[38]="dload_0"						
	m[39]="dload_1"						
	m[40]="dload_2"						
	m[41]="dload_3"						
	m[42]="aload_0"						
	m[43]="aload_1"						
	m[44]="aload_2"						
	m[45]="aload_3"						
	m[46]="iaload"					
	m[47]="laload"						
	m[48]="faload"						
	m[49]="daload"						
	m[50]="aaload"						
	m[51]="baload"						
	m[52]="caload"						
	m[53]="saload"						
	m[54]="istore"						
	m[55]="lstore"						
	m[56]="fstore"						
	m[57]="dstore"						
	m[58]="astore"						
	m[59]="istore_0"					
	m[60]="istore_1"					
	m[61]="istore_2"					
	m[62]="istore_3"					
	m[63]="lstore_0"					
	m[64]="lstore_1"					
	m[65]="lstore_2"					
	m[66]="lstore_3"					
	m[67]="fstore_0"					
	m[68]="fstore_1"					
	m[69]="fstore_2"					
	m[70]="fstore_3"					
	m[71]="dstore_0"					
	m[72]="dstore_1"					
	m[73]="dstore_2"					
	m[74]="dstore_3"					
	m[75]="astore_0"					
	m[76]="astore_1"					
	m[77]="astore_2"					
	m[78]="astore_3"					
	m[79]="iastore"						
	m[80]="lastore"						
	m[81]="fastore"						
	m[82]="dastore"						
	m[83]="aastore"						
	m[84]="bastore"						
	m[85]="castore"						
	m[86]="sastore"						
	m[87]="pop"						
	m[88]="pop2"						
	m[89]="dup"							
	m[90]="dup_x1"						
	m[91]="dup_x2"						 
	m[92]="dup2"						
	m[93]="dup2_x1"						
	m[94]="dup2_x2"						
	m[95]="swap"						
	m[96]="iadd"						
	m[97]="ladd"						
	m[98]="fadd"						
	m[99]="dadd"						
	m[100]="isub"						
	m[101]="lsub"						
	m[102]="fsub"						
	m[103]="dsub"						
	m[104]="imul"						
	m[105]="lmul"						
	m[106]="fmul"						
	m[107]="dmul"						
	m[108]="idiv"						
	m[109]="ldiv"						
	m[110]="fdiv"						
	m[111]="ddiv"						
	m[112]="irem"						
	m[113]="lrem"						
	m[114]="frem"						
	m[115]="drem"						
	m[116]="ineg"						
	m[117]="lneg"						
	m[118]="fneg"						
	m[119]="dneg"						
	m[120]="ishl"						
	m[121]="lshl"						
	m[122]="ishr"						
	m[123]="lshr"						
	m[124]="iushr"						
	m[125]="lushr"						
	m[126]="iand"						
	m[127]="land"						
	m[128]="ior"							
	m[129]="lor"							
	m[130]="ixor"						
	m[131]="lxor"						
	m[132]="iinc"						
	m[133]="i2l"							
	m[134]="i2f"							
	m[135]="i2d"							
	m[136]="l2i"							
	m[137]="l2f"							
	m[138]="l2d"							
	m[139]="f2i"							
	m[140]="f2l"							
	m[141]="f2d"							
	m[142]="d2i"							
	m[143]="d2l"							
	m[144]="d2f"							
	m[145]="i2b"							
	m[146]="i2c"							
	m[147]="i2s"							
	m[148]="lcmp"						
	m[149]="fcmpl"						
	m[150]="fcmpg"						
	m[151]="dcmpl"						
	m[152]="dcmpg"						
	m[153]="ifeq"					
	m[154]="ifne"						
	m[155]="iflt"						
	m[156]="ifge"						
	m[157]="ifgt"						
	m[158]="ifle"						
	m[159]="if_icmpeq"					
	m[160]="if_icmpne"					
	m[161]="if_icmplt"					
	m[162]="if_icmpge"					
	m[163]="if_icmpgt"					
	m[164]="if_icmple"					
	m[165]="if_acmpeq"					
	m[166]="if_acmpne"					
	m[167]="Goto"						
	m[168]="jsr"							
	m[169]="ret"							
	m[170]="tableswitch"					
	m[171]="lookupswitch"				
	m[172]="ireturn"
	m[173]="lreturn"						
	m[174]="freturn"						
	m[175]="dreturn"						
	m[176]="areturn"						
	m[177]="Return"					
	m[178]="getstatic"					
	m[179]="putstatic"					
	m[180]="getfield"					
	m[181]="putfield"					
	m[182]="invokevirtual"				
	m[183]="invokespecial"				
	m[184]="invokestatic"				
	m[185]="invokeinterface"			
	m[186]="invokedynamic"				
	m[187]="new"							
	m[188]="newarray"					
	m[189]="anewarray"					
	m[190]="arraylength"					
	m[191]="athrow"						
	m[192]="checkcast"					
	m[193]="instanceof"					
	m[194]="monitorenter"				
	m[195]="monitorexit"					
	m[196]="wide"												
	m[197]="multianewarray"				
	m[198]="ifnull"						
	m[199]="ifnonnull"					
	m[200]="goto_w"						
	m[201]="jsr_w"						
	m[202]="breakpoint"					
for k:=203 ; k<254; k++{
	m[k]="noname"
}														
	m[254]="impdep1"						
	m[255]="impdep2"	

	checkTank := 0		
for j:=uint32(0) ;j<length ;j++{
	if checkTank==1{
		keepindex := ca[j:j+1]
		 checkTank=0
		 j=j
		 fmt.Println("keepindex 1 : = ",keepindex)
	}else if checkTank==2{
		keepindex := ca[j:j+2]
		 checkTank=0
		 j=j+1
		 fmt.Println("keepindex 2 : = ",keepindex)
	}else{
		t := ca[j:j+1]
		s := make([]interface{}, len(t))
		for i, v := range t {
			s[i]=v
			value,_:=m[int(v)]
				switch int(v){
					case 183:fallthrough
					case 182:fallthrough
					case 181:
						fmt.Println(j,": ",value)
						checkTank=2
					case 18:
						fmt.Println(j,": ",value)
						checkTank=1
					default: 
						fmt.Println(j,": ",value)
						checkTank=0
				}
   		}
	}
}
   		
   		
   		
   
   		
 }


func (d *decoder) readAttribute() {
		binary.Read(d.file, d.bo, &(d.cf.attributes_count))
		fmt.Printf("attribute count : %d\n", d.cf.attributes_count)
		d.cf.attributes = make([]attribute_info, d.cf.attributes_count)
		for i := uint16(0); i < d.cf.attributes_count; i++ {
			var name_index uint16
			var length uint32
			binary.Read(d.file, d.bo, &name_index)
			binary.Read(d.file, d.bo, &length)
				info := make([]uint8, length)
				binary.Read(d.file, d.bo, &info)
					d.cf.attributes[i] = attribute_info{ attribute_name_index:name_index, attribute_length:length, info:info }
					att := d.cf.constant_pool[name_index]
					fmt.Println(string(att.info[2:]))
		}
}

		 	
		
func checkSize(f *os.File) (){
	state,_:=f.Stat()
	fmt.Printf("size = %d bytes\n", state.Size())
}


func openfile(filename string,cf *classfile) {
	f , err := os.Open(filename) // For read file.
		if err != nil {
			fmt.Printf("%v", err)	
		}
	defer f.Close()
	d := decoder{file:f, bo:binary.BigEndian ,cf: cf}

	checkSize(f)
    d.readMagic()
    d.readVersion()
    d.readConstantPool()
    d.readFlag()
    d.readThis()
    d.readInterface()
    d.readField()
    d.readMethod()
    d.readAttribute()
    
}

func findMethod(name string, cf *classfile) (ca code_attribute) {
	fmt.Println("Hey")
	for i := uint16(0); i < cf.method_count; i++ {
		met := cf.constant_pool[cf.method[i].name_index]
		if string(met.info[2:]) == name {
			fmt.Println("Now found main")
			for j := uint16(0); j < cf.method[i].attributes_count; j++ {
				attName := cf.constant_pool[cf.method[i].attributes[j].attribute_name_index]
				if string(attName.info[2:]) == "Code" {
					return cf.method[i].attributes[j]
				}
			}
		}
	}
	return
}
type stack struct{
	data	[]uint32
	tos		int
}
func (s *stack) init(size int) {
	s.data = make([]uint32,size)
	s.tos = -1
}
func (s *stack) push(u uint32) {
	s.tos = s.tos+1
	s.data[s.tos] = u
}
func (s *stack) pop()(u uint32) {
	u = s.data[s.tos]
	s.tos=s.tos-1
	return
}
func execuse(ca code_attribute,cf *classfile) {
	code:=ca.code
	s := &stack{}
	s.init(int(ca.max_stack))
	locals := make([]uint32, ca.max_locals)
	pc:=0
	for {
		op := code[pc]
			switch op{
			case iconst_1:
					s.push(1)
					pc++
			case iconst_2:
					s.push(2)
					pc++
			case iconst_3:
					s.push(3)
					pc++
			case iconst_4:
					s.push(4)
					pc++
			case iconst_5:
					s.push(5)
					pc++
			case istore:
					pc++
					index:=code[pc]
					locals[int(index)] =s.pop()
					pc++
			case istore_1:
					locals[1] =s.pop()
					pc++
			case istore_2:
					locals[2] =s.pop()
					pc++
			case istore_3:
					locals[3] =s.pop()
					pc++
			case iload_1:
					s.push(locals[1])
					pc++
			case iload_2:
					s.push(locals[2])
					pc++
			case iload_3:
					s.push(locals[3])
					pc++
			case iadd:
					o1:=s.pop()
					o2:=s.pop()
					result:=o2+o1
					s.push(result)
					pc++
			case isub:
					o1:=s.pop()
					o2:=s.pop()
					result:=o2-o1
					s.push(result)
					pc++
			case imul:
					o1 := s.pop()
					o2 := s.pop()
					result := o2*o1
					s.push(result)
					pc++
			case idiv: // Need to fixed Divided By Zero
					o1 := s.pop()
					o2 := s.pop()
						result := o2/o1
						s.push(result)
						pc++
			case bipush:
					pc++
					val:=code[pc]
					s.push(uint32(val))
					pc++
			case sipush:
						pc++
					va:=make([]byte,4)
						va[1]=code[pc] //Hight Byte
							pc++
						va[0]=code[pc] //Low Byte
					var sum uint32
					sum = binary.LittleEndian.Uint32(va[:4]) //mustbe 4 bytes
					s.push(sum)
						pc++
			case return_x:
				fmt.Println(locals)
					//fmt.Printf("3: %d\n",int(locals[3])) //special for Sub! set index to last
				pc++
				return
			}	
						//case invokevirtual:
						//	pc++
						//	ib1:=code[pc]
						//	pc++
						//	ib2:=code[pc]
						//	ib:=(ib1 << 8)+ib2
						//	cf.constant_pool[ib]
						//case getstatic:
	

		}
}

func main() {

		var	filename string	
	 	cf := &classfile{}
	
	
	if len(os.Args)==1 {
			filename ="file not found"
	}else{
		filename=os.Args[1]
		openfile(filename,cf)
		ca:=findMethod("main",cf)
		//fmt.Println("Code Attribute:",ca)
		fmt.Printf("\ncode:%x \n",ca.code)
		execuse(ca, cf)
				//lookupcode(ca.code,ca.code_length)	
			 

 	}

}

