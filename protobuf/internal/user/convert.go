package user

import pb "github.com/tndgoat/protobuf/pb"

func ToProto(u User) *pb.UserProtobuf {
	contacts := make([]*pb.ContactProtobuf, len(u.Contact))

	for i, c := range u.Contact {
		contacts[i] = &pb.ContactProtobuf{
			PhoneNumber: c.PhoneNumber,
			Country:     c.Country,
		}
	}

	return &pb.UserProtobuf{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Contact:   contacts,
	}
}

func FromProto(p *pb.UserProtobuf) User {
	contacts := make([]Contact, len(p.Contact))

	for i, c := range p.Contact {
		contacts[i] = Contact{
			PhoneNumber: c.PhoneNumber,
			Country:     c.Country,
		}
	}

	return User{
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
		Contact:   contacts,
	}
}
