package volumes

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "path/filepath"
    "github.com/reboss/docker-shield/types"
)

const whiteListFile = "/usr/lib/docker/volume-wl.conf"

func ValidateBindMounts(body *types.ConfigWrapper) bool {
    whiteList, err := parseWhiteList()
    if err != nil {
        fmt.Println(err.Error())
        return false
    }
    for _, bind := range body.HostConfig.Binds {
        fromBind := strings.Split(bind, ":")[0]
        if isWhiteListed(fromBind, whiteList) {
            continue
        }
        if bind[0] == '/' {
            return false
        }
    }
    return true
}

func parseWhiteList() ([]string, error){
    file, err := os.Open(whiteListFile)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var whiteList []string

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        whiteList = append(whiteList, line)
    }
    if err := scanner.Err(); err != nil {
        return nil, err
    }
    return whiteList, nil
}

func isWhiteListed(bind string, whiteList []string) bool {
    for _, s := range whiteList {
        matched, err := filepath.Match(s, bind)
        if err != nil {
            fmt.Println(err.Error())
            return false
        }
        if matched {
            return true
        }
    }
    return false
}
