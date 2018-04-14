package main

import ( // "encoding/json"

	"Library-Analysis-API/datamodel"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	port = "8066"

	index                      []datamodel.IPIndex
	nameindex                  []datamodel.NAMEIndex
	departmentindexforanalysis []datamodel.NAMEIndex
	companyIndex               []datamodel.CompanyIndex
	companynameindex           []datamodel.NAMEIndex
	healthy                    datamodel.Health
	departmentName             = "NameIndex"
	departmentNameForAnalysis  = "NameDepartment"
	filenameIndex              = "IPIndex"
	filecompanyIndex           = "CompanyIndex"
	companyNameIndex           = "CompanyNameIndex"
	dates                      string
	times                      string
	filename                   string
	database                   string
	panicFlag                  bool
)

func init() {
	// open IPindex in memory
	indexs, err := CreateIndex(filenameIndex)
	if err != nil {
		fmt.Println("Create IPIndex Error!!", err)
	}
	index = indexs
	// open Nameindex in memory
	name, err := CreateNAME(departmentName)
	if err != nil {
		fmt.Println("Create IPIndex Error!!", err)
	}
	nameindex = name
	// open DepartmentName in memory
	nameforanalysis, err := CreateNAME(departmentNameForAnalysis)
	if err != nil {
		fmt.Println("Create IPIndex Error!!", err)
	}
	departmentindexforanalysis = nameforanalysis
	//open Companyindex in memory
	companyindexs, err := CreateCompanyIndex(filecompanyIndex)
	if err != nil {
		panic("Create NameIndex Error!!" + err.Error())
	}
	companyIndex = companyindexs
	//open CompanyNameIndex in memory
	companyname, err := CreateCompanyNAME(companyNameIndex)
	if err != nil {
		panic("Create NameIndex Error!!" + err.Error())
	}
	companynameindex = companyname

	//test function
}
func main() {

	//http server
	myFunction := func() {
		//handle
		http.HandleFunc("/GetData", DataGet)
		http.HandleFunc("/GetData/Analysis", DataAnalysis)
		http.HandleFunc("/GetData/YearAnalysis", DataYearAnalysis)
		http.HandleFunc("/GetData/DepartmentAnalysis", DepartmentYearAnalysis)
		http.HandleFunc("/FileIndex", FileIndexGet)
		http.HandleFunc("/DataDelete", DeleteTable)
		http.HandleFunc("/DataDateString", DataPost1)
		http.HandleFunc("/DataDateInt", DataPost2)
		http.HandleFunc("/DataIPandMonth", DataPost3)
		http.HandleFunc("/Log", DataPost4)
		http.HandleFunc("/healthy", Health)

		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			panic("Connect Fail:" + err.Error())
		}
	}
	fmt.Println("API-Start")
	go myFunction()
	// use go channel to continous code
	endChannel := make(chan os.Signal)
	signal.Notify(endChannel)
	sig := <-endChannel
	fmt.Println("END!:", sig)
}

// http function
// get db data with json format
func DataGet(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if req.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		healthy.GETDataFAIL++
		return
	}
	//open Timer
	year, month, day := time.Now().Date()
	hour, min, sec := time.Now().Clock()
	dates = fmt.Sprintf("%d-%d-%d", year, int(month), day)
	times = fmt.Sprintf("%d:%d:%d", hour, min, sec)
	//set filename
	database := req.FormValue("database")
	Filename := req.FormValue("filename")
	fmt.Printf("SearchFile file:%s database:%s %s %s\n", filename, database, dates, times)
	//call read function
	Data, err := ReadResult(database, Filename)
	if err != nil {
		fmt.Println("ReadResult error!!", err, dates, times)
		healthy.GETDataFAIL++
	}
	//encoding json
	b, err := json.Marshal(Data)
	if err != nil {
		fmt.Println("encoding json error!!", err, dates, times)
		healthy.GETDataFAIL++
	}
	fmt.Fprint(w, string(b))
	healthy.GETDataSUCCESS++
}

