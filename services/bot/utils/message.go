package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Random(list []string, blackList ...string) string {
	rand.Seed(time.Now().Unix())

	if len(blackList) == 0 {
		return list[rand.Intn(len(list))]
	}

	item := ""
	for {
		item = list[rand.Intn(len(list))]
		if !Find(blackList, item) {
			break
		}
	}
	return item
}

func Replace(str string, repMap map[string]string, randomRep map[string][]string) string {
	for key, val := range repMap {
		str = strings.ReplaceAll(str, key, val)
	}

	return str
}

var rangeRegex = regexp.MustCompile(`\{random.[0-9]{1,}\-[0-9]{1,}\}`)

func RandomRange(str string) string {
	foundRandom := rangeRegex.FindAllString(str, -1)

	for _, k := range foundRandom {
		nums := strings.Split(k[8:(len(k)-1)], "-")
		min, err := strconv.Atoi(nums[0])
		if err != nil {
			continue
		}
		max, err := strconv.Atoi(nums[1])
		if err != nil {
			continue
		}

		str = strings.Replace(str, k, fmt.Sprintf("%d", RandomInt(min, max)), 1)
	}
	return str
}

func RandomInt(min, max int) int {
	rand.Seed(time.Now().Unix() + rand.Int63())
	return rand.Intn(max-min) + min
}

func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
