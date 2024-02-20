package main

import (
	"fmt"
	"os"
)

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

func writeVecF(x Vector3) {

	f, err := os.OpenFile("vertors.vec", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	switch x.(type) {
	case arrayVector3:
		_, err := f.WriteString("{A " + fmt.Sprintf("%f", x.X()) + " " + fmt.Sprintf("%f", x.Y()) + " " + fmt.Sprintf("%f", x.Z()) + "}\n")
		if err != nil {
			panic(err)
		}
		return
	case structVector3:
		_, err := f.WriteString("{S " + fmt.Sprintf("%f", x.X()) + " " + fmt.Sprintf("%f", x.Y()) + " " + fmt.Sprintf("%f", x.Z()) + "}\n")
		if err != nil {
			panic(err)
		}
		return
	default:
		return
	}

}

func main() {
	vec := arrayVector3{0: 1, 1: 2, 2: 3}
	vec1 := structVector3{
		x: 3,
		y: 2,
		z: -1,
	}
	vec2 := structVector3{
		x: 7,
		y: 21,
		z: -11,
	}
	writeVecF(vec)
	writeVecF(vec1)
	writeVecF(vec2)
}
