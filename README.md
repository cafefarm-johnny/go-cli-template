# Go CLI Template

## 소개

Go 언어와 [spf13/Cobra](https://github.com/spf13/cobra) 라이브러리를 이용해 CLI 프로그램을 만들기 위한 기본적인 템플릿 구조의 프로젝트입니다.

## structure

```text
go-cli-template
 ├──     README.md
 ├──     go-cli-template.iml
 ├──     go.mod
 ├──     go.sum
 └──     src/
 │  └────     app/
 │  │  └────     app.go  // app 정보
 │  └────     cmd/
 │  │  ├────     template
 │  │  │  └────     help.go  // help 커맨드의 템플릿  
 │  │  ├────     flags.go  // 커맨드 공통 플래그 변수
 │  │  ├────     root.go  // root 커맨드
 │  │  └────     version.go  // 앱 버전 출력 커맨드
 │  └────     main.go // 메인 함수
```