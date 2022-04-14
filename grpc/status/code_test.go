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

	s, ok := status.FromError(err)
	if ok {
		fmt.Println(s.Code())
		fmt.Println(s.Err())
		fmt.Println(s.String())
		fmt.Println(s.Message())
		fmt.Println(s.Details())
	}

	// Output:
	// 该用户已邀请
	// AlreadyExists
	// rpc error: code = AlreadyExists desc = 该博主已招募
	// rpc error: code = AlreadyExists desc = 该博主已招募
	// 该博主已招募
	// []
}
