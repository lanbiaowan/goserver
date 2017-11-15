package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const(

	HEADER = "header"
	HEADERLENGTH = 6
	DATALENGTH = 4



)

func Pack(data []byte)[]byte{

	return append(append([]byte(HEADER),IntToByte(len(data))...),data...)

}


func Unpack(data []byte,readChan chan []byte)[]byte{
	length := len(data)
	fmt.Println(length)
	i:=0
	for;i<length;i++{

		if length < i + HEADERLENGTH + DATALENGTH{
			break
		}

		if string(data[i:i+HEADERLENGTH]) == HEADER{
			messageLen := ByteToInt(data[i+HEADERLENGTH:i+HEADERLENGTH+DATALENGTH])
			if length < messageLen{
				break
			}
			message := data[i+HEADERLENGTH+DATALENGTH:i+HEADERLENGTH+DATALENGTH+messageLen]
			readChan <- message
			i = i + HEADERLENGTH + DATALENGTH -1
		}
	}
	if i == length{
		return make([]byte,0)
	}

	return data[i:]

}

func IntToByte(i int) []byte{
	x := int32(i)
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, x)

	return buf.Bytes()
}

func ByteToInt(b []byte) int{

	buf := bytes.NewBuffer(b)
	var  i int32
	binary.Read(buf,binary.BigEndian,&i)
	return int(i)
}




