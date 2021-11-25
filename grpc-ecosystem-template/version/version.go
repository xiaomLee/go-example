package version

import "fmt"

// Version information.
var (
	BuildTS   = "None"
	GitHash   = "None"
	GitBranch = "None"
	Version   = "None"
	App       = "None"
)

func GetApp() string {
	if App != "None" {
		return fmt.Sprintf("%s-%s", App, GitBranch)
	}
	return App
}

// GetVersion Printer print build version
func GetVersion() string {
	if GitHash != "None" {
		h := GitHash
		if len(h) > 7 {
			h = h[:7]
		}
		return fmt.Sprintf("%s-%s", GitBranch, h)
	}
	return Version
}

func PrintFullVersionInfo() {
	fmt.Println("Application:      ", App)
	fmt.Println("Version:          ", GetVersion())
	fmt.Println("Git Branch:       ", GitBranch)
	fmt.Println("Git Commit:       ", GitHash)
	fmt.Println("Build Time:       ", BuildTS)
}
