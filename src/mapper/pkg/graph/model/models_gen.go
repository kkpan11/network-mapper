// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type AWSOperation struct {
	Resource string          `json:"resource"`
	Actions  []string        `json:"actions"`
	SrcIP    *string         `json:"srcIp,omitempty"`
	Client   *NamespacedName `json:"client,omitempty"`
}

type AzureOperation struct {
	Scope           string   `json:"scope"`
	Actions         []string `json:"actions"`
	DataActions     []string `json:"dataActions"`
	ClientName      string   `json:"clientName"`
	ClientNamespace string   `json:"clientNamespace"`
}

type CaptureResults struct {
	Results []RecordedDestinationsForSrc `json:"results"`
}

type CaptureTCPResults struct {
	Results []RecordedDestinationsForSrc `json:"results"`
}

type Destination struct {
	Destination     string    `json:"destination"`
	DestinationIP   *string   `json:"destinationIP,omitempty"`
	DestinationPort *int64    `json:"destinationPort,omitempty"`
	TTL             *int64    `json:"TTL,omitempty"`
	LastSeen        time.Time `json:"lastSeen"`
}

type GroupVersionKind struct {
	Group   *string `json:"group,omitempty"`
	Version string  `json:"version"`
	Kind    string  `json:"kind"`
}

type HTTPResource struct {
	Path    string       `json:"path"`
	Methods []HTTPMethod `json:"methods,omitempty"`
}

type IdentityResolutionData struct {
	Host           *string `json:"host,omitempty"`
	PodHostname    *string `json:"podHostname,omitempty"`
	ProcfsHostname *string `json:"procfsHostname,omitempty"`
	Port           *int64  `json:"port,omitempty"`
	IsService      *bool   `json:"isService,omitempty"`
	Uptime         *string `json:"uptime,omitempty"`
	LastSeen       *string `json:"lastSeen,omitempty"`
	ExtraInfo      *string `json:"extraInfo,omitempty"`
}

type Intent struct {
	Client         *OtterizeServiceIdentity `json:"client"`
	Server         *OtterizeServiceIdentity `json:"server"`
	Type           *IntentType              `json:"type,omitempty"`
	ResolutionData *string                  `json:"resolutionData,omitempty"`
	KafkaTopics    []KafkaConfig            `json:"kafkaTopics,omitempty"`
	HTTPResources  []HTTPResource           `json:"httpResources,omitempty"`
	AwsActions     []string                 `json:"awsActions,omitempty"`
}

type IstioConnection struct {
	SrcWorkload          string       `json:"srcWorkload"`
	SrcWorkloadNamespace string       `json:"srcWorkloadNamespace"`
	DstWorkload          string       `json:"dstWorkload"`
	DstServiceName       string       `json:"dstServiceName"`
	DstWorkloadNamespace string       `json:"dstWorkloadNamespace"`
	Path                 string       `json:"path"`
	Methods              []HTTPMethod `json:"methods"`
	LastSeen             time.Time    `json:"lastSeen"`
}

type IstioConnectionResults struct {
	Results []IstioConnection `json:"results"`
}

type KafkaConfig struct {
	Name       string           `json:"name"`
	Operations []KafkaOperation `json:"operations,omitempty"`
}

type KafkaMapperResult struct {
	SrcIP           string    `json:"srcIp"`
	ServerPodName   string    `json:"serverPodName"`
	ServerNamespace string    `json:"serverNamespace"`
	Topic           string    `json:"topic"`
	Operation       string    `json:"operation"`
	LastSeen        time.Time `json:"lastSeen"`
}

type KafkaMapperResults struct {
	Results []KafkaMapperResult `json:"results"`
}

type Mutation struct {
}

type NamespacedName struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type OtterizeServiceIdentity struct {
	Name           string                  `json:"name"`
	Namespace      string                  `json:"namespace"`
	Labels         []PodLabel              `json:"labels,omitempty"`
	ResolutionData *IdentityResolutionData `json:"resolutionData,omitempty"`
	// If the service identity was resolved from a pod owner, the GroupVersionKind of the pod owner.
	PodOwnerKind *GroupVersionKind `json:"podOwnerKind,omitempty"`
	// If the service identity was resolved from a Kubernetes service, its name.
	KubernetesService *string `json:"kubernetesService,omitempty"`
}

type PodLabel struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Query struct {
}

type RecordedDestinationsForSrc struct {
	SrcIP        string        `json:"srcIp"`
	SrcHostname  string        `json:"srcHostname"`
	Destinations []Destination `json:"destinations"`
}

