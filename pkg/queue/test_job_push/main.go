package main

import (

	que "planet/pkg/queue"
)

func main() {

	que.NewConfig([]string{"10.10.18.130","5672","guest","guest"}).Push("TestJob","xxxxxx")

}