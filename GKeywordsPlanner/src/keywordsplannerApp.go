package main

import (
	"allcsvfiles"
	"domains"
	"elaboratecsvfile"
	"extphrases"
	"flag"
	"fmt"
	"log"
	"log/syslog"
	"updatephrasesdb"
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
	golog.Info("Start " + locale + " " + themes)

	allphrase := make(map[string]domains.GoogleAdsVal)

	allcsvfilesarr := allcsvfiles.GetAll(*golog,"",locale,themes)

	for _, csvfile := range allcsvfilesarr {

		golog.Info(csvfile)

		intmap := elaboratecsvfile.Elaborate(*golog, csvfile)

		for key, phraseval := range intmap {

			allphrase[key] = phraseval

		}

	}

	fmt.Println("1-> ", len(allphrase))

	extmap := extphrases.GetPhraseFromDB(*golog, locale, themes)

	fmt.Println("ext ", len(extmap))

	for key, phraseval := range extmap {

		allphrase[key] = phraseval

	}

	fmt.Println("2-> ", len(allphrase))

	if len(allphrase) > len(extmap) {

		fmt.Println("OK updatedb")
		updatephrasesdb.UpdateDB(*golog,locale, themes,allphrase)

	} else {
	
		fmt.Println("Hmm. nothing to do?")
	
	}

}
