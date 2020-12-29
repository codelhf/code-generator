package main

import (
	"code-generator-go/server/common"
	"code-generator-go/server/controller"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

func main() {
	HandleRouter()
	http.HandleFunc("/user/logout", ExitSystem)
	http.Handle("/", http.FileServer(http.Dir("./dist")))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	CleanUpload("./file")
	//OpenBrowser(fmt.Sprintf("http://localhost:%s", port))
	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func HandleRouter() {
	//project
	http.HandleFunc("/project/list", controller.ProjectList)
	http.HandleFunc("/project/select", controller.ProjectSelect)
	http.HandleFunc("/project/insert", controller.ProjectInsert)
	http.HandleFunc("/project/update", controller.ProjectUpdate)
	http.HandleFunc("/project/delete", controller.ProjectDelete)
	http.HandleFunc("/project/lastId", controller.ProjectLastId)
	//projectDb
	http.HandleFunc("/projectDb/select", controller.ProjectDbSelect)
	http.HandleFunc("/projectDb/insert", controller.ProjectDbInsert)
	http.HandleFunc("/projectDb/update", controller.ProjectDbUpdate)
	//projectTable
	http.HandleFunc("/projectTable/all", controller.ProjectTableAll)
	http.HandleFunc("/projectTable/select", controller.ProjectTableSelect)
	http.HandleFunc("/projectTable/generate", controller.ProjectTableGenerate)
	//projectTemplate
	http.HandleFunc("/projectTemplate/list", controller.ProjectTemplateList)
	http.HandleFunc("/projectTemplate/select", controller.ProjectTemplateSelect)
	http.HandleFunc("/projectTemplate/insert", controller.ProjectTemplateInsert)
	http.HandleFunc("/projectTemplate/update", controller.ProjectTemplateUpdate)
	http.HandleFunc("/projectTemplate/delete", controller.ProjectTemplateDelete)
	//template
	http.HandleFunc("/template/list", controller.TemplateList)
	http.HandleFunc("/template/export", controller.TemplateExport)
	http.HandleFunc("/template/dump", controller.TemplateImport)
	http.HandleFunc("/template/select", controller.TemplateSelect)
	http.HandleFunc("/template/exists", controller.TemplateNameExists)
	http.HandleFunc("/template/insert", controller.TemplateInsert)
	http.HandleFunc("/template/update", controller.TemplateUpdate)
	http.HandleFunc("/template/delete", controller.TemplateDelete)
	//templateField
	http.HandleFunc("/templateField/list", controller.TemplateFieldList)
	http.HandleFunc("/templateField/select", controller.TemplateFieldSelect)
	http.HandleFunc("/templateField/insert", controller.TemplateFieldInsert)
	http.HandleFunc("/templateField/update", controller.TemplateFieldUpdate)
	http.HandleFunc("/templateField/delete", controller.TemplateFieldDelete)
	//templateGroup
	http.HandleFunc("/templateGroup/list", controller.TemplateGroupList)
	http.HandleFunc("/templateGroup/select", controller.TemplateGroupSelect)
	http.HandleFunc("/templateGroup/insert", controller.TemplateGroupInsert)
	http.HandleFunc("/templateGroup/update", controller.TemplateGroupUpdate)
	http.HandleFunc("/templateGroup/delete", controller.TemplateGroupDelete)
	//typeColumn
	http.HandleFunc("/typeColumn/list", controller.TypeColumnList)
	http.HandleFunc("/typeColumn/select", controller.TypeColumnSelect)
	http.HandleFunc("/typeColumn/insert", controller.TypeColumnInsert)
	http.HandleFunc("/typeColumn/update", controller.TypeColumnUpdate)
	http.HandleFunc("/typeColumn/delete", controller.TypeColumnDelete)
	//typeDB
	http.HandleFunc("/typeDb/allDbType", controller.AllDBType)
	//typeLanguage
	http.HandleFunc("/typeLanguage/list", controller.TypeLanguageList)
	http.HandleFunc("/typeLanguage/select", controller.TypeLanguageSelect)
	http.HandleFunc("/typeLanguage/insert", controller.TypeLanguageInsert)
	http.HandleFunc("/typeLanguage/update", controller.TypeLanguageUpdate)
	http.HandleFunc("/typeLanguage/delete", controller.TypeLanguageDelete)
}

func ExitSystem(w http.ResponseWriter, r *http.Request) {
	common.Success(w)
	os.Exit(0)
}

func OpenBrowser(uri string) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "start", uri)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("xdg-open", uri)
	} else if runtime.GOOS == "darwin" {
		cmd = exec.Command("open", uri)
	} else {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}
	return cmd.Start()
}

func CleanUpload(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("read dir errr %s", err.Error())
	}
	for _, file := range files {
		fullName := dir + "/" + file.Name()
		if file.Name() == ".gitkeep" {
			continue
		} else if file.IsDir() {
			err = CleanUpload(fullName)
		} else {
			err = os.Remove(fullName)
		}
	}
	return err
}
