// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: reviews/reviews.proto

package reviews

import (
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

type ProductId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProductId) Reset() {
	*x = ProductId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reviews_reviews_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductId) ProtoMessage() {}

func (x *ProductId) ProtoReflect() protoreflect.Message {
	mi := &file_reviews_reviews_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductId.ProtoReflect.Descriptor instead.
func (*ProductId) Descriptor() ([]byte, []int) {
	return file_reviews_reviews_proto_rawDescGZIP(), []int{0}
}

func (x *ProductId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Count struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count string `protobuf:"bytes,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Count) Reset() {
	*x = Count{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reviews_reviews_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Count) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Count) ProtoMessage() {}

func (x *Count) ProtoReflect() protoreflect.Message {
	mi := &file_reviews_reviews_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Count.ProtoReflect.Descriptor instead.
func (*Count) Descriptor() ([]byte, []int) {
	return file_reviews_reviews_proto_rawDescGZIP(), []int{1}
}

func (x *Count) GetCount() string {
	if x != nil {
		return x.Count
	}
	return ""
}

type ReviewId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReviewId) Reset() {
	*x = ReviewId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reviews_reviews_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviewId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewId) ProtoMessage() {}

func (x *ReviewId) ProtoReflect() protoreflect.Message {
	mi := &file_reviews_reviews_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewId.ProtoReflect.Descriptor instead.
func (*ReviewId) Descriptor() ([]byte, []int) {
	return file_reviews_reviews_proto_rawDescGZIP(), []int{2}
}

func (x *ReviewId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Review struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Msg       string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	UserId    string `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
	ProductId string `protobuf:"bytes,4,opt,name=productId,proto3" json:"productId,omitempty"`
	Time      string `protobuf:"bytes,5,opt,name=time,proto3" json:"time,omitempty"`
	Date      string `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *Review) Reset() {
	*x = Review{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reviews_reviews_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Review) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Review) ProtoMessage() {}

func (x *Review) ProtoReflect() protoreflect.Message {
	mi := &file_reviews_reviews_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Review.ProtoReflect.Descriptor instead.
func (*Review) Descriptor() ([]byte, []int) {
	return file_reviews_reviews_proto_rawDescGZIP(), []int{3}
}

func (x *Review) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Review) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Review) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Review) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Review) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *Review) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type ReviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReviewEntry *Review `protobuf:"bytes,1,opt,name=reviewEntry,proto3" json:"reviewEntry,omitempty"`
}

func (x *ReviewRequest) Reset() {
	*x = ReviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reviews_reviews_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewRequest) ProtoMessage() {}

func (x *ReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reviews_reviews_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewRequest.ProtoReflect.Descriptor instead.
func (*ReviewRequest) Descriptor() ([]byte, []int) {
	return file_reviews_reviews_proto_rawDescGZIP(), []int{4}
}

func (x *ReviewRequest) GetReviewEntry() *Review {
	if x != nil {
		return x.ReviewEntry
	}
	return nil
}

type ReviewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *Review `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *ReviewResponse) Reset() {
	*x = ReviewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reviews_reviews_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewResponse) ProtoMessage() {}

func (x *ReviewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reviews_reviews_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewResponse.ProtoReflect.Descriptor instead.
func (*ReviewResponse) Descriptor() ([]byte, []int) {
	return file_reviews_reviews_proto_rawDescGZIP(), []int{5}
}

func (x *ReviewResponse) GetResponse() *Review {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_reviews_reviews_proto protoreflect.FileDescriptor

var file_reviews_reviews_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73,
	0x22, 0x1b, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1d, 0x0a,
	0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x1a, 0x0a, 0x08,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x88, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x22, 0x42, 0x0a, 0x0d, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0b, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x0b, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x3d, 0x0a, 0x0e, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xea, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x12, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x30, 0x01, 0x12,
	0x2f, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12,
	0x0f, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x1a, 0x0e, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x31, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x12, 0x11, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reviews_reviews_proto_rawDescOnce sync.Once
	file_reviews_reviews_proto_rawDescData = file_reviews_reviews_proto_rawDesc
)

func file_reviews_reviews_proto_rawDescGZIP() []byte {
	file_reviews_reviews_proto_rawDescOnce.Do(func() {
		file_reviews_reviews_proto_rawDescData = protoimpl.X.CompressGZIP(file_reviews_reviews_proto_rawDescData)
	})
	return file_reviews_reviews_proto_rawDescData
}

var file_reviews_reviews_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_reviews_reviews_proto_goTypes = []interface{}{
	(*ProductId)(nil),      // 0: reviews.ProductId
	(*Count)(nil),          // 1: reviews.Count
	(*ReviewId)(nil),       // 2: reviews.ReviewId
	(*Review)(nil),         // 3: reviews.Review
	(*ReviewRequest)(nil),  // 4: reviews.ReviewRequest
	(*ReviewResponse)(nil), // 5: reviews.ReviewResponse
}
var file_reviews_reviews_proto_depIdxs = []int32{
	3, // 0: reviews.ReviewRequest.reviewEntry:type_name -> reviews.Review
	3, // 1: reviews.ReviewResponse.response:type_name -> reviews.Review
	4, // 2: reviews.ReviewsService.CreateReview:input_type -> reviews.ReviewRequest
	0, // 3: reviews.ReviewsService.GetReviews:input_type -> reviews.ProductId
	3, // 4: reviews.ReviewsService.UpdateReview:input_type -> reviews.Review
	2, // 5: reviews.ReviewsService.DeleteReview:input_type -> reviews.ReviewId
	5, // 6: reviews.ReviewsService.CreateReview:output_type -> reviews.ReviewResponse
	3, // 7: reviews.ReviewsService.GetReviews:output_type -> reviews.Review
	1, // 8: reviews.ReviewsService.UpdateReview:output_type -> reviews.Count
	1, // 9: reviews.ReviewsService.DeleteReview:output_type -> reviews.Count
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_reviews_reviews_proto_init() }
func file_reviews_reviews_proto_init() {
	if File_reviews_reviews_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reviews_reviews_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductId); i {
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
		file_reviews_reviews_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Count); i {
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
		file_reviews_reviews_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviewId); i {
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
		file_reviews_reviews_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Review); i {
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
		file_reviews_reviews_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviewRequest); i {
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
		file_reviews_reviews_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviewResponse); i {
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
			RawDescriptor: file_reviews_reviews_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reviews_reviews_proto_goTypes,
		DependencyIndexes: file_reviews_reviews_proto_depIdxs,
		MessageInfos:      file_reviews_reviews_proto_msgTypes,
	}.Build()
	File_reviews_reviews_proto = out.File
	file_reviews_reviews_proto_rawDesc = nil
	file_reviews_reviews_proto_goTypes = nil
	file_reviews_reviews_proto_depIdxs = nil
}