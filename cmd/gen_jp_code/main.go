package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/kenkyu392/go-holidays"
	"github.com/kenkyu392/go-holidays/cmd/internal"
)

const (
	appName      = "gen_jp_code"
	jsonDirName  = "data"
	jsonFileName = "jp.json"
	codeDirName  = "jp"
	codeFileName = "jp.gen.go"
	tmpl         = `// Code generated by go-holidays. DO NOT EDIT.
// Last Modified on %s
package jp
import (
	"time"
	"github.com/kenkyu392/go-holidays"
)
// Holidays in Japan
// %s ~ %s
var Holidays = holidays.Holidays{
	%s
}
`
)

func main() {
	log.SetPrefix(appName + " ")
	jsonPath, err := internal.FilePath(jsonDirName, jsonFileName)
	if err != nil {
		log.Fatal(err)
	}

	outPath, err := internal.FilePath(codeDirName, codeFileName)
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := os.ReadFile(jsonPath)
	if err != nil {
		log.Fatal(err)
	}

	data := new(internal.Data)
	if json.Unmarshal(jsonData, data); err != nil {
		log.Fatal(err)
	}

	code := genCode(data.Holidays, data.UpdateTime)
	if err := ioutil.WriteFile(outPath, []byte(code), os.ModePerm); err != nil {
		log.Fatal(err)
	}
	log.Printf("Generated '%s' from '%s'", outPath, jsonPath)
}

func genCode(hs holidays.Holidays, ut time.Time) string {
	sort.Slice(hs, func(i, j int) bool {
		return hs[i].Time.Before(hs[j].Time)
	})
	list := make([]string, 0)
	for _, h := range hs {
		list = append(list,
			fmt.Sprintf(
				`{Name: "%s", Time: time.Date(%d, %d, %d, 0, 0, 0, 0, JST)},`,
				h.Name, h.Time.Year(), h.Time.Month(), h.Time.Day(),
			),
		)
	}
	return fmt.Sprintf(tmpl,
		ut.Format("2006/01/02"),
		hs[0].Time.Format("2006/01/02"), hs[len(hs)-1].Time.Format("2006/01/02"),
		strings.TrimSpace(strings.Join(list, "\n\t")),
	)
}
