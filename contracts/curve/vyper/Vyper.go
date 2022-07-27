// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vyper

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// VyperMetaData contains all meta data concerning the Vyper contract.
var VyperMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_base_pool\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[5]\"},{\"name\":\"_min_mint_amount\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[5]\"},{\"name\":\"_min_mint_amount\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange_underlying\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"_dx\",\"type\":\"uint256\"},{\"name\":\"_min_dy\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"exchange_underlying\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"_dx\",\"type\":\"uint256\"},{\"name\":\"_min_dy\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_min_amounts\",\"type\":\"uint256[5]\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_min_amounts\",\"type\":\"uint256[5]\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"_token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"_min_amount\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"_token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"_min_amount\",\"type\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dy_underlying\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"_dx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":19203},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_token_amount\",\"inputs\":[{\"name\":\"_amounts\",\"type\":\"uint256[5]\"},{\"name\":\"_is_deposit\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":9705},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_withdraw_one_coin\",\"inputs\":[{\"name\":\"token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":14299},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"coins\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2643},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"underlying_coins\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2673},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2658},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"base_pool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2688},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2718}]",
	Bin: "0x6f7fffffffffffffffffffffffffffffff6040526040611d99610140396020611d9960c03960c05160a01c611d945760206020611d990160c03960c05160a01c611d94576101405160095561016051600a5560206101e0600463fc0c546a6101805261019c610140515afa15611d9457601f3d1115611d94576000506101e051600b5561018060006003818352015b6020610240602463b9947eb06101c052610180516101e0526101dc610160515afa15611d9457601f3d1115611d9457600050610240516101a0526101a0516001610180516005811015611d9457026004015560006004610220527f095ea7b3000000000000000000000000000000000000000000000000000000006102405261022060048060208461028001018260208501600060045af1505080518201915050610160516020826102800101526020810190507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff602082610280010152602081019050806102805261028090508051602001806103208284600060045af115611d9457505060206103e06103205161034060006101a0515af115611d945760203d808211156101be57806101c0565b815b905090506103c0526103c08051602001806101c08284600060045af115611d9457505060006101c0511815610222576101c08060200151600082518060209013611d945780919012611d9457806020036101000a820490509050905015611d94575b5b815160010180835281141561008e575b505061018060006003818352015b6020610240602463c66106576101c052610180516101e0526101dc610140515afa15611d9457601f3d1115611d9457600050610240516101a0526101a0516001610180516003811015611d9457026001015560006004610220527f095ea7b3000000000000000000000000000000000000000000000000000000006102405261022060048060208461028001018260208501600060045af1505080518201915050610140516020826102800101526020810190507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff602082610280010152602081019050806102805261028090508051602001806103208284600060045af115611d9457505060206103e06103205161034060006101a0515af115611d945760203d808211156103715780610373565b815b905090506103c0526103c08051602001806101c08284600060045af115611d9457505060006101c05118156103d5576101c08060200151600082518060209013611d945780919012611d9457806020036101000a820490509050905015611d94575b6000610180511815610591576020610280600463b16a19de6102205261023c6101a0515afa15611d9457601f3d1115611d9457600050610280516101a0526101a05160016101805160028181830110611d9457808201905090506005811015611d9457026004015560006004610220527f095ea7b3000000000000000000000000000000000000000000000000000000006102405261022060048060208461028001018260208501600060045af1505080518201915050738dff5e27ea6b7ac08ebfdf9eb090f32ee9a30fcf6020826102800101526020810190507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff602082610280010152602081019050806102805261028090508051602001806103208284600060045af115611d9457505060206103e06103205161034060006101a0515af115611d945760203d8082111561052c578061052e565b815b905090506103c0526103c08051602001806101c08284600060045af115611d9457505060006101c0511815610590576101c08060200151600082518060209013611d945780919012611d9457806020036101000a820490509050905015611d94575b5b5b8151600101808352811415610241575b5050611d7c56600436101561000d576117c8565b600035601c526f7fffffffffffffffffffffffffffffff604052600051346117ce576384738499811415610045573361014052610070565b63cf2b51b881141561006b5760c43560a01c6117ce57602060c461014037600050610070565b61061b565b60e0366101603761024060006003818352015b60046102405160058110156117ce5760200201356102605260006102605118156102405760016102405160058110156117ce5702600401546102805260006004610300527f23b872dd000000000000000000000000000000000000000000000000000000006103205261030060048060208461036001018260208501600060045af1505080518201915050336020826103600101526020810190503060208261036001015260208101905061026051602082610360010152602081019050806103605261036090508051602001806104208284600060045af1156117ce5750506020610500610420516104406000610280515af1156117ce5760203d8082111561018d578061018f565b815b905090506104e0526104e08051602001806102a08284600060045af1156117ce57505060006102a05118156101f1576102a080602001516000825180602090136117ce57809190126117ce57806020036101000a8204905090509050156117ce575b602061038060246370a0823161030052306103205261031c610280515afa156117ce57601f3d11156117ce57600050610380516101606102405160038110156117ce5760200201526001610220525b5b8151600101808352811415610083575b505061022051156102b457602061034060a4632b6e993a61024052610160516102605261018051610280526101a0516102a05260006102c05260016102e05261025c6000600a545af1156117ce57601f3d11156117ce57600050610340516101c0525b6000546102405261026060036002818352015b60046102605160058110156117ce5760200201356102805260006102805118156105425760016102605160058110156117ce5702600401546102a05260006004610320527f23b872dd000000000000000000000000000000000000000000000000000000006103405261032060048060208461038001018260208501600060045af1505080518201915050336020826103800101526020810190503060208261038001015260208101905061028051602082610380010152602081019050806103805261038090508051602001806104408284600060045af1156117ce57505060206105206104405161046060006102a0515af1156117ce5760203d808211156103d157806103d3565b815b90509050610500526105008051602001806102c08284600060045af1156117ce57505060006102c0511815610435576102c080602001516000825180602090136117ce57809190126117ce57806020036101000a8204905090509050156117ce575b60006004610320527fe8eda9df000000000000000000000000000000000000000000000000000000006103405261032060048060208461038001018260208501600060045af15050805182019150506102a051602082610380010152602081019050610280516020826103800101526020810190503060208261038001015260208101905061024051602082610380010152602081019050806103805261038090508051602001806104608284600060045af1156117ce57505060006000610460516104806000738dff5e27ea6b7ac08ebfdf9eb090f32ee9a30fcf5af1156117ce57610280516101c06102605160028082106117ce578082039050905060038110156117ce5760200201525b5b81516001018083528114156102c7575b50506009543b156117ce57600060006084634515cef3610260526101c051610280526101e0516102a052610200516102c05260a4356102e05261027c60006009545af1156117ce57600b5461026052602061032060246370a082316102a052306102c0526102bc610260515afa156117ce57601f3d11156117ce5760005061032051610280526020610340604463a9059cbb6102a052610140516102c052610280516102e0526102bc6000610260515af1156117ce57601f3d11156117ce5760005061034050005b6365b2489b81141561063157336101405261065c565b63e2ad025a8114156106575760843560a01c6117ce57602060846101403760005061065c565b610c26565b600060046101c0527f23b872dd000000000000000000000000000000000000000000000000000000006101e0526101c060048060208461022001018260208501600060045af15050805182019150503360208261022001015260208101905030602082610220010152602081019050604435602082610220010152602081019050806102205261022090508051602001806102e08284600060045af1156117ce57505060206103c06102e0516103006000600160043560058110156117ce5702600401545af1156117ce5760203d80821115610738578061073a565b815b905090506103a0526103a08051602001806101608284600060045af1156117ce575050600061016051181561079c5761016080602001516000825180602090136117ce57809190126117ce57806020036101000a8204905090509050156117ce575b6044356101c0526040366101e0376003602435106107cc5760243560028082106117ce5780820390509050610200525b6003600435101561085257606036610220376101c05161022060043560038110156117ce576020020152602061038060a4632b6e993a61028052610220516102a052610240516102c052610260516102e05260006103005260016103205261029c6000600a545af1156117ce57601f3d11156117ce57600050610380516101c05261095b565b60043560028082106117ce57808203905090506101e05260006004610220527fe8eda9df000000000000000000000000000000000000000000000000000000006102405261022060048060208461028001018260208501600060045af1505080518201915050600160043560058110156117ce5702600401546020826102800101526020810190506101c05160208261028001015260208101905030602082610280010152602081019050600054602082610280010152602081019050806102805261028090508051602001806103608284600060045af1156117ce57505060006000610360516103806000738dff5e27ea6b7ac08ebfdf9eb090f32ee9a30fcf5af1156117ce575b60006101e05161020051808210156109735780610975565b815b9050905011156109c5576009543b156117ce57600060006084635b41b908610220526101e0516102405261020051610260526101c0516102805260006102a05261023c60006009545af1156117ce575b60206102c060246370a0823161024052306102605261025c60016102005160038110156117ce5702600101545afa156117ce57601f3d11156117ce576000506102c0516102205261020051610ba5576020610320608463517a55a361024052610220516102605260243560405181116117ce57610280526064356102a05260016102c05261025c6000600a545af1156117ce57601f3d11156117ce57600050610320516102205260006004610240527fa9059cbb00000000000000000000000000000000000000000000000000000000610260526102406004806020846102a001018260208501600060045af1505080518201915050610140516020826102a0010152602081019050610220516020826102a0010152602081019050806102a0526102a090508051602001806103408284600060045af1156117ce5750506020610400610340516103606000600160243560058110156117ce5702600401545af1156117ce5760203d80821115610b3c5780610b3e565b815b905090506103e0526103e08051602001806101608284600060045af1156117ce5750506000610160511815610ba05761016080602001516000825180602090136117ce57809190126117ce57806020036101000a8204905090509050156117ce575b610c24565b60643561022051106117ce57738dff5e27ea6b7ac08ebfdf9eb090f32ee9a30fcf3b156117ce576000600060646369328dec61024052600160243560058110156117ce570260040154610260526102205161028052610140516102a05261025c6000738dff5e27ea6b7ac08ebfdf9eb090f32ee9a30fcf5af1156117ce575b005b63e3bff5ce811415610c3c573361014052610c67565b634f626a31811415610c625760c43560a01c6117ce57602060c461014037600050610c67565b611009565b602061022060646323b872dd610160523361018052306101a0526004356101c05261017c6000600b545af1156117ce57601f3d11156117ce57600050610220506000610160526084356101805260a4356101a0526009543b156117ce5760006000608463ecb586a56101c0526004356101e052610160516102005261018051610220526101a051610240526101dc60006009545af1156117ce57602061026060246370a082316101e05230610200526101fc6001545afa156117ce57601f3d11156117ce57600050610260516101c0526024356101e052604435610200526064356102205260606103a060a463fce647366102a0526101c0516102c0526101e0516102e052610200516103005261022051610320526001610340526102bc6000600a545af1156117ce57605f3d11156117ce576000506103a0805161024052806020015161026052806040015161028052506102a060006003818352015b60006004610320527fa9059cbb000000000000000000000000000000000000000000000000000000006103405261032060048060208461038001018260208501600060045af1505080518201915050610140516020826103800101526020810190506102406102a05160038110156117ce576020020151602082610380010152602081019050806103805261038090508051602001806104208284600060045af1156117ce57505060206104e061042051610440600060016102a05160058110156117ce5702600401545af1156117ce5760203d80821115610ea75780610ea9565b815b905090506104c0526104c08051602001806102c08284600060045af1156117ce57505060006102c0511815610f0b576102c080602001516000825180602090136117ce57809190126117ce57806020036101000a8204905090509050156117ce575b5b8151600101808352811415610dc5575b50506102a060036002818352015b602061034060246370a082316102c052306102e0526102dc60016102a05160028082106117ce578082039050905060038110156117ce5702600101545afa156117ce57601f3d11156117ce57600050610340516101c052738dff5e27ea6b7ac08ebfdf9eb090f32ee9a30fcf3b156117ce576000600060646369328dec6102c05260016102a05160058110156117ce5702600401546102e0526101c0516103005261014051610320526102dc6000738dff5e27ea6b7ac08ebfdf9eb090f32ee9a30fcf5af1156117ce575b8151600101808352811415610f2a575b5050005b63f1dc3cc981141561101f57336101405261104a565b630fbcee6e8114156110455760643560a01c6117ce57602060646101403760005061104a565b61134f565b602061022060646323b872dd610160523361018052306101a0526004356101c05261017c6000600b545af1156117ce57601f3d11156117ce57600050610220506000610160526003602435106110b25760243560028082106117ce5780820390509050610160525b6009543b156117ce5760006000606463f1dc3cc9610180526004356101a052610160516101c05260006101e05261019c60006009545af1156117ce57602061022060246370a082316101a052306101c0526101bc60016101605160038110156117ce5702600101545afa156117ce57601f3d11156117ce576000506102205161018052610160516112ce576020610280608463517a55a36101a052610180516101c05260243560405181116117ce576101e052604435610200526001610220526101bc6000600a545af1156117ce57601f3d11156117ce57600050610280516101805260006004610200527fa9059cbb000000000000000000000000000000000000000000000000000000006102205261020060048060208461026001018260208501600060045af15050805182019150506101405160208261026001015260208101905061018051602082610260010152602081019050806102605261026090508051602001806103008284600060045af1156117ce57505060206103c0610300516103206000600160243560058110156117ce5702600401545af1156117ce5760203d808211156112655780611267565b815b905090506103a0526103a08051602001806101a08284600060045af1156117ce57505060006101a05118156112c9576101a080602001516000825180602090136117ce57809190126117ce57806020036101000a8204905090509050156117ce575b61134d565b60443561018051106117ce57738dff5e27ea6b7ac08ebfdf9eb090f32ee9a30fcf3b156117ce576000600060646369328dec6101a052600160243560058110156117ce5702600401546101c052610180516101e05261014051610200526101bc6000738dff5e27ea6b7ac08ebfdf9eb090f32ee9a30fcf5af1156117ce575b005b6385f11d1e811415611552576003600435602435808210156113715780611373565b815b9050905010156113da5760206102006064635e0d443f6101405260043560405181116117ce576101605260243560405181116117ce57610180526044356101a05261015c600a545afa156117ce57601f3d11156117ce576000506102005160005260206000f35b604435610140526040366101603760036024351061140a5760243560028082106117ce5780820390509050610180525b60036004351015611488576060366101a037610140516101a060043560038110156117ce57602002015260206102e06084633883e119610200526101a051610220526101c051610240526101e0516102605260016102805261021c600a545afa156117ce57601f3d11156117ce576000506102e051610140526114a0565b60043560028082106117ce5780820390509050610160525b6020610280606463556d6e9f6101c052610160516101e052610180516102005261014051610220526101dc6009545afa156117ce57601f3d11156117ce57600050610280516101a05261018051611543576020610260604463cc2b27d76101c0526101a0516101e05260243560405181116117ce57610200526101dc600a545afa156117ce57601f3d11156117ce576000506102605160005260206000f3611550565b6101a05160005260206000f35b005b637ede89c58114156116385760a43560011c6117ce5760043561014052602435610160526044356101805260206102a06084633883e1196101c052610140516101e0526101605161020052610180516102205260a435610240526101dc600a545afa156117ce57601f3d11156117ce576000506102a0516101a0526101a0516101c0526064356101e0526084356102005260206103006084633883e119610220526101c051610240526101e05161026052610200516102805260a4356102a05261023c6009545afa156117ce57601f3d11156117ce576000506103005160005260206000f35b634fb08c5e81141561172e576003602435106116a25760206101e06044634fb08c5e610140526004356101605260243560028082106117ce57808203905090506101805261015c6009545afa156117ce57601f3d11156117ce576000506101e05160005260206000f35b60206102006044634fb08c5e610160526004356101805260006101a05261017c6009545afa156117ce57601f3d11156117ce5760005061020051610140526020610200604463cc2b27d761016052610140516101805260243560405181116117ce576101a05261017c600a545afa156117ce57601f3d11156117ce576000506102005160005260206000f35b63c661065781141561175657600160043560038110156117ce57026001015460005260206000f35b63b9947eb081141561177e57600160043560058110156117ce57026004015460005260206000f35b6316f0115b8114156117965760095460005260206000f35b635d6362bb8114156117ae57600a5460005260206000f35b63fc0c546a8114156117c657600b5460005260206000f35b505b60006000fd5b600080fd5b6105a9611d7c036105a96000396105a9611d7c036000f35b600080fd00000000000000000000000092577943c7ac4accb35288ab2cc84d75fec330af000000000000000000000000445fe580ef8d70ff569ab36e80c647af338db351",
}

