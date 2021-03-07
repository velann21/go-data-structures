package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	titles := []string{"duel","dule","speed","spede","deul","cars"}
	query := "sldk"
	filterdTitles := groupSimilarTitles(&titles)
	for _, v := range filterdTitles{
		for _, arrValues := range v {
			if arrValues == query{
				fmt.Println(v)
				break
			}
		}
	}

}

func groupSimilarTitles(titles *[]string)map[string][]string{
	kv := make(map[string][]string, 0)
	for _, v := range *titles{
		charArray := make([]int, 26)
		chArr := []rune(v)
		for _, chars := range chArr{
			index := chars - 'a'
			charArray[index]++
		}
		builder := strings.Builder{}
		for i := 0; i < 26; i++ {
			builder.WriteString("#")
			builder.WriteString(strconv.Itoa(charArray[i]))
		}
		result := builder.String()
		if len(kv[result]) <= 0{
			values := make([]string,0)
			kv[result] = values
		}
		existArr := kv[result]
		existArr = append(existArr, v)
		kv[result] = existArr
	}
	return kv
}
