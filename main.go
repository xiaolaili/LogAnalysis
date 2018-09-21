package main

import (
	"LogAnalysis/analysis"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	analysis.Run("data/QtHYDisplayClientDTY_20180913.log")
}
