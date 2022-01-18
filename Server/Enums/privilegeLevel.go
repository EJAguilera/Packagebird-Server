package Enums

type PrivilegeLevel int64

const (
	LowTier PrivilegeLevel = iota
	MidTier
	HighTier
	AdminTier
)

func IsTier(tier PrivilegeLevel) bool {
	switch tier {
	case LowTier:
		return true
	case MidTier:
		return true
	case HighTier:
		return true
	case AdminTier:
		return true
	}
	return false
}
