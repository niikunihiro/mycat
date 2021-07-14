package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// -n オプションを取得する
	boolNumber := flag.Bool("n", false, "a bool")

	flag.Parse()

	// プログラム引数を取得
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Please specify the file path.")
		os.Exit(1)
	}

	// プログラム引数はファイルパス
	// そのファイルを読み込んで標準出力にそのまま出力する
	for _, src := range args {
		err := readFile(src, *boolNumber)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// lineNumber 行番号
var lineNumber int

// readFile ファイルを読み込んで内容を出力する
func readFile(src string, line bool) error {
	sf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sf.Close()

	scanner := bufio.NewScanner(sf)
	for scanner.Scan() {
		if line {
			lineNumber++
			fmt.Printf("%d: ", lineNumber)
		}
		fmt.Println(scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return err
	}

	return nil
}
