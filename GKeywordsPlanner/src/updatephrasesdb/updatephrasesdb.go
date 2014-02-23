package updatephrasesdb

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"domains"
	"log/syslog"
	"os"
)

func UpdateDB(golog syslog.Writer, locale string, themes string, allphrases map[string]domains.GoogleAdsVal) {

	db, err := sql.Open("sqlite3", "/home/juno/git/goFastCgi/goFastCgi/singo.db")
	defer db.Close()
	if err != nil {

		golog.Err("updatephrasesdb:UpdateDB " + err.Error())
	} else {

		dbtable := locale + "_" + themes + "_phrases"

		tx, err := db.Begin()
		if err != nil {

			golog.Crit(err.Error())
			os.Exit(1)
		}
		stmt, err := tx.Prepare("delete from " + dbtable)

		if _, err = stmt.Exec(); err != nil {

			golog.Crit(err.Error())
			os.Exit(1)

		}
		stmt.Close()

		for key, val := range allphrases {

			stmt, err = tx.Prepare("insert into " + dbtable + "('Keyword' ,'Avg. monthly searches' ,'Competition' ,'Suggested bid') values(?,?,?,?)")
			if err != nil {

				golog.Err(err.Error())
			}
			defer stmt.Close()

			keyword := key
			avgmonthlysearches := val.Avgmonthlysearch
			competition := val.Competition
			suggestedbid := val.Suggestedbid
			if _, err = stmt.Exec(keyword, avgmonthlysearches, competition, suggestedbid); err != nil {
				golog.Crit(err.Error())
				os.Exit(1)

			}

		}

		tx.Commit()

	}

}
