package getAllKeywords

import (
	"log/syslog"

	"database/sql"
	_ "github.com/mxk/go-sqlite/sqlite3"
)

func GetAllKeywordsFromProdDB(golog syslog.Writer, locale string, themes string)map[string]struct{} {
	
	outmap := make(map[string]struct{})

	db, err := sql.Open("sqlite3", "/home/juno/git/goFastCgi/goFastCgi/singo.db")
	defer db.Close()
	if err != nil {

		golog.Err("getAllKeywords:GetAllKeywordsFromProdDB " + err.Error())
	} else {

		sqlstr := "select keyword from keywords where Locale='" + locale + "' and Themes='" + themes + "'"

		rows, err := db.Query(sqlstr)
		if err != nil {

			golog.Err("getAllKeywords:GetAllKeywordsFromProdDB  " + err.Error())
		}
		defer rows.Close()
		for rows.Next() {
			var keyword string
			rows.Scan(&keyword)
			
			outmap[keyword] =struct{}{}
			
			
		}
		

	}
	db.Close()
	
	return outmap

}
