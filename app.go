package main

import (
	"context"
	"fmt"
)

type App struct {
	ctx context.Context
}

// NewApp NewApp创建一个新的应用程序应用程序结构
func NewApp() *App {
	return &App{}
}

// 应用程序启动时调用startup。保存上下文
// 所以我们可以调用运行时方法
func (a *App) startup(ctx context.Context) {
	fmt.Println("程序启动")
	a.ctx = ctx
}
func (a *App) shutdown(ctx context.Context) {
	fmt.Println("程序关闭")
}
func (a *App) domReady(ctx context.Context) {
	fmt.Println("dom加载完成")
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello123123123 %s, It's show time!", name)
}

func (a *App) Test() string {
	return fmt.Sprintf("test, 测试测试测试")
}
