package main 

import (
    "flag"
    "fmt"
    "log"
    "log/syslog"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
    flag.Parse() // Scan the arguments list 

    if *versionFlag {
        fmt.Println("Version:", APP_VERSION)
    }
    
    	golog, err := syslog.New(syslog.LOG_ERR, "golog")

		defer golog.Close()
		if err != nil {
			log.Fatal("error writing syslog!!")
		}
    
    
    
}

