# Library-System-API (For Window OS)    [![Foundation](https://img.shields.io/badge/Golang-Foundation-green.svg)](http://golangfoundation.org)


## Downloads
### Golang
* **[Golang Official Website](https://golang.org/dl/)**

### Glide
* **If you have install glide , just run:**

> glide install

* **If you didn't install glide yet,please reference
[Glide website](https://github.com/Masterminds/glide).**

## Library-website
### (https://140.124.183.37/lib/)
## Library-API
### (https://140.124.183.37:8066)
## HTTP 切入點 Function 說明 

### 1.Function InputResult1  **(/DataDateString?database=&filename=)**
* 若輸入資料日期格式為 2017/01/01,  利用“/”分割字串並取月進行份計算。   
`EX: "2017/01/01" ->> 2017 01 01`

### 2.Function InputResult2 **(/DataDateInt?database=&filename=)** 
* 若輸入資料日期格式為 201701, 先將String轉換為Int,透過除100取餘數得到月份進行計算。  
`EX: "201701" ->> 201701 % 100 ->> 1`

### 3.Function InputResult3 **(/DataIPandMonth?database=&filename=)** 
* 若輸入資料日期格式以分為各個月份,則直接進行統計。

### 4.Function InputResult4 **(/log?database=&month=&filename=)** 
* 統計Log (依年、月進行統計。)   

### 5.FileIndexGet **(/FileIndex?database=)** 
* 查詢目前在DB存在的廠商資料。

### 6.DataGet **(/GetData?database=&filename=)** 
* 查詢選擇廠商的報表資料。


### 7.DataAnalysis **(/GetData/Analysis?&filename=)** 
* 統計該廠商每一年的使用。

### 8.DataYearAnalysis **(/GetData/YearAnalysis)** 
* 統計每一年所有的使用量。

### 9.DataYearAnalysis **(/GetData/DepartmentAnalysis?database=)** 
* 統計當年各個系所占所有用量中的比例(用量)。

### 10.DeleteTable **(/DataDelete?database=&filename=)** 
* 刪除選擇廠商的報表資料。


## Attention
* Function 1,2 的日期與 IP 欄位必須放在CSV檔的1,2行。
* Function 3 的資料格式必須是從第一欄第一列開始。


-----
### Issues
* `need to change http post body format , now is binary, need change to from-data.`
* `need code reviews.`

### Website Record History （Website更改紀錄）
* `2017/12/01 Success integrate http_client in the project.(need to update the latest version.) --by Neil`
* `2017/12/04 Update http_client to latest version, and fix the error response. --by Neil`
* `2017/12/05 change website domain. ( 140.124.183.37 -> 140.124.183.37/lib/ ) --by Neil`
* `2017/12/06 Add reload function of delete button. --by Jessie`
* `2017/12/08 Add sort operation of statistic table. Add loading function while ajax loading before send. Add line chart of factory database increasing. --by Jessie`
* `2017/12/11 Add line chart of factory usage in every years.  --by Jessie`
* `2017/12/12 Add year analysis api ajax.  --by Jessie`
* `2017/12/14 Add Button of download example files. Add line chart of usage according to select factory in every years. --by Jessie`
* `2017/12/22 Modify chart and refactor the architecture.  --by Jessie`
* `2017/12/22 Add pie chart by year in particular Section, add navbar tag of old page.  --by Jessie`
* `2017/12/26 Fix bug of pie chart and add download file option.  --by Jessie`
* `2017/12/27 Recover line chart and add percentage of pie chart. --by Jessie`
* `2017/12/27 Change the chart name and add go top tag. --by Jessie`
* `2018/01/03 Add navbar toggle button. --by Jessie`
* `2018/01/08 Add checkbox of statistic table and change the name of third chart. Add download button of every chart to transform chart into csv or pdf type. --by Jessie`
* `2018/01/10 Modify percentage and add total row in the table button. --by Jessie`
* `2018/03/01 Update file upload function call from FactoryStatistics.js. --by Neil`

### API Record History （API更改紀錄）
* `2017/9/18 First time push and write Readme. --by Neil`
* `2017/9/23 Create http Getdata function form database and check OK. --by Neil`
* `2017/9/24 create fileindex in db and test it OK . --by Neil `
* `2017/9/26 add http post function and test OK. --by Neil`
* `2017/10/6 add test function and test OK. --by Neil`
* `2017/10/13 removed httpHeader check ,and update the IP index file. --by Neil`
* `2017/10/14 gotest check OK. --by Neil`
* `2017/10/24 update http error response. --by Neil`
* `2017/10/28 Computer center log's analysis function create and testing in local OK. --by Neil`
* `2017/11/02 change the step of log's analysis and the method of write in database. --by Neil` 
* `2017/11/07 Log function test OK. --by Neil`
* `2017/11/08 Log function debug. --by Neil`
* `2017/11/13 fix main error message. --by Neil`
* `2017/11/15 update filestore readResult function. --by Neil`
* `2017/11/17 change sql and post function so let client can choose the database which they want save. --by Neil`
* `2017/11/20 filestore.go debug. --by Neil`
* `2017/11/21 change get function modify databases. --by Neil`
* `2017/11/22 create delete table function and change header to fromvalues(POST). --by Neil`
* `2017/11/30 test all API and change defer Close() and defer Flush() and add its before analysis function. --by Neil`
* `2017/12/01 Solve the problem about the panic via recover function. --by Neil`
* `2017/12/06 add the total counter in the CountEachMonth function in func.go  --by Neil`
* `2017/12/08 fix new sql command can not build problem, and add a function of analysis each years. --by Neil`
* `2017/12/09 Solve the problem of can't change or add function( because of the copy paste, we didn't change the include path), and add a key[YEARS] for the function of analysis each years. --by Neil`
* `2017/12/10 Update the WriteIDB function (it make sure that the different format of the company data can write or update in the same table.) --by Neil`
* `2017/12/11 Decided to remove the total row in every table. Total will count in analysis func,not count in input func. --by Neil`
* `2017/12/12 check all system status. --by Neil`
* `2017/12/14 check all system status and add healthc. --by Neil`
* `2017/12/18 remove month's counter in analysis func,just only count total. Add the department analysis func ,it just count all years used of each departments.  --by Neil`
* `2017/12/20 fixed problem of the sql connection over the max range. add sql.Close() to solved it. --by Neil`
* `2017/12/28 restore month's counter in analysis func. --by Neil`
* `2018/01/06 add a new line of exception IP and save in DB. --by Neil`