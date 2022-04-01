package simplifiedchinese_test

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func Example() {
	GB18030 := simplifiedchinese.GB18030
	data, _ := ioutil.ReadAll(transform.NewReader(strings.NewReader("中国话"), GB18030.NewEncoder()))
	fmt.Printf("%s\n", data)

	ss := []string{"中", "世纪", "啊", "日本", "米雪", "管理"}
	sort.Strings(ss)
	fmt.Println(ss)
	// Output:
}
