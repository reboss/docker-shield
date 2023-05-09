package securityopts

import (
    "strings"
    "github.com/reboss/docker-shield/types"
)

func validateSecOpts(body *types.ConfigWrapper) bool {
    for _, secopt := range body.HostConfig.SecurityOpt {
        apparmor := strings.Contains(secopt, "apparmor")
        seccomp := strings.Contains(secopt, "seccomp")
        if apparmor || seccomp {
            return false
        }
    }
    return true
}

func ValidateSecOpts(body *types.ConfigWrapper) bool {
    return validateSecOpts(body)
}

func ValidatePrivileges(body *types.ConfigWrapper) bool {
    return !(body.Privileged || body.HostConfig.Privileged)
}
