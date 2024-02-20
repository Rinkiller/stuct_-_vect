package main

import (
	"fmt"
	"math"
)

type Vector3 interface {
	Add(v Vector3) Vector3
	Subtract(u Vector3) Vector3
	Multiply(scalar float64) Vector3
	Dot(u Vector3) Vector3
	Length() float64
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

// Сложение векторов и возврат нового вектора {a.X + v.X, a.Y + v.Y, a.Z + v.Z}
func (a arrayVector3) Add(v Vector3) Vector3 {
	return arrayVector3{0: a.X() + v.X(), 1: a.Y() + v.Y(), 2: a.Z() + v.Z()}
}

// Вычитаение векторов и возврат нового вектора {a.X - u.X, a.Y - u.Y, a.Z - u.Z}
func (a arrayVector3) Subtract(v Vector3) Vector3 {
	return arrayVector3{0: a.X() - v.X(), 1: a.Y() - v.Y(), 2: a.Z() - v.Z()}
}

// Умножение вектора на число. Умножьте кажду каоордианту на число и верните веткор
func (a arrayVector3) Multiply(scalar float64) Vector3 {
	return arrayVector3{0: a.X() * scalar, 1: a.Y() * scalar, 2: a.Z() * scalar}
}

// Скалярное произведение векторов. Перемножьте координаты вектора v на координаты вектора u
// Пример a.x * u.x
func (a arrayVector3) Dot(v Vector3) Vector3 {
	return arrayVector3{0: a.X() * v.X(), 1: a.Y() * v.Y(), 2: a.Z() * v.Z()}
}

// Вычисление длины вектора math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
func (a arrayVector3) Length() float64 {
	return math.Sqrt(a.X()*a.X() + a.Y()*a.Y() + a.Z()*a.Z())
}

// Аналогично сделайте для вектора на основе структуры. Формулы аналогичные
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

func (s structVector3) Add(v Vector3) Vector3 {
	return structVector3{x: s.X() + v.X(), y: s.Y() + v.Y(), z: s.Z() + v.Z()}
}

func (s structVector3) Subtract(v Vector3) Vector3 {
	return structVector3{x: s.X() - v.X(), y: s.Y() - v.Y(), z: s.Z() - v.Z()}
}

func (s structVector3) Multiply(scalar float64) Vector3 {
	return structVector3{x: s.X() * scalar, y: s.Y() * scalar, z: s.Z() * scalar}
}

func (s structVector3) Dot(v Vector3) Vector3 {
	return structVector3{x: s.X() * v.X(), y: s.Y() * v.Y(), z: s.Z() * v.Z()}
}

func (s structVector3) Length() float64 {
	return math.Sqrt(s.X()*s.X() + s.Y()*s.Y() + s.Z()*s.Z())
}

func printVec(v, v1 Vector3) {
	fmt.Println(v.Add(v1))
	fmt.Println(v.Subtract(v1))
	fmt.Println(v.Multiply(5))
	fmt.Println(v.Dot(v1))
	fmt.Println(v.Length())
	fmt.Println(v1.Length())
}

func typeOf(x interface{}) int {
	switch x.(type) {
	case arrayVector3:
		return 1
	case structVector3:
		return 2
	default:
		return 0
	}
}

func main() {
	vec := arrayVector3{0: 1, 1: 1, 2: 1}
	vec1 := structVector3{
		x: 0,
		y: 0,
		z: 0,
	}

	fmt.Println(typeOf(vec))
	fmt.Println(typeOf(vec1))
	fmt.Println(typeOf(35))
	fmt.Println(vec)
	fmt.Println(vec1)
	fmt.Println(vec1.X())
	printVec(vec, vec1)
}
