package datamodel

type IPIndex struct {
	IPAddress string
	Name      string
}
type NAMEIndex struct {
	Name string
}

type FileIndex struct {
	ID   int
	Name string
	Date string
	Time string
}

type CompanyIndex struct {
	CompanyName string
	IPAddress   []int
}

type Data struct {
	Date    string
	LoginIP string
}

type ResultOfMonth struct {
	IPAddress string
	Month     string
	Count     int
}

type DataResult struct {
	NAME  string
	IP    string
	JAN   int
	FEB   int
	MAR   int
	APR   int
	MAY   int
	JUN   int
	JUL   int
	AUG   int
	SEP   int
	OCT   int
	NOV   int
	DEC   int
	Total int
}

type LogOfComputerCenter struct {
	Date               string
	SourceAddress      string
	DestinationAddress string
	CompanyName        string
}

type SQLResultFormat struct {
	NAME  string
	JAN   int
	FEB   int
	MAR   int
	APR   int
	MAY   int
	JUN   int
	JUL   int
	AUG   int
	SEP   int
	OCT   int
	NOV   int
	DEC   int
	Total int
}

type AnalysisResult struct {
	YEARS int
	NAME  string
	JAN   int
	FEB   int
	MAR   int
	APR   int
	MAY   int
	JUN   int
	JUL   int
	AUG   int
	SEP   int
	OCT   int
	NOV   int
	DEC   int
	Total int
}

type Health struct {
	GETDataSUCCESS            int
	GETDataFAIL               int
	GETFileIndexSUCCESS       int
	GETFileIndexFAIL          int
	GETDataDeleteSUCCESS      int
	GETDataDeleteFAIL         int
	POSTDataDateStringSUCCESS int
	POSTDataDateStringFAIL    int
	POSTDataDateIntSUCCESS    int
	POSTDataDateIntFAIL       int
	POSTDataIPandMonthSUCCESS int
	POSTDataIPandMonthFAIL    int
	POSTLogSUCCESS            int
	POSTLogFAIL               int
}
