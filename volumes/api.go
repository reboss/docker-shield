package volumes

import (
    "github.com/reboss/docker-shield/types"
)

func ValidateBindMounts(body *types.ConfigWrapper) bool {
    for _, bind := range body.HostConfig.Binds {
        if bind[0] == '/' {
            return false
        }
    }
    return true
}
