package fileoperations

import (
	"Test1/dao"
	"Test1/model"
	"fmt"
	"strconv"
	"strings"
)

type FileService struct {
	Filestore dao.IFileDao
}

type IFileService interface {
	ServiceOperation(test string) map[string]int
	StringServiceOperations(test string) map[string]int
}

// used db to sort
func (F FileService) ServiceOperation(test string) map[string]int {
	var wordcount []model.SelectData

	sql := fmt.Sprintf("select * from ( \n")
	union := fmt.Sprintf("\nunion\n")
	sqlend := "\n) a order by col2 desc"
	var Select []string

	for word, count := range stringToMap(test) {
		Select = append(Select, "select '"+word+"' col1, "+strconv.Itoa(count)+" col2 ")
	}

	finalsql := sql + strings.Join(Select, union) + sqlend

	wordcount = F.Filestore.StringOperations(finalsql)

	return StructMap(wordcount[:10])
}

func stringToMap(test string) map[string]int {
	rawArray := strings.Fields(test)

	wc := make(map[string]int)
	for _, word := range rawArray {
		_, matched := wc[word]
		if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}

	return wc
}

func StructMap(wordcount []model.SelectData) map[string]int {
	data := make(map[string]int)
	for _, v := range wordcount {
		data[v.Col1] = v.Col2
	}
	return data
}

// sort using go code not used db
func (F FileService) StringServiceOperations(test string) map[string]int {
	var wordcount []model.SelectData

	for word, count := range stringToMap(test) {
		wordcount = append(wordcount, model.SelectData{Col1: word, Col2: count})
	}

	for i := 0; i < len(wordcount); i++ {
		for j := 0; j < len(wordcount); j++ {
			if wordcount[i].Col2 > wordcount[j].Col2 {
				wordcount[i], wordcount[j] = wordcount[j], wordcount[i]
			}
		}
	}

	return StructMap(wordcount[:10])
}
