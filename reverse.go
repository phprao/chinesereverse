package chinesereverse

import (
	"bufio"
	"log"
	"os"
	"path"
	"runtime"
	"sync"
)

var dt *dict

// 简体 <=> 繁体对照表，格式为一行简体接一行繁体，可以有多行
var dictname = "/dict.txt"

func init() {
	dt = &dict{
		datas2t: make(map[rune]rune),
		datat2s: make(map[rune]rune),
	}

	buildDict()
}

func buildDict() {
	_, filename, _, _ := runtime.Caller(1)
	// log.Println(filename)
	dictfile := path.Dir(filename) + dictname
	file, err := os.Open(dictfile)
	if err != nil {
		log.Println(err)
		return
	}
	buf := bufio.NewScanner(file)
	var i int
	var simplified []rune
	var traditional []rune
	for buf.Scan() {
		text := buf.Text()
		if text == "" {
			continue
		}
		switch i % 2 {
		case 0:
			simplified = []rune(text)
		case 1:
			traditional = []rune(text)
		}
		i++
	}
	if len(simplified) != len(traditional) {
		log.Println("simplified length is not equal to traditional")
		return
	}
	for i := 0; i < len(simplified); i++ {
		dt.set(simplified[i], traditional[i])
	}
}

// 简体 => 繁体
func SimplifiedToTraditional(val string) string {
	r := []rune(val)
	for i := 0; i < len(r); i++ {
		r[i] = dt.get(r[i], 1)
	}
	return string(r)
}

// 繁体 => 简体
func TraditionalToSimplified(val string) string {
	r := []rune(val)
	for i := 0; i < len(r); i++ {
		r[i] = dt.get(r[i], 2)
	}
	return string(r)
}

type dict struct {
	rwmu    sync.RWMutex
	datas2t map[rune]rune
	datat2s map[rune]rune
}

func (d *dict) set(v1 rune, v2 rune) {
	d.datas2t[v1] = v2
	d.datat2s[v2] = v1
}

func (d *dict) get(key rune, mode uint8) rune {
	if mode == 1 {
		if s, ok := d.datas2t[key]; ok {
			return s
		}
	} else {
		if s, ok := d.datat2s[key]; ok {
			return s
		}
	}
	return key
}
