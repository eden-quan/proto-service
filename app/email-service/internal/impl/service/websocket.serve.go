package serves

import (
	"github.com/go-kratos/kratos/v2/transport/http"

	errorv1 "gitlab.lainuoniao.cn/eden-quan/proto-service/api/proto-common/v1/errors"
	repos "gitlab.lainuoniao.cn/eden-quan/proto-service/app/email-service/internal/domain/repo"

	errorpkg "gitlab.lainuoniao.cn/eden-quan/go-kratos-pkg/error"

	stdhttp "net/http"
)

// WebsocketServer ...
type WebsocketServer interface {
	TestWebsocket(w http.ResponseWriter, r *http.Request)
}

// websocketServer ...
type websocketServer struct {
	websocketService repos.WebsocketService
}

// NewWebsocketServer ...
func NewWebsocketServer(wsRepo repos.WebsocketService) WebsocketServer {
	return &websocketServer{
		websocketService: wsRepo,
	}
}

// TestWebsocket ...
func (s *websocketServer) TestWebsocket(w http.ResponseWriter, r *http.Request) {
	if r.Method != stdhttp.MethodGet {
		err := errorpkg.InternalServer(errorv1.ERROR_METHOD_NOT_ALLOWED.String(), "ERROR_METHOD_NOT_ALLOWED")
		w.WriteHeader(stdhttp.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	err := s.websocketService.Ws(r.Context(), w, r)
	if err != nil {
		w.WriteHeader(stdhttp.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	return
}
