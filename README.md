# ASBG 수료 인증서 시스템

![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)
![Network: Sepolia](https://img.shields.io/badge/Network-Sepolia-blue)
![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go)
![Solidity](https://img.shields.io/badge/Solidity-0.8.24-363636?logo=solidity)

AWS Student Builders Group(ASBG) 동아리의 출석을 블록체인에 기록하고,
NFT 수료 인증서를 자동 발급하는 시스템입니다.

- **네트워크**: Ethereum Sepolia 테스트넷 전용
- **목적**: 유틸리티·기록 목적
- **인원**: 20명 이내

---

## 아키텍처

```
운영진 브라우저
  └─ HTTPS ──► CloudFront ──► S3 (dashboard.html, certificate.html)
                    │
                    └─ API 요청 ──► API Gateway (asbg-api)
                                        │
                          ┌─────────────┴──────────────┐
                          ▼                            ▼
                 Lambda (asbg-dashboard)      Lambda (asbg-job)
                 GET /members                 수료 기준 확인 후
                 GET /certificates            NFT 인증서 자동 발급
                 POST /mint
                          │
                          ▼
                 AWS SSM Parameter Store
                 (PRIVATE_KEY 등 민감 정보)
                          │
           ┌──────────────┼──────────────┐
           ▼              ▼              ▼
     DynamoDB          S3 Bucket     Pinata IPFS
   (asbg-tx-history) (asbg-data)   (NFT 메타데이터)
                          │
                          ▼
                  Ethereum Sepolia
                  (Alchemy RPC)
```

---

## 기술 스택

| 항목 | 내용 |
|---|---|
| 언어 | Go 1.22+ |
| 블록체인 라이브러리 | go-ethereum v1.15+ |
| 스마트 컨트랙트 | Solidity 0.8.24 + OpenZeppelin 5.x |
| 컨트랙트 개발 / 테스트 | Foundry (forge, cast, anvil) |
| ABI 바인딩 | abigen |
| 테스트넷 | Ethereum Sepolia (Chain ID: 11155111) |
| RPC | Alchemy Sepolia |
| NFT 메타데이터 | IPFS (Pinata API) |
| 클라우드 | AWS 서버리스 (Lambda, API Gateway, S3, CloudFront, DynamoDB, SSM Parameter Store) |
| IaC | AWS SAM |
| CI/CD | GitHub Actions (OIDC) |

---

## 스마트 컨트랙트

### AttendanceTracker

출석을 기록하는 컨트랙트입니다. ERC-20이 아닌 순수 기록용으로, 멤버 ID(CSV 행 순서)를 키로 사용합니다.

- 배포 주소: `0xd6fAA528670178D2C27883772ef265C9314e3c96`
- `mintAttendance(memberIds[], eventId)` — 출석자 일괄 기록 (MINTER_ROLE 전용)
- `attendance(memberId)` — 멤버 출석 횟수 조회

### CertificateNFT (ERC-721)

수료 인증서 NFT 컨트랙트입니다. 모든 NFT는 운영진 지갑에 발행되며, 멤버 지갑은 불필요합니다.

- 배포 주소: `0xF5E07d95461f076BF7567666eAf354ba4460953a`
- 수료 기준: `attendance(memberId) >= ceil(totalSessions × 0.9)`
- `issueCertificate(memberId, ipfsCid, memberName)` — 인증서 발행 (ISSUER_ROLE 전용)
- `hasCertificate(memberId)` — 수료 여부 조회
- `tokenURI(tokenId)` — `ipfs://CID` 형식의 메타데이터 URI 반환

### 수료 기준

```
threshold = ceil(TOTAL_SESSIONS × 0.9)

예시: TOTAL_SESSIONS=10 → threshold=9 (1회 결석 허용)
```

---

## 디렉토리 구조

```
.
├── cmd/
│   ├── dashboard/      # Lambda: GET /members, GET /certificates, POST /mint
│   ├── job/            # Lambda: 수료 기준 달성자 인증서 자동 발급
│   ├── deploy/         # 로컬 CLI: 스마트 컨트랙트 배포
│   └── issue/          # 로컬 CLI: 수료 인증서 수동 발급 (데모용)
├── contracts/
│   ├── src/
│   │   ├── AttendanceTracker.sol
│   │   └── CertificateNFT.sol
│   └── test/
├── bindings/           # abigen 자동 생성 — 직접 수정 금지
│   ├── attendance_tracker/
│   └── certificate_nft/
├── dashboard/
│   ├── dashboard.html  # 출석 체크 UI
│   └── certificate.html # 수료 인증서 조회 UI
├── internal/
│   ├── blockchain/     # go-ethereum 클라이언트 래퍼
│   ├── ipfs/           # Pinata 메타데이터 업로드
│   ├── members/        # members.csv 로드 (로컬 / S3)
│   ├── metadata/       # NFT 메타데이터 빌더
│   ├── secret/         # AWS SSM Parameter Store 조회
│   └── store/          # DynamoDB tx_history 기록
├── infra/sam/
│   ├── template.yaml   # SAM 인프라 정의
│   └── samconfig.toml  # SAM 배포 설정
├── .github/workflows/
│   └── deploy.yml      # CI/CD (OIDC → AWS Lambda 자동 배포)
├── members.csv         # 멤버 목록 (name, role)
└── .env
```

---

## AWS 리소스

| 리소스 | 이름 | 역할 |
|---|---|---|
| S3 Bucket | `asbg-static` | dashboard.html, certificate.html 호스팅 |
| S3 Bucket | `asbg-data` | members.csv, 결과 파일 저장 |
| CloudFront | — | HTTPS CDN, API 라우팅 |
| API Gateway | `asbg-api` | HTTP API |
| Lambda | `asbg-dashboard` | 출석 민팅 API |
| Lambda | `asbg-job` | 수료 확인 + 인증서 발행 |
| EventBridge | `asbg-yearly` | 매년 12월 31일 Job 자동 실행 |
| SSM Parameter Store | `asbg/*` | 민감 환경변수 |
| DynamoDB | `asbg-tx-history` | 트랜잭션 로그 |
| IAM Role | `asbg-lambda-role` | Lambda 실행 권한 |

---

## 환경 설정

### `.env` (로컬 전용, git 커밋 금지)

```env
RPC_URL=https://eth-sepolia.g.alchemy.com/v2/<ALCHEMY_KEY>
CLOUDFRONT_URL=https://<CF_DISTRIBUTION>.cloudfront.net

PRIVATE_KEY=<운영진_지갑_프라이빗_키>
DEPLOYER_ADDRESS=<운영진_지갑_주소>

ATTENDANCE_TRACKER_ADDRESS=<배포된_컨트랙트_주소>
CERTIFICATE_NFT_ADDRESS=<배포된_컨트랙트_주소>

TOTAL_SESSIONS=<세션 횟수>

PINATA_JWT=<Pinata_JWT>
```

### AWS SSM Parameter Store (Lambda 런타임 주입)

| 키 | 설명 |
|---|---|
| `asbg/PRIVATE_KEY` | 운영진 지갑 프라이빗 키 |
| `asbg/RPC_URL` | Alchemy Sepolia RPC URL |
| `asbg/PINATA_JWT` | Pinata IPFS JWT |
| `asbg/ATTENDANCE_TRACKER_ADDRESS` | 배포된 AttendanceTracker 주소 |
| `asbg/CERTIFICATE_NFT_ADDRESS` | 배포된 CertificateNFT 주소 |

---

## 시작하기

### 1. 컨트랙트 테스트

```bash
cd contracts
forge test -vv
```

### 2. 컨트랙트 배포 (최초 1회)

```bash
# .env 작성 후 실행
go run cmd/deploy/main.go
# → .env에 ATTENDANCE_TRACKER_ADDRESS, CERTIFICATE_NFT_ADDRESS 기입
```

### 3. AWS 인프라 배포 (최초 1회)

```bash
# Lambda 바이너리 빌드
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap ./cmd/dashboard/main.go
zip dashboard.zip bootstrap && rm bootstrap

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap ./cmd/job/main.go
zip job.zip bootstrap && rm bootstrap

# SAM 배포
cd infra/sam
sam deploy --guided
```

첫 배포 후 `samconfig.toml`의 `CloudFrontURL` 값을 Outputs에서 확인한 뒤 재배포:

```toml
parameter_overrides = "TotalSessions=\"10\" CloudFrontURL=\"https://xxxx.cloudfront.net\""
```

### 4. GitHub Actions 설정

GitHub Repository → Settings → Secrets에 등록:

| Secret | 설명 |
|---|---|
| `AWS_ROLE_ARN` | GitHub Actions OIDC용 IAM Role ARN |
| `CF_DISTRIBUTION_ID` | CloudFront Distribution ID (캐시 무효화용, 선택) |

### 5. members.csv → S3 업로드

```bash
aws s3 cp members.csv s3://asbg-data/members.csv --region ap-northeast-2
```

### 6. 코드 배포 (이후 자동)

`main` 브랜치에 push하면 GitHub Actions가 자동으로 Lambda 배포 + S3 정적 파일 업로드를 수행합니다.

---

## 사용 방법

### 출석 체크 (매 모임마다)

1. `https://<CF_URL>/dashboard.html` 접속
2. 모든 멤버 출석 / 결석 체크
3. `event_id` 입력 (예: `2026-05-26_5월-정기-모임`)
4. 제출 → 출석자 전원 일괄 민팅 + 수료 기준 달성자 자동 인증서 발행

### 수료 인증서 조회

`https://<CF_URL>/certificate.html`

- Leader / Core Member / Member 구역별 카드 표시
- 카드 클릭 → 상세 정보 + LinkedIn 추가 버튼 + Etherscan 검증 링크
- 개인 공유 URL: `certificate.html?member=<ID>`

### abigen 바인딩 재생성 (컨트랙트 수정 시)

```bash
cd contracts && forge build

abigen --abi out/AttendanceTracker.sol/AttendanceTracker.abi.json \
       --bin out/AttendanceTracker.sol/AttendanceTracker.bin \
       --pkg attendance_tracker \
       --out ../bindings/attendance_tracker/attendance_tracker.go

abigen --abi out/CertificateNFT.sol/CertificateNFT.abi.json \
       --bin out/CertificateNFT.sol/CertificateNFT.bin \
       --pkg certificate_nft \
       --out ../bindings/certificate_nft/certificate_nft.go
```

> `bindings/` 폴더는 abigen으로만 재생성. 직접 수정 금지.

---

## 인증서 확인 수단

| 방법 | 설명 |
|---|---|
| 공유 URL | `certificate.html?member=<ID>` — 누구나 접근 가능한 수료증 페이지 |
| Etherscan | `https://sepolia.etherscan.io/token/<CONTRACT>?a=<TOKEN_ID>` — 온체인 영구 기록 |
| LinkedIn | 인증서 페이지의 **LinkedIn에 추가** 버튼으로 자격증 항목 자동 입력 |

---

## NFT 메타데이터 구조

```json
{
  "name": "ASBG 수료 인증서",
  "description": "AWS Student Builders Group 정규 과정을 수료하였음을 증명합니다.",
  "image": "ipfs://<IMAGE_CID>",
  "attributes": [
    { "trait_type": "certificate_name", "value": "ASBG 수료 인증서" },
    { "trait_type": "role",             "value": "Core Member" },
    { "trait_type": "issuer",           "value": "AWS Student Builders Group" },
    { "trait_type": "issued_date",      "value": "2026-05-26" },
    { "trait_type": "expiry",           "value": "영구 (만료 없음)" },
    { "trait_type": "credential_url",   "value": "https://<CF_URL>/certificate.html?member=0" },
    { "trait_type": "attendance",       "value": 9 },
    { "trait_type": "total_sessions",   "value": 10 }
  ]
}
```

---

## CI/CD 흐름

```
main 브랜치에 push
  └─ 1. forge test -vv         (컨트랙트 테스트)
  └─ 2. go build ./...         (Go 빌드 확인)
  └─ 3. GOOS=linux go build    (Lambda 바이너리 빌드)
  └─ 4. aws lambda update-function-code  (Lambda 배포)
  └─ 5. aws s3 sync dashboard/ (정적 파일 업로드)
  └─ 6. CloudFront 캐시 무효화 (CF_DISTRIBUTION_ID 설정 시)
```

---

## 주의사항

- 프라이빗 키를 코드·로그·커밋에 절대 포함하지 않습니다
- `.env` 파일은 git에 커밋하지 않습니다
- Lambda에 `.env` 파일을 직접 업로드하지 않습니다 (SSM Parameter Store 사용)
- `bindings/` 폴더를 직접 수정하지 않습니다 (abigen 재생성)
- 메인넷 배포를 하지 않습니다 (Sepolia 전용)
- `members.csv`에 없는 멤버에게 인증서를 발급하지 않습니다
