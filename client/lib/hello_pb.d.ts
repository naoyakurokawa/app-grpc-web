import * as jspb from 'google-protobuf'



export class HelloRequest extends jspb.Message {
  getName(): string;
  setName(value: string): HelloRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HelloRequest.AsObject;
  static toObject(includeInstance: boolean, msg: HelloRequest): HelloRequest.AsObject;
  static serializeBinaryToWriter(message: HelloRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HelloRequest;
  static deserializeBinaryFromReader(message: HelloRequest, reader: jspb.BinaryReader): HelloRequest;
}

export namespace HelloRequest {
  export type AsObject = {
    name: string,
  }
}

export class HelloResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): HelloResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HelloResponse.AsObject;
  static toObject(includeInstance: boolean, msg: HelloResponse): HelloResponse.AsObject;
  static serializeBinaryToWriter(message: HelloResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HelloResponse;
  static deserializeBinaryFromReader(message: HelloResponse, reader: jspb.BinaryReader): HelloResponse;
}

export namespace HelloResponse {
  export type AsObject = {
    message: string,
  }
}

export class User extends jspb.Message {
  getId(): number;
  setId(value: number): User;

  getName(): string;
  setName(value: string): User;

  getScore(): number;
  setScore(value: number): User;

  getPhotourl(): string;
  setPhotourl(value: string): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    id: number,
    name: string,
    score: number,
    photourl: string,
  }
}

export class GetUsersRequest extends jspb.Message {
  getId(): number;
  setId(value: number): GetUsersRequest;

  getName(): string;
  setName(value: string): GetUsersRequest;

  getScore(): number;
  setScore(value: number): GetUsersRequest;

  getPhotourl(): string;
  setPhotourl(value: string): GetUsersRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUsersRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetUsersRequest): GetUsersRequest.AsObject;
  static serializeBinaryToWriter(message: GetUsersRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUsersRequest;
  static deserializeBinaryFromReader(message: GetUsersRequest, reader: jspb.BinaryReader): GetUsersRequest;
}

export namespace GetUsersRequest {
  export type AsObject = {
    id: number,
    name: string,
    score: number,
    photourl: string,
  }
}

export class GetUsersResponse extends jspb.Message {
  getUsersList(): Array<User>;
  setUsersList(value: Array<User>): GetUsersResponse;
  clearUsersList(): GetUsersResponse;
  addUsers(value?: User, index?: number): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUsersResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetUsersResponse): GetUsersResponse.AsObject;
  static serializeBinaryToWriter(message: GetUsersResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUsersResponse;
  static deserializeBinaryFromReader(message: GetUsersResponse, reader: jspb.BinaryReader): GetUsersResponse;
}

export namespace GetUsersResponse {
  export type AsObject = {
    usersList: Array<User.AsObject>,
  }
}

export class CreateUserRequest extends jspb.Message {
  getId(): number;
  setId(value: number): CreateUserRequest;

  getName(): string;
  setName(value: string): CreateUserRequest;

  getScore(): number;
  setScore(value: number): CreateUserRequest;

  getPhotourl(): string;
  setPhotourl(value: string): CreateUserRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateUserRequest): CreateUserRequest.AsObject;
  static serializeBinaryToWriter(message: CreateUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateUserRequest;
  static deserializeBinaryFromReader(message: CreateUserRequest, reader: jspb.BinaryReader): CreateUserRequest;
}

export namespace CreateUserRequest {
  export type AsObject = {
    id: number,
    name: string,
    score: number,
    photourl: string,
  }
}

export class CreateUserResponse extends jspb.Message {
  getTxt(): string;
  setTxt(value: string): CreateUserResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateUserResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateUserResponse): CreateUserResponse.AsObject;
  static serializeBinaryToWriter(message: CreateUserResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateUserResponse;
  static deserializeBinaryFromReader(message: CreateUserResponse, reader: jspb.BinaryReader): CreateUserResponse;
}

export namespace CreateUserResponse {
  export type AsObject = {
    txt: string,
  }
}

export class GetUserByIdRequest extends jspb.Message {
  getId(): number;
  setId(value: number): GetUserByIdRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUserByIdRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetUserByIdRequest): GetUserByIdRequest.AsObject;
  static serializeBinaryToWriter(message: GetUserByIdRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUserByIdRequest;
  static deserializeBinaryFromReader(message: GetUserByIdRequest, reader: jspb.BinaryReader): GetUserByIdRequest;
}

export namespace GetUserByIdRequest {
  export type AsObject = {
    id: number,
  }
}

export class GetUserByIdResponse extends jspb.Message {
  getUserList(): Array<User>;
  setUserList(value: Array<User>): GetUserByIdResponse;
  clearUserList(): GetUserByIdResponse;
  addUser(value?: User, index?: number): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUserByIdResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetUserByIdResponse): GetUserByIdResponse.AsObject;
  static serializeBinaryToWriter(message: GetUserByIdResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUserByIdResponse;
  static deserializeBinaryFromReader(message: GetUserByIdResponse, reader: jspb.BinaryReader): GetUserByIdResponse;
}

export namespace GetUserByIdResponse {
  export type AsObject = {
    userList: Array<User.AsObject>,
  }
}

export class DeleteUserRequest extends jspb.Message {
  getId(): number;
  setId(value: number): DeleteUserRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteUserRequest): DeleteUserRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteUserRequest;
  static deserializeBinaryFromReader(message: DeleteUserRequest, reader: jspb.BinaryReader): DeleteUserRequest;
}

export namespace DeleteUserRequest {
  export type AsObject = {
    id: number,
  }
}

export class DeleteUserResponse extends jspb.Message {
  getId(): number;
  setId(value: number): DeleteUserResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteUserResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteUserResponse): DeleteUserResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteUserResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteUserResponse;
  static deserializeBinaryFromReader(message: DeleteUserResponse, reader: jspb.BinaryReader): DeleteUserResponse;
}

export namespace DeleteUserResponse {
  export type AsObject = {
    id: number,
  }
}

