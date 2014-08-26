package checkinsertkeyword

import (
	"fmt"
	"getAllKeywords"
	"getNewKeywords"
	"inserNewKeywordsinDB"
	"log/syslog"
)

func CheckInsert(golog syslog.Writer, locale string, themes string) {

	mapresAll := getAllKeywords.GetAllKeywordsFromProdDB(golog, locale, themes)

	fmt.Println("mapresAll len --> ", len(mapresAll))

	mapresNew := getNewKeywords.GetNewKewywordForTmpDB(golog, locale, themes)

	fmt.Println("mapresNew len --> ", len(mapresNew))
	
	var wordstoinsert []string

	for key, _ := range mapresNew {

//		fmt.Println("word-> ", key)
		
		if _,ok := mapresAll[key]; !ok {
			
			wordstoinsert  =append(wordstoinsert,key)
			
			
		}		

	}
	
	inserNewKeywordsinDB.InsertIntoDB(golog,locale,themes,wordstoinsert )
		
	

}
