# HR 관리 스랙 봇

효율적인 직원 출석 관리 시스템입니다. Go언어로 개발되었으며, 간편한 체크인/체크아웃 기능과 상세한 출석 분석을 제공합니다.

## 주요 기능

- 직원 등록
- 실시간 체크인/체크아웃
- 근무 시간 자동 계산
- 초과 근무 및 지각 추적

## 기술 스택

- **언어:** Go 1.16+
- **웹 프레임워크:** Gin
- **ORM:** GORM
- **데이터베이스:** PostgreSQL

## 프로젝트 구조

```
HR_SlackBot/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── domain/
│   │   ├── employee/
│   │   └── attendance/
│   ├── application/
│   │   └── services/
│   ├── infrastructure/
│   │   ├── database/
│   │   └── repository/
│   └── interfaces/
│       └── http/
│       │    ├── handlers/
│       └── middleware/
├── pkg/
│   └── utils/
├── configs/
└── README.md
```

## 설치 및 실행 방법

1. 리포지토리 클론

   ```bash
   git clone https://github.com/Theoblanc/hr-slack-bot.git
   ```

2. 의존성 설치

   ```bash
   go mod tidy
   ```

3. 환경 변수 설정

   ```bash
   cp .env.example .env
   ```

   `.env` 파일을 열고 필요한 설정을 입력하세요.

4. 데이터베이스 마이그레이션

   ```bash
   go run cmd/migrate/main.go
   ```

5. 애플리케이션 실행
   ```bash
   go run cmd/api/main.go
   ```

## API 문서

API 문서는 Swagger를 통해 제공됩니다. 서버 실행 후 `/swagger/index.html`에서 확인할 수 있습니다.

## 테스트

단위 테스트 실행:

```bash
go test ./...
```

## 기여 방법

1. 이 저장소를 포크합니다.
2. 새 브랜치를 생성합니다 (`git checkout -b feature/AmazingFeature`).
3. 변경 사항을 커밋합니다 (`git commit -m 'Add some AmazingFeature'`).
4. 브랜치에 푸시합니다 (`git push origin feature/AmazingFeature`).
5. Pull Request를 생성합니다.

## 라이선스

이 프로젝트는 MIT 라이선스 하에 배포됩니다. 자세한 내용은 `LICENSE` 파일을 참조하세요.