//get fileindex from DB
func FileIndexGet(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if req.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		healthy.GETFileIndexFAIL++
		return
	}
	//open Timer
	year, month, day := time.Now().Date()
	hour, min, sec := time.Now().Clock()
	dates = fmt.Sprintf("%d-%d-%d", year, int(month), day)
	times = fmt.Sprintf("%d:%d:%d", hour, min, sec)
	//set filename
	database := req.FormValue("database")
	fmt.Printf("SearchFileIndex database:%s %s %s\n", database, dates, times)
	//call read function
	Data, err := ReadFileIndex(database)
	if err != nil {
		fmt.Println("ReadResult error!!", err, dates, times)
		healthy.GETFileIndexFAIL++
	}
	//encoding json
	b, err := json.Marshal(Data)
	if err != nil {
		fmt.Println("encoding json error!!", err, dates, times)
		healthy.GETFileIndexFAIL++

	}
	fmt.Fprint(w, string(b))
	healthy.GETFileIndexSUCCESS++
}

// Analysis the company of each years
func DataAnalysis(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if req.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		healthy.GETDataFAIL++
		return
	}
	//open Timer
	year, month, day := time.Now().Date()
	hour, min, sec := time.Now().Clock()
	dates = fmt.Sprintf("%d-%d-%d", year, int(month), day)
	times = fmt.Sprintf("%d:%d:%d", hour, min, sec)
	//set filename
	Filename := req.FormValue("filename")
	fmt.Printf("AnalysisCompany CompanyName: %s %s\n", dates, times)
	//call read function
	Data, err := AnalysisForEachCompany(Filename)
	if err != nil {
		fmt.Println("ReadResult error!!", err, dates, times)
		healthy.GETDataFAIL++
	}
	//encoding json
	b, err := json.Marshal(Data)
	if err != nil {
		fmt.Println("encoding json error!!", err, dates, times)
		healthy.GETDataFAIL++
	}
	fmt.Fprint(w, string(b))
	healthy.GETDataSUCCESS++

}

// Analysis the company of each years
func DataYearAnalysis(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if req.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		healthy.GETDataFAIL++
		return
	}
	//open Timer
	year, month, day := time.Now().Date()
	hour, min, sec := time.Now().Clock()
	dates = fmt.Sprintf("%d-%d-%d", year, int(month), day)
	times = fmt.Sprintf("%d:%d:%d", hour, min, sec)
	fmt.Printf("AnalysisYear Year/database:%s %s %s\n", database, dates, times)
	//call read function
	Data, err := AnalysisForEachYear()
	if err != nil {
		fmt.Println("ReadResult error!!", err, dates, times)
		healthy.GETDataFAIL++
	}
	//encoding json
	b, err := json.Marshal(Data)
	if err != nil {
		fmt.Println("encoding json error!!", err, dates, times)
		healthy.GETDataFAIL++
	}
	fmt.Fprint(w, string(b))
	healthy.GETDataSUCCESS++
}

//Analysis the department of each years
func DepartmentYearAnalysis(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if req.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		healthy.GETDataFAIL++
		return
	}
	//open Timer
	year, month, day := time.Now().Date()
	hour, min, sec := time.Now().Clock()
	dates = fmt.Sprintf("%d-%d-%d", year, int(month), day)
	times = fmt.Sprintf("%d:%d:%d", hour, min, sec)
	database := req.FormValue("database")
	fmt.Printf("AnalysisYearDepartment Year/database:%s %s %s\n", database, dates, times)
	//call read function
	Data, err := AnalysisForEachDepartment(database, departmentindexforanalysis)
	if err != nil {
		fmt.Println("ReadResult error!!", err, dates, times)
		healthy.GETDataFAIL++
	}
	//encoding json
	b, err := json.Marshal(Data)
	if err != nil {
		fmt.Println("encoding json error!!", err, dates, times)
		healthy.GETDataFAIL++
	}
	fmt.Fprint(w, string(b))
	healthy.GETDataSUCCESS++
}

func DeleteTable(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if req.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		healthy.GETDataDeleteFAIL++
		return
	}
	//open Timer
	year, month, day := time.Now().Date()
	hour, min, sec := time.Now().Clock()
	dates = fmt.Sprintf("%d-%d-%d", year, int(month), day)
	times = fmt.Sprintf("%d:%d:%d", hour, min, sec)
	//set filename
	database := req.FormValue("database")
	tablename := req.FormValue("filename")
	fmt.Printf("DeleteTable file:%s database:%s %s %s\n", filename, database, dates, times)
	//call read function
	err := DeleteData(database, tablename)
	if err != nil {
		fmt.Println("delete error!!", err, dates, times)
		healthy.GETDataDeleteFAIL++
	}
	healthy.GETDataDeleteSUCCESS++
}

