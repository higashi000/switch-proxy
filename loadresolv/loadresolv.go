package loadresolv

import (
	"bufio"
	"os"
	"strings"
)

func LoadResolv(fp string) (string, error) {
	file, err := os.Open(fp)

	if err != nil {
		return "", err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	var searchs string

	for sc.Scan() {
		if strings.Index(sc.Text(), "search") != -1 {
			searchs = sc.Text()
		}
	}

	return searchs, nil
}
