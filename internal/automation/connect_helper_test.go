package automation

import (
	"context"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/dop251/goja"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	idmv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/idm/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func Test_withProtoMessage(t *testing.T) {
	called := false

	engine, err := New("", nil, func(e *Engine) {
		e.Registry.RegisterNativeModule("testlib", func(r *goja.Runtime, o *goja.Object) {
			exports := o.Get("exports").(*goja.Object)

			exports.Set("testProto", withProtoMessage(func(t *idmv1.GetUserRequest) (*idmv1.GetUserResponse, error) {
				called = true
				return &idmv1.GetUserResponse{
					Profile: &idmv1.Profile{
						User: &idmv1.User{
							Id: "test",
						},
					},
				}, nil
			}))
		})

	})

	require.NoError(t, err, "creating a new engine should not fail")

	err = engine.RunScript(`
		const test = require("testlib");

		console.log("foobar", test.testProto)
		const result = test.testProto({
			name: "alice"
		})

		console.log(result);

		if (result.profile.user.id !== 'test') {
			throw new Error("expected correct response")
		}
	`)
	require.NoError(t, err, "running the javascript test script should not have failed")

	assert.Equal(t, 0, engine.Stop(), "expected no more running jobs")
	assert.True(t, called, "expected native module function to have been called")
}

func Test_wrapConnectMethod(t *testing.T) {
	called := false

	engine, err := New("", nil, func(e *Engine) {
		e.Registry.RegisterNativeModule("testlib", func(r *goja.Runtime, o *goja.Object) {
			exports := o.Get("exports").(*goja.Object)

			exports.Set("connect", wrapConnectMethod(func(_ context.Context, req *connect.Request[idmv1.GetUserRequest]) (*connect.Response[idmv1.GetUserRequest], error) {
				called = true

				assert.True(
					t,
					proto.Equal(
						&idmv1.GetUserRequest{
							Search: &idmv1.GetUserRequest_Name{
								Name: "alice",
							},
							FieldMask: &fieldmaskpb.FieldMask{
								Paths: []string{"profile.user.id"},
							},
						},
						req.Msg,
					),
					"expected request message",
				)

				return connect.NewResponse(req.Msg), nil
			}))
		})

	})

	require.NoError(t, err, "creating a new engine should not fail")

	err = engine.RunScript(`
		const test = require("testlib");

		const result = test.connect({
			name: "alice",
			fieldMask: "profile.user.id",
		})

		if (result.name !== 'alice' && result.fieldMask === 'profile.user.id') {
			throw new Error("expected correct response")
		}
	`)
	require.NoError(t, err, "running the javascript test script should not have failed")

	assert.Equal(t, 0, engine.Stop(), "expected no more running jobs")
	assert.True(t, called, "expected native module function to have been called")
}
