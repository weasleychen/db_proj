package handler

import (
	"db_proj/define"
	"flag"
	"os"
)

func init() {
	flag.CommandLine.Init(define.ProjectName, flag.ContinueOnError)
}

func ParseCommandLine() {
	flag.IntVar(&define.Port, "p", define.Port, "the port server listen on")
	flag.BoolVar(&define.UseSwagger, "s", define.UseSwagger, "choose whether to use swagger")
	flag.BoolVar(&define.DebugMode, "d", define.DebugMode, "choose whether to open DebugMode")

	help := flag.Bool("help", false, "print help")

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}
}
