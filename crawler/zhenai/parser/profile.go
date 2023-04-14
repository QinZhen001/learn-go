package parser

import (
	"fmt"
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"regexp"
	"strconv"
)

const ageRE = `<td><span class="label">年龄：</span>([\d]+)岁</td>`

func ParseProfile(contents []byte, name string) engine.ParseResult {
	re := regexp.MustCompile(ageRE)
	profile := model.Profile{}
	profile.Name = name
	matches := re.FindSubmatch(contents)
	if len(matches) >= 2 {
		age, err := strconv.Atoi(string(matches[1]))
		if err != nil {
			profile.Age = age
		}
	}

	fmt.Printf("111 %+v\n", profile)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result

}
