package main

import (
	"fmt"
	"regexp"
)

type cmdFunc func()
type registerMap map[string]cmdFunc

var register registerMap

func main() {
	exePrint("FindStringSubmatch")
}
func exePrint(name string) {
	register[name]()
}
func init() {
	register = make(registerMap)
	register["MatchString"] = MatchString
	register["R_MatchString"] = R_MatchString
	register["FindString"] = FindString
	register["FindStringIndex"] = FindStringIndex
	register["FindStringSubmatch"] = FindStringSubmatch
	register["FindAllString"] = FindAllString
	register["FindAllStringSubmatch"] = FindAllStringSubmatch
	register["FindAllStringSubmatchIndex"] = FindAllStringSubmatchIndex
	register["ReplaceAllLiteralString"] = ReplaceAllLiteralString
	register["ReplaceAllString"] = ReplaceAllString
	register["SubexpNames"] = SubexpNames
}

func MatchString() {
	fmt.Println("学习MatchString")
	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
	fmt.Println(validID.MatchString("aaaa[1111]"))
	fmt.Println(validID.MatchString("e[7]"))
	fmt.Println(validID.MatchString("Job[48]"))
	fmt.Println(validID.MatchString("snakey"))
}
func R_MatchString() {
	fmt.Println("学习R_MatchString")
	matched, err := regexp.MatchString("foo.*", "seafoo")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString("bar.*", "seafood")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString("a(b", "seafood")
	fmt.Println(matched, err)
}
func FindString() {
	re := regexp.MustCompile("fo.?")
	fmt.Printf("%q\n", re.FindString("seafood"))
	fmt.Printf("%q\n", re.FindString("meat"))
}

func FindStringIndex() {
	re := regexp.MustCompile("ab?")
	fmt.Println(re.FindStringIndex("tablett"))
	fmt.Println(re.FindStringIndex("foo") == nil)
}

func FindStringSubmatch() {
	fmt.Println("学习FindStringSubmatch")
	ss := "NAME=ora.DATA.dg\nTYPE=ora.diskgroup.type\nTARGET=ONLINE          , ONLINE\nSTATE=ONLINE on rac048, ONLINE on rac049"
	re := regexp.MustCompile("a(x*)b(y|z)c")
	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-"))
	fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))
	fmt.Printf("%q\n", re.FindStringSubmatch("abc"))

	re2 := regexp.MustCompile("TARGET=(.*)")
	fmt.Printf("%q\n", re2.FindStringSubmatch(ss))
}
func FindAllString() {
	fmt.Println("学习FindAllString")
	re := regexp.MustCompile("a.")
	fmt.Println(re.FindAllString("paranormal", -1))
	fmt.Println(re.FindAllString("paranormal", 2))
	fmt.Println(re.FindAllString("graal", -1))
	fmt.Println(re.FindAllString("none", -1))
}

func FindAllStringSubmatch() {
	fmt.Println("学习FindAllStringSubmatch")
	// re := regexp.MustCompile("a(x*)b")
	// fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-", -1))
	// fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-", -1))
	// fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-axb-", -1))
	// fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-ab-", -1))
	re := regexp.MustCompile(`\w-\w`)
	fmt.Printf("%q\n", re.FindAllStringSubmatch("aaaaa-123a", -1))
}

func FindAllStringSubmatchIndex() {
	re := regexp.MustCompile("a(x*)b")
	// Indices:
	//    01234567   012345678
	//    -ab-axb-   -axxb-ab-
	fmt.Println(re.FindAllStringSubmatchIndex("-ab-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-axxb-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-ab-axb-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-axxb-ab-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-foo-", -1))
}

func ReplaceAllLiteralString() {
	re := regexp.MustCompile("a(x*)b")
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "T"))
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "$1"))
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "${1}"))
}

func ReplaceAllString() {
	re := regexp.MustCompile("a(x*)b")
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "T"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1W"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "${1}W"))
}

func SubexpNames() {
	re := regexp.MustCompile("(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)")
	fmt.Println(re.MatchString("A Turing"))
	fmt.Printf("%q\n", re.SubexpNames())
	reversed := fmt.Sprintf("${%s} ${%s}", re.SubexpNames()[2], re.SubexpNames()[1])
	fmt.Println(reversed)
	fmt.Println(re.ReplaceAllString("Alan Turing", reversed))
}
