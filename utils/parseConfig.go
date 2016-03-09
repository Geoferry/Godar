package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type config struct {
	settings map[string]string
}

var Config config

func init() {
	Config.settings = make(map[string]string)
	file, err := os.Open("./conf/config.conf")
	if err != nil {
		fmt.Println("Can't open or find config.conf file in \"conf\" folder.")
		return
	}
	defer file.Close()

	r := bufio.NewReader(file)
	for {
		line, err := r.ReadString('\n')

		if err == io.EOF {
			goto Insert
		}

		if err != nil {
			fmt.Println(err)
			break
		}
	Insert:
		if len(line) == 1 || strings.HasPrefix(line, "#") {
			continue
		}

		line = strings.TrimSuffix(line, "\n")
		m := strings.Split(line, "=")

		Config.settings[strings.TrimSpace(m[0])] = strings.TrimSpace(m[1])

		if err == io.EOF {
			break
		}
	}
}

func (this *config) GetSetting(s string) (value string, ok bool) {
	value, ok = this.settings[s]
	if !ok {
		fmt.Println("Configuration " + s + " Not Exist Please Set It In conf/config.conf !")
	}
	return
}
