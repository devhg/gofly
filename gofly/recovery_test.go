package gofly

import (
	"log"
	"testing"
)

func TestRecovery(t *testing.T) {

}

func TestRecoveryT(t *testing.T) {
	panic(12) // 如果panic在 defer的上方的话，下面defer recover的将不会被开启，必须将defer recover()放到panic的上方
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	panic(12)
}

func Test_trace(t *testing.T) {

}
