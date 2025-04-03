package error

/**
 * @Author nico
 * @Date 2025-04-03
 * @File: error.go
 * @Description: 错误处理
 * @Description: 参考: https://github.com/xxjwxc/uber_go_guide_cn?tab=readme-ov-file#errors
 */

import (
	"errors"
	"fmt"
)

// 调用其他方法时出现错误, 通常有三种处理方式可以选择:
// 		将原始错误原样返回
// 		使用 fmt.Errorf 搭配 %w 将错误添加进上下文后返回
// 		使用 fmt.Errorf 搭配 %v 将错误添加进上下文后返回

// -------------------------------------------------------------------------------------------------
// 调用者是否需要匹配错误以便他们可以处理它？
// 如果是，我们必须通过声明顶级错误变量或自定义类型来支持 errors.Is 或 errors.As 函数
// 如果不需要，则一般化的直接使用静态字符串的方式定义返回错误

// 返回的错误，后续不再匹配
func Open1() error {
	return errors.New("could not open")
}

var ErrCouldNotOpen = errors.New("could not open")

// 返回的错误，后续可通过 errors.Is 等匹配
func Open2() error {
	return ErrCouldNotOpen
}

func t1() {
	if err := Open1(); err != nil {
		// 无法匹配、处理该错误
		panic("unknown error")
	}

	if err := Open2(); err != nil {
		if errors.Is(err, ErrCouldNotOpen) {
			// handle the error
		} else {
			panic("unknown error")
		}
	}
}

// -------------------------------------------------------------------------------------------------
// 错误消息是静态字符串，还是需要上下文信息的动态字符串？
// 如果是静态字符串，可以使用 errors.New；如果是动态态字符串，必须使用 fmt.Errorf 或自定义错误类型

// 返回的错误，后续不再匹配
func Open3(file string) error {
	return fmt.Errorf("file %q not found", file)
}

// 自定义错误（实现 error 接口），后续可匹配
type NotFoundError struct {
	File string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("file %q not found", e.File)
}

func Open4(file string) error {
	return &NotFoundError{File: file}
}

func t2() {
	if err := Open3("testfile.txt"); err != nil {
		// Can't handle the error.
		panic("unknown error")
	}

	if err := Open4("testfile.txt"); err != nil {
		var notFound *NotFoundError
		if errors.As(err, &notFound) {
			// handle the error
		} else {
			panic("unknown error")
		}
	}
}

// 错误包装 -------------------------------------------------------------------------------------------------
// 需要传递由下游函数返回的新错误，使用错误包装
// fmt.Errorf 为你的错误添加上下文， 根据调用者是否应该能够匹配和提取根本原因，在 %w 或 %v 动词之间进行选择。
// 		如果调用者应该可以访问底层错误，请使用 %w。 对于大多数包装错误，这是一个很好的默认值， 但请注意，调用者可能会开始依赖此行为。因此，对于包装错误是已知var或类型的情况，请将其作为函数契约的一部分进行记录和测试。
// 		使用 %v 来混淆底层错误。 调用者将无法匹配它，但如果需要，您可以在将来切换到 %w。
// 在为返回的错误添加上下文时，通过避免使用"failed to"之类的短语来保持上下文简洁，当错误通过堆栈向上渗透时，它会一层一层被堆积起来：

// c
func t3() error {
	err := Open1()
	if err != nil {
		// return fmt.Errorf("fail to open: %w", err) // 避免使用
		return fmt.Errorf("open: %w", err)
	}

	return nil
}

// 错误命名 -------------------------------------------------------------------------------------------------
// 对于存储为全局变量的错误值， 根据是否导出，使用前缀 Err 或 err
var (
	// 导出以下错误，以便此包的用户可以将它们与 errors.Is 进行匹配。
	ErrBrokenLink = errors.New("link is broken")

	// 这个错误没有被导出，因为我们不想让它成为我们公共 API 的一部分。 我们可能仍然在带有错误的包内使用它。
	errNotFound = errors.New("not found")
)

// 对于自定义错误类型，请改用后缀 Error

// 这个错误被导出，以便这个包的用户可以将它与 errors.As 匹配。
type FileNotFoundError struct {
	File string
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("file %q not found", e.File)
}

// 这个错误没有被导出，因为我们不想让它成为公共 API 的一部分。 我们仍然可以在带有 errors.As 的包中使用它。
type resolveError struct {
	Path string
}

func (e *resolveError) Error() string {
	return fmt.Sprintf("resolve %q", e.Path)
}

// 当调用方从被调用方接收到错误时，它可以根据对错误的了解，以各种不同的方式进行处理。
// 其中包括但不限于：
// 		如果被调用者约定定义了特定的错误，则将错误与 errors.Is 或 errors.As 匹配，并以不同的方式处理分支
// 		如果错误是可恢复的，则记录错误并正常降级
// 		如果该错误表示特定于域的故障条件，则返回定义明确的错误
// 		返回错误，无论是 wrapped 还是逐字逐句
// 无论调用方如何处理错误，它通常都应该只处理每个错误一次。例如，调用方不应该记录错误然后返回，因为其调用方也可能处理错误。

func handleErrorDemo() error {
	// Bad: 记录错误并将其返回
	// 堆栈中的调用程序可能会对该错误采取类似的操作。这样做会在应用程序日志中造成大量噪音，但收效甚微
	// u, err := getUser(id)
	// if err != nil {
	// 	// BAD: See description
	// 	log.Printf("Could not get user %q: %v", id, err)
	// 	return err
	// }

	// Good: 将错误换行并返回
	// 如果操作不是绝对必要的，我们可以通过从中恢复来提供降级但不间断的体验。
	// u, err := getUser(id)
	// if err != nil {
	// 	return fmt.Errorf("get user %q: %w", id, err)
	// }

	// Good: 记录错误并正常降级
	// 如果操作不是绝对必要的，我们可以通过从中恢复来提供降级但不间断的体验。
	// if err := emitMetrics(); err != nil {
	// 	// Failure to write metrics should not
	// 	// break the application.
	// 	log.Printf("Could not emit metrics: %v", err)
	// }

	// Good: 匹配错误并适当降级
	// 如果被调用者在其约定中定义了一个特定的错误，并且失败是可恢复的，则匹配该错误案例并正常降级。对于所有其他案例，请包装错误并返回。
	// 堆栈中更靠上的调用程序将处理其他错误。
	// tz, err := getUserTimeZone(id)
	// if err != nil {
	// 	if errors.Is(err, ErrUserNotFound) {
	// 		// User doesn't exist. Use UTC.
	// 		tz = time.UTC
	// 	} else {
	// 		return fmt.Errorf("get user %q: %w", id, err)
	// 	}
	// }

	return nil
}

func main() {
	t1()
	t2()
	_ = t3()
}
