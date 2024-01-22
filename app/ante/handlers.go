package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	ante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	ibcante "github.com/cosmos/ibc-go/v6/modules/core/ante"
	evmcosmosante "github.com/evmos/evmos/v12/app/ante/cosmos"
	evmante "github.com/evmos/evmos/v12/app/ante/evm"
	txfeesante "github.com/osmosis-labs/osmosis/v15/x/txfees/ante"
)

func newEthAnteHandler(options HandlerOptions) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		evmante.NewEthSetUpContextDecorator(options.EvmKeeper),                         // outermost AnteDecorator. SetUpContext must be called first
		evmante.NewEthMempoolFeeDecorator(options.EvmKeeper),                           // Check eth effective gas price against minimal-gas-prices
		evmante.NewEthMinGasPriceDecorator(options.FeeMarketKeeper, options.EvmKeeper), // Check eth effective gas price against the global MinGasPrice
		evmante.NewEthValidateBasicDecorator(options.EvmKeeper),
		evmante.NewEthSigVerificationDecorator(options.EvmKeeper),
		evmante.NewEthAccountVerificationDecorator(options.AccountKeeper, options.EvmKeeper),
		evmante.NewCanTransferDecorator(options.EvmKeeper),
		evmante.NewEthGasConsumeDecorator(options.BankKeeper, options.DistributionKeeper, options.EvmKeeper, options.StakingKeeper, options.MaxTxGasWanted),
		evmante.NewEthIncrementSenderSequenceDecorator(options.AccountKeeper), // innermost AnteDecorator.
		evmante.NewGasWantedDecorator(options.EvmKeeper, options.FeeMarketKeeper),
		evmante.NewEthEmitEventDecorator(options.EvmKeeper), // emit eth tx hash and index at the very last ante handler.
	)
}

func newCosmosAnteHandler(options HandlerOptions) sdk.AnteHandler {
	mempoolFeeDecorator := txfeesante.NewMempoolFeeDecorator(*options.TxFeesKeeper)
	deductFeeDecorator := txfeesante.NewDeductFeeDecorator(*options.TxFeesKeeper, options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper)

	return sdk.ChainAnteDecorators(
		evmcosmosante.RejectMessagesDecorator{}, // reject MsgEthereumTxs
		ante.NewSetUpContextDecorator(),
		ante.NewExtensionOptionsDecorator(options.ExtensionOptionChecker),
		// Use Mempool Fee Decorator from our txfees module instead of default one from auth
		mempoolFeeDecorator,
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		// ethante.NewMinGasPriceDecorator(options.FeeMarketKeeper, options.EvmKeeper), //feemarket decorator
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		deductFeeDecorator,
		ante.NewSetPubKeyDecorator(options.AccountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
		ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
		ibcante.NewRedundantRelayDecorator(options.IBCKeeper),
		// evmante.NewGasWantedDecorator(options.EvmKeeper, options.FeeMarketKeeper),
	)
}
