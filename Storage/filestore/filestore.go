package filestore

import (
	"Library-Analysis-API/datamodel"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

type WriteInFile struct {
	filename string
}

func NewWriteInFile(filename string) (WriteInFile, error) {
	return WriteInFile{filename: filename}, nil
}

func (w WriteInFile) Write(b []byte) error {

	err := ioutil.WriteFile(w.filename, b, 0644)
	if err != nil {
		fmt.Println("write file error!!", err)
	}

	return nil
}
func (w WriteInFile) Read() ([]byte, error) {
	b, err := ioutil.ReadFile(w.filename)
	if err != nil {
		fmt.Println("open file error!")
	}
	tranB, _, _ := transform.Bytes(traditionalchinese.Big5.NewDecoder(), b)
	return tranB, nil
}
func (w WriteInFile) ReadIndex() ([]datamodel.IPIndex, error) {

	b, err := ioutil.ReadFile(w.filename)
	if err != nil {
		fmt.Println("open file error!")
	}
	tranB, _, _ := transform.Bytes(traditionalchinese.Big5.NewDecoder(), b)
	str := strings.Split(string(tranB), "\n")
	Indexs := []datamodel.IPIndex{}
	for i := 0; i < len(str); i++ {
		var index datamodel.IPIndex
		str1 := str[i]
		strsplit := strings.Split(str1, ",")
		IPsplit := strings.Split(strsplit[0], ".")
		index.IPAddress = IPsplit[2]
		index.Name = strsplit[1]
		Indexs = append(Indexs, index)
	}
	return Indexs, nil
}
func (w WriteInFile) ReadCompanyIndex() ([]datamodel.CompanyIndex, error) {

	b, err := ioutil.ReadFile(w.filename)
	if err != nil {
		fmt.Println("open file error!")
	}
	//companyindex file coding with utf-8, and bytes should split "\r"（carriage return） not "\n"（line feed）.
	str := strings.Split(string(b), "\r")
	Indexs := []datamodel.CompanyIndex{}
	for i := 0; i < len(str); i++ {
		var index datamodel.CompanyIndex
		str1 := str[i]
		strsplit := strings.Split(str1, ",")
		//split IP
		IPsplit := strings.Split(strsplit[1], ".")
		index.CompanyName = strsplit[0]
		for i := 0; i < 2; i++ {
			IPsplitInt, err := strconv.Atoi(IPsplit[i])
			if err != nil {
				fmt.Println("Company IP transfer to int error!!", err)
			}
			index.IPAddress = append(index.IPAddress, IPsplitInt)
		}
		Indexs = append(Indexs, index)
	}
	return Indexs, nil
}
func (w WriteInFile) ReadName() ([]datamodel.NAMEIndex, error) {

	b, err := ioutil.ReadFile(w.filename)
	if err != nil {
		fmt.Println("open file error!")
	}
	tranB, _, _ := transform.Bytes(traditionalchinese.Big5.NewDecoder(), b)
	str := strings.Split(string(tranB), "\n")
	NAMEIndexs := []datamodel.NAMEIndex{}
	for i := 0; i < len(str); i++ {
		var index datamodel.NAMEIndex
		str1 := str[i]
		index.Name = str1
		NAMEIndexs = append(NAMEIndexs, index)
	}
	return NAMEIndexs, nil
}
func (w WriteInFile) ReadCompanyName() ([]datamodel.NAMEIndex, error) {

	b, err := ioutil.ReadFile(w.filename)
	if err != nil {
		fmt.Println("open file error!")
	}
	str := strings.Split(string(b), "\n")
	NAMEIndexs := []datamodel.NAMEIndex{}
	for i := 0; i < len(str); i++ {
		var index datamodel.NAMEIndex
		str1 := str[i]
		index.Name = str1
		NAMEIndexs = append(NAMEIndexs, index)
	}
	return NAMEIndexs, nil
}
func (w WriteInFile) ReadData() ([]datamodel.Data, error) {

	b, err := ioutil.ReadFile(w.filename)
	if err != nil {
		fmt.Println("open file error!")
	}
	tranB, _, err := transform.Bytes(traditionalchinese.Big5.NewDecoder(), b)
	if err != nil {
		fmt.Println("transfer big5 format error!!", err)
		return nil, err
	}
	str := strings.Split(string(tranB), "\n")
	Data := []datamodel.Data{}
	for i := 1; i < len(str)-1; i++ {
		var data datamodel.Data
		str1 := str[i]
		strsplit := strings.Split(str1, ",")
		data.Date = strsplit[0]
		data.LoginIP = strsplit[1]
		Data = append(Data, data)
	}
	return Data, nil
}
func (w WriteInFile) ReadLogOfComputerCenter() ([]datamodel.LogOfComputerCenter, error) {

	b, err := ioutil.ReadFile(w.filename)
	if err != nil {
		fmt.Println("open file error!")
	}
	str := strings.Split(string(b), "\n")
	Data := []datamodel.LogOfComputerCenter{}
	for i := 1; i < len(str); i++ {
		var data datamodel.LogOfComputerCenter
		str1 := str[i]
		strsplit := strings.Split(str1, ",")
		//run in Windows OS has the problem : some str after split just read 1 elements.
		if len(strsplit) == 3 {
			data.Date = strsplit[0]
			data.SourceAddress = strsplit[1]
			data.DestinationAddress = strsplit[2]
			Data = append(Data, data)
		}
	}
	return Data, nil
}
func (w WriteInFile) ReadResult(index []datamodel.IPIndex) ([]datamodel.DataResult, error) {
	b, err := ioutil.ReadFile(w.filename)
	if err != nil {
		fmt.Println("open file error!")
	}
	str := strings.Split(string(b), "\n")
	Data := []datamodel.DataResult{}
	for i := 1; i < len(str)-1; i++ {
		var data datamodel.DataResult
		str1 := str[i]
		strsplit := strings.Split(str1, ",")
		Ipsplit := strings.Split(strsplit[0], ".")
		//Some input data has mac address >> ignore it !!
		if len(Ipsplit) != 4 {
			continue
		}
		data.IP = Ipsplit[2]
		//Use index model change IP To Name
		for j := 0; j < len(index); j++ {
			dataIP, err := strconv.Atoi(data.IP)
			if err != nil {
				fmt.Println("strconv dataIP error!", err)
			}
			indexIP, err := strconv.Atoi(index[j].IPAddress)
			if err != nil {
				fmt.Println("strconv dataIP error!", err)
			}
			if dataIP == indexIP {
				data.NAME = index[j].Name
				break
			}
		}
		//str month change to int
		for i := 1; i < len(strsplit); i++ {
			var month int
			if strsplit[i] == "" {
				month = 0
			} else {
				x, err := strconv.Atoi(strings.TrimSpace(strsplit[i]))
				if err != nil {
					fmt.Println("change string to int in month error!!", err)
					return nil, err
				}
				month = x
			}
			//distinguish which month and write to datamodel
			if i == 1 {
				data.JAN = month
			}
			if i == 2 {
				data.FEB = month
			}
			if i == 3 {
				data.MAR = month
			}
			if i == 4 {
				data.APR = month
			}
			if i == 5 {
				data.MAY = month
			}
			if i == 6 {
				data.JUN = month
			}
			if i == 7 {
				data.JUL = month
			}
			if i == 8 {
				data.AUG = month
			}
			if i == 9 {
				data.SEP = month
			}
			if i == 10 {
				data.OCT = month
			}
			if i == 11 {
				data.NOV = month
			}
			if i == 12 {
				data.DEC = month
			}
		}
		data.Total = data.JAN + data.FEB + data.MAR + data.APR + data.MAR + data.JUN + data.JUL + data.AUG + data.SEP + data.OCT + data.NOV + data.DEC
		Data = append(Data, data)
	}

	return Data, nil
}
