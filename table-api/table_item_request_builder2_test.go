package tableapi

import (
	"context"
	"errors"
	"testing"

	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
	"github.com/michaeldcanady/servicenow-sdk-go/table-api/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var _ intCore.Sendable = (*mockRequestBuilder)(nil)

type mockRequestBuilder struct {
	mock.Mock
}

func (rB *mockRequestBuilder) Send(ctx context.Context, method intCore.HttpMethod, config intCore.RequestConfiguration) (interface{}, error) {
	args := rB.Called(ctx, method, config)
	return args.Get(0), args.Error(1)
}

func TestNewTableItemRequestBuilder2(t *testing.T) {
	tests := []internal.Test[any]{
		{
			Title: "Successful",
			Input: []interface{}{&mockClient{}, map[string]string{"baseurl": "baseurl", "table": "table", "sysId": "fdafsdfdsa"}},
		},
		{
			Title:       "missing sysId",
			Input:       []interface{}{&mockClient{}, map[string]string{"baseurl": "baseurl", "table": "table"}},
			ExpectedErr: errors.New("missing \"sysId\" parameter"),
		},
		{
			Title:       "missing table",
			Input:       []interface{}{&mockClient{}, map[string]string{"baseurl": "baseurl"}},
			ExpectedErr: errors.New("missing \"table\" parameter"),
		},
		{
			Title:       "missing baseurl",
			Input:       []interface{}{&mockClient{}, map[string]string{}},
			ExpectedErr: errors.New("pathParameters must contain a value for \"baseurl\" for the URL to be built"),
		},
		{
			Title:       "missing client",
			Input:       []interface{}{(*mockClient)(nil), map[string]string{}},
			ExpectedErr: errors.New("client can't be nil"),
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if test.Setup != nil {
				test.Setup()
			}

			inputs := test.Input.([]interface{})

			_, err := newTableItemRequestBuilder2[*TableRecordImpl](inputs[0].(intCore.ClientSendable), inputs[1].(map[string]string))
			assert.Equal(t, test.ExpectedErr, err)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}

//nolint:dupl
func TestTableItemRequestBuilder2_Get(t *testing.T) {
	requestBuilder := &tableItemRequestBuilder2[*TableRecordImpl]{&mockRequestBuilder{}}

	tests := []internal.Test[any]{
		{
			Title: "Successful",
			Setup: func() {
				mockRB := &mockRequestBuilder{}
				mockRB.On("SendGet2", mock.AnythingOfType("*core.RequestConfiguration")).Return(nil)

				requestBuilder.Sendable = mockRB
			},
			ExpectedErr: nil,
		},
		{
			Title: "Error",
			Setup: func() {
				mockRB := &mockRequestBuilder{}
				mockRB.On("SendGet2", mock.AnythingOfType("*core.RequestConfiguration")).Return(errors.New("unable to send request"))

				requestBuilder.Sendable = mockRB
			},
			ExpectedErr: errors.New("unable to send request"),
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if test.Setup != nil {
				test.Setup()
			}

			_, err := requestBuilder.Get(context.Background(), nil)
			assert.Equal(t, test.ExpectedErr, err)
			requestBuilder.Sendable.(*mockRequestBuilder).AssertExpectations(t)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}

func TestTableItemRequestBuilder2_Delete(t *testing.T) {
	requestBuilder := &tableItemRequestBuilder2[*TableRecordImpl]{&mockRequestBuilder{}}

	tests := []internal.Test[any]{
		{
			Title: "Successful",
			Setup: func() {
				mockRB := &mockRequestBuilder{}
				mockRB.On("SendDelete2", mock.AnythingOfType("*core.RequestConfiguration")).Return(nil)

				requestBuilder.Sendable = mockRB
			},
			ExpectedErr: nil,
		},
		{
			Title: "Error",
			Setup: func() {
				mockRB := &mockRequestBuilder{}
				mockRB.On("SendDelete2", mock.AnythingOfType("*core.RequestConfiguration")).Return(errors.New("unable to send request"))

				requestBuilder.Sendable = mockRB
			},
			ExpectedErr: errors.New("unable to send request"),
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if test.Setup != nil {
				test.Setup()
			}

			err := requestBuilder.Delete(context.Background(), nil)
			assert.Equal(t, test.ExpectedErr, err)
			requestBuilder.Sendable.(*mockRequestBuilder).AssertExpectations(t)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}

func TestTableItemRequestBuilder2_Put(t *testing.T) {
	requestBuilder := &tableItemRequestBuilder2[*TableRecordImpl]{&mockRequestBuilder{}}

	tests := []internal.Test[any]{
		{
			Title: "Successful",
			Setup: func() {
				mockRB := &mockRequestBuilder{}
				mockRB.On("SendPut2", mock.AnythingOfType("*core.RequestConfiguration")).Return(nil)

				requestBuilder.Sendable = mockRB
			},
			Input:       NewTableEntry(),
			ExpectedErr: nil,
		},
		{
			Title: "Error",
			Setup: func() {
				mockRB := &mockRequestBuilder{}
				mockRB.On("SendPut2", mock.AnythingOfType("*core.RequestConfiguration")).Return(errors.New("unable to send request"))

				requestBuilder.Sendable = mockRB
			},
			Input:       NewTableEntry(),
			ExpectedErr: errors.New("unable to send request"),
		},
		{
			Title: "Nil tableEntry",
			Setup: func() {
				mockRB := &mockRequestBuilder{}

				requestBuilder.Sendable = mockRB
			},
			Input:       nil,
			ExpectedErr: errors.New("entry is nil"),
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			if test.Setup != nil {
				test.Setup()
			}

			_, err := requestBuilder.Put(context.Background(), test.Input, nil)
			assert.Equal(t, test.ExpectedErr, err)
			requestBuilder.Sendable.(*mockRequestBuilder).AssertExpectations(t)

			if test.Cleanup != nil {
				test.Cleanup()
			}
		})
	}
}
