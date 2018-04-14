package sqlstore

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type WriteToSQL struct {
	serverurl string
	database  string
	db        *sql.DB
}

func NewWriteToSQL(username, password, serverurl, database string) (*WriteToSQL, error) {
	// Create the database handle, confirm driver is present
	DB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4,big5", username, password, serverurl, database))
	if err != nil {
		fmt.Println("Connect Error(panic)!!", err)
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		return nil, err
	}
	DB.SetMaxOpenConns(0)
	return &WriteToSQL{db: DB}, nil
}

func (w *WriteToSQL) read(sqlQuery string, args ...interface{}) (*sql.Rows, error) {
	res, err := w.db.Query(sqlQuery, args...)
	if err != nil {
		return res, err
	}
	return res, err
}
func (w *WriteToSQL) Close() error {
	return w.db.Close()

}

func (w *WriteToSQL) SelectAll(Tablename string) (*sql.Rows, error) {
	return w.read("SELECT * FROM " + Tablename)
}
func (w *WriteToSQL) ReadName(Tablename string) (*sql.Rows, error) {
	return w.read("SELECT NAME FROM " + Tablename)
}
func (w *WriteToSQL) ReadEachName(Department, Tablename string) (*sql.Rows, error) {
	return w.read("SELECT COUNT(`NAME`) FROM " + Tablename + " WHERE NAME = '" + Department + "'")
}
func (w *WriteToSQL) ReadEachLine(Tablename string) (*sql.Rows, error) {
	return w.read("SELECT * FROM " + Tablename + " WHERE NAME = 'TOTAL'")
}
func (w *WriteToSQL) ReadEachDepartmentTotal(Tablename, Department string) (*sql.Rows, error) {
	return w.read("SELECT Total FROM " + Tablename + " WHERE NAME = '" + Department + "'")
}
func (w *WriteToSQL) CreateTable(TableName string) (*sql.Rows, error) {
	return w.read("CREATE TABLE " + TableName + " (NAME longtext CHARACTER SET big5,JAN int(32),FEB int(32),MAR int(32),APR int(32),MAY int(32),JUN int(32),JUL int(32),AUG int(32),SEP int(32),OCT int(32),NOV int(32),DECE int(32),Total int(32))")
}
func (w *WriteToSQL) WriteResult(filename, NAME string, month1, month2, month3, month4, month5, month6, month7, month8, month9, month10, month11, month12, Total int) (*sql.Rows, error) {
	return w.read("INSERT INTO "+filename+" (NAME,JAN,FEB,MAR,APR,MAY,JUN,JUL,AUG,SEP,OCT,NOV,DECE,Total) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)", NAME, month1, month2, month3, month4, month5, month6, month7, month8, month9, month10, month11, month12, Total)
}
func (w *WriteToSQL) UpdateResultforEachMonth(Tablename, Month, Department string, countnumber int) (*sql.Rows, error) {
	return w.read("UPDATE "+Tablename+" SET "+Month+"= ?,Total = Total + ?  WHERE NAME = '"+Department+"'", countnumber, countnumber)
}
func (w *WriteToSQL) UpdateResult(Tablename, Department string, month1, month2, month3, month4, month5, month6, month7, month8, month9, month10, month11, month12, Total int) (*sql.Rows, error) {
	return w.read("UPDATE "+Tablename+" SET JAN = JAN + ?, FEB = FEB + ?, MAR = MAR + ?, APR = APR + ?, MAY = MAY + ?, JUN = JUN + ?, JUL = JUL + ?, AUG = AUG + ?, SEP = SEP + ?, OCT = OCT + ?, NOV = NOV + ?, DECE = DECE + ?, Total = Total + ?  WHERE NAME = '"+Department+"'", month1, month2, month3, month4, month5, month6, month7, month8, month9, month10, month11, month12, Total)
}
func (w *WriteToSQL) DeleteTable(TableName string) (*sql.Rows, error) {
	return w.read("DROP TABLE " + TableName)
}
func (w *WriteToSQL) WriteInFileIndex(NAME, DATE, TIME string) (*sql.Rows, error) {
	return w.read("INSERT INTO FileIndex (NAME,DATE,TIME) VALUES (?,?,?)", NAME, DATE, TIME)
}
func (w *WriteToSQL) DeleteInFileIndex(tablename string) (*sql.Rows, error) {
	return w.read("DELETE FROM FileIndex WHERE NAME = '" + tablename + "'")
}
