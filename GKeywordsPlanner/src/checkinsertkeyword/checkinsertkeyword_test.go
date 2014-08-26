package checkinsertkeyword

import (
	
	"testing"
	
	"log"
	"log/syslog"
//	"getAllKeywords"
//	"getNewKeywords"
//	"fmt"
)

func TestCheckInsert(t *testing.T) {
	
	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
				
	}
	locale :="en_US"
	themes := "finance" 
	
//	mapresAll := getAllKeywords.GetAllKeywordsFromProdDB(*golog,locale,themes)
//	
//	fmt.Println("mapresAll len --> ",len(mapresAll))
//	
//	mapresNew := getNewKeywords.GetNewKewywordForTmpDB(*golog,locale,themes)
//	
//	fmt.Println("mapresNew len --> ",len(mapresNew))
//	
//	for key,_ :=range mapresNew {
//		
//		fmt.Println("word-> ",key)
//		
//	}
	
	CheckInsert(*golog,locale,themes)
	
}