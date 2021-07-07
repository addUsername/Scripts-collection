package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	//"github.com/dlclark/regexp2"
)

//get-filehash -a SHA512 go.mod | Format-list

var oath = getPowershellPath2()
var wDir = getWorkingDirectory2()
var numFiles2 = 1000
var sizeFile2 = 10 * 1024 * 1024
var fileNames2 = ""
var mainProgramFile = "sha512sumWin.exe"

func main() {

	if runtime.GOOS == "windows" {
		fmt.Println("You are on Windows, good job")
	}
	fmt.Println("num files: " + strconv.Itoa(numFiles2) + "\nsize file: " + strconv.Itoa(int(sizeFile2)))
	generateFiles3()

	fmt.Println(goRoutines2())
	fmt.Println(singleThreadGo2())
	fmt.Println(singleThreadWin2())

}

func singleThreadWin2() float64 {

	fmt.Println("single thread WIN")
	command := "get-filehash -a SHA512 -Path '" + wDir + "\\benchmark\\random\\*" + "' | Format-List"
	call := exec.Command(oath, "Measure-Command {"+command+"  | Out-Default}")

	out, err := call.CombinedOutput()

	if err != nil {
		log.Fatalln(err)
	}
	time := extractTime2(string(out))
	return time
}

func singleThreadGo2() float64 {

	fmt.Println("single thread GO")
	command := ".\\" + mainProgramFile + " 1 1 " + fileNames2
	call := exec.Command(oath, "Measure-Command {"+command+"  | Out-Default}")

	out, err := call.CombinedOutput()
	if err != nil {
		log.Fatalln(out)
	}
	//fmt.Println(string(out))
	time := extractTime2(string(out))
	return time
}

func goRoutines2() float64 {

	fmt.Println("GOroutines")
	command := ".\\" + mainProgramFile + " -1 2 " + fileNames2
	call := exec.Command(oath, "Measure-Command {"+command+"  | Out-Default}")

	out, err := call.CombinedOutput()
	if err != nil {
		log.Fatalln(string(out))
	}
	//fmt.Println(string(out))
	time := extractTime2(string(out))
	return time
}

func getPowershellPath2() string {

	path, err := exec.LookPath("powershell.exe")
	if err != nil {
		log.Fatal(err)
	}
	return path
}

func generateFiles3() {
	fmt.Println("generating files")
	var numFilesSlice []string
	var file string
	firstCommand := " $out = new-object byte[] " + strconv.Itoa(int(sizeFile2)) + "; (new-object Random).NextBytes($out);"

	for i := 1; i <= numFiles2; i++ {
		file = wDir + "\\benchmark\\random\\" + strconv.Itoa(i)
		secondCommand := " [IO.File]::WriteAllBytes('" + file + "', $out);"
		_, err := exec.Command(oath, firstCommand, secondCommand).CombinedOutput()
		if err != nil {
			fmt.Println(err)
			log.Fatalln(err)
		}
		numFilesSlice = append(numFilesSlice, file)
	}
	fileNames2 = strings.Join(numFilesSlice, " ")
}
func extractTime2(time string) float64 {
	res := strings.Split(time, "TotalMilliseconds : ")
	if len(res) < 1 {
		return 0.
	}
	trim := strings.TrimSpace(strings.ReplaceAll(res[1], ",", "."))
	milliSec, err := strconv.ParseFloat(trim, 64)
	if err != nil {
		log.Fatalln(err)
	}
	return milliSec
}
func getWorkingDirectory2() string {
	currentRoute, errRoute := os.Getwd()
	if errRoute != nil {
		log.Fatal("Unable to access the test's working directory.")
	}
	return currentRoute
}
