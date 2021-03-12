package validate

import (
	"fmt"
	pb "github.com/naoyakurokawa/app-grpc-web/hello"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	// "log"
	// "os"
)

func CreateError(code codes.Code, errorList []*errdetails.BadRequest_FieldViolation) error {
	st := status.New(codes.InvalidArgument, "エラー発生")
	// add error message detail
	st, err := st.WithDetails(
		&errdetails.BadRequest{
			FieldViolations: errorList,
		},
	)
	// unexpected error
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %+v", err))
	}

	// return error
	return st.Err()
}

func CreateBadRequestFieldViolation(feild string, desc string) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       feild,
		Description: desc,
	}
}

func CheckLoginUserRequest(request pb.LoginRequest) error {
	var errorList []*errdetails.BadRequest_FieldViolation
	if request.Name == "" {
		errorList = append(errorList, CreateBadRequestFieldViolation("Name", "名前は必須です"))
	}
	if request.Password == "" {
		errorList = append(errorList, CreateBadRequestFieldViolation("Password", "必須です"))
	}
	if len(errorList) > 0 {
		// log.Printf("エラー : %s", errorList)
		return CreateError(codes.InvalidArgument, errorList)
	} else {
		return nil
	}
}
