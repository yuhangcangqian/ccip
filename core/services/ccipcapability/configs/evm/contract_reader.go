package evm

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"

	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/ccip_config"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_onramp"
	kcr "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/capabilities_registry"
	evmrelaytypes "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
)

var (
	onrampABI               = evmtypes.MustGetABI(evm_2_evm_multi_onramp.EVM2EVMMultiOnRampABI)
	capabilitiesRegsitryABI = evmtypes.MustGetABI(kcr.CapabilitiesRegistryABI)
	ccipConfigABI           = evmtypes.MustGetABI(ccip_config.CCIPConfigABI)
)

// MustSourceReaderConfig returns a ChainReaderConfig that can be used to read from the onramp.
// The configuration is marshaled into JSON so that it can be passed to the relayer NewContractReader() method.
func MustSourceReaderConfig() []byte {
	rawConfig := SourceReaderConfig()
	encoded, err := json.Marshal(rawConfig)
	if err != nil {
		panic(fmt.Errorf("failed to marshal ChainReaderConfig into JSON: %w", err))
	}

	return encoded
}

// MustDestReaderConfig returns a ChainReaderConfig that can be used to read from the offramp.
// The configuration is marshaled into JSON so that it can be passed to the relayer NewContractReader() method.
func MustDestReaderConfig() []byte {
	rawConfig := DestReaderConfig()
	encoded, err := json.Marshal(rawConfig)
	if err != nil {
		panic(fmt.Errorf("failed to marshal ChainReaderConfig into JSON: %w", err))
	}

	return encoded
}

// DestReaderConfig returns a ChainReaderConfig that can be used to read from the offramp.
func DestReaderConfig() evmrelaytypes.ChainReaderConfig {
	return evmrelaytypes.ChainReaderConfig{
		Contracts: map[string]evmrelaytypes.ChainContractReader{
			consts.ContractNameOffRamp: {
				ContractABI: evm_2_evm_multi_offramp.EVM2EVMMultiOffRampABI,
				ContractPollingFilter: evmrelaytypes.ContractPollingFilter{
					GenericEventNames: []string{
						mustGetEventName(consts.EventNameExecutionStateChanged, offrampABI),
						mustGetEventName(consts.EventNameCommitReportAccepted, offrampABI),
					},
				},
				Configs: map[string]*evmrelaytypes.ChainReaderDefinition{
					consts.MethodNameGetExecutionState: {
						ChainSpecificName: mustGetMethodName("getExecutionState", offrampABI),
						ReadType:          evmrelaytypes.Method,
					},
					consts.MethodNameGetMerkleRoot: {
						ChainSpecificName: mustGetMethodName("getMerkleRoot", offrampABI),
						ReadType:          evmrelaytypes.Method,
					},
					consts.MethodNameIsBlessed: {
						ChainSpecificName: mustGetMethodName("isBlessed", offrampABI),
						ReadType:          evmrelaytypes.Method,
					},
					consts.MethodNameGetLatestPriceSequenceNumber: {
						ChainSpecificName: mustGetMethodName("getLatestPriceSequenceNumber", offrampABI),
						ReadType:          evmrelaytypes.Method,
					},
					consts.MethodNameOfframpGetStaticConfig: {
						ChainSpecificName: mustGetMethodName("getStaticConfig", offrampABI),
						ReadType:          evmrelaytypes.Method,
					},
					consts.MethodNameOfframpGetDynamicConfig: {
						ChainSpecificName: mustGetMethodName("getDynamicConfig", offrampABI),
						ReadType:          evmrelaytypes.Method,
					},
					consts.MethodNameGetSourceChainConfig: {
						ChainSpecificName: mustGetMethodName("getSourceChainConfig", offrampABI),
						ReadType:          evmrelaytypes.Method,
					},
					consts.EventNameCommitReportAccepted: {
						ChainSpecificName: mustGetEventName(consts.EventNameCommitReportAccepted, offrampABI),
						ReadType:          evmrelaytypes.Event,
					},
					consts.EventNameExecutionStateChanged: {
						ChainSpecificName: mustGetEventName(consts.EventNameExecutionStateChanged, offrampABI),
						ReadType:          evmrelaytypes.Event,
					},
				},
			},
		},
	}
}

