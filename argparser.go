package main

import "flag"

type ArgParser struct {
	Insert    *bool
	Completed *bool
	Update    *bool
	Find      *bool

	Tag *string
	Url *string
}

func GetArgs() *ArgParser {
	argparser := &ArgParser{
		Insert:    flag.Bool("i", false, "Create a task"),
		Completed: flag.Bool("c", false, "List completed tasks"),
		Update:    flag.Bool("u", false, "Update a task"),
		Find:      flag.Bool("f", false, "List not completed tasks or search for one by tag"),

		Tag: flag.String("T", "", "Tag"),
		Url: flag.String("U", "", "URL"),
	}
	flag.Parse()
	return argparser
}