// VyperABI is the input ABI used to generate the binding from.
// Deprecated: Use VyperMetaData.ABI instead.
var VyperABI = VyperMetaData.ABI

// VyperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VyperMetaData.Bin instead.
var VyperBin = VyperMetaData.Bin

// DeployVyper deploys a new Ethereum contract, binding an instance of Vyper to it.
func DeployVyper(auth *bind.TransactOpts, backend bind.ContractBackend, _pool common.Address, _base_pool common.Address) (common.Address, *types.Transaction, *Vyper, error) {
	parsed, err := VyperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VyperBin), backend, _pool, _base_pool)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vyper{VyperCaller: VyperCaller{contract: contract}, VyperTransactor: VyperTransactor{contract: contract}, VyperFilterer: VyperFilterer{contract: contract}}, nil
}

// Vyper is an auto generated Go binding around an Ethereum contract.
type Vyper struct {
	VyperCaller     // Read-only binding to the contract
	VyperTransactor // Write-only binding to the contract
	VyperFilterer   // Log filterer for contract events
}

// VyperCaller is an auto generated read-only Go binding around an Ethereum contract.
type VyperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VyperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VyperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VyperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VyperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VyperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VyperSession struct {
	Contract     *Vyper            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VyperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VyperCallerSession struct {
	Contract *VyperCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VyperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VyperTransactorSession struct {
	Contract     *VyperTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VyperRaw is an auto generated low-level Go binding around an Ethereum contract.
type VyperRaw struct {
	Contract *Vyper // Generic contract binding to access the raw methods on
}

// VyperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VyperCallerRaw struct {
	Contract *VyperCaller // Generic read-only contract binding to access the raw methods on
}

// VyperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VyperTransactorRaw struct {
	Contract *VyperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVyper creates a new instance of Vyper, bound to a specific deployed contract.
func NewVyper(address common.Address, backend bind.ContractBackend) (*Vyper, error) {
	contract, err := bindVyper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vyper{VyperCaller: VyperCaller{contract: contract}, VyperTransactor: VyperTransactor{contract: contract}, VyperFilterer: VyperFilterer{contract: contract}}, nil
}

// NewVyperCaller creates a new read-only instance of Vyper, bound to a specific deployed contract.
func NewVyperCaller(address common.Address, caller bind.ContractCaller) (*VyperCaller, error) {
	contract, err := bindVyper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VyperCaller{contract: contract}, nil
}

// NewVyperTransactor creates a new write-only instance of Vyper, bound to a specific deployed contract.
func NewVyperTransactor(address common.Address, transactor bind.ContractTransactor) (*VyperTransactor, error) {
	contract, err := bindVyper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VyperTransactor{contract: contract}, nil
}

// NewVyperFilterer creates a new log filterer instance of Vyper, bound to a specific deployed contract.
func NewVyperFilterer(address common.Address, filterer bind.ContractFilterer) (*VyperFilterer, error) {
	contract, err := bindVyper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VyperFilterer{contract: contract}, nil
}

// bindVyper binds a generic wrapper to an already deployed contract.
func bindVyper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VyperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vyper *VyperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vyper.Contract.VyperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vyper *VyperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vyper.Contract.VyperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vyper *VyperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vyper.Contract.VyperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vyper *VyperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vyper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vyper *VyperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vyper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vyper *VyperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vyper.Contract.contract.Transact(opts, method, params...)
}

// BasePool is a free data retrieval call binding the contract method 0x5d6362bb.
//
// Solidity: function base_pool() view returns(address)
func (_Vyper *VyperCaller) BasePool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vyper.contract.Call(opts, &out, "base_pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BasePool is a free data retrieval call binding the contract method 0x5d6362bb.
//
// Solidity: function base_pool() view returns(address)
func (_Vyper *VyperSession) BasePool() (common.Address, error) {
	return _Vyper.Contract.BasePool(&_Vyper.CallOpts)
}

// BasePool is a free data retrieval call binding the contract method 0x5d6362bb.
//
// Solidity: function base_pool() view returns(address)
func (_Vyper *VyperCallerSession) BasePool() (common.Address, error) {
	return _Vyper.Contract.BasePool(&_Vyper.CallOpts)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x7ede89c5.
//
// Solidity: function calc_token_amount(uint256[5] _amounts, bool _is_deposit) view returns(uint256)
func (_Vyper *VyperCaller) CalcTokenAmount(opts *bind.CallOpts, _amounts [5]*big.Int, _is_deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _Vyper.contract.Call(opts, &out, "calc_token_amount", _amounts, _is_deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x7ede89c5.
//
// Solidity: function calc_token_amount(uint256[5] _amounts, bool _is_deposit) view returns(uint256)
func (_Vyper *VyperSession) CalcTokenAmount(_amounts [5]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _Vyper.Contract.CalcTokenAmount(&_Vyper.CallOpts, _amounts, _is_deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x7ede89c5.
//
// Solidity: function calc_token_amount(uint256[5] _amounts, bool _is_deposit) view returns(uint256)
func (_Vyper *VyperCallerSession) CalcTokenAmount(_amounts [5]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _Vyper.Contract.CalcTokenAmount(&_Vyper.CallOpts, _amounts, _is_deposit)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_Vyper *VyperCaller) CalcWithdrawOneCoin(opts *bind.CallOpts, token_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vyper.contract.Call(opts, &out, "calc_withdraw_one_coin", token_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_Vyper *VyperSession) CalcWithdrawOneCoin(token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _Vyper.Contract.CalcWithdrawOneCoin(&_Vyper.CallOpts, token_amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_Vyper *VyperCallerSession) CalcWithdrawOneCoin(token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _Vyper.Contract.CalcWithdrawOneCoin(&_Vyper.CallOpts, token_amount, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_Vyper *VyperCaller) Coins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vyper.contract.Call(opts, &out, "coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_Vyper *VyperSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _Vyper.Contract.Coins(&_Vyper.CallOpts, arg0)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 arg0) view returns(address)
func (_Vyper *VyperCallerSession) Coins(arg0 *big.Int) (common.Address, error) {
	return _Vyper.Contract.Coins(&_Vyper.CallOpts, arg0)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x85f11d1e.
//
// Solidity: function get_dy_underlying(uint256 i, uint256 j, uint256 _dx) view returns(uint256)
func (_Vyper *VyperCaller) GetDyUnderlying(opts *bind.CallOpts, i *big.Int, j *big.Int, _dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vyper.contract.Call(opts, &out, "get_dy_underlying", i, j, _dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x85f11d1e.
//
// Solidity: function get_dy_underlying(uint256 i, uint256 j, uint256 _dx) view returns(uint256)
func (_Vyper *VyperSession) GetDyUnderlying(i *big.Int, j *big.Int, _dx *big.Int) (*big.Int, error) {
	return _Vyper.Contract.GetDyUnderlying(&_Vyper.CallOpts, i, j, _dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x85f11d1e.
//
// Solidity: function get_dy_underlying(uint256 i, uint256 j, uint256 _dx) view returns(uint256)
func (_Vyper *VyperCallerSession) GetDyUnderlying(i *big.Int, j *big.Int, _dx *big.Int) (*big.Int, error) {
	return _Vyper.Contract.GetDyUnderlying(&_Vyper.CallOpts, i, j, _dx)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Vyper *VyperCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vyper.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Vyper *VyperSession) Pool() (common.Address, error) {
	return _Vyper.Contract.Pool(&_Vyper.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Vyper *VyperCallerSession) Pool() (common.Address, error) {
	return _Vyper.Contract.Pool(&_Vyper.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Vyper *VyperCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vyper.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Vyper *VyperSession) Token() (common.Address, error) {
	return _Vyper.Contract.Token(&_Vyper.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Vyper *VyperCallerSession) Token() (common.Address, error) {
	return _Vyper.Contract.Token(&_Vyper.CallOpts)
}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 arg0) view returns(address)
func (_Vyper *VyperCaller) UnderlyingCoins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vyper.contract.Call(opts, &out, "underlying_coins", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 arg0) view returns(address)
func (_Vyper *VyperSession) UnderlyingCoins(arg0 *big.Int) (common.Address, error) {
	return _Vyper.Contract.UnderlyingCoins(&_Vyper.CallOpts, arg0)
}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 arg0) view returns(address)
func (_Vyper *VyperCallerSession) UnderlyingCoins(arg0 *big.Int) (common.Address, error) {
	return _Vyper.Contract.UnderlyingCoins(&_Vyper.CallOpts, arg0)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x84738499.
//
// Solidity: function add_liquidity(uint256[5] _amounts, uint256 _min_mint_amount) returns()
func (_Vyper *VyperTransactor) AddLiquidity(opts *bind.TransactOpts, _amounts [5]*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Vyper.contract.Transact(opts, "add_liquidity", _amounts, _min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x84738499.
//
// Solidity: function add_liquidity(uint256[5] _amounts, uint256 _min_mint_amount) returns()
func (_Vyper *VyperSession) AddLiquidity(_amounts [5]*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Vyper.Contract.AddLiquidity(&_Vyper.TransactOpts, _amounts, _min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x84738499.
//
// Solidity: function add_liquidity(uint256[5] _amounts, uint256 _min_mint_amount) returns()
func (_Vyper *VyperTransactorSession) AddLiquidity(_amounts [5]*big.Int, _min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Vyper.Contract.AddLiquidity(&_Vyper.TransactOpts, _amounts, _min_mint_amount)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xcf2b51b8.
//
// Solidity: function add_liquidity(uint256[5] _amounts, uint256 _min_mint_amount, address _receiver) returns()
func (_Vyper *VyperTransactor) AddLiquidity0(opts *bind.TransactOpts, _amounts [5]*big.Int, _min_mint_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Vyper.contract.Transact(opts, "add_liquidity0", _amounts, _min_mint_amount, _receiver)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xcf2b51b8.
//
// Solidity: function add_liquidity(uint256[5] _amounts, uint256 _min_mint_amount, address _receiver) returns()
func (_Vyper *VyperSession) AddLiquidity0(_amounts [5]*big.Int, _min_mint_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Vyper.Contract.AddLiquidity0(&_Vyper.TransactOpts, _amounts, _min_mint_amount, _receiver)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0xcf2b51b8.
//
// Solidity: function add_liquidity(uint256[5] _amounts, uint256 _min_mint_amount, address _receiver) returns()
func (_Vyper *VyperTransactorSession) AddLiquidity0(_amounts [5]*big.Int, _min_mint_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Vyper.Contract.AddLiquidity0(&_Vyper.TransactOpts, _amounts, _min_mint_amount, _receiver)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0x65b2489b.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 _dx, uint256 _min_dy) returns()
func (_Vyper *VyperTransactor) ExchangeUnderlying(opts *bind.TransactOpts, i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _Vyper.contract.Transact(opts, "exchange_underlying", i, j, _dx, _min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0x65b2489b.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 _dx, uint256 _min_dy) returns()
func (_Vyper *VyperSession) ExchangeUnderlying(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _Vyper.Contract.ExchangeUnderlying(&_Vyper.TransactOpts, i, j, _dx, _min_dy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0x65b2489b.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 _dx, uint256 _min_dy) returns()
func (_Vyper *VyperTransactorSession) ExchangeUnderlying(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int) (*types.Transaction, error) {
	return _Vyper.Contract.ExchangeUnderlying(&_Vyper.TransactOpts, i, j, _dx, _min_dy)
}

// ExchangeUnderlying0 is a paid mutator transaction binding the contract method 0xe2ad025a.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 _dx, uint256 _min_dy, address _receiver) returns()
func (_Vyper *VyperTransactor) ExchangeUnderlying0(opts *bind.TransactOpts, i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Vyper.contract.Transact(opts, "exchange_underlying0", i, j, _dx, _min_dy, _receiver)
}

// ExchangeUnderlying0 is a paid mutator transaction binding the contract method 0xe2ad025a.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 _dx, uint256 _min_dy, address _receiver) returns()
func (_Vyper *VyperSession) ExchangeUnderlying0(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Vyper.Contract.ExchangeUnderlying0(&_Vyper.TransactOpts, i, j, _dx, _min_dy, _receiver)
}

// ExchangeUnderlying0 is a paid mutator transaction binding the contract method 0xe2ad025a.
//
// Solidity: function exchange_underlying(uint256 i, uint256 j, uint256 _dx, uint256 _min_dy, address _receiver) returns()
func (_Vyper *VyperTransactorSession) ExchangeUnderlying0(i *big.Int, j *big.Int, _dx *big.Int, _min_dy *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Vyper.Contract.ExchangeUnderlying0(&_Vyper.TransactOpts, i, j, _dx, _min_dy, _receiver)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xe3bff5ce.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[5] _min_amounts) returns()
func (_Vyper *VyperTransactor) RemoveLiquidity(opts *bind.TransactOpts, _amount *big.Int, _min_amounts [5]*big.Int) (*types.Transaction, error) {
	return _Vyper.contract.Transact(opts, "remove_liquidity", _amount, _min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xe3bff5ce.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[5] _min_amounts) returns()
func (_Vyper *VyperSession) RemoveLiquidity(_amount *big.Int, _min_amounts [5]*big.Int) (*types.Transaction, error) {
	return _Vyper.Contract.RemoveLiquidity(&_Vyper.TransactOpts, _amount, _min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xe3bff5ce.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[5] _min_amounts) returns()
func (_Vyper *VyperTransactorSession) RemoveLiquidity(_amount *big.Int, _min_amounts [5]*big.Int) (*types.Transaction, error) {
	return _Vyper.Contract.RemoveLiquidity(&_Vyper.TransactOpts, _amount, _min_amounts)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x4f626a31.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[5] _min_amounts, address _receiver) returns()
func (_Vyper *VyperTransactor) RemoveLiquidity0(opts *bind.TransactOpts, _amount *big.Int, _min_amounts [5]*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Vyper.contract.Transact(opts, "remove_liquidity0", _amount, _min_amounts, _receiver)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x4f626a31.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[5] _min_amounts, address _receiver) returns()
func (_Vyper *VyperSession) RemoveLiquidity0(_amount *big.Int, _min_amounts [5]*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Vyper.Contract.RemoveLiquidity0(&_Vyper.TransactOpts, _amount, _min_amounts, _receiver)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0x4f626a31.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[5] _min_amounts, address _receiver) returns()
func (_Vyper *VyperTransactorSession) RemoveLiquidity0(_amount *big.Int, _min_amounts [5]*big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Vyper.Contract.RemoveLiquidity0(&_Vyper.TransactOpts, _amount, _min_amounts, _receiver)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, uint256 i, uint256 _min_amount) returns()
func (_Vyper *VyperTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, _token_amount *big.Int, i *big.Int, _min_amount *big.Int) (*types.Transaction, error) {
	return _Vyper.contract.Transact(opts, "remove_liquidity_one_coin", _token_amount, i, _min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, uint256 i, uint256 _min_amount) returns()
func (_Vyper *VyperSession) RemoveLiquidityOneCoin(_token_amount *big.Int, i *big.Int, _min_amount *big.Int) (*types.Transaction, error) {
	return _Vyper.Contract.RemoveLiquidityOneCoin(&_Vyper.TransactOpts, _token_amount, i, _min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, uint256 i, uint256 _min_amount) returns()
func (_Vyper *VyperTransactorSession) RemoveLiquidityOneCoin(_token_amount *big.Int, i *big.Int, _min_amount *big.Int) (*types.Transaction, error) {
	return _Vyper.Contract.RemoveLiquidityOneCoin(&_Vyper.TransactOpts, _token_amount, i, _min_amount)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x0fbcee6e.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, uint256 i, uint256 _min_amount, address _receiver) returns()
func (_Vyper *VyperTransactor) RemoveLiquidityOneCoin0(opts *bind.TransactOpts, _token_amount *big.Int, i *big.Int, _min_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Vyper.contract.Transact(opts, "remove_liquidity_one_coin0", _token_amount, i, _min_amount, _receiver)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x0fbcee6e.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, uint256 i, uint256 _min_amount, address _receiver) returns()
func (_Vyper *VyperSession) RemoveLiquidityOneCoin0(_token_amount *big.Int, i *big.Int, _min_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Vyper.Contract.RemoveLiquidityOneCoin0(&_Vyper.TransactOpts, _token_amount, i, _min_amount, _receiver)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x0fbcee6e.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, uint256 i, uint256 _min_amount, address _receiver) returns()
func (_Vyper *VyperTransactorSession) RemoveLiquidityOneCoin0(_token_amount *big.Int, i *big.Int, _min_amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Vyper.Contract.RemoveLiquidityOneCoin0(&_Vyper.TransactOpts, _token_amount, i, _min_amount, _receiver)
}
