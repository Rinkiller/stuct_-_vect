package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//"strings"
)

var baseV = make([]Vector3, 0, 0)

type Vector3 interface {
	X() float64
	Y() float64
	Z() float64
}

var _ Vector3 = arrayVector3{}

type arrayVector3 [3]float64

func (a arrayVector3) X() float64 {
	return a[0]
}

func (a arrayVector3) Y() float64 {
	return a[1]
}

func (a arrayVector3) Z() float64 {
	return a[2]
}

var _ Vector3 = structVector3{}

type structVector3 struct {
	x float64
	y float64
	z float64
}

func (s structVector3) X() float64 {
	return s.x
}

func (s structVector3) Y() float64 {
	return s.y
}

func (s structVector3) Z() float64 {
	return s.z
}

func main() {

	f, err := os.Open("vertors.vec")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()[1 : len(scanner.Text())-1]
		arr := strings.Split(s, " ")
		if len(arr) != 4 {
			panic("Error load file db vectors")
		}
		x, err := strconv.ParseFloat(arr[1], 64)
		if err != nil {
			panic(err)
		}
		y, err := strconv.ParseFloat(arr[2], 64)
		if err != nil {
			panic(err)
		}
		z, err := strconv.ParseFloat(arr[3], 64)
		if err != nil {
			panic(err)
		}
		if arr[0] == "A" {
			baseV = append(baseV, arrayVector3{0: x, 1: y, 2: z})
		} else if arr[0] == "S" {
			baseV = append(baseV, structVector3{x: x, y: y, z: z})
		} else {
			panic("Error load file db vectors")
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(baseV)
}
