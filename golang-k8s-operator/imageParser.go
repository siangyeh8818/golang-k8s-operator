package main

import (
	"strings"
)

func ImagenameSplit(rawimage string) (string, string, string) {
	var parsername []string
	var tag []string

	var return_image string
	var return_image_tag string
	var hub string

	parsername = strings.Split(rawimage, "/")
	if len(parsername) == 1 {
		for i := range parsername {
			//fmt.Println(parsername[i])
			if i == 0 {
				tag = strings.Split(parsername[i], ":")
				/*			for j := range tag {
								fmt.Println(tag[j])
							}
				*/
			}
		}
		hub = ""
		return_image = tag[0]
		return_image_tag = tag[1]
	} else if len(parsername) == 2 {
		for i := range parsername {
			//fmt.Println(parsername[i])
			if i == 1 {
				tag = strings.Split(parsername[i], ":")
				/*			for j := range tag {
								fmt.Println(tag[j])
							}
				*/
			}
		}
		return_image = tag[0]
		return_image_tag = tag[1]
		hub = parsername[1]
	} else if len(parsername) == 3 {
		for i := range parsername {
			//fmt.Println(parsername[i])
			if i == 2 {
				tag = strings.Split(parsername[i], ":")
				/*			for j := range tag {
								fmt.Println(tag[j])
							}
				*/
			}
		}
		return_image = tag[0]
		return_image_tag = tag[1]
		hub = parsername[1]
	}

	return hub, return_image, return_image_tag
}
