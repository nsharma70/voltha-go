/*
 * Copyright 2018-present Open Networking Foundation

 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at

 * http://www.apache.org/licenses/LICENSE-2.0

 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"flag"
	"fmt"
	"github.com/opencord/voltha-lib-go/v3/pkg/adapters/common"
	"time"
)

// RW Core service default constants
const (
	ConsulStoreName                  = "consul"
	EtcdStoreName                    = "etcd"
	defaultGrpcAddress               = ":50057"
	defaultKafkaAdapterAddress       = "127.0.0.1:9092"
	defaultKafkaClusterAddress       = "127.0.0.1:9094"
	defaultKVStoreType               = EtcdStoreName
	defaultKVStoreTimeout            = 5                //in seconds
	defaultKVStoreAddress            = "127.0.0.1:2379" // Port: Consul = 8500; Etcd = 2379
	defaultKVTxnKeyDelTime           = 60
	defaultKVStoreDataPrefix         = "service/voltha"
	defaultLogLevel                  = "WARN"
	defaultBanner                    = false
	defaultDisplayVersionOnly        = false
	defaultCoreTopic                 = "rwcore"
	defaultRWCoreEndpoint            = "rwcore"
	defaultRWCoreKey                 = "pki/voltha.key"
	defaultRWCoreCert                = "pki/voltha.crt"
	defaultRWCoreCA                  = "pki/voltha-CA.pem"
	defaultAffinityRouterTopic       = "affinityRouter"
	defaultInCompetingMode           = true
	defaultLongRunningRequestTimeout = 2000 * time.Millisecond
	defaultDefaultRequestTimeout     = 1000 * time.Millisecond
	defaultCoreTimeout               = 1000 * time.Millisecond
	defaultCoreBindingKey            = "voltha_backend_name"
	defaultCorePairTopic             = "rwcore_1"
	defaultMaxConnectionRetries      = -1 // retries forever
	defaultConnectionRetryInterval   = 2 * time.Second
	defaultLiveProbeInterval         = 60 * time.Second
	defaultNotLiveProbeInterval      = 5 * time.Second // Probe more frequently when not alive
	defaultProbeAddress              = ":8080"
)

type stringValue string

func (i *stringValue) Set(s string) error {
	if err := common.ValidateAddress(s); err != nil {
		return err
	}
	*i = stringValue(s)

	return nil
}

func (i *stringValue) String() string {
	if *i == "" {
		return defaultKafkaAdapterAddress
	}
	return string(*i)
}

// RWCoreFlags represents the set of configurations used by the read-write core service
type RWCoreFlags struct {
	// Command line parameters
	RWCoreEndpoint            string
	GrpcAddress               string
	KafkaAdapterAddress       stringValue
	KafkaClusterAddress       string
	KVStoreType               string
	KVStoreTimeout            int // in seconds
	KVStoreAddress            string
	KVTxnKeyDelTime           int
	KVStoreDataPrefix         string
	CoreTopic                 string
	LogLevel                  string
	Banner                    bool
	DisplayVersionOnly        bool
	RWCoreKey                 string
	RWCoreCert                string
	RWCoreCA                  string
	AffinityRouterTopic       string
	InCompetingMode           bool
	LongRunningRequestTimeout time.Duration
	DefaultRequestTimeout     time.Duration
	DefaultCoreTimeout        time.Duration
	CoreBindingKey            string
	CorePairTopic             string
	MaxConnectionRetries      int
	ConnectionRetryInterval   time.Duration
	LiveProbeInterval         time.Duration
	NotLiveProbeInterval      time.Duration
	ProbeAddress              string
}

// NewRWCoreFlags returns a new RWCore config
func NewRWCoreFlags() *RWCoreFlags {
	var rwCoreFlag = RWCoreFlags{ // Default values
		RWCoreEndpoint:            defaultRWCoreEndpoint,
		GrpcAddress:               defaultGrpcAddress,
		KafkaAdapterAddress:       defaultKafkaAdapterAddress,
		KafkaClusterAddress:       defaultKafkaClusterAddress,
		KVStoreType:               defaultKVStoreType,
		KVStoreTimeout:            defaultKVStoreTimeout,
		KVStoreAddress:            defaultKVStoreAddress,
		KVStoreDataPrefix:         defaultKVStoreDataPrefix,
		KVTxnKeyDelTime:           defaultKVTxnKeyDelTime,
		CoreTopic:                 defaultCoreTopic,
		LogLevel:                  defaultLogLevel,
		Banner:                    defaultBanner,
		DisplayVersionOnly:        defaultDisplayVersionOnly,
		RWCoreKey:                 defaultRWCoreKey,
		RWCoreCert:                defaultRWCoreCert,
		RWCoreCA:                  defaultRWCoreCA,
		AffinityRouterTopic:       defaultAffinityRouterTopic,
		InCompetingMode:           defaultInCompetingMode,
		DefaultRequestTimeout:     defaultDefaultRequestTimeout,
		LongRunningRequestTimeout: defaultLongRunningRequestTimeout,
		DefaultCoreTimeout:        defaultCoreTimeout,
		CoreBindingKey:            defaultCoreBindingKey,
		CorePairTopic:             defaultCorePairTopic,
		MaxConnectionRetries:      defaultMaxConnectionRetries,
		ConnectionRetryInterval:   defaultConnectionRetryInterval,
		LiveProbeInterval:         defaultLiveProbeInterval,
		NotLiveProbeInterval:      defaultNotLiveProbeInterval,
		ProbeAddress:              defaultProbeAddress,
	}
	return &rwCoreFlag
}

// ParseCommandArguments parses the arguments when running read-write core service
func (cf *RWCoreFlags) ParseCommandArguments() {

	help := fmt.Sprintf("RW core endpoint address")
	flag.StringVar(&(cf.RWCoreEndpoint), "vcore-endpoint", defaultRWCoreEndpoint, help)

	help = fmt.Sprintf("GRPC server - address")
	flag.StringVar(&(cf.GrpcAddress), "grpc_address", defaultGrpcAddress, help)

	help = fmt.Sprintf("Kafka - Adapter messaging address")
	flag.Var(&(cf.KafkaAdapterAddress), "kafka_adapter_address", help)

	help = fmt.Sprintf("Kafka - Cluster messaging address")
	flag.StringVar(&(cf.KafkaClusterAddress), "kafka_cluster_address", defaultKafkaClusterAddress, help)

	help = fmt.Sprintf("RW Core topic")
	flag.StringVar(&(cf.CoreTopic), "rw_core_topic", defaultCoreTopic, help)

	help = fmt.Sprintf("Affinity Router topic")
	flag.StringVar(&(cf.AffinityRouterTopic), "affinity_router_topic", defaultAffinityRouterTopic, help)

	help = fmt.Sprintf("In competing Mode - two cores competing to handle a transaction ")
	flag.BoolVar(&cf.InCompetingMode, "in_competing_mode", defaultInCompetingMode, help)

	help = fmt.Sprintf("KV store type")
	flag.StringVar(&(cf.KVStoreType), "kv_store_type", defaultKVStoreType, help)

	help = fmt.Sprintf("The default timeout when making a kv store request")
	flag.IntVar(&(cf.KVStoreTimeout), "kv_store_request_timeout", defaultKVStoreTimeout, help)

	help = fmt.Sprintf("KV store address")
	flag.StringVar(&(cf.KVStoreAddress), "kv_store_address", defaultKVStoreAddress, help)

	help = fmt.Sprintf("The time to wait before deleting a completed transaction key")
	flag.IntVar(&(cf.KVTxnKeyDelTime), "kv_txn_delete_time", defaultKVTxnKeyDelTime, help)

	help = fmt.Sprintf("KV store data prefix")
	flag.StringVar(&(cf.KVStoreDataPrefix), "kv_store_data_prefix", defaultKVStoreDataPrefix, help)

	help = fmt.Sprintf("Log level")
	flag.StringVar(&(cf.LogLevel), "log_level", defaultLogLevel, help)

	help = fmt.Sprintf("Timeout for long running request")
	// TODO:  Change this code once all the params and helm charts have been changed to use the different type
	var temp int64
	flag.Int64Var(&temp, "timeout_long_request", defaultLongRunningRequestTimeout.Milliseconds(), help)
	cf.LongRunningRequestTimeout = time.Duration(temp) * time.Millisecond

	help = fmt.Sprintf("Default timeout for regular request")
	flag.Int64Var(&temp, "timeout_request", defaultDefaultRequestTimeout.Milliseconds(), help)
	cf.DefaultRequestTimeout = time.Duration(temp) * time.Millisecond

	help = fmt.Sprintf("Default Core timeout")
	flag.Int64Var(&temp, "core_timeout", defaultCoreTimeout.Milliseconds(), help)
	cf.DefaultCoreTimeout = time.Duration(temp) * time.Millisecond

	help = fmt.Sprintf("Show startup banner log lines")
	flag.BoolVar(&cf.Banner, "banner", defaultBanner, help)

	help = fmt.Sprintf("Show version information and exit")
	flag.BoolVar(&cf.DisplayVersionOnly, "version", defaultDisplayVersionOnly, help)

	help = fmt.Sprintf("The name of the meta-key whose value is the rw-core group to which the ofagent is bound")
	flag.StringVar(&(cf.CoreBindingKey), "core_binding_key", defaultCoreBindingKey, help)

	help = fmt.Sprintf("Core pairing group topic")
	flag.StringVar(&cf.CorePairTopic, "core_pair_topic", defaultCorePairTopic, help)

	help = fmt.Sprintf("The number of retries to connect to a dependent component")
	flag.IntVar(&(cf.MaxConnectionRetries), "max_connection_retries", defaultMaxConnectionRetries, help)

	help = fmt.Sprintf("The number of seconds between each connection retry attempt")
	flag.DurationVar(&(cf.ConnectionRetryInterval), "connection_retry_interval", defaultConnectionRetryInterval, help)

	help = fmt.Sprintf("The number of seconds between liveness probes while in a live state")
	flag.DurationVar(&(cf.LiveProbeInterval), "live_probe_interval", defaultLiveProbeInterval, help)

	help = fmt.Sprintf("The number of seconds between liveness probes while in a not live state")
	flag.DurationVar(&(cf.NotLiveProbeInterval), "not_live_probe_interval", defaultNotLiveProbeInterval, help)

	help = fmt.Sprintf("The address on which to listen to answer liveness and readiness probe queries over HTTP.")
	flag.StringVar(&(cf.ProbeAddress), "probe_address", defaultProbeAddress, help)

	flag.Parse()
}
