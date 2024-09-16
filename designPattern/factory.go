package main

/**
 * @Author elastic·H
 * @Date 2024-09-16
 * @File: factory.go
 * @Description: 工厂模式
 */

import (
	"fmt"
)

type Shape interface {
	Draw()
}

type Circle struct{}

func (c *Circle) Draw() {
	fmt.Println("Drawing Circle")
}

type Square struct{}

func (s *Square) Draw() {
	fmt.Println("Drawing Square")
}

type ShapeFactory struct{}

func (sf *ShapeFactory) GetShape(shapeType string) Shape {
	if shapeType == "CIRCLE" {
		return &Circle{}
	} else if shapeType == "SQUARE" {
		return &Square{}
	}
	return nil
}

func main() {
	factory := ShapeFactory{}
	shape1 := factory.GetShape("CIRCLE")
	shape1.Draw()
	shape2 := factory.GetShape("SQUARE")
	shape2.Draw()
}
