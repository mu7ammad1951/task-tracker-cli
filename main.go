package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("no command provided. usage - '$ <command> [args]'")
	}

	switch strings.ToLower(args[0]) {
	case "add":
		if len(args) < 2 {
			log.Fatal("missing arguments. usage - '$ add [task description]")
		}
		AddTask(args[1])

	case "update":
		if len(args) < 3 {
			log.Fatal("missing arguments. usage - '$ update [task id] [new task description]")
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal("invalid id")
		}
		UpdateTask(id, args[2])

	case "list":
		if len(args) == 2 {
			ListTasksFilter(args[1])
		} else {
			ListTasks()
		}

	case "mark-done":
		if len(args) < 2 {
			log.Fatal("missing arguments. usage - '$ mark-done [id]'")
		}

		id, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal("invalid id")
		}

		MarkDone(id)

	case "mark-in-progress":

		if len(args) < 2 {
			log.Fatal("missing arguments. usage - '$ mark-done [id]'")
		}

		id, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal("invalid id")
		}

		MarkInProgress(id)
	}

}
