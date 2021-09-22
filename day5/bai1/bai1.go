package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type iDockerClient interface {
	printListContainer()
	getPrintListMode() int
	setPrintListMode(mode int)
	startContainer(containerID string)
	stopContainer(containerID string)
}
type DockerClient struct {
	mode int
}

var myDockerClient *DockerClient

func GetInstanceDockerClient() iDockerClient {
	if myDockerClient == nil {
		myDockerClient = &DockerClient{mode: 1}
	}
	return myDockerClient
}

const containerRunningState = 1
const constainerExitedState = 2
const containerAllState = 3

func main() {
	myDockerClient := GetInstanceDockerClient()
	argsWithProg := os.Args

	if canStartContainer(argsWithProg) {
		myDockerClient.startContainer(getContainerID(argsWithProg[2]))
		return
	}

	if canStopContainer(argsWithProg) {
		myDockerClient.stopContainer(getContainerID(argsWithProg[2]))
		return
	}

	printDefaultActionResult()
}

/*
check start command can run
*/
func canStartContainer(argsWithProg []string) bool {
	return len(argsWithProg) >= 3 && strings.Contains("start", argsWithProg[1])
}

/*
check stop command can run
*/
func canStopContainer(argsWithProg []string) bool {
	return len(argsWithProg) >= 3 && strings.Contains("stop", argsWithProg[1])
}

/*
print default result when not input any parameter
*/
func printDefaultActionResult() {
	ch := make(chan string)
	go func(ch chan string) {
		// disable input buffering
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
		// do not display entered characters on the screen
		var b = make([]byte, 1)
		for {
			os.Stdin.Read(b)
			ch <- string(b)
		}
	}(ch)

	for {
		myDockerClient := GetInstanceDockerClient()
		myDockerClient.printListContainer()
		stdin, _ := <-ch

		if stdin == "l" {
			c := exec.Command("clear")
			c.Stdout = os.Stdout
			c.Run()
			mode := myDockerClient.getPrintListMode()
			mode = mode + 1
			myDockerClient.setPrintListMode(mode)
		}
		if myDockerClient.getPrintListMode() > 3 {
			mode := myDockerClient.getPrintListMode()
			myDockerClient.setPrintListMode(mode % 3)
		}
	}
}

func (dockerClient *DockerClient) getPrintListMode() int {
	return dockerClient.mode
}

func (dockerClient *DockerClient) setPrintListMode(mode int) {
	dockerClient.mode = mode
}

/*
print list of container
*/
func (dockerClient *DockerClient) printListContainer() {
	fmt.Println("----------------------------------------------------------------------------------------------------------")
	fmt.Println("Press keyboard l to switch mode, first time is running, second time is exited, third time is all")
	fmt.Println()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), getContainerFilter())
	if err != nil {
		panic(err)
	}

	fmt.Printf("%-62s | %-20s | %-10s | %-30s | %s", "Container ID", "Container Name", "Image Name", "Port (host:container)", "State")
	fmt.Println()
	for _, container := range containers {
		fmt.Printf("%s %s %s %s %s\n", container.ID, strings.TrimPrefix(container.Names[0], "/"), container.Image, getContainerPortInfo(container.Ports), container.State)
	}
}

/*
get container filter by selected mode
*/
func getContainerFilter() types.ContainerListOptions {
	myDockerClient := GetInstanceDockerClient()

	switch myDockerClient.getPrintListMode() {
	case containerRunningState:
		filters := filters.NewArgs()
		filters.Add("status", "running")
		return types.ContainerListOptions{Filters: filters}
	case constainerExitedState:
		filters := filters.NewArgs()
		filters.Add("status", "exited")
		return types.ContainerListOptions{Filters: filters}
	case containerAllState:
		return types.ContainerListOptions{All: true}
	default:
		filters := filters.NewArgs()
		filters.Add("status", "running")
		return types.ContainerListOptions{Filters: filters}
	}
}

/*
format port information
*/
func getContainerPortInfo(ports []types.Port) string {
	containerPortInfo := ""
	for _, port := range ports {
		containerPortInfo = containerPortInfo + port.IP + ":" + strconv.Itoa(int(port.PublicPort)) + "->" + strconv.Itoa(int(port.PrivatePort)) + "/tcp" + ", "
	}
	return strings.TrimRight(containerPortInfo, ", ")
}

/*
stop container by container id
*/
func (dockerClient *DockerClient) stopContainer(containerID string) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStop(context.Background(), containerID, nil); err != nil {
		log.Printf("Unable to stop container %s: %s", containerID, err)
	}
}

/*
start container by containerId
*/
func (dockerClient *DockerClient) startContainer(containerID string) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{}); err != nil {
		log.Printf("Unable to start container %s: %s", containerID, err)
	}
}

/*
get containerID from container name
*/
func getContainerID(containerName string) string {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if container.Names[0] == containerName {
			return container.ID
		}
	}
	return containerName
}
