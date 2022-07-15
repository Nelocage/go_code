package quizPart1

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
)

func read_csv(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("无法读取文件 "+filePath, err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("无法解析 "+filePath, err)
	}
	return records
}

const (
	Literal = iota
	Operator
)

type Token struct {
	Tok  string
	Type int
}

type Parse struct {
	Source string
	ch     byte
	offset int
	err    error // 扫描过程出现的错误收集
}

func (p *Parse) nextTok() *Token {

	//如果一个函数中有返回值，则所有的代码分支，比如所有的if 都要带return
	if p.offset >= len(p.Source) || p.err != nil {
		return nil
	}

	var err error

	for p.isWhitespace(p.ch) && err == nil {
	}

	//start:=

}

func (p *Parse) nextCh() error {

	p.offset++
	if p.offset < len(p.Source) {
		p.ch = p.Source[p.offset]
		return nil
	}

	return errors.New("EOF")

}

func (p *Parse) isWhitespace(c byte) bool {
	return c == ' ' ||
		c == '\t' ||
		c == '\n' ||
		c == '\v' ||
		c == '\f' ||
		c == '\r'

}

//前面的这个括号，相当与类，将这个函数变成该种结构体的成员函数
func (p *Parse) parse() []*Token {
	println("ddd")
}

func Part01() {
	records := read_csv("1.csv")
	for index := 0; index < len(records); index++ {
		//fmt.Println(records[index])
		fmt.Printf("我是题目:%s  ", records[index][0])
		fmt.Printf("我是答案:%s  ", records[index][1])
		fmt.Println()
	}
}
