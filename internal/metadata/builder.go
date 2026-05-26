package metadata

import "math"

type CertParams struct {
	MemberName      string
	Role            string
	IssuedDate      string
	CredentialURL   string
	MemberAddress   string
	Attendance      int
	TotalSessions   int
	ContractAddress string
	TokenID         uint64
}

func CalcThreshold(totalSessions int) int {
	return int(math.Ceil(float64(totalSessions) * 0.9))
}

func Build(p CertParams) map[string]any {
	return map[string]any{
		"name":        "ASBG 수료 인증서",
		"description": "AWS Student Builders Group 정규 과정을 수료하였음을 증명합니다.",
		"image":       "ipfs://",
		"attributes": []map[string]any{
			{"trait_type": "certificate_name", "value": "ASBG 수료 인증서"},
			{"trait_type": "role", "value": p.Role},
			{"trait_type": "issuer", "value": "AWS Student Builders Group"},
			{"trait_type": "issued_date", "value": p.IssuedDate},
			{"trait_type": "expiry", "value": "영구 (만료 없음)"},
			{"trait_type": "credential_url", "value": p.CredentialURL},
			{"trait_type": "member_address", "value": p.MemberAddress},
			{"trait_type": "attendance", "value": p.Attendance},
			{"trait_type": "total_sessions", "value": p.TotalSessions},
		},
	}
}
