package gvm
//type stack struct{
//	data	[]interface{} //[]uint32
//	tos		int
//}
//func (s *stack) init(size int) {
//	s.data = make([]interface{},size)
//	s.tos = -1
//}
//func (s *stack) push(u interface{}) {
//	s.tos = s.tos+1
//	s.data[s.tos] = u
//}
//func (s *stack) pop()(u interface{}) {
//	u = s.data[s.tos]
//	s.tos=s.tos-1
//	return
//}
func Execuse(ca code_attribute,cf *classfile) {
	code:=ca.code
	s := &stack{}
	s.init(int(ca.max_stack))
	locals := make([]interface{}, ca.max_locals)
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
					s.push(o2.(int)+o1.(int))
					pc++
			case isub:
					o1:=s.pop()
					o2:=s.pop()
					s.push(o2.(int)-o1.(int))
					pc++
			case imul:
					o1 := s.pop()
					o2 := s.pop()
					s.push(o2.(int)*o1.(int))
					pc++
			case idiv: // Need to fixed Divided By Zero
					o1 := s.pop()
					o2 := s.pop()
					s.push(o2.(int)/o1.(int))
					pc++
			case bipush:
					pc++
					val:=code[pc]
					s.push(int(val))
					pc++
			case sipush:
					pc++
					va:=make([]byte,4)
						va[1]=code[pc] //Hight Byte
							pc++
						va[0]=code[pc] //Low Byte
					var sum uint32
					sum = (binary.LittleEndian.Uint32(va[:4])) //mustbe 4 bytes
					s.push(int(sum))
					pc++
			case ldc:
				pc++
					index := code[pc]
			//	fmt.Println("index:",index)
				d := cf.constant_pool[index]
				nextindex := binary.BigEndian.Uint16(d.info[:])
			//	fmt.Println(string(cf.constant_pool[nextindex].info[2:])) 
				s.push(int(nextindex))
				pc++
			case getstatic:
				pc++
				value:=make([]byte,4)
					value[1]=code[pc]
						pc++
					value[0]=code[pc]
			//	var all uint32
			//	all = binary.LittleEndian.Uint32(value[:4])
   			//	s.push(all)
   				pc++
			case invokevirtual :
				pc++
				value:=make([]byte,4)
					value[1]=code[pc]
						pc++
					value[0]=code[pc]
				var all uint32
				ccon := cf.constant_pool
				all = binary.LittleEndian.Uint32(value[:4])	//4
			//	fmt.Println(all)
					firststep := ccon[all]	//27,28
					numfirst := binary.BigEndian.Uint16(firststep.info[:2])	//27
					numsec := binary.BigEndian.Uint16(firststep.info[2:4]) 	//28
						innumfirst := ccon[numfirst].info[:] 		//36
						innumsec :=	ccon[numsec].info[:2] 			//37
						innumsec1:= ccon[numsec].info[2:4]			//38
							innerfirst := binary.BigEndian.Uint16(innumfirst)
							innersec := binary.BigEndian.Uint16(innumsec)
							innersec1 := binary.BigEndian.Uint16(innumsec1)
			//	fmt.Println(string(cf.constant_pool[innerfirst].info[2:]), string(cf.constant_pool[innersec].info[2:]), string(cf.constant_pool[innersec1].info[2:]))
				printstream := string(ccon[innerfirst].info[2:])
				printline := string(ccon[innersec].info[2:])
				datatype := string(ccon[innersec1].info[2:])
				if (printstream=="java/io/PrintStream")&&(printline=="println")&&(datatype=="(Ljava/lang/String;)V"){
					retrived:=s.pop().(int)
					fmt.Println(string(ccon[retrived].info[2:]))
				}
				if (printstream=="java/io/PrintStream")&&(printline=="println")&&(datatype=="(I)V"){
					retrived:=s.pop().(int)
					fmt.Println(int(retrived))
				}
   				pc++	
			case return_x:
				fmt.Println(locals)
				pc++
				return
			}	
		}
}