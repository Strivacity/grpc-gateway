package runtime

import (
	"google.golang.org/grpc/metadata"
)

// LocalTransportStream provides
type LocalTransportStream struct {
	method string
	md     ServerMetadata
}

// NewLocalTransportStream creates a transport stream
// to provide Metadata handling in local requests
// it implements grpc.ServerTransportStream
func NewLocalTransportStream(method string) *LocalTransportStream {
	return &LocalTransportStream{
		method: method,
		md: ServerMetadata{
			HeaderMD:  make(metadata.MD),
			TrailerMD: make(metadata.MD),
		},
	}
}

// Metadata returns the underlying ServerMetadata of the stream
func (t LocalTransportStream) Metadata() ServerMetadata {
	return t.md
}

// Method returns the called gRPC method
func (t LocalTransportStream) Method() string { return t.method }

// SetHeader appends md to the Header metadara
func (t LocalTransportStream) SetHeader(md metadata.MD) error {
	for k, v := range md {
		t.md.HeaderMD.Append(k, v...)
	}
	return nil
}

// SendHeader does nothing as local metadata will be extracted by the generated code
func (t LocalTransportStream) SendHeader(md metadata.MD) error {
	return nil
}

// SetTrailer appends md to the Trailer metadara
func (t LocalTransportStream) SetTrailer(md metadata.MD) error {
	for k, v := range md {
		t.md.TrailerMD.Append(k, v...)
	}
	return nil
}
