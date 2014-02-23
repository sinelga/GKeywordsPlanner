package allcsvfiles

import (
	"io/ioutil"
	"log/syslog"
	"strings"
	
)

func GetAll(golog syslog.Writer, dir string,locale string, themes string) []string {

	var filesarr []string
	var rootdir string

	if dir == "" {

		rootdir = "csvfiles/"+locale+"_"+themes+"/"

	} else {

		rootdir = dir

	}
	if dr, err := ioutil.ReadDir(rootdir); err != nil {

		golog.Crit(err.Error())

	} else {

		for _, fl := range dr {

			flname := fl.Name()
			
			if strings.HasSuffix(flname, ".csv") {
			
				filesarr = append(filesarr,rootdir+flname)
			
			}

		}

	}

	return filesarr

}
