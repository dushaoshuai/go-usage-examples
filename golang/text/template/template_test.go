package template_test

import (
	"fmt"
	"strings"
	"text/template"
)

func ExampleTemplate_Execute() {
	stringBuilder := new(strings.Builder)
	tmpl := template.Must(template.New("").Parse("您的活动{{.CampaignName}}已通过审核，将在{{.StartTime}}开始"))

	err := tmpl.Execute(stringBuilder, struct {
		CampaignName string
		StartTime    string
	}{
		CampaignName: "金秋赏月",
		StartTime:    "2022-12-22 09:11:45",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(stringBuilder.String())

	// Output:
	// 您的活动金秋赏月已通过审核，将在2022-12-22 09:11:45开始
}
