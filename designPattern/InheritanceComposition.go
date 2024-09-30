package main

/**
 * @Author elastic·H
 * @Date 2024-09-30
 * @File: InheritanceComposition.go
 * @Description:
 */

import (
	"fmt"
)

// Go 在设计上，选择了组合优于继承的编程设计模式

// 继承 vs 组合
// 继承：
// 继承父类的属性和方法：子类会自动获得父类的所有属性和方法，因此父类的更改会直接影响子类。这使得子类的行为依赖于父类，这被称为强耦合。
// 适合描述 "is-a" 关系。例如，Dog 继承自 Animal，因为 Dog 本质上也是一种 Animal。
//
// 组合：
// 包含另一个类作为属性：组合则是通过将一个类的实例作为另一个类的属性来实现功能。例如，Dog 可以包含 SoundBehavior 作为属性，而不直接继承 Animal。这样可以根据需要灵活地组合不同的类。
// 适合描述 "has-a" 关系。例如，Dog 可以有一个 SoundBehavior，使其具有发声功能，而不需要继承所有 Animal 的行为。

// is-a 和 has-a 代表面向对象编程中的两种关系：
// is-a 关系：
// 描述继承关系，表示一个类是另一类的子类。
// 例如，Dog 是 Animal 的子类，这意味着 Dog is a type of Animal。
// 继承用于实现 is-a 关系，强调一种类型是另一种类型的特化。
//
// has-a 关系：
// 描述组合关系，表示一个类拥有另一个类的实例。
// 例如，Car 有一个 Engine，表示 Car has an Engine。
// 组合用于实现 has-a 关系，强调对象由多个独立部分组成。
// is-a 强调类之间的层次结构，而 has-a 强调对象之间的组成关系。这两种关系有助于清晰地表达类与对象的关系和职责分离，指导系统设计时的代码复用与扩展。

// ParentAnimal 父类
type ParentAnimal struct {
	Name string
}

func (a ParentAnimal) Speak() {
	fmt.Println(a.Name, "makes a sound")
}

// ChildDog 子类
type ChildDog struct {
	ParentAnimal
}

func inheritanceDemo() {
	_dog := ChildDog{ParentAnimal{Name: "Buddy"}}
	_dog.Speak() // 输出: Buddy makes a sound
}

// /////////////////////////////////////////////////////////////////////////////////////

// SoundBehavior 可复用的结构体
type SoundBehavior struct{}

func (s SoundBehavior) Speak(name string) {
	fmt.Println(name, "makes a sound")
}

// ChildDog2 组合而非继承
type ChildDog2 struct {
	Name          string
	SoundBehavior // 嵌入了 SoundBehavior 以复用其功能
}

func compositionDemo() {
	_dog := ChildDog2{Name: "Buddy"}
	_dog.Speak(_dog.Name) // 输出: Buddy makes a sound
}

func main() {
	inheritanceDemo()
	compositionDemo()
}
