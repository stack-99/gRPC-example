// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: quiz.proto

package models

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// QuizClient is the client API for Quiz service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuizClient interface {
	GetQuizQuestions(ctx context.Context, in *QuestionRequest, opts ...grpc.CallOption) (*QuestionResponse, error)
	AnswerQuiz(ctx context.Context, in *AnswerRequest, opts ...grpc.CallOption) (*AnswerResponse, error)
	GetUserScore(ctx context.Context, in *UserScoreRequest, opts ...grpc.CallOption) (*UserScoreResponse, error)
}

type quizClient struct {
	cc grpc.ClientConnInterface
}

func NewQuizClient(cc grpc.ClientConnInterface) QuizClient {
	return &quizClient{cc}
}

func (c *quizClient) GetQuizQuestions(ctx context.Context, in *QuestionRequest, opts ...grpc.CallOption) (*QuestionResponse, error) {
	out := new(QuestionResponse)
	err := c.cc.Invoke(ctx, "/quiz.Quiz/GetQuizQuestions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizClient) AnswerQuiz(ctx context.Context, in *AnswerRequest, opts ...grpc.CallOption) (*AnswerResponse, error) {
	out := new(AnswerResponse)
	err := c.cc.Invoke(ctx, "/quiz.Quiz/AnswerQuiz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quizClient) GetUserScore(ctx context.Context, in *UserScoreRequest, opts ...grpc.CallOption) (*UserScoreResponse, error) {
	out := new(UserScoreResponse)
	err := c.cc.Invoke(ctx, "/quiz.Quiz/GetUserScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuizServer is the server API for Quiz service.
// All implementations must embed UnimplementedQuizServer
// for forward compatibility
type QuizServer interface {
	GetQuizQuestions(context.Context, *QuestionRequest) (*QuestionResponse, error)
	AnswerQuiz(context.Context, *AnswerRequest) (*AnswerResponse, error)
	GetUserScore(context.Context, *UserScoreRequest) (*UserScoreResponse, error)
	mustEmbedUnimplementedQuizServer()
}

// UnimplementedQuizServer must be embedded to have forward compatible implementations.
type UnimplementedQuizServer struct {
}

func (UnimplementedQuizServer) GetQuizQuestions(context.Context, *QuestionRequest) (*QuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuizQuestions not implemented")
}
func (UnimplementedQuizServer) AnswerQuiz(context.Context, *AnswerRequest) (*AnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnswerQuiz not implemented")
}
func (UnimplementedQuizServer) GetUserScore(context.Context, *UserScoreRequest) (*UserScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserScore not implemented")
}
func (UnimplementedQuizServer) mustEmbedUnimplementedQuizServer() {}

// UnsafeQuizServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuizServer will
// result in compilation errors.
type UnsafeQuizServer interface {
	mustEmbedUnimplementedQuizServer()
}

func RegisterQuizServer(s grpc.ServiceRegistrar, srv QuizServer) {
	s.RegisterService(&Quiz_ServiceDesc, srv)
}

func _Quiz_GetQuizQuestions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizServer).GetQuizQuestions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quiz.Quiz/GetQuizQuestions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizServer).GetQuizQuestions(ctx, req.(*QuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Quiz_AnswerQuiz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizServer).AnswerQuiz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quiz.Quiz/AnswerQuiz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizServer).AnswerQuiz(ctx, req.(*AnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Quiz_GetUserScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizServer).GetUserScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quiz.Quiz/GetUserScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizServer).GetUserScore(ctx, req.(*UserScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Quiz_ServiceDesc is the grpc.ServiceDesc for Quiz service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Quiz_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "quiz.Quiz",
	HandlerType: (*QuizServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuizQuestions",
			Handler:    _Quiz_GetQuizQuestions_Handler,
		},
		{
			MethodName: "AnswerQuiz",
			Handler:    _Quiz_AnswerQuiz_Handler,
		},
		{
			MethodName: "GetUserScore",
			Handler:    _Quiz_GetUserScore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quiz.proto",
}
