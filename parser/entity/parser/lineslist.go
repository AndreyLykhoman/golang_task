package parser

import (
	"math/rand"
)

type LinesList struct {
	Lines [] Line
}

func (ls *LinesList) GetNBiggest( N int) (biggest LinesList){
	if(len(ls.Lines) < N){
		N = len(ls.Lines)
	}
	for index := 0 ; index < N; index++{
		biggest.Lines =  append(biggest.Lines, ls.Lines[index])
	}
	return
}

func (ls LinesList) getLinesValue() (value string) {
	var byteResult[] byte;
	for _, line := range ls.Lines  {
		var strLine string = line.Value + "\n"
		byteResult = append(byteResult, []byte(strLine)...)
	}
	value = string(byteResult)
	return
}

func (ls LinesList) SortLinesBiggestFromLess() LinesList {
	length := len(ls.Lines)
	if length <= 1 {
		LinesListCopy := LinesList{
			make([]Line, length),
		}
		copy(LinesListCopy.Lines, ls.Lines)
		return LinesListCopy
	}
	m := ls.Lines[rand.Intn(length)]
	less := LinesList{
		make([]Line, 0, length),
	}
	middle := LinesList{
		make([]Line, 0, length),
	}
	more := LinesList{
		make([]Line, 0, length),
	}
	for _, item := range ls.Lines {
		switch {
		case item.Size < m.Size:
			less.Lines = append(less.Lines, item)
		case item.Size == m.Size:
			middle.Lines = append(middle.Lines, item)
		case item.Size > m.Size:
			more.Lines = append(more.Lines, item)
		}
	}
	less, more = less.SortLinesBiggestFromLess(), more.SortLinesBiggestFromLess()
	more.Lines = append(more.Lines, middle.Lines...)
	more.Lines = append(more.Lines, less.Lines...)
	return more
}