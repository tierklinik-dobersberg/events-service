package automation

import (
	"context"

	"github.com/elazarl/goproxy"
	"github.com/olebedev/gojax/fetch"
	idmv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/idm/v1"
	pbx3cxv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/pbx3cx/v1"
	rosterv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/roster/v1"
	tasksv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/tasks/v1"
)

func WithFetchModule() EngineOption {
	return func(e *Engine) {
		fetch.Enable(e.loop, goproxy.NewProxyHttpServer())
	}
}

func WithTaskModule(ctx context.Context, ep string) EngineOption {
	return WithConnectService("tasks", ep, tasksv1.File_tkd_tasks_v1_tasks_proto.Services().Get(0))
}

func WithBoardModule(ctx context.Context, ep string) EngineOption {
	return WithConnectService("boards", ep, tasksv1.File_tkd_tasks_v1_boards_proto.Services().Get(0))
}

func WithUsersModule(ctx context.Context, ep string) EngineOption {
	return WithConnectService("users", ep, idmv1.File_tkd_idm_v1_user_service_proto.Services().Get(0))
}

func WithRolesModule(ctx context.Context, ep string) EngineOption {
	return WithConnectService("roles", ep, idmv1.File_tkd_idm_v1_role_service_proto.Services().Get(0))
}

func WithNotifyModule(ctx context.Context, ep string) EngineOption {
	return WithConnectService("notify", ep, idmv1.File_tkd_idm_v1_notify_service_proto.Services().Get(0))
}

func WithCallModule(ctx context.Context, ep string) EngineOption {
	return WithConnectService("calls", ep, pbx3cxv1.File_tkd_pbx3cx_v1_calllog_proto.Services().Get(0))
}

func WithVoiceMailModule(ctx context.Context, ep string) EngineOption {
	return WithConnectService("calls", ep, pbx3cxv1.File_tkd_pbx3cx_v1_voicemail_proto.Services().Get(0))
}

func WithRosterModule(ctx context.Context, ep string) EngineOption {
	return WithConnectService("calls", ep, rosterv1.File_tkd_roster_v1_roster_proto.Services().Get(0))
}
