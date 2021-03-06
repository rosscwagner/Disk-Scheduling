/*
Author: Ross Wagner
Date: 11/17/2017

I Ross Wagner (ro520462) affirm that
this program is entirely my own work and that I have neither developed my code with any
another person, nor copied any code from any other person, nor permitted my code to be copied
or otherwise used by any other person, nor have I copied, modified, or otherwise used programs
created by others. I acknowledge that any violation of the above terms will be treated as
academic dishonesty.

This program's goal is to implement 6 disk searching algorithms
a) fcfs
b) sstf
c) SCAN
d) C-SCAN
e) LOOK
f) C-LOOK
and compare their performance

*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"container/list"
)

/*
Data structure to hold the input parameters for disk searching alg
*/
type parameters struct {
	alg      string
	lowerCYL int
	upperCYL int
	initCYL  int
	requests []int
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*
Given a series of initial position and stops in the form of an int array,
calculates the total traversal distance and accompanying output indicating each
stop
*/
func CalcTrav(req []int, start int) (int, string) {
	var at int = start
	var trav int = 0
	var out string

	for _, cyl := range req {
		out += fmt.Sprintf("Servicing %5d\n", cyl)

		trav += Abs(cyl - at)
		at = cyl
	}
	return trav, out
	//return -1, "-1"
}

func Clook (p parameters)string{
	//var current *Element 
	//var next int = -1
	var traversal int = 0  
	//var index int = 0
	var out string = ""
	var reqList = list.New()
	reqList.Init()

	// sort requests
	cpy := p.requests
	sort.Ints(cpy)



	// find next elem larger than init. start at copy[0]
	for _, req := range cpy{
		reqList.PushBack(req)
	} 
	
	current := reqList.Back()

	for e := reqList.Front(); e != nil; e = e.Next() {
		if e.Value.(int) >= p.initCYL{
			current = e
			//traversal += 
			break
		}
	}

	//calculate initial traversal from start to next req
	out += fmt.Sprintf("Servicing %5d\n", current.Value)

	if current.Value.(int) < p.initCYL{
		traversal += Abs(current.Value.(int)-p.initCYL) + p.upperCYL-p.initCYL
	}else{
		traversal += Abs(current.Value.(int)-p.initCYL)
	}


	

	// remove elements larger than init. once all larger items remove add max - prev + max - next to traversal
	// find next and subtract distance
	
	for reqList.Len() > 1{
		old := current
		
			if current == reqList.Back() {
				// reached top of disk
				
				
				current = reqList.Front()
				
			}else{ // normal uppward operation
				current= current.Next()
				
			}
		
		traversal += Abs(current.Value.(int)-old.Value.(int))
		out += fmt.Sprintf("Servicing %5d\n", current.Value)
		reqList.Remove(old)
		
	}
	
	out += fmt.Sprintf("C-LOOK traversal count = %5d\n", traversal)
	return out
}

/*
similar to scan but jumps to begining if top of disk reached
*/
func CSCAN(p parameters)string{
	//var current *Element 
	//var next int = -1
	var traversal int = 0  
	//var up bool = true // direction that scan is going. true for up false for down
	//var index int = 0
	var out string = ""
	var reqList = list.New()
	reqList.Init()

	// sort requests
	cpy := p.requests
	sort.Ints(cpy)



	// find next elem larger than init. start at copy[0]
	for _, req := range cpy{
		reqList.PushBack(req)
	} 
	
	current := reqList.Back()

	for e := reqList.Front(); e != nil; e = e.Next() {
		if e.Value.(int) >= p.initCYL{
			current = e
			
			break
		}
	}

	//calculate inital traversal from start to next req
	out += fmt.Sprintf("Servicing %5d\n", current.Value)

	if current.Value.(int) < p.initCYL{
		traversal += Abs(current.Value.(int)-p.initCYL) + p.upperCYL-p.initCYL
	}else{
		traversal += Abs(current.Value.(int)-p.initCYL)
	}


	

	// find next and subtract distance
	
	for reqList.Len() > 1{
		old := current
		
			if current == reqList.Back() {
				// reached top of disk go to bottem
				traversal += 2* Abs(p.upperCYL - current.Value.(int))
				current = reqList.Front()
				traversal += 2* Abs(current.Value.(int)-p.lowerCYL)
				
			}else{ // normal uppward operation
				current= current.Next()
				
			}
		
		traversal += Abs(current.Value.(int)-old.Value.(int))
		out += fmt.Sprintf("Servicing %5d\n", current.Value)
		reqList.Remove(old)
		
	}
	
	out += fmt.Sprintf("C-SCAN traversal count = %5d\n", traversal)
	return out
}


/*
Equal tells whether a and b contain the same elements.
A nil argument is equivalent to an empty slice.
*/
func Equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

//
func FCFS(p parameters) string {
	var out string = ""
	var trav int = 0
	//var at int = p.initCYL
	var o string
	trav, o = CalcTrav(p.requests, p.initCYL)

	out += o
	out += fmt.Sprintf("FCFS traversal count = %5d\n", trav)

	return out
}

/*
This function takes in a file and returns its contents as a string
Parameters File
Returns string
*/
func FileToStr(file *os.File) string {
	scanner := bufio.NewScanner(file)
	txt := ""
	for scanner.Scan() {
		txt += scanner.Text() + "\n"

	}
	return txt
}

/*
given a sorted slice of integers retruns the value and index of element closest
to target
*/
func FindClosest(nums []int, target int) (int, int) {

	// binary search untill 2 items remain. calc diff from target. take item with
	// smallest diff

	var temp []int = nums
	// to keep track of index through binary search. should be the corosponding
	// original index of the 0th item in temp
	var i int = 0
	var val int = -1
 
	for len(temp) > 2 {
		var mid int = (len(temp) - 1) / 2
		//fmt.Printf("%d : %d\n", mid, temp[mid])
		if temp[mid] == target {
			return target, (i + mid)
		} else if temp[mid] > target {
			temp = temp[:mid+1] // +1 to include the last one
		} else {
			i = i + mid
			temp = temp[mid:]
		}
	}

	//fmt.Printf("%d\n", len(temp))
	if len(temp) <= 0 {
		val = -1
		i = -1
	} else if len(temp) == 1 {
		val = temp[0]
	} else { // len of temp must be 2 at this point
		var d0 int = Abs(temp[0] - target)
		var d1 int = Abs(temp[1] - target)

		if d0 < d1 {
			val = temp[0]
			// i is unchanged
			//i = i - 1
		} else { // take the larger value on tie
			val = temp[1]
			i = i + 1
		}
	}

	return val, i
}

func Look(p parameters)string{

	//var current *Element 
	//var next int = -1
	var traversal int = 0  
	var up bool = true // direction that scan is going. true for up false for down
	//var index int = 0
	var out string = ""
	var reqList = list.New()
	reqList.Init()

	// sort requests
	cpy := p.requests
	sort.Ints(cpy)



	// find next elem larger than init. start at copy[0]
	for _, req := range cpy{
		reqList.PushBack(req)
	} 
	
	current := reqList.Back()

	for e := reqList.Front(); e != nil; e = e.Next() {
		if e.Value.(int) >= p.initCYL{
			current = e
			//traversal += 
			break
		}
	}

	//calculate initial traversal from start to next req
	out += fmt.Sprintf("Servicing %5d\n", current.Value)

	if current.Value.(int) < p.initCYL{
		traversal += Abs(current.Value.(int)-p.initCYL) + p.upperCYL-p.initCYL
	}else{
		traversal += Abs(current.Value.(int)-p.initCYL)
	}


	

	// remove elements larger than init. once all larger items remove add max - prev + max - next to traversal
	// find next and subtract distance
	
	for reqList.Len() > 1{
		old := current
		if up{
			if current == reqList.Back() {
				// reached top of disk
				up = !up
				
				current = current.Prev()
				
			}else{ // normal uppward operation
				current= current.Next()
				
			}
		}else{
			if current == reqList.Front(){ // reached bottem 
				up = !up
				
				current = current.Next()
			}else{ // normal downward operation
				current = current.Prev()

			}
		}
		traversal += Abs(current.Value.(int)-old.Value.(int))
		out += fmt.Sprintf("Servicing %5d\n", current.Value)
		reqList.Remove(old)
		
	}
	
	out += fmt.Sprintf("LOOK traversal count = %5d\n", traversal)
	return out
}

/*
returns the contents of paramiters struct as a string
*/
func ParamsToString(params parameters) string {
	var out string

	out += fmt.Sprintf("Seek algorithm: %s\n", strings.ToUpper(params.alg))
	out += fmt.Sprintf("\tLower cylinder: %5d\n", params.lowerCYL)
	out += fmt.Sprintf("\tUpper cylinder: %5d\n", params.upperCYL)
	out += fmt.Sprintf("\tInit cylinder:  %5d\n", params.initCYL)
	out += fmt.Sprintf("\tCylinder requests:\n")
	for _, req := range params.requests {
		out += fmt.Sprintf("\t\tCylinder %5d\n", req)
	}
	return out
}

/*
reads input line by line and puts the data into a parameters struct
returns error code -1 if Abort error detected
*/
func Parse(lines []string) (parameters, int) {

	var alg string = ""
	var lower int = -1
	var upper int = -1
	var init int = -1
	var req []int = nil
	var p parameters
	var state int = 0
	var prevTok string = ""
	var prevTokLine int = 0

	// finite state machine
	for l, line := range lines {
		if state == -1 {
			break
		}
		state = 0
		// tokenize based on spaces
		tokens := strings.FieldsFunc(line, Split)
		for _, token := range tokens {
			if state == -1 {
				break
			}
			switch state {
			// start
			case 0:
				if token == "use" {
					state = 1
				} else if token == "lowerCYL" {
					state = 2
				} else if token == "upperCYL" {
					state = 3
				} else if token == "initCYL" {
					state = 4
				} else if token == "cylreq" {
					state = 5
				} else if token == "end" {
					state = 6
				} else {
					// unrecognized token
					state = 7
				}			
			case 1:
				if ValidAlg(token) {
					alg = token
				}
			case 2:
				i, err := strconv.Atoi(token)

				if err != nil {
					log.Fatal(err)
				}

				lower = i
			case 3:
				i, err := strconv.Atoi(token)

				if err != nil {
					log.Fatal(err)
				}

				upper = i
			case 4:
				i, err := strconv.Atoi(token)

				if err != nil {
					log.Fatal(err)
				}

				init = i
			case 5: // cylinder request
				i, err := strconv.Atoi(token)

				if err != nil {
					log.Fatal(err)
				}

				if i <= upper && i >= lower{
					// add to requests array
					req = append(req, i)
				}else{
					fmt.Printf("ERROR(15):Request out of bounds: req (%d) > upper (%d) or  < lower (%d)\n",i,upper,lower)
				}


				

			case 7:
				state = -1
				// get previous token and line

				fmt.Printf("ABORT(1): unrecognized token: %s on line %d\n", prevTok, prevTokLine)

				// Abort
			}
			prevTok = token
			prevTokLine = l
		}
	}

	//Check inputs are valid
	state = -1
	if init > upper{
		fmt.Printf("ABORT(11):initial (%d) > upper (%d)\n", init, upper)
		
	}else if init < lower{
		fmt.Printf("ABORT(12):initial (%d) < lower (%d)\n", init, lower)

	}else if upper < lower{
		fmt.Printf("ABORT(13):upper (%d) < lower (%d)\n", upper, lower)
	}else { // all clear
		state = 0
	}

	if state == -1 {
		// error encountered
		return p, -1
	}

	

	//assign values to p
	p.alg = alg
	p.lowerCYL = lower
	p.upperCYL = upper
	p.initCYL = init
	p.requests = req
	// return p with normal exit code
	return p, 0
}

func RemoveComments(lines []string) []string {
	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
				lines[i] = line[0:j]
				break
			}
		}
	}
	return lines
}

