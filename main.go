package main

import(
	"os"
	"io/ioutil"
)

func getperms(path string) os.FileMode {
	info, _ := os.Stat(os.Args[0])
	perms := info.Mode()
	return perms
}

func copymyself(path string) {
	dat, _ := ioutil.ReadFile(os.Args[0])
	err := ioutil.WriteFile("_"+os.Args[0]+,dat,getperms(os.Args[0]))
	if err != nil {
		os.Exit(1)
	}
}


func main() {
	copymyself("test.go")
}
