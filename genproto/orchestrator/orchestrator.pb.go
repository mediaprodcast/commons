// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: orchestrator.proto

package orchestrator

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

type JobStatus int32

const (
	JobStatus_JOB_STATUS_UNSPECIFIED JobStatus = 0
	JobStatus_QUEUED                 JobStatus = 1
	JobStatus_PROCESSING             JobStatus = 2
	JobStatus_COMPLETED              JobStatus = 3
	JobStatus_FAILED                 JobStatus = 4
	JobStatus_CANCELLED              JobStatus = 5
)

// Enum value maps for JobStatus.
var (
	JobStatus_name = map[int32]string{
		0: "JOB_STATUS_UNSPECIFIED",
		1: "QUEUED",
		2: "PROCESSING",
		3: "COMPLETED",
		4: "FAILED",
		5: "CANCELLED",
	}
	JobStatus_value = map[string]int32{
		"JOB_STATUS_UNSPECIFIED": 0,
		"QUEUED":                 1,
		"PROCESSING":             2,
		"COMPLETED":              3,
		"FAILED":                 4,
		"CANCELLED":              5,
	}
)

func (x JobStatus) Enum() *JobStatus {
	p := new(JobStatus)
	*p = x
	return p
}

func (x JobStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_orchestrator_proto_enumTypes[0].Descriptor()
}

func (JobStatus) Type() protoreflect.EnumType {
	return &file_orchestrator_proto_enumTypes[0]
}

func (x JobStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobStatus.Descriptor instead.
func (JobStatus) EnumDescriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{0}
}

type CreateJobRequest struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	OrgId               string                 `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	TranscodingProfiles []*TranscodingProfile  `protobuf:"bytes,2,rep,name=transcoding_profiles,json=transcodingProfiles,proto3" json:"transcoding_profiles,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *CreateJobRequest) Reset() {
	*x = CreateJobRequest{}
	mi := &file_orchestrator_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobRequest) ProtoMessage() {}

func (x *CreateJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobRequest.ProtoReflect.Descriptor instead.
func (*CreateJobRequest) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{0}
}

func (x *CreateJobRequest) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *CreateJobRequest) GetTranscodingProfiles() []*TranscodingProfile {
	if x != nil {
		return x.TranscodingProfiles
	}
	return nil
}

type CreateJobResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobId         string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Status        JobStatus              `protobuf:"varint,2,opt,name=status,proto3,enum=orchestrator.JobStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateJobResponse) Reset() {
	*x = CreateJobResponse{}
	mi := &file_orchestrator_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobResponse) ProtoMessage() {}

func (x *CreateJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobResponse.ProtoReflect.Descriptor instead.
func (*CreateJobResponse) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{1}
}

func (x *CreateJobResponse) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *CreateJobResponse) GetStatus() JobStatus {
	if x != nil {
		return x.Status
	}
	return JobStatus_JOB_STATUS_UNSPECIFIED
}

type UpdateJobRequest struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	JobId               string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Status              JobStatus              `protobuf:"varint,2,opt,name=status,proto3,enum=orchestrator.JobStatus" json:"status,omitempty"`
	TranscodingProfiles []*TranscodingProfile  `protobuf:"bytes,3,rep,name=transcoding_profiles,json=transcodingProfiles,proto3" json:"transcoding_profiles,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *UpdateJobRequest) Reset() {
	*x = UpdateJobRequest{}
	mi := &file_orchestrator_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobRequest) ProtoMessage() {}

func (x *UpdateJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobRequest.ProtoReflect.Descriptor instead.
func (*UpdateJobRequest) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateJobRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *UpdateJobRequest) GetStatus() JobStatus {
	if x != nil {
		return x.Status
	}
	return JobStatus_JOB_STATUS_UNSPECIFIED
}

func (x *UpdateJobRequest) GetTranscodingProfiles() []*TranscodingProfile {
	if x != nil {
		return x.TranscodingProfiles
	}
	return nil
}

type UpdateJobResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Job           *Job                   `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateJobResponse) Reset() {
	*x = UpdateJobResponse{}
	mi := &file_orchestrator_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobResponse) ProtoMessage() {}

func (x *UpdateJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobResponse.ProtoReflect.Descriptor instead.
func (*UpdateJobResponse) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateJobResponse) GetJob() *Job {
	if x != nil {
		return x.Job
	}
	return nil
}

type DestroyJobRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobId         string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DestroyJobRequest) Reset() {
	*x = DestroyJobRequest{}
	mi := &file_orchestrator_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DestroyJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestroyJobRequest) ProtoMessage() {}

func (x *DestroyJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestroyJobRequest.ProtoReflect.Descriptor instead.
func (*DestroyJobRequest) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{4}
}

func (x *DestroyJobRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type DestroyJobResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobId         string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Success       bool                   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DestroyJobResponse) Reset() {
	*x = DestroyJobResponse{}
	mi := &file_orchestrator_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DestroyJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestroyJobResponse) ProtoMessage() {}

func (x *DestroyJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestroyJobResponse.ProtoReflect.Descriptor instead.
func (*DestroyJobResponse) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{5}
}

func (x *DestroyJobResponse) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *DestroyJobResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type TranscodingProfile struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Bitrate       int32                  `protobuf:"varint,1,opt,name=bitrate,proto3" json:"bitrate,omitempty"`
	FrameRate     int32                  `protobuf:"varint,2,opt,name=frame_rate,json=frameRate,proto3" json:"frame_rate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TranscodingProfile) Reset() {
	*x = TranscodingProfile{}
	mi := &file_orchestrator_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TranscodingProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranscodingProfile) ProtoMessage() {}

func (x *TranscodingProfile) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranscodingProfile.ProtoReflect.Descriptor instead.
func (*TranscodingProfile) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{6}
}

func (x *TranscodingProfile) GetBitrate() int32 {
	if x != nil {
		return x.Bitrate
	}
	return 0
}

func (x *TranscodingProfile) GetFrameRate() int32 {
	if x != nil {
		return x.FrameRate
	}
	return 0
}

type ListJobsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrgId         string                 `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	PageSize      int32                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken     int32                  `protobuf:"varint,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListJobsRequest) Reset() {
	*x = ListJobsRequest{}
	mi := &file_orchestrator_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListJobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsRequest) ProtoMessage() {}

func (x *ListJobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobsRequest.ProtoReflect.Descriptor instead.
func (*ListJobsRequest) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{7}
}

func (x *ListJobsRequest) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *ListJobsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListJobsRequest) GetPageToken() int32 {
	if x != nil {
		return x.PageToken
	}
	return 0
}

type ListJobsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Jobs          []*Job                 `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListJobsResponse) Reset() {
	*x = ListJobsResponse{}
	mi := &file_orchestrator_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListJobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsResponse) ProtoMessage() {}

func (x *ListJobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobsResponse.ProtoReflect.Descriptor instead.
func (*ListJobsResponse) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{8}
}

func (x *ListJobsResponse) GetJobs() []*Job {
	if x != nil {
		return x.Jobs
	}
	return nil
}

func (x *ListJobsResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type Job struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	JobId         string                 `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Status        JobStatus              `protobuf:"varint,3,opt,name=status,proto3,enum=orchestrator.JobStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Job) Reset() {
	*x = Job{}
	mi := &file_orchestrator_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{9}
}

func (x *Job) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Job) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *Job) GetStatus() JobStatus {
	if x != nil {
		return x.Status
	}
	return JobStatus_JOB_STATUS_UNSPECIFIED
}

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_orchestrator_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{10}
}

type GetJobRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrgId         string                 `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	JobId         string                 `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetJobRequest) Reset() {
	*x = GetJobRequest{}
	mi := &file_orchestrator_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobRequest) ProtoMessage() {}

func (x *GetJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobRequest.ProtoReflect.Descriptor instead.
func (*GetJobRequest) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{11}
}

func (x *GetJobRequest) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *GetJobRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

var File_orchestrator_proto protoreflect.FileDescriptor

var file_orchestrator_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x22, 0x7e, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x53, 0x0a,
	0x14, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6f, 0x72,
	0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x13, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x22, 0x5b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x2f,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17,
	0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x4a, 0x6f,
	0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0xaf, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6f, 0x72,
	0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x53, 0x0a, 0x14,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6f, 0x72, 0x63,
	0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63,
	0x6f, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x13, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x22, 0x38, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x22, 0x2a, 0x0a, 0x11, 0x44,
	0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x74, 0x72,
	0x6f, 0x79, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a,
	0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a,
	0x6f, 0x62, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x4d,
	0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x61, 0x74, 0x65, 0x22, 0x64, 0x0a,
	0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x5a, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x5d, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x2f, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e,
	0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x4a, 0x6f, 0x62,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3d, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12,
	0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x2a, 0x6d, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x16, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x51, 0x55, 0x45, 0x55, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x50,
	0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43,
	0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c,
	0x4c, 0x45, 0x44, 0x10, 0x05, 0x32, 0xc6, 0x03, 0x0a, 0x13, 0x4f, 0x72, 0x63, 0x68, 0x65, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x1e, 0x2e, 0x6f, 0x72, 0x63,
	0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6f, 0x72, 0x63,
	0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x1e, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0a, 0x44, 0x65, 0x73,
	0x74, 0x72, 0x6f, 0x79, 0x4a, 0x6f, 0x62, 0x12, 0x1f, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x47, 0x65,
	0x74, 0x4a, 0x6f, 0x62, 0x12, 0x1b, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x4a, 0x6f, 0x62, 0x12, 0x3d, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4a, 0x6f,
	0x62, 0x12, 0x1b, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x49, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x12,
	0x1d, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x38,
	0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x70, 0x72, 0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x63, 0x68,
	0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orchestrator_proto_rawDescOnce sync.Once
	file_orchestrator_proto_rawDescData = file_orchestrator_proto_rawDesc
)

func file_orchestrator_proto_rawDescGZIP() []byte {
	file_orchestrator_proto_rawDescOnce.Do(func() {
		file_orchestrator_proto_rawDescData = protoimpl.X.CompressGZIP(file_orchestrator_proto_rawDescData)
	})
	return file_orchestrator_proto_rawDescData
}

var file_orchestrator_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_orchestrator_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_orchestrator_proto_goTypes = []any{
	(JobStatus)(0),             // 0: orchestrator.JobStatus
	(*CreateJobRequest)(nil),   // 1: orchestrator.CreateJobRequest
	(*CreateJobResponse)(nil),  // 2: orchestrator.CreateJobResponse
	(*UpdateJobRequest)(nil),   // 3: orchestrator.UpdateJobRequest
	(*UpdateJobResponse)(nil),  // 4: orchestrator.UpdateJobResponse
	(*DestroyJobRequest)(nil),  // 5: orchestrator.DestroyJobRequest
	(*DestroyJobResponse)(nil), // 6: orchestrator.DestroyJobResponse
	(*TranscodingProfile)(nil), // 7: orchestrator.TranscodingProfile
	(*ListJobsRequest)(nil),    // 8: orchestrator.ListJobsRequest
	(*ListJobsResponse)(nil),   // 9: orchestrator.ListJobsResponse
	(*Job)(nil),                // 10: orchestrator.Job
	(*Empty)(nil),              // 11: orchestrator.Empty
	(*GetJobRequest)(nil),      // 12: orchestrator.GetJobRequest
}
var file_orchestrator_proto_depIdxs = []int32{
	7,  // 0: orchestrator.CreateJobRequest.transcoding_profiles:type_name -> orchestrator.TranscodingProfile
	0,  // 1: orchestrator.CreateJobResponse.status:type_name -> orchestrator.JobStatus
	0,  // 2: orchestrator.UpdateJobRequest.status:type_name -> orchestrator.JobStatus
	7,  // 3: orchestrator.UpdateJobRequest.transcoding_profiles:type_name -> orchestrator.TranscodingProfile
	10, // 4: orchestrator.UpdateJobResponse.job:type_name -> orchestrator.Job
	10, // 5: orchestrator.ListJobsResponse.jobs:type_name -> orchestrator.Job
	0,  // 6: orchestrator.Job.status:type_name -> orchestrator.JobStatus
	1,  // 7: orchestrator.OrchestratorService.CreateJob:input_type -> orchestrator.CreateJobRequest
	3,  // 8: orchestrator.OrchestratorService.UpdateJob:input_type -> orchestrator.UpdateJobRequest
	5,  // 9: orchestrator.OrchestratorService.DestroyJob:input_type -> orchestrator.DestroyJobRequest
	12, // 10: orchestrator.OrchestratorService.GetJob:input_type -> orchestrator.GetJobRequest
	12, // 11: orchestrator.OrchestratorService.CancelJob:input_type -> orchestrator.GetJobRequest
	8,  // 12: orchestrator.OrchestratorService.ListJobs:input_type -> orchestrator.ListJobsRequest
	2,  // 13: orchestrator.OrchestratorService.CreateJob:output_type -> orchestrator.CreateJobResponse
	4,  // 14: orchestrator.OrchestratorService.UpdateJob:output_type -> orchestrator.UpdateJobResponse
	6,  // 15: orchestrator.OrchestratorService.DestroyJob:output_type -> orchestrator.DestroyJobResponse
	10, // 16: orchestrator.OrchestratorService.GetJob:output_type -> orchestrator.Job
	11, // 17: orchestrator.OrchestratorService.CancelJob:output_type -> orchestrator.Empty
	9,  // 18: orchestrator.OrchestratorService.ListJobs:output_type -> orchestrator.ListJobsResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_orchestrator_proto_init() }
func file_orchestrator_proto_init() {
	if File_orchestrator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orchestrator_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orchestrator_proto_goTypes,
		DependencyIndexes: file_orchestrator_proto_depIdxs,
		EnumInfos:         file_orchestrator_proto_enumTypes,
		MessageInfos:      file_orchestrator_proto_msgTypes,
	}.Build()
	File_orchestrator_proto = out.File
	file_orchestrator_proto_rawDesc = nil
	file_orchestrator_proto_goTypes = nil
	file_orchestrator_proto_depIdxs = nil
}