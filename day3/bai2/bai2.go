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

var mode int = 1

const containerRunningState = 1
const constainerExitedState = 2
const containerAllState = 3

func main() {
	argsWithProg := os.Args

	if len(argsWithProg) >= 3 && strings.Contains("start", argsWithProg[1]) {
		startContainer(argsWithProg[1])
		return
	}

	if len(argsWithProg) >= 3 && strings.Contains("stop", argsWithProg[1]) {
		stopContainer(argsWithProg[1])
		return
	}

	printDefaultActionResult()
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
		printListContainer()
		stdin, _ := <-ch

		if stdin == "l" {
			c := exec.Command("clear")
			c.Stdout = os.Stdout
			c.Run()
			mode++
		}
		if mode > 3 {
			mode = mode % 3
		}
	}
}

/*
print list of container
*/
func printListContainer() {
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
	switch mode {
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
func stopContainer(containerID string) {
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
func startContainer(containerID string) {
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
