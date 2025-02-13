package main

import (
	"CocosCreator3DPlayable/fileEmbed"
	"CocosCreator3DPlayable/mix/engine"
	"CocosCreator3DPlayable/mix/res"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Jecced/go-tools/src/ak"
	_ "github.com/Jecced/go-tools/src/ak"
	"github.com/Jecced/go-tools/src/fileutil"
)

var (
	mobileDir  = "E:\\Projects\\Cocos\\Taxi"
	outDir     = ""
	htmlFile   = ""
	tmpOutFile = ""
)

func main() {
	//build, _ := os.Executable()
	build := mobileDir
	//build = filepath.Dir(build)
	buildPath := flag.String("path", build, "build path")
	flag.Parse()
	if !fileutil.PathExists(*buildPath) {
		fmt.Printf("%s Path Error\n", *buildPath)
		return
	}
	mobileDir = *buildPath + ak.PS + "build/web-mobile"
	outDir = *buildPath + ak.PS + "build/PS"
	htmlFile = outDir + ak.PS + "index.html"
	tmpOutFile = outDir + ak.PS + "/index.html"
	fileutil.ClearDir(outDir)
	fmt.Printf("copy project from:%s --->to:%s\n", mobileDir, outDir)
	var err = fileutil.DirCopy(mobileDir, outDir)
	if err != nil {
		return
	}
	index, err := fileEmbed.Template.ReadFile("template/index.js")
	if err != nil {
		return
	}
	fileutil.WriteData(index, outDir+ak.PS+"index.js")
	indexhtml, err := fileEmbed.Template.ReadFile("template/index.html")
	if err != nil {
		return
	}
	fileutil.WriteData(indexhtml, outDir+ak.PS+"index.html")
	application, err := fileEmbed.Template.ReadFile("template/application.js")
	if err != nil {
		return
	}
	fileutil.WriteData(application, outDir+ak.PS+"application.js")
	systembundle, err := fileEmbed.Template.ReadFile("template/src/system.bundle.js")
	if err != nil {
		return
	}
	fileutil.WriteData(systembundle, outDir+ak.PS+"src/system.bundle.js")

	plugin, err := fileEmbed.Plugin.ReadFile("plugin/new-res-loader.js")
	if err != nil {
		return
	}
	fileutil.WriteData(plugin, outDir+ak.PS+"new-res-loader.js")
	path, _ := os.Executable()
	dir := filepath.Dir(path)
	fmt.Printf("OS:%s\n", dir)
	htmlContent, err := fileutil.ReadText(htmlFile)
	if err != nil {
		fmt.Println(err, htmlContent)
	}
	engine.Mix(outDir, &htmlContent)
	res.Mix(outDir, &htmlContent)
	fileutil.ClearDir(outDir)
	fileutil.WriteText(htmlContent, tmpOutFile)
	fileutil.DelEmptyDir(outDir)
}
