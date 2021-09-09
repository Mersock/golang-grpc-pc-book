package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJson(message proto.Message) (string, error) {
	marshaler := protojson.MarshalOptions{
		Indent:          "\t",
		UseProtoNames:   false,
		EmitUnpopulated: true,
	}
	b, err := marshaler.Marshal(message)
	return string(b), err
}
