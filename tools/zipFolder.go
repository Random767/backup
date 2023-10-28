package tools

import (
    "fmt"
    "os"
    "log"
    "io"
    "archive/zip"
    "path/filepath"
    "errors"
)

func ZipFolder(folderName string, outputFilename string) error {
    if outputFilename == "" {
        return errors.New("Um valor no outputFilename é esperado")
    }
    if fileExists(folderName + "/" + outputFilename + ".zip") {
        return errors.New("Arquivo "+ folderName + "/" + outputFilename +" já existe")
    }

    files, err := readDirRecursive(folderName)
    if err != nil {
        return err
    }

    fmt.Println("[Start] Criando arquivivo zip")
    zipFolder, err := os.Create(folderName + "/" + outputFilename + ".zip")

    if err != nil {
        return err
    }

    zipWriter := zip.NewWriter(zipFolder)
    defer zipFolder.Close()
    defer zipWriter.Close()
    fmt.Printf("[Stat] Iniciando compactação de %d arquivos...\n", len(files))
    for i := 0; i < len(files); i++ {
        fmt.Println("[Stat] Compactadando o arquivo " + files[i])
        file, err := os.Open(files[i])
        if err != nil {
            fmt.Println(err) // Revise isso aqui depois
            continue
        }
        defer file.Close()
        if isDir(files[i]) {
            continue
        }
        fileZip, err := zipWriter.Create(files[i])
        if err != nil {
            fmt.Println(err) // Revise isso também mais tarde
            continue
        }

        if _, err := io.Copy(fileZip, file); err != nil {
            fmt.Println(err)// Revisa tudo logo de uma vez
            continue
        }
        fmt.Println("[Success] Arquivo "+ files[i] +" compactado com sucesso")
    }
    return nil
}

func isDir(filename string) bool {
    info, err := os.Stat(filename)
    if err != nil {
        log.Fatal(err)
    }

    if info.IsDir() {
        return true
    }

    return false
}

func fileExists(filename string) bool {
    _, err := os.Stat(filename)

    if os.IsNotExist(err) {
        return false
    } else {
        return true
    }
}

func readDirRecursive(folderName string) ([]string, error){
    var files []string
    err := filepath.Walk(folderName,
        func(path string, info os.FileInfo, err error) error {
            if err != nil {
                return err
            }
            files = append(files, path)
            return nil
        })
    if err != nil {
        return files, err
    }
    return files, nil
}
