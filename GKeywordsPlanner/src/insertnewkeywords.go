package main 

import (
    "flag"
//    "fmt"
    "log"
	"log/syslog"
)

// The flag package provides a default help printer via -h switch
//var versionFlag *bool = flag.Bool("v", false, "Print the version number.")
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
}

