// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: films.proto

package pbfilms

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PeopleData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ExternalReference string `protobuf:"bytes,2,opt,name=external_reference,json=externalReference,proto3" json:"external_reference,omitempty"`
	Age               string `protobuf:"bytes,3,opt,name=age,proto3" json:"age,omitempty"`
	Gender            string `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	HairColor         string `protobuf:"bytes,5,opt,name=hair_color,json=hairColor,proto3" json:"hair_color,omitempty"`
	EyeColor          string `protobuf:"bytes,6,opt,name=eye_color,json=eyeColor,proto3" json:"eye_color,omitempty"`
	Slug              string `protobuf:"bytes,8,opt,name=slug,proto3" json:"slug,omitempty"`
}

func (x *PeopleData) Reset() {
	*x = PeopleData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_films_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeopleData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeopleData) ProtoMessage() {}

func (x *PeopleData) ProtoReflect() protoreflect.Message {
	mi := &file_films_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeopleData.ProtoReflect.Descriptor instead.
func (*PeopleData) Descriptor() ([]byte, []int) {
	return file_films_proto_rawDescGZIP(), []int{0}
}

func (x *PeopleData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PeopleData) GetExternalReference() string {
	if x != nil {
		return x.ExternalReference
	}
	return ""
}

func (x *PeopleData) GetAge() string {
	if x != nil {
		return x.Age
	}
	return ""
}

func (x *PeopleData) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *PeopleData) GetHairColor() string {
	if x != nil {
		return x.HairColor
	}
	return ""
}

func (x *PeopleData) GetEyeColor() string {
	if x != nil {
		return x.EyeColor
	}
	return ""
}

func (x *PeopleData) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type FilmData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title             string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	ExternalReference string `protobuf:"bytes,2,opt,name=external_reference,json=externalReference,proto3" json:"external_reference,omitempty"`
	ReleaseYear       uint64 `protobuf:"varint,3,opt,name=release_year,json=releaseYear,proto3" json:"release_year,omitempty"`
	DirectorName      string `protobuf:"bytes,4,opt,name=director_name,json=directorName,proto3" json:"director_name,omitempty"`
	ProducerName      string `protobuf:"bytes,5,opt,name=producer_name,json=producerName,proto3" json:"producer_name,omitempty"`
	Rating            uint64 `protobuf:"varint,6,opt,name=rating,proto3" json:"rating,omitempty"`
	Id                uint64 `protobuf:"varint,7,opt,name=id,proto3" json:"id,omitempty"`
	Slug              string `protobuf:"bytes,8,opt,name=slug,proto3" json:"slug,omitempty"`
	Description       string `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *FilmData) Reset() {
	*x = FilmData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_films_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilmData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilmData) ProtoMessage() {}