// SourceReaderConfig returns a ChainReaderConfig that can be used to read from the onramp.
func SourceReaderConfig() evmrelaytypes.ChainReaderConfig {
	return evmrelaytypes.ChainReaderConfig{
		Contracts: map[string]evmrelaytypes.ChainContractReader{
			consts.ContractNameOnRamp: {
				ContractABI: evm_2_evm_multi_onramp.EVM2EVMMultiOnRampABI,
				ContractPollingFilter: evmrelaytypes.ContractPollingFilter{
					GenericEventNames: []string{
						mustGetEventName(consts.EventNameCCIPSendRequested, onrampABI),
					},
				},
				Configs: map[string]*evmrelaytypes.ChainReaderDefinition{
					// all "{external|public} view" functions in the onramp except for getFee and getPoolBySourceToken are here.
					// getFee is not expected to get called offchain and is only called by end-user contracts.
					consts.MethodNameGetExpectedNextSequenceNumber: {
						ChainSpecificName: mustGetMethodName("getExpectedNextSequenceNumber", onrampABI),
						ReadType:          evmrelaytypes.Method,
					},
					consts.MethodNameOnrampGetStaticConfig: {
						ChainSpecificName: mustGetMethodName("getStaticConfig", onrampABI),
						ReadType:          evmrelaytypes.Method,
					},
					consts.MethodNameOnrampGetDynamicConfig: {
						ChainSpecificName: mustGetMethodName("getDynamicConfig", onrampABI),
						ReadType:          evmrelaytypes.Method,
					},
					consts.MethodNameGetDestChainConfig: {
						ChainSpecificName: mustGetMethodName("getDestChainConfig", onrampABI),
						ReadType:          evmrelaytypes.Method,
					},
					consts.MethodNameGetPremiumMultiplierWeiPerEth: {
						ChainSpecificName: mustGetMethodName("getPremiumMultiplierWeiPerEth", onrampABI),
						ReadType:          evmrelaytypes.Method,
					},
					consts.MethodNameGetTokenTransferFeeConfig: {
						ChainSpecificName: mustGetMethodName("getTokenTransferFeeConfig", onrampABI),
						ReadType:          evmrelaytypes.Method,
					},
					consts.EventNameCCIPSendRequested: {
						ChainSpecificName: mustGetEventName(consts.EventNameCCIPSendRequested, onrampABI),
						ReadType:          evmrelaytypes.Event,
						EventDefinitions: &evmrelaytypes.EventDefinitions{
							GenericDataWordNames: map[string]uint8{
								consts.EventAttributeSequenceNumber: 5,
							},
						},
					},
				},
			},
		},
	}
}

// HomeChainReaderConfigRaw returns a ChainReaderConfig that can be used to read from the home chain.
func HomeChainReaderConfigRaw() evmrelaytypes.ChainReaderConfig {
	return evmrelaytypes.ChainReaderConfig{
		Contracts: map[string]evmrelaytypes.ChainContractReader{
			consts.ContractNameCapabilitiesRegistry: {
				ContractABI: kcr.CapabilitiesRegistryABI,
				Configs: map[string]*evmrelaytypes.ChainReaderDefinition{
					consts.MethodNameGetCapability: {
						ChainSpecificName: mustGetMethodName("getCapability", capabilitiesRegsitryABI),
					},
				},
			},
			consts.ContractNameCCIPConfig: {
				ContractABI: ccip_config.CCIPConfigABI,
				Configs: map[string]*evmrelaytypes.ChainReaderDefinition{
					consts.MethodNameGetAllChainConfigs: {
						ChainSpecificName: mustGetMethodName("getAllChainConfigs", ccipConfigABI),
					},
					consts.MethodNameGetOCRConfig: {
						ChainSpecificName: mustGetMethodName("getOCRConfig", ccipConfigABI),
					},
				},
			},
		},
	}
}

func mustGetEventName(event string, tabi abi.ABI) string {
	e, ok := tabi.Events[event]
	if !ok {
		panic(fmt.Sprintf("missing event %s in onrampABI", event))
	}
	return e.Name
}
