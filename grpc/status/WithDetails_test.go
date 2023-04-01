package status_test

import (
	"fmt"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ExampleWithDetail() {
	s := status.New(codes.ResourceExhausted, "a per-user quota has been exhausted")

	var err error
	s, err = s.WithDetails(
		&errdetails.LocalizedMessage{
			Locale:  "zh-CN",
			Message: "您已领取，下次再来吧",
		},
		&errdetails.Help{
			Links: []*errdetails.Help_Link{
				{
					Description: "gPRC - Richer error model",
					Url:         "https://grpc.io/docs/guides/error/#richer-error-model",
				},
				{
					Description: "googleapis - Error Model",
					Url:         "https://cloud.google.com/apis/design/errors#error_model",
				},
				{
					Description: "googleapis - google.rpc.Status",
					Url:         "https://github.com/googleapis/googleapis/blob/master/google/rpc/status.proto",
				},
				{
					Description: "googleapis - google.rpc.error_details",
					Url:         "https://github.com/googleapis/googleapis/blob/master/google/rpc/error_details.proto",
				},
				{
					Description: "google.golang.org/grpc/status",
					Url:         "https://pkg.go.dev/google.golang.org/grpc@v1.54.0/status",
				},
				{
					Description: "google.golang.org/genproto/googleapis/rpc/errdetails",
					Url:         "https://pkg.go.dev/google.golang.org/genproto/googleapis/rpc/errdetails",
				},
			},
		},
	)
	if err != nil {
		panic(err)
	}

	for _, detail := range s.Details() {
		switch detail := detail.(type) {
		case *errdetails.LocalizedMessage:
			fmt.Println(detail.GetLocale(), detail.GetMessage())
		case *errdetails.Help:
			for _, link := range detail.GetLinks() {
				fmt.Println(link.GetDescription())
				fmt.Println(link.GetUrl())
			}
		}
	}

	// Output:
	// zh-CN 您已领取，下次再来吧
	// gPRC - Richer error model
	// https://grpc.io/docs/guides/error/#richer-error-model
	// googleapis - Error Model
	// https://cloud.google.com/apis/design/errors#error_model
	// googleapis - google.rpc.Status
	// https://github.com/googleapis/googleapis/blob/master/google/rpc/status.proto
	// googleapis - google.rpc.error_details
	// https://github.com/googleapis/googleapis/blob/master/google/rpc/error_details.proto
	// google.golang.org/grpc/status
	// https://pkg.go.dev/google.golang.org/grpc@v1.54.0/status
	// google.golang.org/genproto/googleapis/rpc/errdetails
	// https://pkg.go.dev/google.golang.org/genproto/googleapis/rpc/errdetails
}
