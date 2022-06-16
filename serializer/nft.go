package serializer

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"spike-blockchain-server/model"
)

type NFT struct {
	TokenAddress string `json:"token_address"`
	TokenId      string `json:"token_id"`
	Amount       string `json:"amount"`
	ContractType string `json:"contract_type"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	TokenUri     string `json:"token_uri"`
	Metadata     string `json:"metadata"`
}

type MoralisNFTs struct {
	Total  int    `json:"total"`
	Cursor string `json:"cursor"`
	Result []NFT  `json:"result"`
}

func BuildNFTsResponse(nfts MoralisNFTs) MoralisNFTs {
	client := resty.New()

	var res []NFT
	for _, nft := range nfts.Result {
		resp, err := client.R().Get(nft.TokenUri)
		if err != nil {

		}

		if resp.IsError() {

		}

		if resp.IsSuccess() {
			var metadata model.SpikeMetadata
			err := json.Unmarshal(resp.Body(), &metadata)
			if err != nil {

			}
			if metadata.SpikeInfo.Validate() {
				res = append(res, nft)
			}
		}
	}
	return MoralisNFTs{
		Total:  len(res),
		Cursor: nfts.Cursor,
		Result: res,
	}
}
