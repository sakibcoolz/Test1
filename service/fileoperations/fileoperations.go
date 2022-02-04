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
	FileServiceOperation(test string) (wordcount []model.SelectData)
}

func (F FileService) FileServiceOperation(test string) (wordcount []model.SelectData) {
	// top := 10

	rawArray := strings.Fields(test)

	sql := fmt.Sprintf("select * from ( \n")
	union := fmt.Sprintf("\nunion\n")
	sqlend := "\n) a order by col2 desc limit 10"
	var Select []string

	wc := make(map[string]int)
	for _, word := range rawArray {
		_, matched := wc[word]
		if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}

	for word, count := range wc {
		Select = append(Select, "select '"+word+"' col1, "+strconv.Itoa(count)+" col2 ")
	}

	finalsql := sql + strings.Join(Select, union) + sqlend

	wordcount = F.Filestore.StringOperations(finalsql)

	return wordcount[:10]
}
