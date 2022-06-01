package buffer

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
)

type ReadBuffer struct {
	reader *bufio.Reader
	buffer []byte
	readIndex int
	len int
}

func NewReadBuffer(reader io.Reader) *ReadBuffer{
	return &ReadBuffer{
		reader: bufio.NewReader(reader),
	}
}
func (bb *ReadBuffer) ReadByte() (res byte,err error){
	if bb.readIndex >= bb.len {
		err = bb.Grow()
		if err != nil {
			return
		}
	}
	res = bb.buffer[bb.readIndex]
	bb.readIndex++
	return
}
func (bb *ReadBuffer) ReadInt64() (int64,error){
	reb,err := bb.Read(8)
	if err != nil {
		return 0,err
	}
	bytesBuffer := bytes.NewBuffer(reb)
	var x int64
	err = binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x,err
}
func (bb *ReadBuffer) Read(lens int) (reb []byte,err error){
	reqLen := lens
	if bb.IsReadable() {
		start,end := bb.readIndex ,bb.readIndex + lens
		if lens > bb.ReadableBytes() {
			reqLen -= bb.ReadableBytes()
			end = bb.len
		} else {
			reqLen = 0
		}
		reb = append(reb,bb.buffer[start:end]...)
		bb.readIndex = end
	}
	if reqLen > 0{
		err = bb.Grow()
		if err != nil {
			return
		}
		bs, errs := bb.Read(reqLen)
		if errs != nil {
			return nil,errs
		}
		reb = append(reb,bs...)
	}
	return
}
func (bb *ReadBuffer) IsReadable() bool {
	return bb.readIndex < bb.len-1
}
func (bb *ReadBuffer) ReadableBytes() int {
	return bb.len - bb.readIndex
}
func (bb *ReadBuffer) Grow() error{
	var buf [1024]byte
	n, err := bb.reader.Read(buf[:])
	if err != nil {
		return err
	}
	bb.readIndex = 0
	bb.len = n
	bb.buffer = buf[:n]
	return nil
}
