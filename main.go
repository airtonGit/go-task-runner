package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/airtongit/monologger"
)

var log *monologger.Log

func check(err error) {
	if err != nil {
		log.Error(err.Error())
	}
}

func main() {
	log, _ = monologger.New(os.Stdout, "runner ", true)

	base, err := os.Getwd()
	check(err)
	baseConvertido := filepath.Dir(base + "/docker/mysql-data-crope/")
	log.Info(baseConvertido)
	cmdDockerMySQLOptionsVolume := fmt.Sprintf("-v%s:/var/lib/mysql", baseConvertido)

	mysqlcmd := exec.Command("docker", "run", cmdDockerMySQLOptionsVolume, "-p3306:3308", "mysql:5.6")

	mysqlcmdOut, err := mysqlcmd.StdoutPipe()
	check(err)
	mysqlcmdErr, err := mysqlcmd.StderrPipe()
	check(err)
	log.Info("MySQL...")
	err = mysqlcmd.Start()
	check(err)
	qtd, err := io.Copy(os.Stdout, mysqlcmdOut)
	check(err)
	_, err = io.Copy(os.Stdout, mysqlcmdErr)
	mysqlcmd.Wait()
	log.Info("Bytes ", qtd)
}
