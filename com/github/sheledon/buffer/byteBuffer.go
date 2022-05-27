package buffer

import (
	"bufio"
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
		bb.Grow()
	}
	res = bb.buffer[bb.readIndex]
	bb.readIndex++
	return
}
func (bb *ReadBuffer) ReadInt64() (int64,error){
	reb := bb.Read(8)
	u := binary.BigEndian.Uint64(reb)
	return int64(u),nil
	//return int64(binary.BigEndian.Uint64(bb.Read(64))),nil
}
func (bb *ReadBuffer) Read(lens int) (reb []byte){
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
		bb.Grow()
		reb = append(reb,bb.Read(reqLen)...)
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
