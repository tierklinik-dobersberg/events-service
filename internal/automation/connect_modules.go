package automation

import (
	"context"

	"github.com/dop251/goja"
	"github.com/elazarl/goproxy"
	"github.com/olebedev/gojax/fetch"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/idm/v1/idmv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/pbx3cx/v1/pbx3cxv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/roster/v1/rosterv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/tasks/v1/tasksv1connect"
)

func WithFetchModule() EngineOption {
	return func(e *Engine) {
		fetch.Enable(e.loop, goproxy.NewProxyHttpServer())
	}
}

func WithTaskModule(ctx context.Context, cli tasksv1connect.TaskServiceClient) EngineOption {
	return func(e *Engine) {
		e.RegisterNativeModuleHelper("tasks", map[string]any{
			"queryView":    wrapConnectMethod(cli.QueryView),
			"updateTask":   wrapConnectMethod(cli.UpdateTask),
			"completeTask": wrapConnectMethod(cli.CompleteTask),
			"deleteTask":   wrapConnectMethod(cli.DeleteTask),
			"assignTask":   wrapConnectMethod(cli.AssignTask),
			"getTimeline":  wrapConnectMethod(cli.GetTimeline),
		})
	}
}

func WithBoardModule(ctx context.Context, cli tasksv1connect.BoardServiceClient) EngineOption {
	return func(e *Engine) {
		e.Registry.RegisterNativeModule("boards", func(r *goja.Runtime, o *goja.Object) {
			e, ok := o.Get("exports").(*goja.Object)
			if !ok {
				panic("failed to add boards module")
			}

			e.Set("listBoards", wrapConnectMethod(cli.ListBoards))
			e.Set("getBoard", wrapConnectMethod(cli.GetBoard))
		})
	}
}

func WithUsersModule(ctx context.Context, cli idmv1connect.UserServiceClient) EngineOption {
	return func(e *Engine) {
		e.Registry.RegisterNativeModule("users", func(r *goja.Runtime, o *goja.Object) {
			e, ok := o.Get("exports").(*goja.Object)
			if !ok {
				panic("failed to add users module")
			}

			e.Set("listUsers", wrapConnectMethod(cli.ListUsers))
			e.Set("getUser", wrapConnectMethod(cli.GetUser))
		})
	}
}

func WithRolesModule(ctx context.Context, cli idmv1connect.RoleServiceClient) EngineOption {
	return func(e *Engine) {
		e.Registry.RegisterNativeModule("roles", func(r *goja.Runtime, o *goja.Object) {
			e, ok := o.Get("exports").(*goja.Object)
			if !ok {
				panic("failed to add users module")
			}

			e.Set("listRoles", wrapConnectMethod(cli.ListRoles))
			e.Set("getRole", wrapConnectMethod(cli.GetRole))
		})
	}
}

func WithNotifyModule(ctx context.Context, cli idmv1connect.NotifyServiceClient) EngineOption {
	return func(e *Engine) {
		e.Registry.RegisterNativeModule("notify", func(r *goja.Runtime, o *goja.Object) {
			e, ok := o.Get("exports").(*goja.Object)
			if !ok {
				panic("failed to add users module")
			}

			e.Set("send", wrapConnectMethod(cli.SendNotification))
		})
	}
}

func WithCallModule(ctx context.Context, cli pbx3cxv1connect.CallServiceClient) EngineOption {
	return func(e *Engine) {
		e.Registry.RegisterNativeModule("calls", func(r *goja.Runtime, o *goja.Object) {
			e, ok := o.Get("exports").(*goja.Object)
			if !ok {
				panic("failed to add users module")
			}

			e.Set("listInboundNumbers", wrapConnectMethod(cli.ListInboundNumber))
			e.Set("createOverwrite", wrapConnectMethod(cli.CreateOverwrite))
			e.Set("deleteOverwrite", wrapConnectMethod(cli.DeleteOverwrite))
			e.Set("getOverwrite", wrapConnectMethod(cli.GetOverwrite))
			e.Set("searchCallLogs", wrapConnectMethod(cli.SearchCallLogs))
			e.Set("getOnCall", wrapConnectMethod(cli.GetOnCall))
		})
	}
}

func WithVoiceMailModule(ctx context.Context, cli pbx3cxv1connect.VoiceMailServiceClient) EngineOption {
	return func(e *Engine) {
		e.Registry.RegisterNativeModule("voicemails", func(r *goja.Runtime, o *goja.Object) {
			e, ok := o.Get("exports").(*goja.Object)
			if !ok {
				panic("failed to add voicemails module")
			}

			e.Set("listMailboxes", wrapConnectMethod(cli.ListMailboxes))
			e.Set("getVoiceMail", wrapConnectMethod(cli.GetVoiceMail))
			e.Set("listVoiceMails", wrapConnectMethod(cli.ListVoiceMails))
			e.Set("markVoiceMails", wrapConnectMethod(cli.MarkVoiceMails))
		})
	}
}

func WithRosterModule(ctx context.Context, cli rosterv1connect.RosterServiceClient) EngineOption {
	return func(e *Engine) {
		e.Registry.RegisterNativeModule("roster", func(r *goja.Runtime, o *goja.Object) {
			e, ok := o.Get("exports").(*goja.Object)
			if !ok {
				panic("failed to add roster module")
			}

			e.Set("analyzeWorkTime", wrapConnectMethod(cli.AnalyzeWorkTime))
			e.Set("approveRoster", wrapConnectMethod(cli.ApproveRoster))
			e.Set("deleteRoster", wrapConnectMethod(cli.DeleteRoster))
			e.Set("exportRoster", wrapConnectMethod(cli.ExportRoster))
			e.Set("getRequiredShifts", wrapConnectMethod(cli.GetRequiredShifts))
			e.Set("getRoster", wrapConnectMethod(cli.GetRoster))
			e.Set("getUserShifts", wrapConnectMethod(cli.GetUserShifts))
			e.Set("getWorkingStaff2", wrapConnectMethod(cli.GetWorkingStaff2))
			e.Set("listRosterTypes", wrapConnectMethod(cli.ListRosterTypes))
			e.Set("listShiftTags", wrapConnectMethod(cli.ListShiftTags))
			e.Set("saveRoster", wrapConnectMethod(cli.SaveRoster))
		})
	}
}
