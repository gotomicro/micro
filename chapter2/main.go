package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"runtime"
)

var (
	buildName        = "unknown"
	buildGitRevision = "unknown"
	buildUser        = "unknown"
	buildHost        = "unknown"
	buildStatus      = "unknown"
	buildTime        = "unknown"
)

var (
	version bool
	run     bool
)

func init() {
	pflag.BoolVarP(&version, "version", "v", false, `查看版本号`)
	pflag.BoolVarP(&run, "run", "r", false, `运行程序`)
	pflag.Parse()
}

func main() {
	if version == true {
		fmt.Println(LongForm())
	}
	if run == true {
		fmt.Println("go to micro")
	}
}

func LongForm() string {
	return fmt.Sprintf(`Name: %v
GitRevision: %v
User: %v@%v
GolangVersion: %v
BuildStatus: %v
BuildTime: %v
`,
		buildName,
		buildGitRevision,
		buildUser,
		buildHost,
		runtime.Version(),
		buildStatus,
		buildTime)
}
