package main

import (
	"log"
	"log/slog"
	"net/http"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/dushaoshuai/go-usage-examples/protobuf/httpx"
	protox "github.com/dushaoshuai/go-usage-examples/protobuf/httpx/proto"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		out, err := proto.Marshal(&protox.AddressBook{
			People: []*protox.Person{
				{
					Name:  "dss",
					Id:    1,
					Email: "okmy@email.com",
					Phones: []*protox.Person_PhoneNumber{
						{
							Number: "11111111111",
							Type:   protox.PhoneType_PHONE_TYPE_HOME,
						},
					},
					LastUpdated: timestamppb.Now(),
				},
			},
		})
		if err != nil {
			slog.Error(err.Error())
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(out)
	})

	log.Fatal(http.ListenAndServe(httpx.Port, nil))
}
