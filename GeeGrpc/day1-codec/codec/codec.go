package codec

import "io"

type Header struct {
	ServiceMethod string //format "Service.Method"
	Seq           uint64 //sequence number chosen by client
	Error         string
}

type Codec interface {
	io.Closer
	ReadHeader(header *Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

type NewCodecFunc func(closer io.ReadWriteCloser) Codec
