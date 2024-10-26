package services

import (
	"context"
	stdhttp "net/http"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"

	errorv1 "gitlab.lainuoniao.cn/eden-quan/proto-service/api/proto-common/v1/errors"
	repos "gitlab.lainuoniao.cn/eden-quan/proto-service/app/sms-service/internal/domain/repo"

	errorpkg "gitlab.lainuoniao.cn/eden-quan/go-kratos-pkg/error"
	websocketpkg "gitlab.lainuoniao.cn/eden-quan/go-kratos-pkg/websocket"
)

// websocketService ...
type websocketService struct {
	log *log.Helper
}

// NewWebsocketService ...
func NewWebsocketService(logger log.Logger) repos.WebsocketService {
	return &websocketService{
		log: log.NewHelper(log.With(logger, "module", "ping/domain/service/websocket")),
	}
}

// WsMessage ws
type WsMessage struct {
	Type    int
	Content []byte
}

// Ws websocket
func (s *websocketService) Ws(ctx context.Context, w stdhttp.ResponseWriter, r *stdhttp.Request) (err error) {
	// 升级连接
	cc, err := websocketpkg.UpgradeConn(w, r, w.Header())
	if err != nil {
		err = errorpkg.InternalServer(errorv1.ERROR_CONNECTION.String(), "upgrade conn failed", err)
		s.log.WithContext(ctx).Error(err)
		return
	}
	defer func() { _ = cc.Close() }()

	// 处理消息
	err = s.processWssMsg(ctx, cc)
	if err != nil {
		return err
	}
	return err
}

// processWssMsg ...
func (s *websocketService) processWssMsg(ctx context.Context, cc *websocket.Conn) error {
	// 读消息
	for {
		messageType, messageBytes, wsErr := cc.ReadMessage()
		if wsErr != nil {
			if websocketpkg.IsCloseError(wsErr) {
				s.log.WithContext(ctx).Infow("websocket close", wsErr.Error())
				break
			}
			err := errorpkg.InternalServer(errorv1.ERROR_CONNECTION.String(), "ws读取信息失败", wsErr)
			s.log.WithContext(ctx).Error(err)
			return err
		}

		// 消息
		wsMessage := &WsMessage{
			Type:    messageType,
			Content: messageBytes,
		}
		// messageChan <- wsMessage

		// 处理
		needCloseConn, err := s.processWsMessage(ctx, wsMessage)
		if err != nil {
			err = errorpkg.InternalServer(errorv1.ERROR_INTERNAL_SERVER.String(), "ws处理信息失败", err)
			s.log.WithContext(ctx).Error(err)
			return err
		}

		// 响应
		err = cc.WriteMessage(messageType, messageBytes)
		if err != nil {
			if websocketpkg.IsCloseError(wsErr) {
				s.log.WithContext(ctx).Infow("websocket close", wsErr.Error())
				break
			}
			err = errorpkg.InternalServer(errorv1.ERROR_INTERNAL_SERVER.String(), "JSON编码错误", err)
			s.log.WithContext(ctx).Error(err)
			return err
		}

		// 关闭
		if needCloseConn {
			return err
		}
	}
	return nil
}

// processWsMessage 处理ws-message
func (s *websocketService) processWsMessage(ctx context.Context, wsMessage *WsMessage) (needCloseConn bool, err error) {
	s.log.WithContext(ctx).Infow("ws-message type", wsMessage.Type)
	s.log.WithContext(ctx).Infow("ws-message message", string(wsMessage.Content))
	if string(wsMessage.Content) == "close" {
		needCloseConn = true
	}
	return needCloseConn, err
}
