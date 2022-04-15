package main

import (
	"fmt"
	"os"
	"regexp"
)

const minlen = 8

const maxlen = 30

func correctword(word string) (res bool, errm string) {
	res = true
	if len(word) < minlen || len(word) > maxlen {
		errm = "length should be at range from 8 to 30"
		res = false
	}
	match, err := regexp.MatchString("[A-z]+", word)
	if err != nil {
		fmt.Println(err)
		res = false
	}
	if !match {
		errm = "word should contain only litters"
		res = false
	}
	return res, errm
}

func initGame(initWord string) func() bool {
	res, errs := correctword(initWord)

	if !res {
		printErrors(errs)
	}
	fstWord := initWord

	fstMap := mapWord(fstWord)

	return func() bool {
		word := input()
		ok, errs := correctword(word)
		if !ok {
			printErrors(errs)
			return false
		}

		curmap := mapWord(word)

		return included(fstMap, curmap)
	}
}

func included(master map[rune]*int, slave map[rune]*int) bool {
	if len(slave) > len(master) {
		return false
	}
	for sk, sv := range slave {
		v, ok := master[sk]
		return ok && *v >= *sv
	}
	return true
}

func mapWord(word string) (resmap map[rune]*int) {
	resmap = map[rune]*int{}
	for _, r := range word {
		count, ok := resmap[r]
		if ok {
			*count++
		} else {

			resmap[r] = new(int)
			*resmap[r] = 1
		}
	}
	return resmap
}

type player struct {
	name  string
	score int
}

func input() (res string) {
	fmt.Println("write word")
	fmt.Fscan(os.Stdin, &res)
	return res
}
func printErrors(err string) {
	fmt.Println(err)
}

func printScores(players []*player) {
	for _, a := range players {
		fmt.Println(a.name, a.score)
	}
}

func main() {
	var round = initGame(input())
	var players = [2]*player{&player{name: "fst"}, &player{name: "snd"}}
	gfinish := false
	for {
		for _, p := range players {
			var ok = round()

			if !ok {
				gfinish = true
				break
			} else {
				p.score++
			}
		}
		if gfinish {
			break
		}
	}

	printScores(players[:])
}
