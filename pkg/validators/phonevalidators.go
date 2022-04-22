package validators

import (
	"go.jumia.org/customers/app/schema"
	"regexp"
)

var (
	MatchCameroon   = regexp.MustCompile(`^\(237\)`)
	MatchEthiopia   = regexp.MustCompile(`^\(251\)`)
	MatchMorocco    = regexp.MustCompile(`^\(212\)`)
	MatchMozambique = regexp.MustCompile(`^\(258\)`)
	MatchUganda     = regexp.MustCompile(`^\(256\)`)

	ValidCameroonPhone   = regexp.MustCompile(`\(237\)\ ?[2368]\d{7,8}$`)
	ValidEthiopiaPhone   = regexp.MustCompile(`\(251\)\ ?[1-59]\d{8}$`)
	ValidMoroccoPhone    = regexp.MustCompile(`\(212\)\ ?[5-9]\d{8}$`)
	ValidMozambiquePhone = regexp.MustCompile(`\(258\)\ ?[28]\d{7,8}$`)
	ValidUgandaPhone     = regexp.MustCompile(`\(256\)\ ?\d{9}$`)
)

func CheckCountry(phoneNumber string) schema.Country {

	if MatchCameroon.MatchString(phoneNumber) {
		return schema.Cameroon
	} else if MatchEthiopia.MatchString(phoneNumber) {
		return schema.Ethiopia
	} else if MatchMorocco.MatchString(phoneNumber) {
		return schema.Morocco
	} else if MatchMozambique.MatchString(phoneNumber) {
		return schema.Mozambique
	} else if MatchUganda.MatchString(phoneNumber) {
		return schema.Uganda
	}
	return ""
}

func CheckCountryCode(phoneNumber string) schema.Country {

	if MatchCameroon.MatchString(phoneNumber) {
		return "+237"
	} else if MatchEthiopia.MatchString(phoneNumber) {
		return "+251"
	} else if MatchMorocco.MatchString(phoneNumber) {
		return "+212"
	} else if MatchMozambique.MatchString(phoneNumber) {
		return "+258"
	} else if MatchUganda.MatchString(phoneNumber) {
		return "+256"
	}
	return ""
}

func CheckPhoneValidity(phoneNumber string) bool {
	country := CheckCountry(phoneNumber)

	switch country {
	case schema.Cameroon:
		return ValidCameroonPhone.MatchString(phoneNumber)
	case schema.Ethiopia:
		return ValidEthiopiaPhone.MatchString(phoneNumber)
	case schema.Morocco:
		return ValidMoroccoPhone.MatchString(phoneNumber)
	case schema.Mozambique:
		return ValidMozambiquePhone.MatchString(phoneNumber)
	case schema.Uganda:
		return ValidUgandaPhone.MatchString(phoneNumber)
	}
	return false
}
