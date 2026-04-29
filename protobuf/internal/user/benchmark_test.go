package user

import (
	"encoding/json"
	"testing"

	pb "github.com/tndgoat/protobuf/pb"
	"google.golang.org/protobuf/proto"
)

func sampleJSONUser() User {
	return User{
		FirstName: "Tung",
		LastName:  "Nguyen",
		Email:     "tung@gmail.com",
		Contact: []Contact{
			{
				PhoneNumber: "0123456789",
				Country:     "Vietnam",
			},
		},
	}
}

func sampleProtoUser() *pb.UserProtobuf {
	return &pb.UserProtobuf{
		FirstName: "Tung",
		LastName:  "Nguyen",
		Email:     "tung@gmail.com",
		Contact: []*pb.ContactProtobuf{
			{
				PhoneNumber: "0123456789",
				Country:     "Vietnam",
			},
		},
	}
}

// JSON benchmark
func BenchmarkJSON(b *testing.B) {
	u := sampleJSONUser()
	var v User

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		data, _ := json.Marshal(u)
		_ = json.Unmarshal(data, &v)
	}
}

// Protobuf benchmark
func BenchmarkProtobuf(b *testing.B) {
	u := sampleProtoUser()
	var v pb.UserProtobuf

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		data, _ := proto.Marshal(u)
		_ = proto.Unmarshal(data, &v)
	}
}
