package generator

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/slicktronix/go-bluetooth/gen/override"
	"github.com/slicktronix/go-bluetooth/gen/types"
)

var TplPath = "gen/generator/tpl/%s.go.tpl"

// rename variable name to avoid collision with Go languages
func renameReserved(varname string) string {
	switch varname {
	case "type":
		return "type1"
	default:
		return varname
	}
}

func getBaseDir() string {
	baseDir := os.Getenv("BASEDIR")
	if baseDir == "" {
		baseDir = "."
	}
	return baseDir
}

func getTplPath() string {
	return fmt.Sprintf("%s/%s", getBaseDir(), TplPath)
}

func loadtpl(name string) *template.Template {
	return template.Must(template.ParseFiles(fmt.Sprintf(getTplPath(), name)))
}

func prepareDocs(src string, skipFirstComment bool, leftpad int) string {
	return src
	// lines := strings.Split(src, "\n")
	// result := []string{}
	// // comment := "// "
	// comment := ""
	// prefixLen := leftpad + len(comment)
	// fmtt := fmt.Sprintf("%%%ds%%s", prefixLen)
	//
	// for _, line := range lines {
	// 	line = strings.Trim(line, " \t\r")
	// 	if len(line) == 0 {
	// 		continue
	// 	}
	//
	// 	result = append(result, fmt.Sprintf(fmtt, comment, line))
	// }
	// if skipFirstComment && len(result) > 0 && len(result[0]) > 3 {
	// 	result[0] = result[0][prefixLen:]
	// }
	// return strings.Join(result, "\n")
}

func getApiPackage(apiGroup *types.ApiGroup) string {
	apiName, ok := override.MapFile(apiGroup.FileName)
	if !ok {
		apiName = apiGroup.FileName
	}
	apiName = strings.ReplaceAll(apiName, "-api.txt", "")
	apiName = strings.ReplaceAll(apiName, "_api.txt", "")
	apiName = strings.ReplaceAll(apiName, "org.bluez.", "")
	apiName = strings.ReplaceAll(apiName, ".rst", "")
	apiName = strings.ReplaceAll(apiName, "-", "_")
	apiName = strings.ReplaceAll(apiName, " [experimental]", "")
	apiName = strings.ToLower(apiName)
	return apiName
}

func appendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}
