package main

import (
        "bufio"
        "fmt"
        "io"
        "os"
)

func main() {
        if len(os.Args) != 2 {
                fmt.Println("引数が正しくありません。入力ファイル名を指定してください。")
                os.Exit(1)
        }

        // 入力ファイルを開く
        inputFile, err := os.Open(os.Args[1])
        if err != nil {
                fmt.Printf("入力ファイルを開けません: %v\n", err)
                os.Exit(1)
        }
        defer inputFile.Close()

        // 出力ファイルを開く
        outputFile, err := os.Create(os.Args[1] + ".hed")
        if err != nil {
                fmt.Printf("出力ファイルを作成できません: %v\n", err)
                os.Exit(1)
        }
        defer outputFile.Close()

        // 入力ファイルの情報を取得
        fileInfo, err := inputFile.Stat()
        if err != nil {
                fmt.Printf("ファイル情報を取得できません: %v\n", err)
                os.Exit(1)
        }
        fileSize := fileInfo.Size()

        // 出力ファイルに情報を書き込む
        _, err = outputFile.WriteString(os.Args[1] + "\n")
        if err != nil {
                fmt.Printf("出力ファイルに書き込めません: %v\n", err)
                os.Exit(1)
        }

        // 行数をカウント
        scanner := bufio.NewScanner(inputFile)
        lineCount := 0
        for scanner.Scan() {
                lineCount++
        }
        if err := scanner.Err(); err != nil {
                fmt.Printf("ファイルを読み取れません: %v\n", err)
                os.Exit(1)
        }

        // 行数を書き込む
        _, err = fmt.Fprintf(outputFile, "%d\n", lineCount-1)
        if err != nil {
                fmt.Printf("出力ファイルに書き込めません: %v\n", err)
                os.Exit(1)
        }

        // ファイルサイズを書き込む (改行なし)
        _, err = fmt.Fprintf(outputFile, "%d", fileSize)
        if err != nil {
                fmt.Printf("出力ファイルに書き込めません: %v\n", err)
                os.Exit(1)
        }
}
