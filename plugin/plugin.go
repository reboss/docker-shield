package plugin

import (
    "bytes"
    "regexp"
    "fmt"
    "net/url"
    "encoding/json"
    "github.com/docker/go-plugins-helpers/authorization"
    "github.com/reboss/docker-shield/types"
    "github.com/reboss/docker-shield/securityopts"
    "github.com/reboss/docker-shield/volumes"
)

var createAPI = regexp.MustCompile(`/v.*/containers/create`)
var execAPI = regexp.MustCompile(`/v.*/containers/.+/exec`)

func NewDockerShield() (*DockerShield) {
    return &DockerShield{}
}

// Implements authorization.Plugin interface
type DockerShield struct {}

func handleCreateAPI(body *types.ConfigWrapper) error {
    if !securityopts.ValidateSecOpts(body) {
        return fmt.Errorf("Not allowed to modify security options")
    }
    if !securityopts.ValidatePrivileges(body) {
        return fmt.Errorf("Privileged containers not allowed")
    }
    if !volumes.ValidateBindMounts(body) {
        return fmt.Errorf("Not allowed to bind mount directories from the host")
    }
    return nil
}

func handleExecAPI(body *types.ConfigWrapper) error {
    if !securityopts.ValidatePrivileges(body) {
        return fmt.Errorf("Privileged containers not allowed")
    }
    return nil
}

func (p *DockerShield) AuthZReq(req authorization.Request) authorization.Response {

    uri, err := url.QueryUnescape(req.RequestURI)
    if err != nil {
        return authorization.Response{Err: err.Error()}
    }

    if req.RequestMethod == "POST" && req.RequestBody != nil {
        body := &types.ConfigWrapper{}
        if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(body); err != nil {
            return authorization.Response{Err: err.Error()}
        }
        if createAPI.MatchString(uri) {
            fmt.Println("Here we are")
            if err = handleCreateAPI(body); err != nil {
                return authorization.Response{Msg: err.Error()}
            }
        } else if execAPI.MatchString(uri) {
            fmt.Println("we are execing in now")
            if err = handleExecAPI(body); err != nil {
                return authorization.Response{Msg: err.Error()}
            }
        }
    }
    return authorization.Response{Allow: true}
}

func (p *DockerShield) AuthZRes(req authorization.Request) authorization.Response {
    return authorization.Response{Allow: true}
}
