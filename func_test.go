package main

import (
	"Library-system-API/datamodel"
	"errors"
	"testing"
)

var (
	ErrIPIndex        = errors.New("open IPIndex file err!!")
	ErrNameIndex      = errors.New("open NameIndex file err!!")
	ErrSplitData      = errors.New("split original Data err!!")
	ErrCountEachMonth = errors.New("count each month err!!")
)

func TestCreateIndex(t *testing.T) {
	var indexname string
	indexname = "IPIndex"
	_, err := CreateIndex(indexname)
	if err != nil {
		t.Error(ErrIPIndex)
	}

}
func TestCreateNameIndex(t *testing.T) {
	var indexname string
	indexname = "NameIndex"
	_, err := CreateNAME(indexname)
	if err != nil {
		t.Error(ErrNameIndex)
	}
}
func TestSplitDateBySlash(t *testing.T) {
	var indexname string
	testdata := []datamodel.Data{}
	indexname = "IPIndex"
	for i := 0; i < 3; i++ {
		data := datamodel.Data{}
		data.Date = "2017/01/01"
		data.LoginIP = "100.100.100.100"
		testdata = append(testdata, data)
	}

	index, err := CreateIndex(indexname)
	if err != nil {
		t.Error(ErrIPIndex)
	}
	_, err = SplitDateBySlash(testdata, index)
	if err != nil {
		t.Error(ErrSplitData)
	}
}
func TestTransferDateToInt(t *testing.T) {
	var indexname string
	testdata := []datamodel.Data{}
	indexname = "IPIndex"
	for i := 0; i < 3; i++ {
		data := datamodel.Data{}
		data.Date = "20170101"
		data.LoginIP = "100.100.100.100"
		testdata = append(testdata, data)
	}

	index, err := CreateIndex(indexname)
	if err != nil {
		t.Error(ErrIPIndex)
	}
	_, err = TransferDateToInt(testdata, index)
	if err != nil {
		t.Error(ErrSplitData)
	}
}
func TestCountEachMonth(t *testing.T) {
	var indexname string
	testdata := []datamodel.DataResult{}
	indexname = "NameIndex"
	indexname = "IPIndex"
	for i := 0; i < 3; i++ {
		data := datamodel.DataResult{}
		data.NAME = "資工系"
		data.IP = "100"
		data.JAN = 1
		data.FEB = 1
		data.MAR = 1
		data.APR = 1
		data.MAY = 1
		data.JUN = 1
		data.JUL = 1
		data.APR = 1
		data.SEQ = 1
		data.OCT = 1
		data.NOV = 1
		data.DEC = 1
		testdata = append(testdata, data)
	}

	testindex, err := CreateNAME(indexname)
	if err != nil {
		t.Error(ErrNameIndex)
	}
	_, err = CountEachMonth(testdata, testindex)
	if err != nil {
		t.Error(ErrCountEachMonth)
	}
}
