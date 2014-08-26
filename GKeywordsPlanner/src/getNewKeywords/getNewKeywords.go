package getNewKeywords

import (
	
		"log/syslog"

	"database/sql"
	_ "github.com/mxk/go-sqlite/sqlite3"
//	"fmt"
	"strings"
	"regexp"

)

func GetNewKewywordForTmpDB(golog syslog.Writer, locale string, themes string)map[string]struct{} {
	
	outmap := make(map[string]struct{})
	
	
	db, err := sql.Open("sqlite3", "/home/juno/git/goFastCgi/goFastCgi/singo.db")
	defer db.Close()
	if err != nil {

		golog.Err("getNewKeywords:GetNewKewywordForTmpDB " + err.Error())
	} else {
			
		dbtable := locale + "_" + themes + "_phrases"

		sqlstr := "select Keyword from " + dbtable

		rows, err := db.Query(sqlstr)
		if err != nil {

			golog.Err("getNewKeywords:GetNewKewywordForTmpDB " + err.Error())
		}
		defer rows.Close()
		
		for rows.Next() {
			var keyword_phrase string
			rows.Scan(&keyword_phrase)
			
			words := strings.Fields(keyword_phrase)
			
//			fmt.Println(words, len(words))
			r, _ := regexp.Compile(`^[a-zA-Z]+$`)
			
			for _,word := range words {
				
				
				if len(word) > 4 {

					
					if r.MatchString(word) {
						
//						fmt.Println("!!!! ",word)
						outmap[word] =struct{}{}
						
					}

					
					
				}
				
			} 
			
			
		}
		
		
	
	}
	
	db.Close()
	return outmap
	
}