// Date is 2017/01/01
func DataPost1(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if req.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		healthy.POSTDataDateStringFAIL++
		return
	}
	//open Timer
	year, month, day := time.Now().Date()
	hour, min, sec := time.Now().Clock()
	dates = fmt.Sprintf("%d-%d-%d", year, int(month), day)
	times = fmt.Sprintf("%d:%d:%d", hour, min, sec)
	//because of some http problem , we split the one header to two string(filename and database)
	filename := req.FormValue("filename")
	database := req.FormValue("database")
	fmt.Printf("Format1 file:%s database:%s %s %s\n", filename, database, dates, times)
	// check the req.body content is csv format

	reader := csv.NewReader(req.Body)
	csvstuff, err := reader.ReadAll()
	if err != nil {
		http.Error(w, "檔案格式不正確 !!  請確認檔案格式為CSV !!", http.StatusBadRequest)
		fmt.Println("Error!! file format is not csv !!", err, dates, times)
		healthy.POSTDataDateStringFAIL++
		return
	}
	//if csv check ok ,Write file in storage.
	file, err := os.Create("Data/" + filename + ".csv")
	if err != nil {
		fmt.Println("create new file error!!", err, dates, times)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		healthy.POSTDataDateStringFAIL++
		return
	}

	writer := csv.NewWriter(file)

	for _, value := range csvstuff {
		err := writer.Write(value)
		if err != nil {
			fmt.Println("write to file error!", err, dates, times)
			http.Error(w, "選擇的格式與資料不符合,請再次確認 !!", http.StatusInternalServerError)
			healthy.POSTDataDateStringFAIL++
			return
		}

	}
	writer.Flush()
	file.Close()

	err = InputResult1(filename, dates, times, database, index, nameindex)
	if err != nil {
		http.Error(w, "Error!! Table is already exists !!", http.StatusInternalServerError)
		fmt.Println("InputResult Error!!", err, dates, times)
		healthy.POSTDataDateStringFAIL++
		return
	}
	fmt.Println("Scuess !!", dates, times)
	healthy.POSTDataDateStringSUCCESS++
}

// Date is 201701
func DataPost2(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if req.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		healthy.POSTDataDateIntFAIL++
		return
	}
	//open Timer
	year, month, day := time.Now().Date()
	hour, min, sec := time.Now().Clock()
	dates = fmt.Sprintf("%d-%d-%d", year, int(month), day)
	times = fmt.Sprintf("%d:%d:%d", hour, min, sec)
	filename := req.FormValue("filename")
	database := req.FormValue("database")
	fmt.Printf("Format2 file:%s database:%s %s %s\n", filename, database, dates, times)
	// check the req.body content is csv format
	reader := csv.NewReader(req.Body)
	csvstuff, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error!! file format is not csv !!", err, dates, times)
		http.Error(w, "檔案格式不正確 !!  請確認檔案格式為CSV !!", http.StatusBadRequest)
		healthy.POSTDataDateIntFAIL++
		return
	}
	//if csv check ok ,Write file in storage.
	file, err := os.Create("Data/" + filename + ".csv")
	if err != nil {
		fmt.Println("create new file error!!", err, dates, times)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		healthy.POSTDataDateIntFAIL++
		return
	}

	writer := csv.NewWriter(file)

	for _, value := range csvstuff {
		err := writer.Write(value)
		if err != nil {
			fmt.Println("write to file error!", err, dates, times)
			http.Error(w, "選擇的格式與資料不符合,請再次確認 !!", http.StatusInternalServerError)
			healthy.POSTDataDateIntFAIL++
			return
		}

	}
	writer.Flush()
	file.Close()

	err = InputResult2(filename, dates, times, database, index, nameindex)
	if err != nil {
		fmt.Println("InputResult Error!!", err, dates, times)
		http.Error(w, "Error!! Table is already exists !!", http.StatusInternalServerError)
		healthy.POSTDataDateIntFAIL++
		return
	}
	fmt.Println("Scuess !!", dates, times)
	healthy.POSTDataDateIntSUCCESS++
	return
}