func (x *FilmData) ProtoReflect() protoreflect.Message {
	mi := &file_films_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilmData.ProtoReflect.Descriptor instead.
func (*FilmData) Descriptor() ([]byte, []int) {
	return file_films_proto_rawDescGZIP(), []int{1}
}

func (x *FilmData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *FilmData) GetExternalReference() string {
	if x != nil {
		return x.ExternalReference
	}
	return ""
}

func (x *FilmData) GetReleaseYear() uint64 {
	if x != nil {
		return x.ReleaseYear
	}
	return 0
}

func (x *FilmData) GetDirectorName() string {
	if x != nil {
		return x.DirectorName
	}
	return ""
}

func (x *FilmData) GetProducerName() string {
	if x != nil {
		return x.ProducerName
	}
	return ""
}

func (x *FilmData) GetRating() uint64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *FilmData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FilmData) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *FilmData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type FilmDataWithPeople struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Film   *FilmData              `protobuf:"bytes,1,opt,name=film,proto3" json:"film,omitempty"`
	People map[uint64]*PeopleData `protobuf:"bytes,2,rep,name=people,proto3" json:"people,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FilmDataWithPeople) Reset() {
	*x = FilmDataWithPeople{}
	if protoimpl.UnsafeEnabled {
		mi := &file_films_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilmDataWithPeople) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilmDataWithPeople) ProtoMessage() {}

func (x *FilmDataWithPeople) ProtoReflect() protoreflect.Message {
	mi := &file_films_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilmDataWithPeople.ProtoReflect.Descriptor instead.
func (*FilmDataWithPeople) Descriptor() ([]byte, []int) {
	return file_films_proto_rawDescGZIP(), []int{2}
}

func (x *FilmDataWithPeople) GetFilm() *FilmData {
	if x != nil {
		return x.Film
	}
	return nil
}

func (x *FilmDataWithPeople) GetPeople() map[uint64]*PeopleData {
	if x != nil {
		return x.People
	}
	return nil
}

type RetrieveFilmsWithPeopleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *RetrieveFilmsWithPeopleRequest) Reset() {
	*x = RetrieveFilmsWithPeopleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_films_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveFilmsWithPeopleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveFilmsWithPeopleRequest) ProtoMessage() {}

func (x *RetrieveFilmsWithPeopleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_films_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveFilmsWithPeopleRequest.ProtoReflect.Descriptor instead.
func (*RetrieveFilmsWithPeopleRequest) Descriptor() ([]byte, []int) {
	return file_films_proto_rawDescGZIP(), []int{3}
}

func (x *RetrieveFilmsWithPeopleRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type RetrieveFilmsWithPeopleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Films map[uint64]*FilmDataWithPeople `protobuf:"bytes,1,rep,name=films,proto3" json:"films,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RetrieveFilmsWithPeopleResponse) Reset() {
	*x = RetrieveFilmsWithPeopleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_films_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveFilmsWithPeopleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveFilmsWithPeopleResponse) ProtoMessage() {}

func (x *RetrieveFilmsWithPeopleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_films_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveFilmsWithPeopleResponse.ProtoReflect.Descriptor instead.
func (*RetrieveFilmsWithPeopleResponse) Descriptor() ([]byte, []int) {
	return file_films_proto_rawDescGZIP(), []int{4}
}

func (x *RetrieveFilmsWithPeopleResponse) GetFilms() map[uint64]*FilmDataWithPeople {
	if x != nil {
		return x.Films
	}
	return nil
}

type RetrieveFilmsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Films map[uint64]*FilmData `protobuf:"bytes,1,rep,name=films,proto3" json:"films,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RetrieveFilmsResponse) Reset() {
	*x = RetrieveFilmsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_films_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveFilmsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveFilmsResponse) ProtoMessage() {}

func (x *RetrieveFilmsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_films_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveFilmsResponse.ProtoReflect.Descriptor instead.
func (*RetrieveFilmsResponse) Descriptor() ([]byte, []int) {
	return file_films_proto_rawDescGZIP(), []int{5}
}

func (x *RetrieveFilmsResponse) GetFilms() map[uint64]*FilmData {
	if x != nil {
		return x.Films
	}
	return nil
}

type RetrieveFilmsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *RetrieveFilmsRequest) Reset() {
	*x = RetrieveFilmsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_films_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveFilmsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveFilmsRequest) ProtoMessage() {}

func (x *RetrieveFilmsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_films_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveFilmsRequest.ProtoReflect.Descriptor instead.
func (*RetrieveFilmsRequest) Descriptor() ([]byte, []int) {
	return file_films_proto_rawDescGZIP(), []int{6}
}

func (x *RetrieveFilmsRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type RetrieveFilmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RetrieveFilmRequest) Reset() {
	*x = RetrieveFilmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_films_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveFilmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveFilmRequest) ProtoMessage() {}

func (x *RetrieveFilmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_films_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveFilmRequest.ProtoReflect.Descriptor instead.
func (*RetrieveFilmRequest) Descriptor() ([]byte, []int) {
	return file_films_proto_rawDescGZIP(), []int{7}
}

func (x *RetrieveFilmRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RetrieveFilmResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *FilmData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RetrieveFilmResponse) Reset() {
	*x = RetrieveFilmResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_films_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveFilmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveFilmResponse) ProtoMessage() {}

func (x *RetrieveFilmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_films_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveFilmResponse.ProtoReflect.Descriptor instead.
func (*RetrieveFilmResponse) Descriptor() ([]byte, []int) {
	return file_films_proto_rawDescGZIP(), []int{8}
}

func (x *RetrieveFilmResponse) GetData() *FilmData {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateFilmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *FilmData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateFilmRequest) Reset() {
	*x = CreateFilmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_films_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFilmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFilmRequest) ProtoMessage() {}

func (x *CreateFilmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_films_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFilmRequest.ProtoReflect.Descriptor instead.
func (*CreateFilmRequest) Descriptor() ([]byte, []int) {
	return file_films_proto_rawDescGZIP(), []int{9}
}

func (x *CreateFilmRequest) GetData() *FilmData {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreatePeopleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *PeopleData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreatePeopleRequest) Reset() {
	*x = CreatePeopleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_films_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePeopleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePeopleRequest) ProtoMessage() {}

func (x *CreatePeopleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_films_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePeopleRequest.ProtoReflect.Descriptor instead.
func (*CreatePeopleRequest) Descriptor() ([]byte, []int) {
	return file_films_proto_rawDescGZIP(), []int{10}
}

func (x *CreatePeopleRequest) GetData() *PeopleData {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateJoinPeopleFilmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilmExternalReference   string `protobuf:"bytes,1,opt,name=film_external_reference,json=filmExternalReference,proto3" json:"film_external_reference,omitempty"`
	PeopleExternalReference string `protobuf:"bytes,2,opt,name=people_external_reference,json=peopleExternalReference,proto3" json:"people_external_reference,omitempty"`
}

func (x *CreateJoinPeopleFilmRequest) Reset() {
	*x = CreateJoinPeopleFilmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_films_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJoinPeopleFilmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJoinPeopleFilmRequest) ProtoMessage() {}

func (x *CreateJoinPeopleFilmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_films_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJoinPeopleFilmRequest.ProtoReflect.Descriptor instead.
func (*CreateJoinPeopleFilmRequest) Descriptor() ([]byte, []int) {
	return file_films_proto_rawDescGZIP(), []int{11}
}

func (x *CreateJoinPeopleFilmRequest) GetFilmExternalReference() string {
	if x != nil {
		return x.FilmExternalReference
	}
	return ""
}

func (x *CreateJoinPeopleFilmRequest) GetPeopleExternalReference() string {
	if x != nil {
		return x.PeopleExternalReference
	}
	return ""
}

var File_films_proto protoreflect.FileDescriptor

var file_films_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x66,
	0x69, 0x6c, 0x6d, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc9, 0x01, 0x0a, 0x0a, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x68, 0x61, 0x69, 0x72, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x68, 0x61, 0x69, 0x72, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x09,
	0x65, 0x79, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x65, 0x79, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75,
	0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x22, 0x9a, 0x02,
	0x0a, 0x08, 0x46, 0x69, 0x6c, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x2d, 0x0a, 0x12, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x59, 0x65,
	0x61, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc6, 0x01, 0x0a, 0x12, 0x46,
	0x69, 0x6c, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x57, 0x69, 0x74, 0x68, 0x50, 0x65, 0x6f, 0x70, 0x6c,
	0x65, 0x12, 0x23, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x66, 0x69, 0x6c, 0x6d, 0x12, 0x3d, 0x0a, 0x06, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x46,
	0x69, 0x6c, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x57, 0x69, 0x74, 0x68, 0x50, 0x65, 0x6f, 0x70, 0x6c,
	0x65, 0x2e, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70,
	0x65, 0x6f, 0x70, 0x6c, 0x65, 0x1a, 0x4c, 0x0a, 0x0b, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x50, 0x65,
	0x6f, 0x70, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x36, 0x0a, 0x1e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x46,
	0x69, 0x6c, 0x6d, 0x73, 0x57, 0x69, 0x74, 0x68, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xbf, 0x01, 0x0a, 0x1f,
	0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x57, 0x69, 0x74,
	0x68, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x47, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x46,
	0x69, 0x6c, 0x6d, 0x73, 0x57, 0x69, 0x74, 0x68, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x1a, 0x53, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x6d,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e,
	0x46, 0x69, 0x6c, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x57, 0x69, 0x74, 0x68, 0x50, 0x65, 0x6f, 0x70,
	0x6c, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa1, 0x01,
	0x0a, 0x15, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x52,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x05, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x1a, 0x49, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x46, 0x69,
	0x6c, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x2c, 0x0a, 0x14, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x46, 0x69, 0x6c,
	0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0x25, 0x0a, 0x13, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x14, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x46, 0x69, 0x6c, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x66,
	0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x38, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x46,
	0x69, 0x6c, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3c, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x50, 0x65, 0x6f, 0x70, 0x6c,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x91, 0x01, 0x0a, 0x1b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65,
	0x46, 0x69, 0x6c, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x17, 0x66,
	0x69, 0x6c, 0x6d, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x66, 0x69,
	0x6c, 0x6d, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x19, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x5f, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x45, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x32,
	0xf1, 0x03, 0x0a, 0x05, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x12, 0x49, 0x0a, 0x0c, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x6d, 0x12, 0x1a, 0x2e, 0x66, 0x69, 0x6c, 0x6d,
	0x73, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x52, 0x65,
	0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0d, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x46, 0x69, 0x6c, 0x6d, 0x73, 0x12, 0x1b, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x52, 0x65,
	0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x6a, 0x0a, 0x17, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x46, 0x69,
	0x6c, 0x6d, 0x73, 0x57, 0x69, 0x74, 0x68, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x12, 0x25, 0x2e,
	0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x46, 0x69,
	0x6c, 0x6d, 0x73, 0x57, 0x69, 0x74, 0x68, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x57, 0x69, 0x74, 0x68, 0x50, 0x65,
	0x6f, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x12, 0x18, 0x2e,
	0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x28, 0x01, 0x12, 0x46, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x6f,
	0x70, 0x6c, 0x65, 0x12, 0x1a, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12, 0x56, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x46,
	0x69, 0x6c, 0x6d, 0x12, 0x22, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x28, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x70, 0x62, 0x66, 0x69, 0x6c, 0x6d, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_films_proto_rawDescOnce sync.Once
	file_films_proto_rawDescData = file_films_proto_rawDesc
)

func file_films_proto_rawDescGZIP() []byte {
	file_films_proto_rawDescOnce.Do(func() {
		file_films_proto_rawDescData = protoimpl.X.CompressGZIP(file_films_proto_rawDescData)
	})
	return file_films_proto_rawDescData
}

var file_films_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_films_proto_goTypes = []interface{}{
	(*PeopleData)(nil),                      // 0: films.PeopleData
	(*FilmData)(nil),                        // 1: films.FilmData
	(*FilmDataWithPeople)(nil),              // 2: films.FilmDataWithPeople
	(*RetrieveFilmsWithPeopleRequest)(nil),  // 3: films.RetrieveFilmsWithPeopleRequest
	(*RetrieveFilmsWithPeopleResponse)(nil), // 4: films.RetrieveFilmsWithPeopleResponse
	(*RetrieveFilmsResponse)(nil),           // 5: films.RetrieveFilmsResponse
	(*RetrieveFilmsRequest)(nil),            // 6: films.RetrieveFilmsRequest
	(*RetrieveFilmRequest)(nil),             // 7: films.RetrieveFilmRequest
	(*RetrieveFilmResponse)(nil),            // 8: films.RetrieveFilmResponse
	(*CreateFilmRequest)(nil),               // 9: films.CreateFilmRequest
	(*CreatePeopleRequest)(nil),             // 10: films.CreatePeopleRequest
	(*CreateJoinPeopleFilmRequest)(nil),     // 11: films.CreateJoinPeopleFilmRequest
	nil,                                     // 12: films.FilmDataWithPeople.PeopleEntry
	nil,                                     // 13: films.RetrieveFilmsWithPeopleResponse.FilmsEntry
	nil,                                     // 14: films.RetrieveFilmsResponse.FilmsEntry
	(*empty.Empty)(nil),                     // 15: google.protobuf.Empty
}
var file_films_proto_depIdxs = []int32{
	1,  // 0: films.FilmDataWithPeople.film:type_name -> films.FilmData
	12, // 1: films.FilmDataWithPeople.people:type_name -> films.FilmDataWithPeople.PeopleEntry
	13, // 2: films.RetrieveFilmsWithPeopleResponse.films:type_name -> films.RetrieveFilmsWithPeopleResponse.FilmsEntry
	14, // 3: films.RetrieveFilmsResponse.films:type_name -> films.RetrieveFilmsResponse.FilmsEntry
	1,  // 4: films.RetrieveFilmResponse.data:type_name -> films.FilmData
	1,  // 5: films.CreateFilmRequest.data:type_name -> films.FilmData
	0,  // 6: films.CreatePeopleRequest.data:type_name -> films.PeopleData
	0,  // 7: films.FilmDataWithPeople.PeopleEntry.value:type_name -> films.PeopleData
	2,  // 8: films.RetrieveFilmsWithPeopleResponse.FilmsEntry.value:type_name -> films.FilmDataWithPeople
	1,  // 9: films.RetrieveFilmsResponse.FilmsEntry.value:type_name -> films.FilmData
	7,  // 10: films.Films.RetrieveFilm:input_type -> films.RetrieveFilmRequest
	6,  // 11: films.Films.RetrieveFilms:input_type -> films.RetrieveFilmsRequest
	3,  // 12: films.Films.RetrieveFilmsWithPeople:input_type -> films.RetrieveFilmsWithPeopleRequest
	9,  // 13: films.Films.CreateFilms:input_type -> films.CreateFilmRequest
	10, // 14: films.Films.CreatePeople:input_type -> films.CreatePeopleRequest
	11, // 15: films.Films.CreateJoinPeopleFilm:input_type -> films.CreateJoinPeopleFilmRequest
	8,  // 16: films.Films.RetrieveFilm:output_type -> films.RetrieveFilmResponse
	5,  // 17: films.Films.RetrieveFilms:output_type -> films.RetrieveFilmsResponse
	4,  // 18: films.Films.RetrieveFilmsWithPeople:output_type -> films.RetrieveFilmsWithPeopleResponse
	15, // 19: films.Films.CreateFilms:output_type -> google.protobuf.Empty
	15, // 20: films.Films.CreatePeople:output_type -> google.protobuf.Empty
	15, // 21: films.Films.CreateJoinPeopleFilm:output_type -> google.protobuf.Empty
	16, // [16:22] is the sub-list for method output_type
	10, // [10:16] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_films_proto_init() }
func file_films_proto_init() {
	if File_films_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_films_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeopleData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_films_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilmData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_films_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilmDataWithPeople); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_films_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveFilmsWithPeopleRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_films_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveFilmsWithPeopleResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_films_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveFilmsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_films_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveFilmsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_films_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveFilmRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_films_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveFilmResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_films_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFilmRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_films_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePeopleRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_films_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJoinPeopleFilmRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_films_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_films_proto_goTypes,
		DependencyIndexes: file_films_proto_depIdxs,
		MessageInfos:      file_films_proto_msgTypes,
	}.Build()
	File_films_proto = out.File
	file_films_proto_rawDesc = nil
	file_films_proto_goTypes = nil
	file_films_proto_depIdxs = nil
}
