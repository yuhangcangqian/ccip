// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {OptimismL1BridgeAdapter} from "../bridge-adapters/OptimismL1BridgeAdapter.sol";

/// @dev to generate abi's for the OptimismL1BridgeAdapter's various payload types.
/// @dev for usage examples see core/scripts/ccip/rebalancer/opstack/prove_withdrawal.go
/// @dev or core/scripts/ccip/rebalancer/opstack/finalize.go.
abstract contract OptimismL1BridgeAdapterEncoder {
  function encodeFinalizeWithdrawalERC20Payload(
    OptimismL1BridgeAdapter.FinalizeWithdrawERC20Payload memory payload
  ) public pure {}

  function encodeOptimismProveWithdrawalPayload(
    OptimismL1BridgeAdapter.OptimismProveWithdrawalPayload memory payload
  ) public pure {}

  function encodeOptimismFinalizationPayload(
    OptimismL1BridgeAdapter.OptimismFinalizationPayload memory payload
  ) public pure {}
}
