import * as jspb from 'google-protobuf'



export class RegisterRequest extends jspb.Message {
  getName(): string;
  setName(value: string): RegisterRequest;

  getEmail(): string;
  setEmail(value: string): RegisterRequest;

  getPassword(): string;
  setPassword(value: string): RegisterRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterRequest): RegisterRequest.AsObject;
  static serializeBinaryToWriter(message: RegisterRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterRequest;
  static deserializeBinaryFromReader(message: RegisterRequest, reader: jspb.BinaryReader): RegisterRequest;
}

export namespace RegisterRequest {
  export type AsObject = {
    name: string,
    email: string,
    password: string,
  }
}

export class LoginRequest extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): LoginRequest;

  getPassword(): string;
  setPassword(value: string): LoginRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LoginRequest): LoginRequest.AsObject;
  static serializeBinaryToWriter(message: LoginRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginRequest;
  static deserializeBinaryFromReader(message: LoginRequest, reader: jspb.BinaryReader): LoginRequest;
}

export namespace LoginRequest {
  export type AsObject = {
    email: string,
    password: string,
  }
}

export class AuthResponse extends jspb.Message {
  getToken(): string;
  setToken(value: string): AuthResponse;

  getUserid(): string;
  setUserid(value: string): AuthResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AuthResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AuthResponse): AuthResponse.AsObject;
  static serializeBinaryToWriter(message: AuthResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AuthResponse;
  static deserializeBinaryFromReader(message: AuthResponse, reader: jspb.BinaryReader): AuthResponse;
}

export namespace AuthResponse {
  export type AsObject = {
    token: string,
    userid: string,
  }
}

export class UserIdRequest extends jspb.Message {
  getUserid(): string;
  setUserid(value: string): UserIdRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserIdRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UserIdRequest): UserIdRequest.AsObject;
  static serializeBinaryToWriter(message: UserIdRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserIdRequest;
  static deserializeBinaryFromReader(message: UserIdRequest, reader: jspb.BinaryReader): UserIdRequest;
}

export namespace UserIdRequest {
  export type AsObject = {
    userid: string,
  }
}

export class UpdateProfileRequest extends jspb.Message {
  getUserid(): string;
  setUserid(value: string): UpdateProfileRequest;

  getName(): string;
  setName(value: string): UpdateProfileRequest;

  getGoal(): string;
  setGoal(value: string): UpdateProfileRequest;

  getHeight(): number;
  setHeight(value: number): UpdateProfileRequest;

  getWeight(): number;
  setWeight(value: number): UpdateProfileRequest;

  getAddress(): string;
  setAddress(value: string): UpdateProfileRequest;

  getPhone(): string;
  setPhone(value: string): UpdateProfileRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateProfileRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateProfileRequest): UpdateProfileRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateProfileRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateProfileRequest;
  static deserializeBinaryFromReader(message: UpdateProfileRequest, reader: jspb.BinaryReader): UpdateProfileRequest;
}

export namespace UpdateProfileRequest {
  export type AsObject = {
    userid: string,
    name: string,
    goal: string,
    height: number,
    weight: number,
    address: string,
    phone: string,
  }
}

export class UserResponse extends jspb.Message {
  getUserid(): string;
  setUserid(value: string): UserResponse;

  getName(): string;
  setName(value: string): UserResponse;

  getEmail(): string;
  setEmail(value: string): UserResponse;

  getGoal(): string;
  setGoal(value: string): UserResponse;

  getHeight(): number;
  setHeight(value: number): UserResponse;

  getWeight(): number;
  setWeight(value: number): UserResponse;

  getAddress(): string;
  setAddress(value: string): UserResponse;

  getPhone(): string;
  setPhone(value: string): UserResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UserResponse): UserResponse.AsObject;
  static serializeBinaryToWriter(message: UserResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserResponse;
  static deserializeBinaryFromReader(message: UserResponse, reader: jspb.BinaryReader): UserResponse;
}

export namespace UserResponse {
  export type AsObject = {
    userid: string,
    name: string,
    email: string,
    goal: string,
    height: number,
    weight: number,
    address: string,
    phone: string,
  }
}

export class Empty extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Empty.AsObject;
  static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
  static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Empty;
  static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
  export type AsObject = {
  }
}

