// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/api/serviceconfig/service.proto
// DO NOT EDIT!

package google_api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/label"
import google_api2 "google.golang.org/genproto/googleapis/api/metric"
import google_api3 "google.golang.org/genproto/googleapis/api/monitoredres"
import _ "github.com/golang/protobuf/ptypes/any"
import google_protobuf4 "google.golang.org/genproto/protobuf"
import google_protobuf3 "google.golang.org/genproto/protobuf"
import google_protobuf5 "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// `Service` is the root object of the configuration schema. It
// describes basic information like the name of the service and the
// exposed API interfaces, and delegates other aspects to configuration
// sub-sections.
//
// Example:
//
//     type: google.api.Service
//     config_version: 1
//     name: calendar.googleapis.com
//     title: Google Calendar API
//     apis:
//     - name: google.calendar.Calendar
//     backend:
//       rules:
//       - selector: "*"
//         address: calendar.example.com
type Service struct {
	// The version of the service configuration. The config version may
	// influence interpretation of the configuration, for example, to
	// determine defaults. This is documented together with applicable
	// options. The current default for the config version itself is `3`.
	ConfigVersion *google_protobuf5.UInt32Value `protobuf:"bytes,20,opt,name=config_version,json=configVersion" json:"config_version,omitempty"`
	// The DNS address at which this service is available,
	// e.g. `calendar.googleapis.com`.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// A unique ID for a specific instance of this message, typically assigned
	// by the client for tracking purpose. If empty, the server may choose to
	// generate one instead.
	Id string `protobuf:"bytes,33,opt,name=id" json:"id,omitempty"`
	// The product title associated with this service.
	Title string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	// The id of the Google developer project that owns the service.
	// Members of this project can manage the service configuration,
	// manage consumption of the service, etc.
	ProducerProjectId string `protobuf:"bytes,22,opt,name=producer_project_id,json=producerProjectId" json:"producer_project_id,omitempty"`
	// A list of API interfaces exported by this service. Only the `name` field
	// of the [google.protobuf.Api][google.protobuf.Api] needs to be provided by the configuration
	// author, as the remaining fields will be derived from the IDL during the
	// normalization process. It is an error to specify an API interface here
	// which cannot be resolved against the associated IDL files.
	Apis []*google_protobuf4.Api `protobuf:"bytes,3,rep,name=apis" json:"apis,omitempty"`
	// A list of all proto message types included in this API service.
	// Types referenced directly or indirectly by the `apis` are
	// automatically included.  Messages which are not referenced but
	// shall be included, such as types used by the `google.protobuf.Any` type,
	// should be listed here by name. Example:
	//
	//     types:
	//     - name: google.protobuf.Int32
	Types []*google_protobuf3.Type `protobuf:"bytes,4,rep,name=types" json:"types,omitempty"`
	// A list of all enum types included in this API service.  Enums
	// referenced directly or indirectly by the `apis` are automatically
	// included.  Enums which are not referenced but shall be included
	// should be listed here by name. Example:
	//
	//     enums:
	//     - name: google.someapi.v1.SomeEnum
	Enums []*google_protobuf3.Enum `protobuf:"bytes,5,rep,name=enums" json:"enums,omitempty"`
	// Additional API documentation.
	Documentation *Documentation `protobuf:"bytes,6,opt,name=documentation" json:"documentation,omitempty"`
	// API backend configuration.
	Backend *Backend `protobuf:"bytes,8,opt,name=backend" json:"backend,omitempty"`
	// HTTP configuration.
	Http *Http `protobuf:"bytes,9,opt,name=http" json:"http,omitempty"`
	// Auth configuration.
	Authentication *Authentication `protobuf:"bytes,11,opt,name=authentication" json:"authentication,omitempty"`
	// Context configuration.
	Context *Context `protobuf:"bytes,12,opt,name=context" json:"context,omitempty"`
	// Configuration controlling usage of this service.
	Usage *Usage `protobuf:"bytes,15,opt,name=usage" json:"usage,omitempty"`
	// Configuration for network endpoints.  If this is empty, then an endpoint
	// with the same name as the service is automatically generated to service all
	// defined APIs.
	Endpoints []*Endpoint `protobuf:"bytes,18,rep,name=endpoints" json:"endpoints,omitempty"`
	// Configuration for the service control plane.
	Control *Control `protobuf:"bytes,21,opt,name=control" json:"control,omitempty"`
	// Defines the logs used by this service.
	Logs []*LogDescriptor `protobuf:"bytes,23,rep,name=logs" json:"logs,omitempty"`
	// Defines the metrics used by this service.
	Metrics []*google_api2.MetricDescriptor `protobuf:"bytes,24,rep,name=metrics" json:"metrics,omitempty"`
	// Defines the monitored resources used by this service. This is required
	// by the [Service.monitoring][google.api.Service.monitoring] and [Service.logging][google.api.Service.logging] configurations.
	MonitoredResources []*google_api3.MonitoredResourceDescriptor `protobuf:"bytes,25,rep,name=monitored_resources,json=monitoredResources" json:"monitored_resources,omitempty"`
	// Logging configuration of the service.
	Logging *Logging `protobuf:"bytes,27,opt,name=logging" json:"logging,omitempty"`
	// Monitoring configuration of the service.
	Monitoring *Monitoring `protobuf:"bytes,28,opt,name=monitoring" json:"monitoring,omitempty"`
	// Configuration for system parameters.
	SystemParameters *SystemParameters `protobuf:"bytes,29,opt,name=system_parameters,json=systemParameters" json:"system_parameters,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *Service) GetConfigVersion() *google_protobuf5.UInt32Value {
	if m != nil {
		return m.ConfigVersion
	}
	return nil
}

func (m *Service) GetApis() []*google_protobuf4.Api {
	if m != nil {
		return m.Apis
	}
	return nil
}

func (m *Service) GetTypes() []*google_protobuf3.Type {
	if m != nil {
		return m.Types
	}
	return nil
}

func (m *Service) GetEnums() []*google_protobuf3.Enum {
	if m != nil {
		return m.Enums
	}
	return nil
}

func (m *Service) GetDocumentation() *Documentation {
	if m != nil {
		return m.Documentation
	}
	return nil
}

func (m *Service) GetBackend() *Backend {
	if m != nil {
		return m.Backend
	}
	return nil
}

func (m *Service) GetHttp() *Http {
	if m != nil {
		return m.Http
	}
	return nil
}

func (m *Service) GetAuthentication() *Authentication {
	if m != nil {
		return m.Authentication
	}
	return nil
}

func (m *Service) GetContext() *Context {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Service) GetUsage() *Usage {
	if m != nil {
		return m.Usage
	}
	return nil
}

func (m *Service) GetEndpoints() []*Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *Service) GetControl() *Control {
	if m != nil {
		return m.Control
	}
	return nil
}

func (m *Service) GetLogs() []*LogDescriptor {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *Service) GetMetrics() []*google_api2.MetricDescriptor {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *Service) GetMonitoredResources() []*google_api3.MonitoredResourceDescriptor {
	if m != nil {
		return m.MonitoredResources
	}
	return nil
}

func (m *Service) GetLogging() *Logging {
	if m != nil {
		return m.Logging
	}
	return nil
}

func (m *Service) GetMonitoring() *Monitoring {
	if m != nil {
		return m.Monitoring
	}
	return nil
}

func (m *Service) GetSystemParameters() *SystemParameters {
	if m != nil {
		return m.SystemParameters
	}
	return nil
}

func init() {
	proto.RegisterType((*Service)(nil), "google.api.Service")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/api/serviceconfig/service.proto", fileDescriptor13)
}

var fileDescriptor13 = []byte{
	// 778 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x95, 0xed, 0x6e, 0xfb, 0x34,
	0x14, 0xc6, 0xd5, 0xfe, 0xdb, 0xff, 0x36, 0xef, 0x85, 0xcd, 0x7b, 0xc1, 0x2b, 0x43, 0x1a, 0x08,
	0x69, 0x95, 0xd0, 0x12, 0xa9, 0x93, 0x26, 0x21, 0x84, 0x50, 0xbb, 0x0d, 0x56, 0x31, 0xa0, 0xf2,
	0xd8, 0xc4, 0xb7, 0xca, 0x4d, 0xbc, 0xcc, 0x90, 0xd8, 0x96, 0xed, 0x0c, 0x7a, 0x3b, 0x5c, 0x0b,
	0x17, 0x86, 0x62, 0x3b, 0x6b, 0xda, 0x0e, 0x18, 0x19, 0x1f, 0xb6, 0x26, 0x79, 0x7e, 0xcf, 0x93,
	0x73, 0x62, 0xf9, 0x18, 0x5c, 0x25, 0x42, 0x24, 0x29, 0x0d, 0x12, 0x91, 0x12, 0x9e, 0x04, 0x42,
	0x25, 0x61, 0x42, 0xb9, 0x54, 0xc2, 0x88, 0xd0, 0x49, 0x44, 0x32, 0x1d, 0x12, 0xc9, 0x42, 0x4d,
	0xd5, 0x13, 0x8b, 0x68, 0x24, 0xf8, 0x03, 0x4b, 0xca, 0xbb, 0xc0, 0xa2, 0x10, 0xf8, 0x18, 0x22,
	0x59, 0x67, 0x58, 0x37, 0x92, 0x70, 0x2e, 0x0c, 0x31, 0x4c, 0x70, 0xed, 0x62, 0x3b, 0x83, 0xda,
	0x51, 0xb9, 0x79, 0xf4, 0x19, 0xb5, 0x3b, 0x9c, 0x90, 0xe8, 0x57, 0xca, 0xe3, 0xb7, 0xc6, 0x44,
	0x82, 0x1b, 0xfa, 0xbb, 0xf9, 0x3f, 0x62, 0x94, 0x48, 0x7d, 0xcc, 0x77, 0x75, 0x63, 0x62, 0x11,
	0xe5, 0x19, 0xe5, 0xee, 0x33, 0xfb, 0xb0, 0x6f, 0xea, 0x86, 0x51, 0x1e, 0x4b, 0xc1, 0xb8, 0x79,
	0xeb, 0x6a, 0x3d, 0x1a, 0x23, 0x7d, 0xc6, 0x97, 0xaf, 0xcf, 0x48, 0xc9, 0x84, 0xa6, 0xee, 0xbf,
	0x37, 0xf7, 0xeb, 0x16, 0x90, 0x8a, 0xe4, 0xad, 0xeb, 0x93, 0x8a, 0x24, 0x61, 0xbc, 0x8c, 0xf9,
	0xea, 0xf5, 0x31, 0x19, 0x35, 0x8a, 0x45, 0xfe, 0xc7, 0xdb, 0x7f, 0xfc, 0x0f, 0x76, 0xc1, 0x99,
	0x11, 0x8a, 0xc6, 0x8a, 0xea, 0xd9, 0xcd, 0x58, 0x51, 0x2d, 0x72, 0x55, 0xee, 0xcf, 0xce, 0x75,
	0xdd, 0xb6, 0x7c, 0xe2, 0xac, 0xb3, 0x1f, 0x6a, 0x0f, 0x8c, 0xa9, 0x36, 0x34, 0x1b, 0x4b, 0xa2,
	0x48, 0x46, 0x0d, 0x55, 0x3e, 0xef, 0xa2, 0x6e, 0x5e, 0xae, 0x49, 0x52, 0xb6, 0x17, 0x26, 0xcc,
	0x3c, 0xe6, 0x93, 0x20, 0x12, 0x59, 0xe8, 0x82, 0x42, 0x2b, 0x4c, 0xf2, 0x87, 0x50, 0x9a, 0xa9,
	0xa4, 0x3a, 0x24, 0x7c, 0x5a, 0xfc, 0x79, 0xc3, 0xe9, 0x3f, 0xbc, 0xf5, 0xd9, 0x49, 0x24, 0xf3,
	0x78, 0xf0, 0x1a, 0xbc, 0x78, 0x8f, 0xe7, 0xbf, 0xf8, 0xf7, 0x7a, 0x7e, 0x53, 0x44, 0x4a, 0xaa,
	0x66, 0x17, 0xce, 0xfa, 0xe9, 0x9f, 0xab, 0x60, 0xe5, 0xd6, 0x35, 0x0a, 0x2f, 0xc0, 0x96, 0x6b,
	0x76, 0xfc, 0x44, 0x95, 0x66, 0x82, 0xa3, 0xbd, 0xe3, 0x46, 0x77, 0xbd, 0x77, 0x54, 0xd6, 0x53,
	0x86, 0x06, 0x77, 0x43, 0x6e, 0xce, 0x7a, 0xf7, 0x24, 0xcd, 0x29, 0xde, 0x74, 0x9e, 0x7b, 0x67,
	0x81, 0x10, 0xb4, 0x38, 0xc9, 0x28, 0x6a, 0x1c, 0x37, 0xba, 0x6b, 0xd8, 0x5e, 0xc3, 0x2d, 0xd0,
	0x64, 0x31, 0xfa, 0xc4, 0x3e, 0x69, 0xb2, 0x18, 0xee, 0x81, 0xb6, 0x61, 0x26, 0xa5, 0xa8, 0x69,
	0x1f, 0xb9, 0x1b, 0x18, 0x80, 0x5d, 0xa9, 0x44, 0x9c, 0x47, 0x54, 0x8d, 0xa5, 0x12, 0xbf, 0xd0,
	0xc8, 0x8c, 0x59, 0x8c, 0x0e, 0x2c, 0xb3, 0x53, 0x4a, 0x23, 0xa7, 0x0c, 0x63, 0xd8, 0x05, 0xad,
	0x62, 0xad, 0xd0, 0xbb, 0xe3, 0x77, 0xdd, 0xf5, 0xde, 0xde, 0x52, 0x91, 0x7d, 0xc9, 0xb0, 0x25,
	0xe0, 0xe7, 0xa0, 0x6d, 0xbf, 0x02, 0x6a, 0x59, 0x74, 0x7f, 0x09, 0xfd, 0x69, 0x2a, 0x29, 0x76,
	0x4c, 0x01, 0x53, 0x9e, 0x67, 0x1a, 0xb5, 0xff, 0x06, 0xbe, 0xe2, 0x79, 0x86, 0x1d, 0x03, 0xbf,
	0x06, 0x9b, 0x73, 0x23, 0x0e, 0xbd, 0xb7, 0x5f, 0xec, 0x30, 0x98, 0x1d, 0x50, 0xc1, 0x65, 0x15,
	0xc0, 0xf3, 0x3c, 0x3c, 0x05, 0x2b, 0x7e, 0xf0, 0xa3, 0x55, 0x6b, 0xdd, 0xad, 0x5a, 0x07, 0x4e,
	0xc2, 0x25, 0x03, 0x3f, 0x03, 0xad, 0x62, 0x7a, 0xa1, 0x35, 0xcb, 0x6e, 0x57, 0xd9, 0x6b, 0x63,
	0x24, 0xb6, 0x2a, 0x1c, 0x80, 0xad, 0xe2, 0x44, 0xa2, 0xdc, 0xb0, 0xc8, 0x95, 0xb5, 0x6e, 0xf9,
	0x4e, 0x95, 0xef, 0xcf, 0x11, 0x78, 0xc1, 0x51, 0x14, 0xe6, 0x8f, 0x12, 0xb4, 0xb1, 0x5c, 0xd8,
	0x85, 0x93, 0x70, 0xc9, 0xc0, 0x13, 0xd0, 0xb6, 0x3b, 0x04, 0x7d, 0x60, 0xe1, 0x9d, 0x2a, 0x7c,
	0x57, 0x08, 0xd8, 0xe9, 0xb0, 0x07, 0xd6, 0xca, 0x39, 0xae, 0x11, 0x9c, 0x5f, 0xba, 0x02, 0xbe,
	0xf2, 0x22, 0x9e, 0x61, 0x65, 0x2d, 0x4a, 0xa4, 0x68, 0xff, 0xe5, 0x5a, 0x94, 0x48, 0x71, 0xc9,
	0xc0, 0x53, 0xd0, 0x4a, 0x45, 0xa2, 0xd1, 0x87, 0x36, 0x7d, 0x6e, 0x2d, 0x6e, 0x44, 0x72, 0x49,
	0x75, 0xa4, 0x98, 0x34, 0x42, 0x61, 0x8b, 0xc1, 0x73, 0xb0, 0xe2, 0xa6, 0xa1, 0x46, 0xc8, 0x3a,
	0x8e, 0xaa, 0x8e, 0xef, 0xad, 0x54, 0x31, 0x95, 0x30, 0xfc, 0x19, 0xec, 0x2e, 0x0f, 0x40, 0x8d,
	0x0e, 0x6d, 0xc6, 0xc9, 0x5c, 0x46, 0x89, 0x61, 0x4f, 0x55, 0xe2, 0x60, 0xb6, 0x28, 0xda, 0x7e,
	0xfd, 0x7c, 0x47, 0x1f, 0x2d, 0xf7, 0x7b, 0xe3, 0x24, 0x5c, 0x32, 0xf0, 0x1c, 0x80, 0xd9, 0xdc,
	0x44, 0x47, 0xd6, 0x71, 0xf0, 0xc2, 0xfb, 0x0b, 0x53, 0x85, 0x84, 0x43, 0xb0, 0xb3, 0x38, 0x25,
	0x35, 0xfa, 0x78, 0x7e, 0xcb, 0x17, 0xf6, 0x5b, 0x0b, 0x8d, 0x9e, 0x19, 0xbc, 0xad, 0x17, 0x9e,
	0x0c, 0x4e, 0x8a, 0xd1, 0x91, 0x55, 0x4c, 0x83, 0x0d, 0x3f, 0x55, 0x46, 0xc5, 0xb6, 0x19, 0x35,
	0xfe, 0x68, 0xb6, 0xbe, 0xed, 0x8f, 0x86, 0x93, 0xf7, 0x76, 0x1b, 0x9d, 0xfd, 0x15, 0x00, 0x00,
	0xff, 0xff, 0xfc, 0x22, 0x08, 0x2f, 0x09, 0x0a, 0x00, 0x00,
}
