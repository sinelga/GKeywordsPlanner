package extphrases

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"domains"
	"log/syslog"
)

func GetPhraseFromDB(golog syslog.Writer, locale string, themes string) map[string]domains.GoogleAdsVal {

	var outmap map[string]domains.GoogleAdsVal

	db, err := sql.Open("sqlite3", "/home/juno/git/goFastCgi/goFastCgi/singo.db")
	defer db.Close()
	if err != nil {

		golog.Err("extphrases:GetPhraseFromDB " + err.Error())
	} else {

		dbtable := locale + "_" + themes + "_phrases"

		sqlstr := "select * from " + dbtable

		rows, err := db.Query(sqlstr)
		if err != nil {

			golog.Err("extphrases:GetPhraseFromDB " + err.Error())
		}
		defer rows.Close()

		outmap = make(map[string]domains.GoogleAdsVal)

		for rows.Next() {
			var phrase string
			var avgmonthlysearch int
			var competition float64
			var suggestedbid float64
			
			rows.Scan(&phrase, &avgmonthlysearch, &competition, &suggestedbid)

			var googleAdsVal domains.GoogleAdsVal
			googleAdsVal.Avgmonthlysearch = avgmonthlysearch
			googleAdsVal.Competition = competition
			googleAdsVal.Suggestedbid = suggestedbid

			outmap[phrase] = googleAdsVal

		}
		rows.Close()
		db.Close()

	}

	return outmap
}
