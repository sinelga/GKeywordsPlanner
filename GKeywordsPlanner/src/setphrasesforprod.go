package main

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"log/syslog"
	"os"
	"time"
)

var localeFlag = flag.String("locale", "", "must be fi_FI/en_US/it_IT")
var themesFlag = flag.String("themes", "", "must be porno/finance/fortune...")

func main() {
	flag.Parse() // Scan the arguments list

	locale := *localeFlag
	themes := *themesFlag

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}
	golog.Info("Start setphrasesforprod" + locale + " " + themes)

	var allphrasesmap map[string]struct{}
	
	timenow := int64(time.Now().Unix())*1000

	db, err := sql.Open("sqlite3", "/home/juno/git/goFastCgi/goFastCgi/singo.db")
	defer db.Close()
	if err != nil {

		golog.Err("extphrases:GetPhraseFromDB " + err.Error())
	} else {

		allphrasesmap = make(map[string]struct{})

		dbtable := locale + "_" + themes + "_phrases"

		sqlstr := "select Keyword from " + dbtable

		rows, err := db.Query(sqlstr)
		if err != nil {
			golog.Err("setphrasesforprod " + err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var phrase string
			rows.Scan(&phrase)
			allphrasesmap[phrase] = struct{}{}

		}

		fmt.Println(len(allphrasesmap))
		rows.Close()

		sqlstr = "select Phrase from phrases where Locale='" + locale + "' and Themes='" + themes + "'"

		rows, err = db.Query(sqlstr)
		if err != nil {
			golog.Err("setphrasesforprod " + err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var phrase string
			rows.Scan(&phrase)
			allphrasesmap[phrase] = struct{}{}

		}

		fmt.Println(len(allphrasesmap))
		rows.Close()

		tx, err := db.Begin()
		if err != nil {

			golog.Crit(err.Error())
			os.Exit(1)
		}
		stmt, err := tx.Prepare("delete from phrases where Locale='" + locale + "' and Themes='" + themes + "'")

		if _, err = stmt.Exec(); err != nil {

			golog.Crit(err.Error())
			os.Exit(1)

		}
		stmt.Close()

		for key := range allphrasesmap {

			stmt, err = tx.Prepare("insert into phrases('Phrase' ,'Locale','Themes','Created','Updated') values(?,?,?,?,?)")
			if err != nil {
				golog.Err(err.Error())
			}
			defer stmt.Close()

			if _, err = stmt.Exec(key, locale, themes,timenow,timenow); err != nil {
				golog.Crit(err.Error())
				os.Exit(1)

			}
		}

		tx.Commit()

		db.Close()

	}

}
