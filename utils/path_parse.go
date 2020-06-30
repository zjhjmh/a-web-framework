package utils

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Get_path_params(path string) (id int, found bool) {
	//获取path中的数字部分
	path_list := strings.Split(path, "/")
	log.Print(path_list)
	for _, path := range path_list {
		log.Print(path)
		matched, err := regexp.MatchString(`^[1-9]\d*|0$`, path)
		if err != nil {
			log.Fatal(err)
		}
		if matched {
			id, err = strconv.Atoi(path)
			if err != nil {
				log.Fatal(err)
			}
			return id, true
		}
	}
	return -1, false
}
