package main

import (
	"Library-Analysis-API/Storage/filestore"
	"Library-Analysis-API/Storage/sqlstore"
	"Library-Analysis-API/datamodel"
	"fmt"
	"strconv"
	"strings"
)

var (
	exceptionIP int
)

func CreateIndex(filenameIndex string) ([]datamodel.IPIndex, error) {
	//Read File
	file, err := filestore.NewWriteInFile("IndexData/" + filenameIndex + ".csv")
	if err != nil {
		fmt.Println("read file error!")
	}
	Indexs, err := file.ReadIndex()
	if err != nil {
		fmt.Println("read file error!")
	}
	return Indexs, nil
}
func CreateNAME(FileNameIndex string) ([]datamodel.NAMEIndex, error) {
	//Read File
	file, err := filestore.NewWriteInFile("IndexData/" + FileNameIndex + ".csv")
	if err != nil {
		fmt.Println("read file error!")
	}
	IndexsName, err := file.ReadName()
	if err != nil {
		fmt.Println("read file error!")
	}
	return IndexsName, nil
}
func CreateCompanyNAME(FileNameIndex string) ([]datamodel.NAMEIndex, error) {
	//Read File
	file, err := filestore.NewWriteInFile("IndexData/" + FileNameIndex + ".csv")
	if err != nil {
		fmt.Println("read file error!")
	}
	IndexsName, err := file.ReadCompanyName()
	if err != nil {
		fmt.Println("read file error!")
	}
	return IndexsName, nil
}
func CreateCompanyIndex(fileNameIndex string) ([]datamodel.CompanyIndex, error) {
	//Read File
	file, err := filestore.NewWriteInFile("IndexData/" + fileNameIndex + ".csv")
	if err != nil {
		fmt.Println("read file error!")
	}
	CompanyIndexs, err := file.ReadCompanyIndex()
	if err != nil {
		fmt.Println("read file error!")
	}
	return CompanyIndexs, nil
}
func SplitDateBySlash(data []datamodel.Data, index []datamodel.IPIndex) ([]datamodel.DataResult, error) {
	Data := []datamodel.DataResult{}
	exceptionIP = 0
	for i := 0; i < len(data); i++ {
		var newdata datamodel.DataResult
		//var result datamodel.ResultOfMonth
		splitIP := strings.Split(data[i].LoginIP, ".")
		//splitDate := strings.Split(data[i].Date, "/")
		splitDate := strings.FieldsFunc(data[i].Date, func(r rune) bool { return r == '/' || r == '-' })
		if len(splitDate) < 3 || splitIP[0] != "140" {
			//count the IP which can not classify
			exceptionIP++
			continue
		}
		newdata.IP = splitIP[2]
		//split month
		if splitDate[1] == "1" {
			newdata.JAN++
		}
		if splitDate[1] == "2" {
			newdata.FEB++
		}
		if splitDate[1] == "3" {
			newdata.MAR++
		}
		if splitDate[1] == "4" {
			newdata.APR++
		}
		if splitDate[1] == "5" {
			newdata.MAY++
		}
		if splitDate[1] == "6" {
			newdata.JUN++
		}
		if splitDate[1] == "7" {
			newdata.JUL++
		}
		if splitDate[1] == "8" {
			newdata.AUG++
		}
		if splitDate[1] == "9" {
			newdata.SEP++
		}
		if splitDate[1] == "10" {
			newdata.OCT++
		}
		if splitDate[1] == "11" {
			newdata.NOV++
		}
		if splitDate[1] == "12" {
			newdata.DEC++
		}
		for j := 0; j < len(index); j++ {
			dataIP, err := strconv.Atoi(newdata.IP)
			if err != nil {
				fmt.Println("strconv dataIP error!", err)
			}
			indexIP, err := strconv.Atoi(index[j].IPAddress)
			if err != nil {
				fmt.Println("strconv dataIP error!", err)
			}
			if dataIP == indexIP {
				newdata.NAME = index[j].Name
				break
			}
		}

		//append
		Data = append(Data, newdata)
	}
	return Data, nil
}
func TransferDateToInt(data []datamodel.Data, index []datamodel.IPIndex) ([]datamodel.DataResult, error) {
	//Transfer date string to int. ex: "20170101" to 20170101
	Data := []datamodel.DataResult{}
	for i := 1; i < len(data); i++ {
		var newdata datamodel.DataResult
		//var result datamodel.ResultOfMonth
		splitIP := strings.Split(data[i].LoginIP, ".")
		//Transfer date string to int format
		date, err := strconv.Atoi(data[i].Date)
		if err != nil {
			fmt.Println("transfer int error !", err)
			return nil, err
		}

		date = date % 100
		newdata.IP = splitIP[2]
		//split month
		if date == 1 {
			newdata.JAN++
		}
		if date == 2 {
			newdata.FEB++
		}
		if date == 3 {
			newdata.MAR++
		}
		if date == 4 {
			newdata.APR++
		}
		if date == 5 {
			newdata.MAY++
		}
		if date == 6 {
			newdata.JUN++
		}
		if date == 7 {
			newdata.JUL++
		}
		if date == 8 {
			newdata.AUG++
		}
		if date == 9 {
			newdata.SEP++
		}
		if date == 10 {
			newdata.OCT++
		}
		if date == 11 {
			newdata.NOV++
		}
		if date == 12 {
			newdata.DEC++
		}
		for j := 0; j < len(index); j++ {
			dataIP, err := strconv.Atoi(newdata.IP)
			if err != nil {
				fmt.Println("strconv dataIP error!", err)
			}
			indexIP, err := strconv.Atoi(index[j].IPAddress)
			if err != nil {
				fmt.Println("strconv dataIP error!", err)
			}
			if dataIP == indexIP {
				newdata.NAME = index[j].Name
				break
			}
		}

		//append
		Data = append(Data, newdata)
	}
	return Data, nil
}
func CountEachMonth(data []datamodel.DataResult, nameindex []datamodel.NAMEIndex) ([]datamodel.DataResult, error) {
	// check if have same Department
	NewData := []datamodel.DataResult{}
	for i := 0; i < len(nameindex); i++ {
		var Data datamodel.DataResult
		for j := 0; j < len(data); j++ {
			if strings.TrimSpace(nameindex[i].Name) == strings.TrimSpace(data[j].NAME) {
				Data.NAME = strings.TrimSpace(nameindex[i].Name)
				Data.IP = data[j].IP
				Data.JAN = Data.JAN + data[j].JAN
				Data.FEB = Data.FEB + data[j].FEB
				Data.MAR = Data.MAR + data[j].MAR
				Data.APR = Data.APR + data[j].APR
				Data.MAY = Data.MAY + data[j].MAY
				Data.JUN = Data.JUN + data[j].JUN
				Data.JUL = Data.JUL + data[j].JUL
				Data.AUG = Data.AUG + data[j].AUG
				Data.SEP = Data.SEP + data[j].SEP
				Data.OCT = Data.OCT + data[j].OCT
				Data.NOV = Data.NOV + data[j].NOV
				Data.DEC = Data.DEC + data[j].DEC
				Data.Total = Data.JAN + Data.FEB + Data.MAR + Data.APR + Data.MAY + Data.JUN + Data.JUL + Data.AUG + Data.SEP + Data.OCT + Data.NOV + Data.DEC
			}
		}
		NewData = append(NewData, Data)
	}
	var exceptionData datamodel.DataResult
	exceptionData.NAME = "無法辨識的IP"
	exceptionData.Total = exceptionIP
	NewData = append(NewData, exceptionData)
	return NewData, nil
}
func ReadResult(database, filename string) ([]datamodel.SQLResultFormat, error) {
	//open DB
	sql, err := sqlstore.NewWriteToSQL("root", "123456", "localhost", database)
	if err != nil {
		return nil, err
	}
	//Read
	Data := []datamodel.SQLResultFormat{}
	res, err := sql.SelectAll(filename)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var data datamodel.SQLResultFormat
		err = res.Scan(&data.NAME, &data.JAN, &data.FEB, &data.MAR, &data.APR, &data.MAY, &data.JUN, &data.JUL, &data.AUG, &data.SEP, &data.OCT, &data.NOV, &data.DEC, &data.Total)
		if err != nil {
			fmt.Println("error!!!", err)
		}
		Data = append(Data, data)
	}
	res.Close()
	sql.Close()
	return Data, nil
}
func WirteInDB(Tablename, database, Date, Time string, res []datamodel.DataResult) error {
	//connect DB
	sql, err := sqlstore.NewWriteToSQL("root", "123456", "localhost", database)
	if err != nil {
		return err
	}
	//CreateTable
	option := 0
	sqlres, err := sql.CreateTable(Tablename)
	if err != nil {
		if err.Error() == "Error 1050: Table '"+Tablename+"' already exists" {
			option++
		} else {
			return err
		}
	} else {
		sqlres.Close()
	}
	//switch case0: no table, case1: table exist.
	switch option {
	case 0:
		//write in sql table
		for i := 0; i < len(res); i++ {
			//check  if result name is null
			if res[i].NAME != "" {
				res, err := sql.WriteResult(Tablename, res[i].NAME, res[i].JAN, res[i].FEB, res[i].MAR, res[i].APR, res[i].MAY, res[i].JUN, res[i].JUL, res[i].AUG, res[i].SEP, res[i].OCT, res[i].NOV, res[i].DEC, res[i].Total)
				if err != nil {
					return err
				}
				res.Close()
			}
		}
		// write fileinfo in DB fileIndex
		res, err := sql.WriteInFileIndex(Tablename, Date, Time)
		if err != nil {
			fmt.Println("WriteInFileIndex error!!", err)
			return err
		}
		res.Close()

	case 1:
		//select table's all department,if the department exist in the table,update the information. else, insert the information.
		for i := 0; i < len(res); i++ {
			if res[i].Total != 0 {
				//check department exist or not in the table
				var selectResult int
				sqlres, err := sql.ReadEachName(res[i].NAME, Tablename)
				if err != nil {
					return err
				}
				for sqlres.Next() {
					err = sqlres.Scan(&selectResult)
				}
				sqlres.Close()

				//update
				if selectResult == 1 {
					res, err := sql.UpdateResult(Tablename, res[i].NAME, res[i].JAN, res[i].FEB, res[i].MAR, res[i].APR, res[i].MAY, res[i].JUN, res[i].JUL, res[i].AUG, res[i].SEP, res[i].OCT, res[i].NOV, res[i].DEC, res[i].Total)
					if err != nil {
						fmt.Println("update in db error!", err)
						return err
					}
					res.Close()
				} else {
					//insert
					if res[i].Total != 0 {
						res, err := sql.WriteResult(Tablename, res[i].NAME, res[i].JAN, res[i].FEB, res[i].MAR, res[i].APR, res[i].MAY, res[i].JUN, res[i].JUL, res[i].AUG, res[i].SEP, res[i].OCT, res[i].NOV, res[i].DEC, res[i].Total)
						if err != nil {
							fmt.Println("write in db error!", err)
							return err
						}
						res.Close()
					}

				}

			}
		}
	}
	sql.Close()
	return nil
}
func ReadFileIndex(database string) ([]datamodel.FileIndex, error) {
	//open DB
	sql, err := sqlstore.NewWriteToSQL("root", "123456", "localhost", database)
	if err != nil {
		return nil, err
	}
	//Read
	FileIndex := []datamodel.FileIndex{}
	res, err := sql.SelectAll("FileIndex")
	if err != nil {
		return nil, err
	}
	for res.Next() {
		var fileindex datamodel.FileIndex
		err = res.Scan(&fileindex.ID, &fileindex.Name, &fileindex.Date, &fileindex.Time)
		if err != nil {
			fmt.Println("error!!!", err)
		}
		FileIndex = append(FileIndex, fileindex)
	}
	res.Close()
	sql.Close()
	return FileIndex, nil
}
func DeleteData(database, tablename string) error {
	//open DB
	sql, err := sqlstore.NewWriteToSQL("root", "123456", "localhost", database)
	if err != nil {
		return err
	}
	//Read
	res, err := sql.DeleteTable(tablename)
	if err != nil {
		return err
	}
	res.Close()
	res, err = sql.DeleteInFileIndex(tablename)
	if err != nil {
		return err
	}
	res.Close()
	sql.Close()
	return nil
}
func AnalysisForEachCompany(filename string) ([]datamodel.AnalysisResult, error) {

	//Read
	Data := []datamodel.AnalysisResult{}
	//open DB
	for i := 2010; i < 2020; i++ {
		sql, err := sqlstore.NewWriteToSQL("root", "123456", "localhost", strconv.Itoa(i))
		if err != nil {
			continue
		}
		res, err := sql.SelectAll(filename)
		if err != nil {
			sql.Close()
			continue
		}
		var data datamodel.AnalysisResult
		data.YEARS = i
		for res.Next() {
			var temp datamodel.SQLResultFormat
			err = res.Scan(&temp.NAME, &temp.JAN, &temp.FEB, &temp.MAR, &temp.APR, &temp.MAY, &temp.JUN, &temp.JUL, &temp.AUG, &temp.SEP, &temp.OCT, &temp.NOV, &temp.DEC, &temp.Total)
			if err != nil {
				fmt.Println("error!!!", err)
				return nil, err
			}
			data.JAN = temp.JAN + data.JAN
			data.FEB = temp.FEB + data.FEB
			data.MAR = temp.MAR + data.MAR
			data.APR = temp.APR + data.APR
			data.MAY = temp.MAY + data.MAY
			data.JUN = temp.JUN + data.JUN
			data.JUL = temp.JUL + data.JUL
			data.AUG = temp.AUG + data.AUG
			data.SEP = temp.SEP + data.SEP
			data.OCT = temp.OCT + data.OCT
			data.NOV = temp.NOV + data.NOV
			data.DEC = temp.DEC + data.DEC
			data.Total = temp.Total + data.Total
		}
		Data = append(Data, data)
		res.Close()
		sql.Close()
	}
	return Data, nil
}
func AnalysisForEachYear() ([]datamodel.AnalysisResult, error) {

	TotalOfYear := []datamodel.AnalysisResult{}

	for years := 2010; years < 2020; years++ {
		var TableIndex []string
		sql, err := sqlstore.NewWriteToSQL("root", "123456", "localhost", strconv.Itoa(years))
		if err != nil {
			continue
		}
		//search all table in the database
		res, err := sql.ReadName("FileIndex")
		if err != nil {
			sql.Close()
			return nil, err
		}

		for res.Next() {
			var Tablename string
			err = res.Scan(&Tablename)
			if err != nil {
				fmt.Println("error!!!", err)
				sql.Close()
				return nil, err
			}
			TableIndex = append(TableIndex, Tablename)
		}
		res.Close()
		sql.Close()

		//append all tables totals
		var temp datamodel.AnalysisResult
		temp.YEARS = years
		temp.NAME = "TOTAL"
		for i := 0; i < len(TableIndex); i++ {
			eachCompanyTotal, _ := AnalysisForEachCompany(TableIndex[i])

			for j := 0; j < len(eachCompanyTotal); j++ {
				if eachCompanyTotal[j].YEARS == years {
					temp.Total = temp.Total + eachCompanyTotal[j].Total
				}
			}
		}
		TotalOfYear = append(TotalOfYear, temp)
	}
	return TotalOfYear, nil
}
func AnalysisForEachDepartment(years string, department []datamodel.NAMEIndex) ([]datamodel.AnalysisResult, error) {

	TotalOfYear := []datamodel.AnalysisResult{}
	var TableIndex []string
	sql, err := sqlstore.NewWriteToSQL("root", "123456", "localhost", years)
	if err != nil {
		return nil, err
	}
	//search all table in the database
	res, err := sql.ReadName("FileIndex")
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var Tablename string
		err = res.Scan(&Tablename)
		if err != nil {
			fmt.Println("select fileindex error!!!", err)
			return nil, err
		}
		TableIndex = append(TableIndex, Tablename)
	}
	res.Close()

	//append all tables totals
	for i := 0; i < len(department); i++ {
		var temp datamodel.AnalysisResult
		for j := 0; j < len(TableIndex); j++ {
			res, err := sql.ReadEachDepartmentTotal(TableIndex[j], strings.TrimSpace(department[i].Name))
			if err != nil {
				fmt.Println(err)
				continue
			}
			//defer res.Close()
			if res.Next() {
				var sqlrestemp datamodel.SQLResultFormat
				err = res.Scan(&sqlrestemp.Total)
				if err != nil {
					fmt.Println("error!!!", err)
				}
				temp.Total = temp.Total + sqlrestemp.Total
			}
			res.Close()
		}
		temp.NAME = department[i].Name
		TotalOfYear = append(TotalOfYear, temp)
	}
	sql.Close()

	return TotalOfYear, nil
}
func InputResult1(filename, Date, Time, Database string, index []datamodel.IPIndex, Nameindex []datamodel.NAMEIndex) error {
	// Input datas Date is string , and split by "/".  ex: "2017/01/01" -> 2017,01,01
	//Read File
	file, err := filestore.NewWriteInFile("Data/" + filename + ".csv")
	if err != nil {
		return err
	}

	data, err := file.ReadData()
	if err != nil {
		return err
	}
	Data, err := SplitDateBySlash(data, index)
	if err != nil {
		return err
	}
	res, err := CountEachMonth(Data, Nameindex)
	if err != nil {
		return err
	}
	_ = WirteInDB(filename, Database, Date, Time, res)

	return nil
}
func InputResult2(filename, Date, Time, Database string, index []datamodel.IPIndex, Nameindex []datamodel.NAMEIndex) error {
	// Input data  Dates is string and only year and month. Change string to int first and divide 100, used remainder to count. ex:201701 % 100 = 1
	//Read File
	file, err := filestore.NewWriteInFile("Data/" + filename + ".csv")
	if err != nil {
		return err
	}
	data, err := file.ReadData()
	if err != nil {
		return err
	}
	Data, err := TransferDateToInt(data, index)
	if err != nil {
		return err
	}
	res, err := CountEachMonth(Data, Nameindex)
	if err != nil {
		return err
	}
	//CreateTable
	_ = WirteInDB(filename, Database, Date, Time, res)
	return nil
}
func InputResult3(filename, Date, Time, Database string, index []datamodel.IPIndex, Nameindex []datamodel.NAMEIndex) error {
	// Input data is already have each month. Just count result.
	//Read File
	file, err := filestore.NewWriteInFile("Data/" + filename + ".csv")
	if err != nil {
		return err
	}
	data, err := file.ReadResult(index)
	if err != nil {
		return err
	}
	res, err := CountEachMonth(data, Nameindex)
	// Change result to json format
	if err != nil {
		return err
	}
	//CreateTable
	_ = WirteInDB(filename, Database, Date, Time, res)
	return nil
}
func InputResult4(filename, Date, Time, month, Database string, index []datamodel.IPIndex, Nameindex []datamodel.NAMEIndex, CompanyIndex []datamodel.CompanyIndex, CompanyNameindex []datamodel.NAMEIndex) error {
	// Input data is already have each month. Just count result.
	//open DB
	sql, err := sqlstore.NewWriteToSQL("root", "123456", "localhost", Database)
	if err != nil {
		return err
	}
	//Read File
	file, err := filestore.NewWriteInFile("Data/" + filename + ".csv")
	if err != nil {
		fmt.Println("read file error!")
	}
	//read log
	data, err := file.ReadLogOfComputerCenter()
	if err != nil {
		fmt.Println("read file error!")
	}

	//split log destination ip. ex: [111.111.111.111] ->[111.111] just take first and second.
	for i := 0; i < len(data); i++ {
		Ipsplit := strings.Split(data[i].DestinationAddress, ".")
		IPsplitInt1, _ := strconv.Atoi(Ipsplit[0])
		IPsplitInt2, _ := strconv.Atoi(Ipsplit[1])
		//match CompanyName and write in datamodel
		for j := 0; j < len(CompanyIndex); j++ {
			if IPsplitInt1 == CompanyIndex[j].IPAddress[0] {
				if IPsplitInt2 == CompanyIndex[j].IPAddress[1] {
					data[i].CompanyName = CompanyIndex[j].CompanyName
				}

			}
		}

	}
	//split each company and count it result.
	//split
	for i := 0; i < len(CompanyNameindex); i++ {
		EachCompany := []datamodel.Data{}
		var Company string
		for j := 0; j < len(data); j++ {
			if strings.TrimSpace(CompanyNameindex[i].Name) == strings.TrimSpace(data[j].CompanyName) {
				Company = CompanyNameindex[i].Name
				var eachcompany datamodel.Data
				eachcompany.Date = data[j].Date
				eachcompany.LoginIP = data[j].SourceAddress
				EachCompany = append(EachCompany, eachcompany)
			}
		}
		//count.
		//if TableName == null,mean did't match any companyName,and don't do any count function.
		if Company != "" {
			tempResult, err := SplitDateBySlash(EachCompany, index)
			if err != nil {
				fmt.Println("SplitDateBySlash error!", err)
				return err
			}
			res, err := CountEachMonth(tempResult, Nameindex)
			if err != nil {
				fmt.Println("CountEachMonth error!", err)
				return err
			}

			//create table name
			TableName := Company
			//CreateTable
			option := 0
			sqlresponse, err := sql.CreateTable(TableName)
			if err != nil {
				if err.Error() == "Error 1050: Table '"+TableName+"' already exists" {
					option++
				} else {
					return err
				}
			} else {
				sqlresponse.Close()
			}
			//switch case0: no table,case1: table exist.
			switch option {
			case 0:
				//write in sql table
				for i := 0; i < len(res); i++ {
					//check  if result name is null
					if res[i].NAME != "" {
						sqlresponse, err = sql.WriteResult(TableName, res[i].NAME, res[i].JAN, res[i].FEB, res[i].MAR, res[i].APR, res[i].MAY, res[i].JUN, res[i].JUL, res[i].AUG, res[i].SEP, res[i].OCT, res[i].NOV, res[i].DEC, res[i].Total)
						if err != nil {
							fmt.Println("write in db error!", err)
							return err
						}
						sqlresponse.Close()

					}
				}
				// write fileinfo in DB fileIndex
				sqlresponse, err = sql.WriteInFileIndex(TableName, Date, Time)
				if err != nil {
					fmt.Println("WriteInFileIndex error!!", err)
					return err
				}
				sqlresponse.Close()

			case 1:
				//select table's all department
				//write in table
				//if the department exist in the table
				//if the department not exist in th e table
				for i := 0; i < len(res); i++ {
					if res[i].Total != 0 {
						//check department exist or not in th e table
						var selectResult int
						sqlres, err := sql.ReadEachName(res[i].NAME, TableName)
						if err != nil {
							return err
						}
						for sqlres.Next() {
							err = sqlres.Scan(&selectResult)
						}
						sqlres.Close()
						//update
						if selectResult == 1 {
							switch month {
							case "1":
								sqlresponse, err = sql.UpdateResultforEachMonth(TableName, "JAN", res[i].NAME, res[i].JAN)
								if err != nil {
									fmt.Println("write in db error!", err)
									return err
								}
								sqlresponse.Close()
							case "2":
								sqlresponse, err = sql.UpdateResultforEachMonth(TableName, "FEB", res[i].NAME, res[i].FEB)
								if err != nil {
									fmt.Println("write in db error!", err)
									return err
								}
								sqlresponse.Close()
							case "3":
								sqlresponse, err = sql.UpdateResultforEachMonth(TableName, "MAR", res[i].NAME, res[i].MAR)
								if err != nil {
									fmt.Println("write in db error!", err)
									return err
								}
								sqlresponse.Close()
							case "4":
								sqlresponse, err = sql.UpdateResultforEachMonth(TableName, "APR", res[i].NAME, res[i].APR)
								if err != nil {
									fmt.Println("write in db error!", err)
									return err
								}
								sqlresponse.Close()
							case "5":
								sqlresponse, err = sql.UpdateResultforEachMonth(TableName, "MAY", res[i].NAME, res[i].MAY)
								if err != nil {
									fmt.Println("write in db error!", err)
									return err
								}
								sqlresponse.Close()
							case "6":
								sqlresponse, err = sql.UpdateResultforEachMonth(TableName, "JUN", res[i].NAME, res[i].JUN)
								if err != nil {
									fmt.Println("write in db error!", err)
									return err
								}
								sqlresponse.Close()
							case "7":
								sqlresponse, err = sql.UpdateResultforEachMonth(TableName, "JUL", res[i].NAME, res[i].JUL)
								if err != nil {
									fmt.Println("write in db error!", err)
									return err
								}
								sqlresponse.Close()
							case "8":
								sqlresponse, err = sql.UpdateResultforEachMonth(TableName, "AUG", res[i].NAME, res[i].AUG)
								if err != nil {
									fmt.Println("write in db error!", err)
									return err
								}
								sqlresponse.Close()
							case "9":
								sqlresponse, err = sql.UpdateResultforEachMonth(TableName, "SEP", res[i].NAME, res[i].SEP)
								if err != nil {
									fmt.Println("write in db error!", err)
									return err
								}
								sqlresponse.Close()
							case "10":
								sqlresponse, err = sql.UpdateResultforEachMonth(TableName, "OCT", res[i].NAME, res[i].OCT)
								if err != nil {
									fmt.Println("write in db error!", err)
									return err
								}
								sqlresponse.Close()
							case "11":
								sqlresponse, err = sql.UpdateResultforEachMonth(TableName, "NOV", res[i].NAME, res[i].NOV)
								if err != nil {
									fmt.Println("write in db error!", err)
									return err
								}
								sqlresponse.Close()
							case "12":
								sqlresponse, err = sql.UpdateResultforEachMonth(TableName, "DECE", res[i].NAME, res[i].DEC)
								if err != nil {
									fmt.Println("write in db error!", err)
									return err
								}
								sqlresponse.Close()
							}

						} else {
							//insert
							if res[i].Total != 0 {
								sqlresponse, err = sql.WriteResult(TableName, res[i].NAME, res[i].JAN, res[i].FEB, res[i].MAR, res[i].APR, res[i].MAY, res[i].JUN, res[i].JUL, res[i].AUG, res[i].SEP, res[i].OCT, res[i].NOV, res[i].DEC, res[i].Total)
								if err != nil {
									fmt.Println("write in db error!", err)
									return err
								}
								sqlresponse.Close()
							}

						}

					}
				}
			}

		}
	}
	sql.Close()
	return nil
}
