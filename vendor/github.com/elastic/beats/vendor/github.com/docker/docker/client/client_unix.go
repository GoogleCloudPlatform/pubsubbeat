// +build linux freebsd solaris openbsd darwin netbsd

package client

// DefaultDockerHost defines os specific default if DOCKER_HOST is unset
const DefaultDockerHost = "unix:///var/run/docker.sock"
