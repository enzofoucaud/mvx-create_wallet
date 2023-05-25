package mvx_api

/*
	API Error Response
*/

type ErrorMvx struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

/*
	Accounts
*/

type Accounts struct {
	Address                  string `json:"address"`
	Nonce                    int    `json:"nonce"`
	Balance                  string `json:"balance"`
	RootHash                 string `json:"rootHash"`
	TxCount                  int    `json:"txCount"`
	ScrCount                 int    `json:"scrCount"`
	Username                 string `json:"username"`
	Shard                    int    `json:"shard"`
	DeveloperReward          string `json:"developerReward"`
	OwnerAddress             string `json:"ownerAddress"`
	DeployedAt               int    `json:"deployedAt"`
	IsUpgradeable            bool   `json:"isUpgradeable"`
	IsReadable               bool   `json:"isReadable"`
	IsPayable                bool   `json:"isPayable"`
	IsPayableBySmartContract bool   `json:"isPayableBySmartContract"`
}
