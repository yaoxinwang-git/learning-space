/**
 * @Author xinwang_yao
 * @Date 2022/12/1 14:28
 **/

package main

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

// 计算斐波那契队列
func Fibonacci(n uint) (uint64, error) {
	if n <= 1 {
		return uint64(n), nil
	}

	var n2, n1 uint64 = 0, 1
	for i := uint(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}

	return n2 + n1, nil
}

// 这里模拟服务之间的掉用的效果.
func Run(ctx context.Context) {
	_, span := otel.Tracer("nihao").Start(ctx, "two")
	defer span.End()
	fibonacci, err := Fibonacci(20)
	if err != nil {
		return
	}
	fmt.Println(fibonacci)
}

// newExporter returns a console exporter.
// 创建导出器,对于链路追踪服务器的选择就是在这里配置
// 下面是基于jaeger的导出器.
// 这里需要注意,在使用jaeger导出器的时候传输格式是jaeger.thrift over HTTP
// 所以在配置url时请使用14268端口 与/api/traces
func newExporter(url string) (*jaeger.Exporter, error) {
	return jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
}

// 资源是一种特殊类型的属性，适用于进程生成的所有跨度。这些应该用于表示有关非临时进程的底层元数据
// 例如，进程的主机名或其实例 ID
func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("fib"),
			semconv.ServiceVersionKey.String("v0.1.0"),
			attribute.String("environment", "demo"),
		),
	)
	return r
}

// 主函数
func main() {
	//url := "http://collector.tl.com:31498/api/traces"
	url := "http://127.0.0.1:31498/api/traces"
	// 创建导出器
	exp, _ := newExporter(url)
	// 创建链路生成器,这里将导出器与资源信息配置进去.
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(newResource()),
	)
	// 结束时关闭链路生成器
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			fmt.Println(err)
		}
	}()

	// 将创建的生成器设置为全局变量.
	// 这样直接使用otel.Tracer就可以创建链路.
	// 否则 就要使用 tp.Tracer的形式创建链路.
	otel.SetTracerProvider(tp)
	newCtx, span := otel.Tracer("nihao").Start(context.Background(), "one")
	Run(newCtx)
	span.End()
}
