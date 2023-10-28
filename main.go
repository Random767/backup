package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "flag"

    "backup/tools"
    //"time"
)

var debugState bool = true
var backupFolder string = "./backup"
var defaultBackupDirectory string = "./"

func main() {
    fmt.Println("[Starting] Servidor SCB está iniciando")
    debug("Checking", "Checando se a pasta "+ backupFolder +" existe")
    if checkFileExists("./backup") {
        debug("Check", "Aquivo "+ backupFolder +" existe. Continuando...")
    } else {
        debug("Creating", "Arquivo "+ backupFolder +" não existe. Criando arquivo...")
        createFolder(backupFolder)
        debug("Created", "Pasta "+ backupFolder +" foi criada com sucesso. Continuando...")
    }

    listCmd := flag.NewFlagSet("list", flag.ExitOnError)
    listFilename := listCmd.String("filename", "", "Nome da pasta")

    makeCmd := flag.NewFlagSet("make", flag.ExitOnError)
    makeFilename := makeCmd.String("filename", "", "Nome da pasta")

    restoreCmd := flag.NewFlagSet("restore", flag.ExitOnError)
    restoreFilename := restoreCmd.String("filename", "", "Nome da pasta")

    deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
    deleteFilename := deleteCmd.String("filename", "", "Nome da pasta")
    if len(os.Args) < 2 {
        fmt.Println("[Info] Esperado um desses subcomandos: list, get, make, make-all, restore ou delete")
        os.Exit(1)
    }

    switch os.Args[1] {
        case "list":
            debug("Subcommand", "Executando o subcomando list...")
            listCmd.Parse(os.Args[2:])
            if *listFilename == "" {
                debug("Listing", "Iniciando listagem de arquivos para "+ backupFolder)
                fmt.Println("[List] Listando os backups no diretório "+ backupFolder)
                files := readDir(backupFolder)
                for i := 0; i < len(files); i++ {
                    fmt.Println("      [File]: "+ files[i])
                }
            } else {
                debug("Checking", "Checando a existencia da pasta " + *listFilename)
                // Checar também se o arquivo é uma pasta
                if !checkFileExists(backupFolder + "/" + *listFilename) {
                    fmt.Println("[Error] Nome de arquivo inválido")
                    return
                } else if !isDir(backupFolder + "/" + *listFilename) {
                    fmt.Println("[Error] O arquivo não é um diretório")
                    return
                }
                debug("Listing", "Iniciando listagem de arquivos para " + backupFolder + "/" + *listFilename)
                files := readDir(backupFolder + "/" + *listFilename)
                for i := 0; i < len(files); i++ {
                    fmt.Println("      [File]: "+ files[i])
                }
            }
        case "make":
            debug("Subcommand", "Executando o subcomando make...")
            makeCmd.Parse(os.Args[2:])
            if *makeFilename == "" {
                debug("Start", "Iniciando backup do diretório " + defaultBackupDirectory)
            }
        case "restore":
            debug("Subcommand", "Executando o subcomando restore...")
            restoreCmd.Parse(os.Args[2:])
            if *restoreFilename == "" {
                fmt.Println("[Error] Uso do subcomando restore inválido. Utilize restore -h para saber mais")
            }
        case "delete":
            debug("Subcommand", "Executando o subcomando delete")
            deleteCmd.Parse(os.Args[2:])
            if *deleteFilename == "" {
                fmt.Println("[Error] Uso do subcomando delete inválido. Utilize delete -h para saber mais")
            }

        default:
            fmt.Println("[Info] Esperado um desses subcomandos: list, get, make, make-all, restore ou delete")
    }


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

func checkFileExists(filename string) bool {
    _, err := os.Stat(filename)

    if os.IsNotExist(err) {
        return false
    } else {
        return true
    }
}

func createFolder(filename string) {
    if err := os.Mkdir(filename, os.ModePerm); err != nil {
        log.Fatal(err)
    }
}

func readDir(dirname string) []string{
    files, err := ioutil.ReadDir(dirname)
    if err != nil {
        log.Fatal(err)
    }
    dirContent := make([]string, len(files)-len(files)) // Gambiarra momento
    fmt.Println("[Info] Quantidade de aquivos no diretório "+ dirname +":", len(files))
    for _, file := range files {
        dirContent = append(dirContent, file.Name())
    }
    return dirContent
}

func debug(action string, message string) {
    if debugState {
        fmt.Printf("- Debug [%s]: %s\n", action, message)
    }
}
