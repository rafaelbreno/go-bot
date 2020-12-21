package err

import "log"

type Err struct{}

func (_ Err) Log(e error) {
	log.Fatalln(e)
}
