package server

type Tier string

const (
	TierFree Tier = "free"
	TierPro  Tier = "pro"
)

type Limits struct {
	Tier        Tier
	Description string
}

func LimitsFor(tier string) Limits {
	if tier == "pro" {
		return Limits{Tier: TierPro, Description: "Unlimited accounts and transactions"}
	}
	return Limits{Tier: TierFree, Description: "2 accounts, 500 transactions"}
}

func (l Limits) IsPro() bool {
	return l.Tier == TierPro
}
