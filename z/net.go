package z

import (
	"github.com/zhang201702/zhang/zconfig"
	"os/exec"
	"runtime"
)

func GetUrl() string {
	return "http://localhost:" + zconfig.Conf.GetString("port", "80")
}

func OpenBrowse(url string) {
	sysType := runtime.GOOS
	if sysType == "windows" {
		exec.Command(`cmd`, `/c`, `start`, url).Start()

	} else {
		exec.Command(`open`, url).Start()
	}
}
