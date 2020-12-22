package setproxy

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func SetProxy(proxyAddr string) {
	ShellProxy(proxyAddr)

	err := GitProxy(proxyAddr)
	if err != nil {
		log.Println(err)
	}
}

func ShellProxy(proxyAddr string) {
	err := os.Setenv("http_proxy", proxyAddr)
	if err != nil {
		fmt.Println(err)
	}

	err = os.Setenv("https_proxy", proxyAddr)
	if err != nil {
		fmt.Println(err)
	}

	err = os.Setenv("ftp_proxy", proxyAddr)
	if err != nil {
		fmt.Println(err)
	}

	err = os.Setenv("all_proxy", proxyAddr)
	if err != nil {
		fmt.Println(err)
	}

}

func GitProxy(proxyAddr string) error {
	err := exec.Command("git", "config", "--global", "http.proxy", proxyAddr).Run()
	if err != nil {
		return err
	}

	err = exec.Command("git", "config", "--global", "https.proxy", proxyAddr).Run()
	if err != nil {
		return err
	}

	return nil
}
