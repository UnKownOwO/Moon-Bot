package statsviz

import (
	"net/http"

	"github.com/arl/statsviz"
)

// Run 运行性能监测服务
// 原生pprof /debug/pprof
// 可视化图表 /debug/statsviz
func Run(addr string) error {
	err := statsviz.RegisterDefault()
	if err != nil {
		return err
	}
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		return err
	}
	return nil
}
