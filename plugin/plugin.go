package plugin

import (
    "bytes"
//    "net/url"
    "encoding/json"
    "github.com/docker/go-plugins-helpers/authorization"
    "github.com/reboss/docker-shield/types"
    "github.com/reboss/docker-shield/securityopts"
)

//var createAPI = regexp.MustCompile(`/v.*/containers/create`)
//var execAPI = regexp.MustCompile(`/v.*/containers/.+/exec`)

func NewDockerShield() (*DockerShield) {
    return &DockerShield{}
}

// Implements authorization.Plugin interface
type DockerShield struct {}

func (p *DockerShield) AuthZReq(req authorization.Request) authorization.Response {

    //uri, err := url.QueryUnescape(req.RequestURI)
    //if err != nil {
    //    return authorization.Response{Err: err.Error()}
    //}

    if req.RequestMethod == "POST" && req.RequestBody != nil {
        body := &types.ConfigWrapper{}
        if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(body); err != nil {
            return authorization.Response{Err: err.Error()}
        }

        if !securityopts.ValidateSecOpts(body) {
            return authorization.Response{Msg: "Not allowed to modify security options"}
        }
        if !securityopts.ValidatePrivileges(body) {
            return authorization.Response{Msg: "Privileged containers not allowed"}
        }
    }
    return authorization.Response{Allow: true}
}

func (p *DockerShield) AuthZRes(req authorization.Request) authorization.Response {
    return authorization.Response{Allow: true}
}
