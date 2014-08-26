package inserNewKeywordsinDB

import (
	"database/sql"
	//	"fmt"
	_ "github.com/mxk/go-sqlite/sqlite3"
	"log/syslog"
	"time"
)

func InsertIntoDB(golog syslog.Writer, locale string, themes string, wordstoinsert []string) {

	timenow := int64(time.Now().Unix())*1000
	
	db, err := sql.Open("sqlite3", "/home/juno/git/goFastCgi/goFastCgi/singo.db")
	defer db.Close()
	if err != nil {

		golog.Err("getNewKeywords:GetNewKewywordForTmpDB " + err.Error())
	} else {

		tx, err := db.Begin()
		if err != nil {

			golog.Crit(err.Error())
			//			os.Exit(1)
		}

		for _, newword := range wordstoinsert {

			stmt, err := tx.Prepare("insert into keywords(Locale,Themes,Keyword,Created,Updated) values(?,?,?,?,?)")

			if err != nil {

				golog.Err(err.Error())
			}
			defer stmt.Close()

			if _, err = stmt.Exec(locale, themes, newword,timenow,timenow); err != nil {
				golog.Crit(err.Error())
				//				os.Exit(1)

			}

		}
		tx.Commit()
		db.Close()

	}

}
