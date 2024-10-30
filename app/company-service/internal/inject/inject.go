package inject

import "gitlab.lainuoniao.cn/eden-quan/go-biz-kit/injection"


var injector = injection.NewInjector()

func Injection() *injection.Injector {
	return &injector
}
