package staking

// Description
type Description struct {
	Moniker         string `bson:"moniker"`
	Identity        string `bson:"identity"`
	Website         string `bson:"website"`
	SecurityContact string `bson:"security_contact"`
	Details         string `bson:"details"`
}

type CommissionRates struct {
	Rate          string `bson:"rate"`            // the commission rate charged to delegators
	MaxRate       string `bson:"max_rate"`        // maximum commission rate which validator can ever charge
	MaxChangeRate string `bson:"max_change_rate"` // maximum daily increase of the validator commission
}
