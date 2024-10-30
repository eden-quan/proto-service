package serves

import (
	"sync"
	"fmt"

	messagecenterservicev1 "gitlab.lainuoniao.cn/eden-quan/proto-service/api/messagecenter-service/v1/services"
)

var _singleMessagecenter struct {
	messagecenter messagecenterservicev1.MessagecenterV1Server
	lock sync.Once
}

func NewSingletonMessagecenterServer(arg NewMessagecenterArgs) messagecenterservicev1.MessagecenterV1Server {
	_singleMessagecenter.lock.Do(func() {
	    var err error
		_singleMessagecenter.messagecenter, err = NewMessagecenterServer(arg)
		if err != nil {
		    panic(fmt.Sprintf("create Messagecenter.messagecenter with error %s", err))
		}
	})

	return _singleMessagecenter.messagecenter
}
