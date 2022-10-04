package bal

import (
	"database/sql"
	"log"
)

const (
	NL         = "\n"
	BAL_RID    = 1
	BAL_REC    = "bal_rec"
	TIME_REC   = "time_rec"
	StringType = DataType("StringType")
	IntType    = DataType("IntType")
)

func route() {

}

type DataType string

func format(target string, dataType DataType) string {
	if dataType == IntType {
		if target == "" {
			target = "0"
		}
		return target
	} else {
		if target == "" {
			target = "None."
		}
		return target
	}
}

func nl(str *string) {
	if *str != "" {
		*str += NL
	}
}

func hf(str *string) {
	*str += NL
	*str += "---------------"
	*str += NL
}

func ScanRows(rows *sql.Rows, err error, dataType DataType) string {
	ret := ""

	if err != nil {
		log.Println(err)
		return err.Error()
	}

	for rows.Next() {
		nl(&ret)

		var tmp string
		if err := rows.Scan(&tmp); err != nil {
			log.Println(err)
		}
		ret += tmp
	}
	return format(ret, dataType)
}

func ScanRows2(rows *sql.Rows, err error) string {
	ret := ""

	if err != nil {
		log.Println(err)
		return err.Error()
	}

	for rows.Next() {
		nl(&ret)

		var left, right string
		if err := rows.Scan(&left, &right); err != nil {
			log.Println(err)
		}
		ret += left + ": " + right
	}
	return format(ret, StringType)
}

func MonthSum(db *sql.DB, msg string) string {
	// Get Balance records group
	s1 := "SELECT main.name, bal.sum FROM (SELECT mid, SUM(value) sum FROM bal_rec WHERE date LIKE '" + msg + "%'" + " GROUP BY bal_rec.mid ORDER BY mid) bal LEFT JOIN main ON bal.mid = main.mid"
	balGrpRow, err := db.Query(s1)
	balGrp := ScanRows2(balGrpRow, err)

	// Get Balance records
	balSumRow, err := db.Query("SELECT SUM(value) FROM bal_rec WHERE date LIKE '" + msg + "%'")
	balSum := "Total:" + ScanRows(balSumRow, err, IntType)

	// Get Time records
	timeRows, err := db.Query("SELECT type, SUM(duration) FROM time_rec WHERE date LIKE '" + msg + "%' GROUP BY type")
	timeGrp := ScanRows2(timeRows, err)

	ret := balGrp
	nl(&ret)
	ret += balSum
	hf(&ret)
	ret += timeGrp
	return ret
}

func TodaySum(db *sql.DB) string {
	// Get Balance records group
	balGrpRow, err := db.Query("SELECT main.name, bal.sum FROM (SELECT mid, SUM(value) sum FROM bal_rec WHERE date = DATE_FORMAT(NOW(),'%Y%m%d') GROUP BY bal_rec.mid ORDER BY mid) bal LEFT JOIN main ON bal.mid = main.mid")
	balGrp := ScanRows2(balGrpRow, err)

	// Get Balance records sum
	balSumRow, err := db.Query("SELECT SUM(value) FROM bal_rec WHERE date = DATE_FORMAT(NOW(),'%Y%m%d')")
	balSum := "Total: " + ScanRows(balSumRow, err, IntType)

	// Get Time records
	timeRows, err := db.Query("SELECT type, SUM(duration) FROM time_rec WHERE date = DATE_FORMAT(NOW(),'%Y%m%d') GROUP BY type")
	timeGrp := ScanRows2(timeRows, err)

	ret := balGrp
	nl(&ret)
	ret += balSum
	hf(&ret)
	ret += timeGrp
	return ret
}

func YesterdaySum(db *sql.DB) string {
	// Get Balance records group
	balGrpRow, err := db.Query("SELECT main.name, bal.sum FROM (SELECT mid, SUM(value) sum FROM bal_rec WHERE date = DATE_FORMAT(SUBDATE(NOW(),1),'%Y%m%d') GROUP BY bal_rec.mid ORDER BY mid) bal LEFT JOIN main ON bal.mid = main.mid")
	balGrp := ScanRows2(balGrpRow, err)

	// Get Balance records sum
	balSumRow, err := db.Query("SELECT SUM(value) FROM bal_rec WHERE date = DATE_FORMAT(SUBDATE(NOW(),1),'%Y%m%d')")
	balSum := "Total: " + ScanRows(balSumRow, err, IntType)

	// Get Time records
	timeRows, err := db.Query("SELECT type, SUM(duration) FROM time_rec WHERE date = DATE_FORMAT(SUBDATE(NOW(),1),'%Y%m%d') GROUP BY type")
	timeGrp := ScanRows2(timeRows, err)

	ret := balGrp
	nl(&ret)
	ret += balSum
	hf(&ret)
	ret += timeGrp
	return ret
}

func Catalog(db *sql.DB) string {
	var result string

	rows, err := db.Query("SELECT mid, sid, name FROM sub ORDER BY mid, sid")
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	for rows.Next() {
		if len(result) != 0 {
			result += "\n"
		}

		var mid, sid, name string
		if err := rows.Scan(&mid, &sid, &name); err != nil {
			log.Println(err)
			return err.Error()
		}

		result += mid + "," + sid + "," + name
	}

	return result
}
