package setproxy

import "os/exec"

func SetProxy(proxyAddr string) {

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
