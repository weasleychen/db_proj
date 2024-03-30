package cmdlinehandler

import "flag"
import "db_proj/define"

func init() {
	flag.CommandLine.Init(define.ProjectName, flag.ContinueOnError)
}

func ParseCommandLine() int {
	port := flag.Int("p", define.DefaultPort, "the port server listen on")

	flag.Parse()
	
	return *port
}