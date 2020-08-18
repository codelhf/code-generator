package util

import (
	"encoding/json"
	"github.com/tealeg/xlsx"
	"io"
	"net/http"
	"os"
)

func WriteExcel(w http.ResponseWriter, fileName, sheetName, heads string, obj interface{}) {
	// 设置响应头
	w.Header().Set("Content-Type", "application/vnd.ms-excel;charset=UTF-8")
	w.Header().Add("Content-Disposition", "attachment;filename="+fileName+".xls")
	file := xlsx.NewFile()
	// 创建一个工作表
	sheet, err := file.AddSheet(sheetName)
	checkErr(err)
	// 设置标题
	// json str to map
	tHeads := make([]map[string]string, 0)
	err = json.Unmarshal([]byte(heads), &tHeads)
	checkErr(err)
	row, err := sheet.AddRowAtIndex(0)
	checkErr(err)
	row.SetHeightCM(0.5)
	style := xlsx.NewStyle()
	style.Fill = xlsx.Fill{PatternType: "none", FgColor: "", BgColor: "#eee"}
	for i := 0; i < len(tHeads); i++ {
		head := row.AddCell()
		head.SetStyle(style)
		head.Value = tHeads[i]["name"]
	}
	// struct to json str
	jsonStr, err := json.Marshal(obj)
	checkErr(err)
	// 设置单元格
	// json str to map
	tBodys := make([]map[string]interface{}, 0)
	err = json.Unmarshal(jsonStr, &tBodys)
	checkErr(err)
	for i := 0; i < len(tBodys); i++ {
		row, err := sheet.AddRowAtIndex(i + 1)
		checkErr(err)
		for j := 0; j < len(tHeads); j++ {
			field := tHeads[j]["field"]
			cell := row.AddCell()
			cell.SetValue(tBodys[i][field])
		}
	}
	// 根据指定路径保存文件
	err = file.Write(w)
	checkErr(err)
}

func ReadExcel(r *http.Request, uploadFile, sheetName, heads string, obj interface{}) {
	// 默认能上传文件的大小为32MB，可以通过 iris.WithPostMaxMemory(maxSize) 设置，
	// 比如10MB = 10 * 1024 * 1024 =maxSize，iris.WithPostMaxMemory(maxSize)
	// 3个参数，1是文件句柄，2是文件头，3是错误信息
	mFile, mFileReader, err := r.FormFile(uploadFile)
	checkErr(err)
	defer mFile.Close()
	// 创建1个相同名字的文件，存放在upload目录里面
	// 假定本地已经有名字为 upload 的目录，没有的话会报错
	fileName := mFileReader.Filename
	out, err := os.OpenFile("./dist/upload/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	checkErr(err)
	defer out.Close()
	io.Copy(out, mFile)
	// 打开文件
	file, err := xlsx.OpenFile("./dist/upload/" + fileName)
	checkErr(err)
	sheet := file.Sheet[sheetName]
	rows := sheet.Rows
	cell0 := rows[0].Cells
	// 读取标题并修改为对应字段
	// json str to map
	tHeads := make([]map[string]string, 0)
	err = json.Unmarshal([]byte(heads), &tHeads)
	for i := 0; i < len(tHeads); i++ {
		name := tHeads[i]["name"]
		field := tHeads[i]["field"]
		for j := 0; j < len(cell0); j++ {
			head := cell0[j].Value
			if head == name {
				cell0[j].Value = field
			}
		}
	}
	// 读取内容
	objects := make([]map[string]interface{}, 0)
	for i := 0; i < len(rows)-1; i++ {
		cells := rows[i+1].Cells
		object := make(map[string]interface{})
		for j := 0; j < len(cells); j++ {
			head := cell0[j].Value
			if head == "type" {
				object[head], err = cells[j].Int()
			} else {
				object[head] = cells[j].Value
			}
		}
		objects = append(objects, object)
	}
	// map to json str
	jsonStr, err := json.Marshal(objects)
	checkErr(err)
	// json str to struct
	err = json.Unmarshal(jsonStr, obj)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
		return
	}
}
