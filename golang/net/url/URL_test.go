package url_test

import (
	"fmt"
	"net/url"
)

func ExampleURL() {
	localFile := "file://localhost/home/shaouai/%E8%A7%86%E9%A2%91/obs/Life%20Is%20Strange%20-%20Episode%201%20-%20Chrysalis.mp4?age=24&u=shaouai#foo_fragment"

	u, err := url.Parse(localFile)
	if err != nil {
		panic(err)
	}

	fmtURL(u)

	// Output:
	// (*url.URL).IsAbs() ==> true
	// (*url.URL).Hostname() ==> "localhost"
	// (*url.URL).Port() ==> ""
	// (*url.URL).EscapedPath() ==> "/home/shaouai/%E8%A7%86%E9%A2%91/obs/Life%20Is%20Strange%20-%20Episode%201%20-%20Chrysalis.mp4"
	// (*url.URL).Query().Encode() ==> "age=24&u=shaouai"
	// (*url.URL).EscapedFragment() ==> "foo_fragment"
	// (*url.URL).String() ==> "file://localhost/home/shaouai/%E8%A7%86%E9%A2%91/obs/Life%20Is%20Strange%20-%20Episode%201%20-%20Chrysalis.mp4?age=24&u=shaouai#foo_fragment"
	// (*url.URL).MarshalBinary() ==> "file://localhost/home/shaouai/%E8%A7%86%E9%A2%91/obs/Life%20Is%20Strange%20-%20Episode%201%20-%20Chrysalis.mp4?age=24&u=shaouai#foo_fragment"
}

func fmtURL(u *url.URL) {
	fmt.Printf("(*url.URL).IsAbs() ==> %v\n", u.IsAbs())

	fmt.Printf("(*url.URL).Hostname() ==> %q\n", u.Hostname())
	fmt.Printf("(*url.URL).Port() ==> %q\n", u.Port())
	fmt.Printf("(*url.URL).EscapedPath() ==> %q\n", u.EscapedPath())
	fmt.Printf("(*url.URL).Query().Encode() ==> %q\n", u.Query().Encode())
	fmt.Printf("(*url.URL).EscapedFragment() ==> %q\n", u.EscapedFragment())

	fmt.Printf("(*url.URL).String() ==> %q\n", u.String())
	bURL, _ := u.MarshalBinary()
	fmt.Printf("(*url.URL).MarshalBinary() ==> %q\n", bURL)
}
