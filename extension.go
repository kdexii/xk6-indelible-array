package myextension

import (
	"context"
	"fmt"

	"go.k6.io/k6/js/modules"
)

// Определение структуры, которая будет экспортироваться в k6
type MyExtension struct{}

// Метод, который можно вызывать из k6
func (m *MyExtension) Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// Регистрация расширения в k6
func init() {
	modules.Register("k6/x/xk6-static-array", new(MyExtension))
}
