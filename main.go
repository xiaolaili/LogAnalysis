package main

import (
	"log"
	"os"

	"LogAnalysis/analysis"
	_ "LogAnalysis/matchers"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	analysis.Run("data/QtHYDisplayClientDTY_20180913.log")

	log.Println("end")
}
