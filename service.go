package supervisor_service

import (
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	supervisorv1 "github.com/submaline/supervisor-service/gen/supervisor/v1"
)

type SupervisorService struct{}

func (ss *SupervisorService) CreateAccount(
	_ context.Context, request *connect.Request[supervisorv1.CreateAccountRequest]) (
	*connect.Response[supervisorv1.CreateAccountResponse], error) {
	return nil, connect.NewError(
		connect.CodeUnimplemented,
		fmt.Errorf(fmt.Sprintf("%s is unimplemented", request.Spec().Procedure)))
}

func (ss *SupervisorService) CreateProfile(
	_ context.Context, request *connect.Request[supervisorv1.CreateProfileRequest]) (
	*connect.Response[supervisorv1.CreateProfileResponse], error) {
	return nil, connect.NewError(
		connect.CodeUnimplemented,
		fmt.Errorf(fmt.Sprintf("%s is unimplemented", request.Spec().Procedure)))
}

func (ss *SupervisorService) RecordOperation(
	_ context.Context, request *connect.Request[supervisorv1.RecordOperationRequest]) (
	*connect.Response[supervisorv1.RecordOperationResponse], error) {
	return nil, connect.NewError(
		connect.CodeUnimplemented,
		fmt.Errorf(fmt.Sprintf("%s is unimplemented", request.Spec().Procedure)))
}
