package status_test

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ExampleCode() {
	err := status.Error(codes.AlreadyExists, "该博主已招募")
	if status.Code(err) == codes.AlreadyExists {
		fmt.Println("该用户已邀请")
	}

	if s, ok := status.FromError(err); ok {
		fmt.Println("Code() ==> ", s.Code())
		fmt.Println("Err() ==> ", s.Err())
		fmt.Println("String() ==> ", s.String())
		fmt.Println("Message() ==> ", s.Message())
		fmt.Println("Details() ==> ", s.Details())
	}

	// Output:
	// 该用户已邀请
	// Code() ==>  AlreadyExists
	// Err() ==>  rpc error: code = AlreadyExists desc = 该博主已招募
	// String() ==>  rpc error: code = AlreadyExists desc = 该博主已招募
	// Message() ==>  该博主已招募
	// Details() ==>  []
}
