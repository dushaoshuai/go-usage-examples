package url_test

import (
	"net/url"
)

func Example_request_URI() {
	requestURI := "https://github.com/gin-gonic/gin/issues?q=is%3Aissue+is%3Aclose"

	u, err := url.ParseRequestURI(requestURI)
	if err != nil {
		panic(err)
	}

	fmtURL(u)

	// Output:
	// (*url.URL).IsAbs() ==> true
	// (*url.URL).Hostname() ==> "github.com"
	// (*url.URL).Port() ==> ""
	// (*url.URL).EscapedPath() ==> "/gin-gonic/gin/issues"
	// (*url.URL).Query().Encode() ==> "q=is%3Aissue+is%3Aclose"
	// (*url.URL).EscapedFragment() ==> ""
	// (*url.URL).String() ==> "https://github.com/gin-gonic/gin/issues?q=is%3Aissue+is%3Aclose"
	// (*url.URL).MarshalBinary() ==> "https://github.com/gin-gonic/gin/issues?q=is%3Aissue+is%3Aclose"
}