func RemoveElementInt(sl []int,i int)[]int{
	// Remove the element at index i from cpy.
	copy(sl[i:], sl[i+1:]) // Shift a[i+1:] left one index.
	sl[len(sl)-1] = -1     // Erase last element (write zero value).
	sl=sl[:len(sl)-1]   // Truncate slice.
	return sl
}

func Run(p parameters) string {

	var output string = ""
	// determine which alg to use
	// maybe use code of alg 
	if p.alg == "fcfs" {
		output = FCFS(p)
	} else if p.alg == "sstf" {
		output = SSTF(p)
	} else if p.alg =="scan"{
		output = SCAN(p)
	}else if p.alg =="c-scan"{
		output = CSCAN(p)
	}else if p.alg =="look"{
		output = Look(p)
	} else if p.alg == "c-look"{
		output = Clook(p)
	}

	return output
}

/*
like an elevator. starts at init and goes up all the way to the top then back down
*/
func SCAN(p parameters) string{
	

	//var current *Element 
	//var next int = -1
	var traversal int = 0  
	var up bool = true // direction that scan is going. true for up false for down
	//var index int = 0
	var out string = ""
	var reqList = list.New()
	reqList.Init()

	// sort requests
	cpy := p.requests
	sort.Ints(cpy)



	// find next elem larger than init. start at copy[0]
	for _, req := range cpy{
		reqList.PushBack(req)
	} 
	
	current := reqList.Back()

	for e := reqList.Front(); e != nil; e = e.Next() {
		if e.Value.(int) >= p.initCYL{
			current = e
			//traversal += 
			break
		}
	}

	//calculate initial traversal from start to next req
	out += fmt.Sprintf("Servicing %5d\n", current.Value)

	if current.Value.(int) < p.initCYL{
		traversal += Abs(current.Value.(int)-p.initCYL) + p.upperCYL-p.initCYL
	}else{
		traversal += Abs(current.Value.(int)-p.initCYL)
	}


	

	// remove elements larger than init. once all larger items remove add max - prev + max - next to traversal
	// find next and subtract distance
	
	for reqList.Len() > 1{
		old := current
		if up{
			if current == reqList.Back() {
				// reached top of disk
				up = !up
				traversal += 2*Abs(current.Value.(int) - p.upperCYL)
				current = current.Prev()
				
			}else{ // normal uppward operation
				current= current.Next()
				
			}
		}else{
			if current == reqList.Front(){ // reached bottem 
				up = !up
				traversal += 2*Abs(current.Value.(int) - p.lowerCYL)
				current = current.Next()
			}else{ // normal downward operation
				current = current.Prev()

			}
		}
		traversal += Abs(current.Value.(int)-old.Value.(int))
		out += fmt.Sprintf("Servicing %5d\n", current.Value)
		reqList.Remove(old)
		
	}
	
	out += fmt.Sprintf("SCAN traversal count = %5d\n", traversal)
	return out
}
/**/
func Split(r rune) bool {
	return r == ' ' || r == '\n' || r == '\t' || r == '\r' //|| r == '\n\r'
}

