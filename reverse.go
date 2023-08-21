package chinesereverse

import (
	"bufio"
	"errors"
	"log"
	"os"
	"sync"
)

var dt *dict
var once sync.Once

func init() {
	dt = &dict{
		datas2t: make(map[rune]rune),
		datat2s: make(map[rune]rune),
	}

	if err := buildDict(""); err != nil {
		log.Println(err)
	}
}

// 在现有对照表的基础上追加自定的对照表，如果有相同的字，那么会在相应位置上覆盖掉原来的。
// 格式为一行简体接一行繁体，可以有多行。
//
// 考虑到它只是个简单的工具，不应该对现有系统有任何锁的影响，所以此处没有考虑加锁，因此，此函数应该放在项目初始化的地方。
func WithExtraDictFile(filepath string) {
	if filepath != "" {
		once.Do(func() {
			if err := buildDict(filepath); err != nil {
				log.Println(err)
			}
		})
	}
}

func buildDict(dictfile string) error {
	var simplified []rune
	var traditional []rune
	if dictfile != "" {
		file, err := os.Open(dictfile)
		if err != nil {
			return err
		}
		defer file.Close()
		buf := bufio.NewScanner(file)
		var i int
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
	} else {
		simplified = DefaultSimplified
		traditional = DefaultTraditional
	}

	if len(simplified) != len(traditional) {
		return errors.New("simplified length is not equal to traditional")
	}

	for i := 0; i < len(simplified); i++ {
		dt.set(simplified[i], traditional[i])
	}

	return nil
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
