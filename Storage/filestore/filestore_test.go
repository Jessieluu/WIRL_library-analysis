package filestore

import (
	"errors"
	"fmt"
	"testing"
)

var (
	ErrWrite     = errors.New("write file err!!")
	ErrRead      = errors.New("read file err!!")
	ErrReadIndex = errors.New("read index err!!")
	ErrReadData  = errors.New("Read Data err!!")
)

func TestWrite(t *testing.T) {
	var filename string
	var testbyte string
	testbyte = "abcd"
	filename = "testdata/test.csv"
	w, _ := NewWriteInFile(filename)
	err := w.Write([]byte(testbyte))
	if err != nil {
		t.Error(ErrWrite)
	}
}
func TestRead(t *testing.T) {
	var filename string
	filename = "testdata/test.csv"
	w, _ := NewWriteInFile(filename)
	_, err := w.Read()
	if err != nil {
		t.Error(ErrRead)
	}
}
func TestIndex(t *testing.T) {
	var filename string
	filename = "testdata/testIPIndex.csv"
	w, _ := NewWriteInFile(filename)
	_, err := w.ReadIndex()
	if err != nil {
		fmt.Println(err)
		t.Error(ErrReadIndex)
	}
}
func TestReadName(t *testing.T) {
	var filename string
	filename = "testdata/testNameIndex.csv"
	w, _ := NewWriteInFile(filename)
	_, err := w.ReadName()
	if err != nil {
		fmt.Println(err)
		t.Error(ErrReadIndex)
	}
}
func TestReadData(t *testing.T) {
	var filename string
	filename = "testdata/testData.csv"
	w, _ := NewWriteInFile(filename)
	_, err := w.ReadData()
	if err != nil {
		fmt.Println(err)
		t.Error(ErrReadIndex)
	}
}
func TestReadResult(t *testing.T) {
	var filename string
	var Indexname string
	Indexname = "testdata/testIPIndex.csv"
	I, _ := NewWriteInFile(Indexname)
	testIndex, err := I.ReadIndex()
	if err != nil {
		fmt.Println(err)
		t.Error(ErrReadData)
	}
	filename = "testdata/testDataresult.csv"
	w, _ := NewWriteInFile(filename)
	_, err = w.ReadResult(testIndex)
	if err != nil {
		fmt.Println(err)
		t.Error(ErrReadData)
	}
}
