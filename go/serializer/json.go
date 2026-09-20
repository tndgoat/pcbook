package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ProtobufToJSON serializes a protobuf message to a JSON string.
func ProtobufToJSON(message proto.Message) (string, error) {
	opts := protojson.MarshalOptions{
		UseEnumNumbers:  false, // true = serialize enum values as numbers, false = serialize enum values as strings
		EmitUnpopulated: false, // true = include fields with zero values, false = omit fields with zero values
		UseProtoNames:   true,  // true = use the original proto field names, false = use the JSON field names
		Indent:          "  ",  // pretty print the JSON output with indentation
	}

	bytes, err := opts.Marshal(message)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
