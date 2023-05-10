# DockerShield

Docker Shield is an authorization plugin for Docker that addresses common security configuration issues with Docker.  Essentially, it creates a locked-down version of Docker that removes a lot of the attack surface, while still allowing for Docker's feature dense usages.

Docker Shield ensures the following:
1. Custom security profiles cannot be added with `--security-opts apparmor=unconfined`, for example
2. Disallows bind mounts from the host
3. Prevents privileged mode

Work in progress:
1. Enforcing that users start containers and exec into them with their own user ids, and not the default root user (may not be possible with a plugin)
2. All containers are started with `--no-new-privileges` option, which ensures that users can't escalate privileges within a container, for example, with a setuid binary (this can be achieved by default in the /etc/daemon.json file
3. Create a white list for bind mounts
4. Create a white list for apparmor/seccomp/selinux profiles
5. Prevent `--cap-add` option from being used

## Installation 

Installation is pretty straight forward:

```
make
sudo make install
```

Then we need to update our systemd service file for docker.  Update the following line in /lib/systemd/system/docker.service:
```
- ExecStart=/usr/bin/dockerd --debug -H fd:// --containerd=/run/containerd/containerd.sock
+ ExecStart=/usr/bin/dockerd --debug -H fd:// --containerd=/run/containerd/containerd.sock --authorization-plugin=docker-shield
```

to uninstall:
```
sudo make uninstall
```