// Date is each month
func DataPost3(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if req.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		healthy.POSTDataIPandMonthFAIL++
		return
	}
	//open Timer
	year, month, day := time.Now().Date()
	hour, min, sec := time.Now().Clock()
	dates = fmt.Sprintf("%d-%d-%d", year, int(month), day)
	times = fmt.Sprintf("%d:%d:%d", hour, min, sec)
	filename := req.FormValue("filename")
	database := req.FormValue("database")
	fmt.Printf("Format3 file:%s database:%s %s %s\n", filename, database, dates, times)
	// check the req.body content is csv format
	reader := csv.NewReader(req.Body)
	csvstuff, err := reader.ReadAll()
	if err != nil {
		http.Error(w, "檔案格式不正確 !!  請確認檔案格式為CSV !!", http.StatusBadRequest)
		fmt.Println("file format error!!", err, dates, times)
		healthy.POSTDataIPandMonthFAIL++
		return
	}
	//if csv check ok ,Write file in storage.
	file, err := os.Create("Data/" + filename + ".csv")
	if err != nil {
		fmt.Println("create new file error!!", err, dates, times)
		http.Error(w, "InputResult Error!!", http.StatusInternalServerError)
		healthy.POSTDataIPandMonthFAIL++
		return
	}

	writer := csv.NewWriter(file)

	for _, value := range csvstuff {
		err := writer.Write(value)
		if err != nil {
			fmt.Println("write to file error!", err, dates, times)
			http.Error(w, "檔案格式不正確 !!  請確認檔案格式為CSV !!", http.StatusInternalServerError)
			healthy.POSTDataIPandMonthFAIL++
			return
		}

	}
	writer.Flush()
	file.Close()
	//recover, catch the panic message.
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error Panic!!!", err)
			http.Error(w, "選擇的格式與資料不符合,請再次確認 !!", http.StatusNotAcceptable)
			healthy.POSTDataIPandMonthFAIL++
			return
		}
	}()

	err = InputResult3(filename, dates, times, database, index, nameindex)
	if err != nil {
		fmt.Println(err, dates, times)
		http.Error(w, "Server Error!!", http.StatusNotAcceptable)
		healthy.POSTDataIPandMonthFAIL++
		return
	}
	fmt.Println("Scuess !!", dates, times)
	healthy.POSTDataIPandMonthSUCCESS++
}

//Logs
func DataPost4(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if req.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		healthy.POSTLogFAIL++
		return
	}
	//open Timer
	year, mon, day := time.Now().Date()
	hour, min, sec := time.Now().Clock()
	dates = fmt.Sprintf("%d-%d-%d", year, int(mon), day)
	times = fmt.Sprintf("%d:%d:%d", hour, min, sec)

	filename := req.FormValue("filename")
	database := req.FormValue("database")
	month := req.FormValue("month")
	fmt.Printf("Format4 file:%s database:%s  month:%s  %s %s\n", filename, database, month, dates, times)
	// check the req.body content is csv format
	reader := csv.NewReader(req.Body)
	csvstuff, err := reader.ReadAll()
	if err != nil {
		http.Error(w, "檔案格式不正確 !!  請確認檔案格式為CSV !!", http.StatusBadRequest)
		fmt.Println("Error!! file format is not csv !!", err, dates, times)
		healthy.POSTLogFAIL++
		return
	}
	//if csv check ok ,Write file in storage.
	file, err := os.Create("Data/" + filename + ".csv")
	if err != nil {
		fmt.Println("create new file error!!", err, dates, times)
		http.Error(w, "Error!! file exist !!", http.StatusInternalServerError)
		healthy.POSTLogFAIL++
		return
	}

	writer := csv.NewWriter(file)
	for _, value := range csvstuff {
		err := writer.Write(value)
		if err != nil {
			fmt.Println("write to file error!", err, dates, times)
			http.Error(w, "write error !!", http.StatusInternalServerError)
			healthy.POSTLogFAIL++
			return
		}
	}

	writer.Flush()
	file.Close()
	//Log function call
	err = InputResult4(filename, dates, times, month, database, index, nameindex, companyIndex, companynameindex)
	if err != nil {
		fmt.Println("InputResult Error!!", err, dates, times)
		http.Error(w, "Server Error!!", http.StatusNotAcceptable)
		healthy.POSTLogFAIL++
		return
	}
	fmt.Println("Scuess !!", dates, times)
	healthy.POSTLogSUCCESS++
}

//health
func Health(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if req.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	b, err := json.Marshal(healthy)
	if err != nil {
		fmt.Println("encoding json error(healthc)!!", err, dates, times)
	}
	fmt.Fprint(w, string(b))
}