type ServerFilter struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type ServiceIntents struct {
	Client  *OtterizeServiceIdentity  `json:"client"`
	Intents []OtterizeServiceIdentity `json:"intents"`
}

type SocketScanResults struct {
	Results []RecordedDestinationsForSrc `json:"results"`
}

type HTTPMethod string

const (
	HTTPMethodGet     HTTPMethod = "GET"
	HTTPMethodPost    HTTPMethod = "POST"
	HTTPMethodPut     HTTPMethod = "PUT"
	HTTPMethodDelete  HTTPMethod = "DELETE"
	HTTPMethodOptions HTTPMethod = "OPTIONS"
	HTTPMethodTrace   HTTPMethod = "TRACE"
	HTTPMethodPatch   HTTPMethod = "PATCH"
	HTTPMethodConnect HTTPMethod = "CONNECT"
	HTTPMethodAll     HTTPMethod = "ALL"
)

var AllHTTPMethod = []HTTPMethod{
	HTTPMethodGet,
	HTTPMethodPost,
	HTTPMethodPut,
	HTTPMethodDelete,
	HTTPMethodOptions,
	HTTPMethodTrace,
	HTTPMethodPatch,
	HTTPMethodConnect,
	HTTPMethodAll,
}

func (e HTTPMethod) IsValid() bool {
	switch e {
	case HTTPMethodGet, HTTPMethodPost, HTTPMethodPut, HTTPMethodDelete, HTTPMethodOptions, HTTPMethodTrace, HTTPMethodPatch, HTTPMethodConnect, HTTPMethodAll:
		return true
	}
	return false
}

func (e HTTPMethod) String() string {
	return string(e)
}

func (e *HTTPMethod) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = HTTPMethod(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid HttpMethod", str)
	}
	return nil
}

func (e HTTPMethod) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type IntentType string

const (
	IntentTypeHTTP     IntentType = "HTTP"
	IntentTypeKafka    IntentType = "KAFKA"
	IntentTypeDatabase IntentType = "DATABASE"
	IntentTypeAws      IntentType = "AWS"
	IntentTypeS3       IntentType = "S3"
)

var AllIntentType = []IntentType{
	IntentTypeHTTP,
	IntentTypeKafka,
	IntentTypeDatabase,
	IntentTypeAws,
	IntentTypeS3,
}

func (e IntentType) IsValid() bool {
	switch e {
	case IntentTypeHTTP, IntentTypeKafka, IntentTypeDatabase, IntentTypeAws, IntentTypeS3:
		return true
	}
	return false
}

func (e IntentType) String() string {
	return string(e)
}

func (e *IntentType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = IntentType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid IntentType", str)
	}
	return nil
}

func (e IntentType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type KafkaOperation string

const (
	KafkaOperationAll             KafkaOperation = "ALL"
	KafkaOperationConsume         KafkaOperation = "CONSUME"
	KafkaOperationProduce         KafkaOperation = "PRODUCE"
	KafkaOperationCreate          KafkaOperation = "CREATE"
	KafkaOperationAlter           KafkaOperation = "ALTER"
	KafkaOperationDelete          KafkaOperation = "DELETE"
	KafkaOperationDescribe        KafkaOperation = "DESCRIBE"
	KafkaOperationClusterAction   KafkaOperation = "CLUSTER_ACTION"
	KafkaOperationDescribeConfigs KafkaOperation = "DESCRIBE_CONFIGS"
	KafkaOperationAlterConfigs    KafkaOperation = "ALTER_CONFIGS"
	KafkaOperationIdempotentWrite KafkaOperation = "IDEMPOTENT_WRITE"
)

var AllKafkaOperation = []KafkaOperation{
	KafkaOperationAll,
	KafkaOperationConsume,
	KafkaOperationProduce,
	KafkaOperationCreate,
	KafkaOperationAlter,
	KafkaOperationDelete,
	KafkaOperationDescribe,
	KafkaOperationClusterAction,
	KafkaOperationDescribeConfigs,
	KafkaOperationAlterConfigs,
	KafkaOperationIdempotentWrite,
}

func (e KafkaOperation) IsValid() bool {
	switch e {
	case KafkaOperationAll, KafkaOperationConsume, KafkaOperationProduce, KafkaOperationCreate, KafkaOperationAlter, KafkaOperationDelete, KafkaOperationDescribe, KafkaOperationClusterAction, KafkaOperationDescribeConfigs, KafkaOperationAlterConfigs, KafkaOperationIdempotentWrite:
		return true
	}
	return false
}

func (e KafkaOperation) String() string {
	return string(e)
}

func (e *KafkaOperation) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = KafkaOperation(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid KafkaOperation", str)
	}
	return nil
}

func (e KafkaOperation) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
