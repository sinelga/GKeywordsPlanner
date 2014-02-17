package main 

import (
    "flag"
    "fmt"
    "log"
    "log/syslog"
    "allcsvfiles"
    "elaboratecsvfile"
    "extphrases"
    "domains"
)

var localeFlag = flag.String("locale", "", "must be fi_FI/en_US/it_IT")
var themesFlag = flag.String("themes", "", "must be porno/finance/fortune...")

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
    flag.Parse() // Scan the arguments list
    
    locale := *localeFlag
	themes := *themesFlag

    if *versionFlag {
        fmt.Println("Version:", APP_VERSION)
    }
    
    	golog, err := syslog.New(syslog.LOG_ERR, "golog")

		defer golog.Close()
		if err != nil {
			log.Fatal("error writing syslog!!")
		}
    	golog.Info("Start "+locale+" "+themes)
    	
    	allphrase := make(map[string]domains.GoogleAdsVal)
    	
    	allcsvfilesarr := allcsvfiles.GetAll(*golog,"")
    	
    	for _,csvfile := range allcsvfilesarr {
    	
    		golog.Info(csvfile)
    		
    		intmap := elaboratecsvfile.Elaborate(*golog,csvfile)
    		
    		for key,phraseval := range intmap{
    		
    			allphrase[key]=phraseval
    			    		
    		}
    		    	
    	}
    	 
    	fmt.Println("1-> ",len(allphrase))
    	  	
    	extmap := extphrases.GetPhraseFromDB(*golog,locale,themes)
    	
    	fmt.Println("ext ",len(extmap))
    	
    	
    	for key,phraseval :=range extmap {
    	
    		allphrase[key]=phraseval
    	
    	}
    	
    	fmt.Println("2-> ",len(allphrase))
    	
    	if len(allphrase) > len(extmap) {
    	
    		fmt.Println("OK updatedb")
    	    	
    	}
    	    	   		    
    
}

