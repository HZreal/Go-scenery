package main

/**
 * @Author nico
 * @Date 2024-12-10
 * @File: template.go
 * @Description: 模板模式（Template Method Pattern）是一种行为型设计模式，它定义了一个操作的框架，并允许子类通过实现部分步骤来实现该操作的细节。
 * 也就是说，模板模式通过将固定的算法框架留给父类，而将算法中可变的部分交给子类来实现，从而避免重复代码的同时，保持了灵活性。
 */

import (
	"fmt"
)

// 关键要素
// 		抽象类：定义一个模板方法（template method），这个方法是算法的骨架，包含了固定的步骤。它通常调用一系列的步骤方法，其中某些步骤是由子类实现的。
// 		具体子类：子类实现算法中的一些具体步骤。这些步骤在父类的模板方法中被调用，从而完成整个操作。
// 		钩子方法（Hook Method）：模板方法中的某些步骤可以是空的钩子方法，子类可以选择性地覆盖它们以改变某些行为。

// 应用场景
// 		数据处理/转换：处理数据时，可能会有多个步骤，其中一些步骤可能是不同的（如读取文件、处理数据、输出结果等），模板模式允许每个子类实现自己特定的数据处理步骤。
// 		游戏开发中的不同关卡：不同的关卡可能有一些固定的步骤（如初始化场景、加载资源、开始游戏），而每个关卡的具体实现会有所不同。模板模式能很好地将共同的行为抽象出来，子类只需要关注具体关卡的实现。
// 		文件处理流程：读取文件、处理内容、保存文件等步骤的流程可能是固定的，而处理的内容（如解析不同类型的文件）可能不同。模板方法可以将固定的文件处理流程留给父类，而处理内容的部分交给子类实现。

// 模板模式的优缺点
// 优点：
// 		代码复用：模板方法模式将重复的代码提取到父类中，提高了代码复用性。
// 		灵活性：子类可以只实现一些特定的步骤，其他步骤由父类提供，保持了灵活性。
// 		封装变化：将变化的部分（即算法的某些步骤）封装到子类中，从而避免了修改父类代码。
// 缺点：
// 		继承限制：模板方法模式依赖于继承，如果算法的步骤变化较大，可能需要创建多个子类，导致继承层次过深。
// 		扩展困难：模板方法模式并不容易扩展，比如增加新的步骤时，可能需要修改模板方法，这就违背了“开闭原则”。
// 		代码膨胀：如果模板方法中的步骤太多，可能会导致父类的代码过于庞大和复杂，子类则需要实现大量的方法。

// 定义一个抽象类接口
type AbstractClass interface {
	Step1()
	Step2()
	Step3()
	TemplateMethod()
}

// 具体实现1
type ConcreteClass1 struct{}

func (c *ConcreteClass1) Step1() {
	fmt.Println("ConcreteClass1 Step1")
}

func (c *ConcreteClass1) Step2() {
	fmt.Println("ConcreteClass1 Step2")
}

func (c *ConcreteClass1) Step3() {
	fmt.Println("ConcreteClass1 Step3")
}

func (c *ConcreteClass1) TemplateMethod() {
	c.Step1()
	c.Step2()
	c.Step3()
}

// 具体实现2
type ConcreteClass2 struct{}

func (c *ConcreteClass2) Step1() {
	fmt.Println("ConcreteClass2 Step1")
}

func (c *ConcreteClass2) Step2() {
	fmt.Println("ConcreteClass2 Step2")
}

func (c *ConcreteClass2) Step3() {
	fmt.Println("ConcreteClass2 Step3")
}

func (c *ConcreteClass2) TemplateMethod() {
	c.Step1()
	c.Step2()
	c.Step3()
}

func main() {
	class1 := &ConcreteClass1{}
	class1.TemplateMethod()

	class2 := &ConcreteClass2{}
	class2.TemplateMethod()
}