func SSTF(p parameters) string { 
	var output string = ""
	var prev int = p.initCYL
	var trav int
	var i int
	var next int = -1
	var o string = ""
	// create new thing to hold sstf ordered requests
	var sstfOrder []int
	// sort p.req
	cpy := p.requests
	sort.Ints(cpy)
	// use binary searches to find next closest cyl req and its index in cpy
	for len(cpy) > 0 {
		next, i = FindClosest(cpy, prev)
		sstfOrder = append(sstfOrder, next)
		prev = next
		// Remove the element at index i from cpy.
		copy(cpy[i:], cpy[i+1:]) // Shift a[i+1:] left one index.
		cpy[len(cpy)-1] = -1     // Erase last element (write zero value).
		cpy = cpy[:len(cpy)-1]   // Truncate slice.

	}
	// find traversal dist of sstf ordered requests
	trav, o = CalcTrav(sstfOrder, p.initCYL)
	output += o

	output += fmt.Sprintf("SSTF traversal count = %5d\n", trav)
	return output
}

func ValidAlg(alg string) bool {
	var algs []string = []string{
		"fcfs", "sstf", "scan", "c-scan", "look", "c-look",
	}
	for _, str := range algs {
		if alg == str {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: ./pa2 <input file>\n")
		return
	}

	inputFile := os.Args[1]

	file, err := os.Open(inputFile)

	if err != nil {
		file.Close()
		log.Fatal(err)
	}

	txt := FileToStr(file)
	file.Close()

	// print contents
	//fmt.Printf("%s", txt)
	//fmt.Printf("\n\n")
	// split by line
	lines := strings.Split(txt, "\n")

	// remove comments
	lines = RemoveComments(lines)

	/*
		// print without comments
		for _, line := range lines {
			fmt.Printf("%s\n", line)
		}
	*/

	// extract input parameters from the input
	inParam, code := Parse(lines)
	if code == -1 {
		// abort condition triggered
		return
	}

	// check for invalid params
	fmt.Printf("%s", ParamsToString(inParam))
	fmt.Printf("%s", Run(inParam))

}
