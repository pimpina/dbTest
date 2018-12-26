package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"io"
	"log"
	_ "os"

	_ "gopkg.in/goracle.v2"
)

type product struct {
	ID      int64  `json:"id"`
	Descrip string `json:"desc"`
}

type model struct {
	BUID        int64
	BUDesc      string
	ModelID     int64
	ModelDesc   string
	ProductID   int64
	ProductName string
	ProductDesc string
	Owner       string
}

func testCursor() {
	fmt.Println("...Start test Cursor...")
	/*var dbusername = os.Getenv("APP_DB_USERNAME")
	var dbuserpass = os.Getenv("APP_DB_PASSWORD")
	var dbname = os.Getenv("APP_DB_NAME")*/
	var oProduct []product
	//fmt.Println(dbusername + "/" + dbuserpass + "@" + dbname)
	//db, err := sql.Open("goracle", dbusername+"/"+dbuserpass+"@"+dbname)
	db, err := sql.Open("goracle", "bgweb/bgweb#1@//tv-uat62-dq.tvsit.co.th:1521/UAT62")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//sol1
	statement := `begin redibsservice.getlistbu62(:1); end;`
	var resultI driver.Rows

	if _, err := db.Exec(statement, sql.Out{Dest: &resultI}); err != nil {
		log.Fatal(err)
	}

	defer resultI.Close()
	values := make([]driver.Value, len(resultI.Columns()))

	for {
		err = resultI.Next(values)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("error:", err)
		}

		var lProduct product
		lProduct.ID = values[0].(int64)
		lProduct.Descrip = values[1].(string)
		oProduct = append(oProduct, lProduct)
	}
	fmt.Println(oProduct)
	fmt.Println("...End test Cursor...")
}

func testQuery() {
	fmt.Println("...Start test Query...")
	// Connect string format: [username]/[password]@//[hostname]:[port]/[DB name]
	db, err := sql.Open("goracle", "bgweb/bgweb#1@//tv-uat62-dq.tvsit.co.th:1521/UAT62")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select b.id, b.description from businessunit b")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var strVal, strVal2 string
	for rows.Next() {
		err := rows.Scan(&strVal, &strVal2)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(strVal, ",", strVal2)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("...End test Query...")
}

func testProc() {
	fmt.Println("...Start test Proc...")
	var oModel []model
	db, err := sql.Open("goracle", "bgweb/bgweb#1@//tv-uat62-dq.tvsit.co.th:1521/UAT62")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var statement string
	statement = "begin redibsservice.getdatamodel62(:0,:1); end;"
	var resultI driver.Rows
	var mID int64 = 765

	if _, err := db.Exec(statement, mID, sql.Out{Dest: &resultI}); err != nil {
		log.Fatal(err)
	}

	defer resultI.Close()
	values := make([]driver.Value, len(resultI.Columns()))

	for {
		err = resultI.Next(values)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("error:", err)
		}

		var lModel model
		lModel.BUID = values[0].(int64)
		lModel.BUDesc = values[1].(string)
		lModel.ModelID = values[2].(int64)
		lModel.ModelDesc = values[3].(string)
		lModel.ProductID = values[4].(int64)
		lModel.ProductName = values[5].(string)
		lModel.ProductDesc = values[6].(string)
		lModel.Owner = values[7].(string)
		oModel = append(oModel, lModel)
	}
	fmt.Println(oModel)
	fmt.Println("...End test Proc...")

}

func testFunc() {
	fmt.Println("...Start test Func...")
	db, err := sql.Open("goracle", "bgweb/bgweb#1@//tv-uat62-dq.tvsit.co.th:1521/UAT62")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var statement string
	statement = "begin :result := redibsservice.updatewarrantysnicc(:0,:1,:2); end;"
	var resultI int64
	var sn, wdate, resultS string
	sn = "40023458223"
	wdate = /*"01/01/2018"*/ "01/01/1900"

	if _, err := db.Exec(statement, sql.Out{Dest: &resultI}, sn, wdate, sql.Out{Dest: &resultS}); err != nil {
		log.Fatal(err)
	}

	fmt.Println("1:", resultI)
	fmt.Println("2:", resultS)

	fmt.Println("...End test Func...")

}

func main() {
	testQuery()
	testCursor()
	testProc()
	testFunc()

	/*var strVal string
	fmt.Printf("input :")
	fmt.Scan(&strVal)
	fmt.Println(strVal)*/
}
