package mkltree

import (
	"log"
)

// FIXME: not to affect global log config
func init()  {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//log.SetOutput(ioutil.Discard) // discard for product env
}