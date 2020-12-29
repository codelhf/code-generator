package codegen

import (
	"bytes"
	"code-generator-go/server/util"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func GenerateToCode(tmplName, tmplText string, tmplData interface{}) ([]byte, error) {
	tmplCode, err := template.New(tmplName).Funcs(GetFuncMap()).Parse(tmplText)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := tmplCode.Execute(&buf, tmplData); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func GetFuncMap() map[string]interface{} {
	return template.FuncMap{
		"now": func() string {
			return time.Now().Format(time.RFC3339)
		},
		"add": func(a, b int) int {
			return a + b
		},
		"sub": func(a, b int) int {
			return a - b
		},
		"size": func(arr []ColumnInfo) int {
			return len(arr)
		},
		"concat": func(strList ...string) string {
			// 定义一个字节缓冲,快速拼接字符串
			var result bytes.Buffer
			for _, value := range strList {
				result.WriteString(value)
			}
			return result.String()
		},
		"equals": func(str1, str2 string) bool {
			return strings.EqualFold(str1, str2)
		},
		"isTrue": func(str interface{}) bool {
			switch str.(type) {
			case string:
				return str == "true"
			case bool:
				return str == true
			default:
				return false
			}
		},
		"isFalse": func(str interface{}) bool {
			switch str.(type) {
			case string:
				return str == "false"
			case bool:
				return str == false
			default:
				return false
			}
		},
		"firstToUpperCase": func(str string) string {
			return util.FirstToUpperCase(str)
		},
		"subLastCharacter": func(str string) string {
			return util.SubLastCharacter(str)
		},
	}
}

func CreateFile(directory, packageName, fileName, fileEncode, fileContent string, isOverride bool) string {
	// 先创建修改目录分隔符
	directory = strings.ReplaceAll(directory, "\\\\", "/")
	directory = strings.ReplaceAll(directory, "\\", "/")
	directory = getGeneratePath(directory, packageName)
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		// 必须分成两步
		// 先创建文件夹
		err := os.MkdirAll(directory, 0777)
		checkErr(err)
		// 再修改权限
		err = os.Chmod(directory, 0777)
		checkErr(err)
	}
	exists, err := isFileExist(directory + "/" + fileName)
	checkError(err)
	if exists && !isOverride {
		fileName = getUniqueFileName(directory, fileName)
	}
	file, err := os.Create(directory + "/" + fileName)
	checkError(err)
	fileContent, err = UTF8To(Charset(fileEncode), fileContent)
	checkErr(err)
	file.WriteString(fileContent)
	file.Close()
	fmt.Println("[INFO] generating " + directory)
	fmt.Println("[INFO] generating " + fileName)
	resultString := "[INFO] generating " + directory + "\n"
	resultString += "[INFO] generating " + fileName
	return resultString
}

func getGeneratePath(directory, packageName string) string {
	if !strings.HasSuffix(directory, "/") {
		directory = directory + "/"
	}
	return directory + package2Path(packageName)
}

func package2Path(packageName string) string {
	if util.IsBlank(packageName) {
		return util.EMPTY
	}
	packages := strings.Split(packageName, ".")
	return strings.Join(packages, "/")
}

//生成唯一文件名
func getUniqueFileName(directory, fileName string) string {
	var uniqueName string
	// try up to 1000 times to generate a unique file name
	for i := 1; i < 1000; i++ {
		testName := fileName + "." + strconv.Itoa(i)
		exists, err := isFileExist(directory + "/" + testName)
		checkError(err)
		if !exists {
			uniqueName = testName
			break
		}
	}
	if util.IsBlank(uniqueName) {
		panic("Cannot generate unique file name in directory {" + directory + "}")
	}
	return uniqueName
}

//判断文件文件夹是否存在
func isFileExist(path string) (bool, error) {
	fileInfo, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false, nil
	}
	//我这里判断了如果是0也算不存在
	if fileInfo == nil || fileInfo.Size() == 0 {
		return false, nil
	}
	if err == nil {
		return true, nil
	}
	return false, err
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
