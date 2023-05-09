package types

import (
    "github.com/docker/engine-api/types/container"
)

type ConfigWrapper struct {
    HostConfig *container.HostConfig "json:HostConfig,omitempty"
    Privileged bool "json:Privileged,omitempty"
}
