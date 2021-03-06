/**
 * @fileoverview gRPC-Web generated client stub for 
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as hello_pb from './hello_pb';


export class HelloServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoSayHello = new grpcWeb.AbstractClientBase.MethodInfo(
    hello_pb.HelloResponse,
    (request: hello_pb.HelloRequest) => {
      return request.serializeBinary();
    },
    hello_pb.HelloResponse.deserializeBinary
  );

  sayHello(
    request: hello_pb.HelloRequest,
    metadata: grpcWeb.Metadata | null): Promise<hello_pb.HelloResponse>;

  sayHello(
    request: hello_pb.HelloRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: hello_pb.HelloResponse) => void): grpcWeb.ClientReadableStream<hello_pb.HelloResponse>;

  sayHello(
    request: hello_pb.HelloRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: hello_pb.HelloResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/HelloService/SayHello',
        request,
        metadata || {},
        this.methodInfoSayHello,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/HelloService/SayHello',
    request,
    metadata || {},
    this.methodInfoSayHello);
  }

  methodInfoGetUsers = new grpcWeb.AbstractClientBase.MethodInfo(
    hello_pb.GetUsersResponse,
    (request: hello_pb.GetUsersRequest) => {
      return request.serializeBinary();
    },
    hello_pb.GetUsersResponse.deserializeBinary
  );

  getUsers(
    request: hello_pb.GetUsersRequest,
    metadata: grpcWeb.Metadata | null): Promise<hello_pb.GetUsersResponse>;

  getUsers(
    request: hello_pb.GetUsersRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: hello_pb.GetUsersResponse) => void): grpcWeb.ClientReadableStream<hello_pb.GetUsersResponse>;

  getUsers(
    request: hello_pb.GetUsersRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: hello_pb.GetUsersResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/HelloService/GetUsers',
        request,
        metadata || {},
        this.methodInfoGetUsers,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/HelloService/GetUsers',
    request,
    metadata || {},
    this.methodInfoGetUsers);
  }

  methodInfoCreateUser = new grpcWeb.AbstractClientBase.MethodInfo(
    hello_pb.CreateUserResponse,
    (request: hello_pb.CreateUserRequest) => {
      return request.serializeBinary();
    },
    hello_pb.CreateUserResponse.deserializeBinary
  );

  createUser(
    request: hello_pb.CreateUserRequest,
    metadata: grpcWeb.Metadata | null): Promise<hello_pb.CreateUserResponse>;

  createUser(
    request: hello_pb.CreateUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: hello_pb.CreateUserResponse) => void): grpcWeb.ClientReadableStream<hello_pb.CreateUserResponse>;

  createUser(
    request: hello_pb.CreateUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: hello_pb.CreateUserResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/HelloService/CreateUser',
        request,
        metadata || {},
        this.methodInfoCreateUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/HelloService/CreateUser',
    request,
    metadata || {},
    this.methodInfoCreateUser);
  }

}

