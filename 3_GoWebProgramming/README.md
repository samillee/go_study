﻿Go 언어 웹 프로그래밍 철저 입문
======================

## Go 언어, 이렇게 활용한다!

<img src="cover.jpg" href="http://book.naver.com/bookdb/book_detail.nhn?bid=10406887" />

### Go 언어다운 프로그래밍을 배운다
Go는 간결하고 유연한 문법을 지원하며, 고루틴으로 병행 처리 코드를 쉽게 작성할 수 있다. 또한, 상속이 아닌 조합으로 코드를 재사용하여 확장성이 좋고, 패키지화된 소스 코드에서 실제로 사용되는 부분만 컴파일하므로 컴파일 속도가 매우 빠르다. 이런 특징들을 고려하여 Go 언어다운 프로그래밍 방법을 설명한다.

### 확장성이 좋고 유연한 마이크로 프레임워크를 만든다
각 상황에 맞게 Go 기본 라이브러리와 다양한 외부 패키지를 조합하여 마이크로 서비스 형태로 자신만의 프레임워크 제작 방법을 배운다. 이렇게 만든 마이크로 프레임워크로는 여러 사용자가 실시간으로 대화할 수 있는 채팅 애플리케이션을 만든다.

### Revel 프레임워크로 빠르게 웹 애플리케이션을 만든다
풀스택 프레임워크 Revel은 웹 개발에 필요한 기능 대부분을 제공하므로 웹 애플리케이션을 아주 빠르게 제작할 수 있고, 이미 만들어진 웹 애플리케이션에 기능을 추가하기도 쉽다. 이 책에서는 Revel을 활용하여 웹 애플리케이션을 만든다.


----
# 책 요약

## 1장 Go 시작하기: 1.4.2 - > 1.7.1 (2016-08-15) 6개월주기 버전업
### 1.1 Go는 어떤 언어인가? 
  + 간결, 정적, 병행, 동적, 가비지콜렉터, 상속대신 조합, 쉬운협업, 빠른 컴파일
  + 상태를 표현하는 타입 + 동작을 표현하는 메서드
  + 구조체, 인터페이스, 함수,  
### 1.2 Go 설치하기
  1. 설치환경: 윈도, OS X, 리눅스와 FreeBSD: 작업환경위해 gopath 환경변수 설정 필요
  4. 설치 확인: go version
  Chocolatey 이용 간편 설치: chocolatey.org
### 1.3 Go 개발 환경
  1. 작업 공간 구성: GOPATH 환경설정
  2. 에디터: sublime Text 3, VS Code 유명한 플러그인들 설치할 것, LiteIDE
### 1.4 첫 번째 Go 프로그램
  1. 코드 실행: go run, go build 후 실행, go install 후 실행
  2. 코드 분석: 실행가능 프로그램과 라이브러리
### 1.5 Go 참고 문서
    https://golang.org/doc: Effective Go 참고
    https://golang.org/pkg: 패키지
	https://play.golang.org
	godoc -http=:8000   로컬컴에서 웹도움 실행
	godoc fmt Fprintf
    도구: Oracle, Vet, Fix, Test

----
## 2장 기본 문법
### 2.1 Go 문법의 특징
  1. 간결함과 유연함: for, switch-case
  2. 정적 타입 언어, 동적 프로그래밍
  3. 모호한 요소 제거: 증감연산자 후치연산만 가능, 포인터는 사용가능하나 포인터 연산 없음
  4. 세미콜론 생략 가능
  5. 주석: C#과 동일
  6. gofmt로 코드 서식 지정: goimport 툴 이용하면 import 자동으로 해줌
### 2.2 변수와 상수
  1. 변수 선언: 이름 뒤에 타입 지정 - 여러 변수를 한번에 타입 지정 가능, 
  	값 할당하면서 변수선언하면 타입 생략 가능
  	짧은 선언 := 함수안에서만 가능
  2. 변수 이름: 유니코드 가능, 짧고 간결하고 함축적으로, 낙타표기법
  	Getter Get없이 명사로 Setter는 Set+명사
  3. 상수: 불, 숫자, 문자만 가능, 계산 식으로 가능
  4. 열거형: iota 값은 0 이후 값은 1씩 증가, _ 빈식별자를 써서 0 시작을 무하고 1부터 시작 가능
  	비트 조작해서 정의하는 예제들 좋음
### 2.3 프로그램의 기본 흐름 제어
  1. if: 가급적 else 안쓰도록 한다. 초기구문 지정 가능
  2. switch: fallthrough 가급적 쓰지말자 차라리 case 하나에 ,로 여러조건 지정할 것, 초기구문 지정 가능
  3. for: break, continue, 레이블 지정해서 탈출 가능
### 2.4 함수: First-class citizen: 값으로 취급되는 함수
  1. 함수 정의: 매개변수, 가변인자, 다중반환, nil, _ 빈식별자
  2. 매개변수 전달 방식: call by value, call by reference
  3. defer: 정리가 필요한 것 호출 후 바로 지정
  4. 내장 함수: close, len, cap, new, make, copy, append, panic, recover, complex, real, image
  5. 클로저: 외부 선언 변수를 함수 리터럴 내에서 마음대로 접근할 수 있는 코드, 함수변수, 클로저를 사용한 팩토리 함수
  6. 함수를 매개변수로 전달하기
     + f := func(c rune) bool { return unicode.Is(unicode.Hangul, c); }
### 2.5 패키지: 코드 구조화 및 재사용 단위, 패키지 이름 = 폴더이름, 같은 패키지 소스는 같은 폴더에 있어야함, 보통 소문자
  1. 패키지 종류: 실행가능 프로그램 vs 라이브러리
  2. 접근 제어: 소문자로 시작하면 private, 대분자면 public 
  3. 별칭: _ 빈식별자 가능
  4. init() 함수: 패키지 로드될 때 가장 먼저 실행

----
## 3장 데이터 타입
### 3.1 불
### 3.2 숫자
  1. 정수: int, uint, byte, rune = int32 (utf8표현), uintptr = unit
  2. 실수(부동소수점): float32, float64, math 패키지
  3. 복소수: complex64, complex128, math/cmplx 패키지
  4. 숫자 연산: 반드시 같은 타입 변환후 연산, math.Max정수
### 3.3 문자열: len([]rune(s)), strconv.Atoi(s), strconv.Itoa(i)
  1. 문자열과 문자: unicode.Is종류
  2. 문자열 연산: s[n:m], s[n:], s[:m], +, +=, strings.Join(), bytes.Buffer  Write.String(s)
### 3.4 배열과 슬라이스: 고정 vs. 가변길이, 값타입 vs. 참조타입, 등위연산 가능 vs. 불가능, append, copy, s[n:m]은 슬라이스만 가능
  1. 생성과 초기화: 슬라이스 make 이용
  2. 내부 요소에 순차적으로 접근: i, v range 배열|슬라이스
  **3. 부분 슬라이스 추출: s[n:m]**
  4. 슬라이스 변경: append, insert 없어서 특별히 해야함, sort Sort() 인터페이스로 지정 가능
### 3.5 맵: 키와 값의 테이블 형태 컬렉션, 키는 등위비교 가능해야함
  1. 생성과 초기화: map[키타입]값타입, k, v range 맵
  2. 값 찾기
  3. 요소 추가, 수정, 삭제: delete(맵, 키값)
### 3.6 포인터와 참조 타입: * 연산자, & 주소연산자
  1. 포인터 생성과 초기화: & 주소연산자로 메모리 주소 할당, new() 함수로 메모리 초기화후 주소할당
  2. 값 전달: 원본 변경하려면 포인터 사용, 값이 커도 속도 향상위해 포이터 사용 권장

----
## 4장 객체 지향 프로그래밍
### 4.1 객체 표현 방식: 타입 + 메서드 (타입에 바인딩된 함수)
	func (리서버명 리시버타입) 메서드명(매개변수) (반환타입 또는 값) { }
### 4.2 사용자 정의 타입
  1. 사용자 정의 타입의 4종류: 기본타입, 함수서명, stuct, interface (메서드 묶음)
  	 + 기본타입 <-> 사용자타입 변환: int(q)
  2. 메서드: 리시버도 값 또는 참조 타입 (맵, 슬라이스), 리시버 변수 생략 가능
### 4.3 구조체: 실세계 엔터티를 표현, 익명 구조체도 가능,  필드 태그
  1. 생성과 초기화: 타입{초기값}, &타입{} 또는 &타입{초기값}: 초기값 할당된 구조체 포인터 생성, new(타입)
  2. **내부 필드 접근: . 연산자, 필드명과 값 모두출력 fmt.Printf("%#v", v), reflect.TypeOf()로 확인**
  3. 구조체 임베딩: 내부 필드와 같은게 있음 필드타입을 하께 적어줘야함, 임베디드 메서드 재사용 및 오버라이딩 가능
  4. 정보 은닉: 생성함수는 New(...)로 시작하도록 한다.  Getter는 Get으로 시작하지 않고 그냥 명사, Setter는 Set+명사
### 4.4 인터페이스: 덕타이핑 방식으로 객체의 변수나 메서드의 집합이 타입을 결정
  1. 인터페이스 정의: 인터페이스 이름은 메서드 이름에 er 또는 r 붙여서 짓는다. 익명과 빈 인터페이스도 가능
  2. 다형성: 인터페이스로 다형성 지원, is-a 관계, 제너릭콜렉션 type Items []인터페이스
  	기본라이브러리: fmt.Stringer String(), io.Writer Write()
  3. 인터페이스 임베딩: 멀티도 가능
  4. **타입 변환: 타입 assertion, switch, reflection**

----
## 5장 병행 처리
### 5.1 고루틴: go 키워드 + <- chan <- 채널
### 5.2 채널: 값 vs. 참조 채널, 참조채녈은 mutex로 보호 할 필요 있음
	runtime.GOMAXPROCS(runtime.NumCPU())
  1. 채널 방향: <- 수신 chan 송신 <-, 지정 안하면 양방향
  2. 버퍼드 채널: ch := make(chan int, 100) 꽉 찼는데 계속 전송하면 deadlock 에러 발생
  3. close & range: v, ij := <- ch, for i := range c
  4. select: 여러 채널과 통신할 때 사용, case로 채널을 대기시킴, default 채널 사용가능 상태 아닐 경우
### 5.3 저수준 제어
  1. **sync.Mutex: Lock(), UnLock()**  
  2. **sync.RWMutex: 쓰기 Lock(), UnLock() vs 읽기 RLock(), RUnLock()**
  3. sync.Once
  4. **sync.WaitGroup: 모든 고루틴 종ㄽ 대기, Add(1), Done(), Wait()**
  5. **원자성을 보장하는 연산: sync.atomic 패키지 AddT, LoadT, StoreT, SwapT, CompareAndSwapT**
   		atomic.AddOmt64(&c.i, 100)
### 5.4 활용 **예제 잘 쓸 것** 
  1. **타임아웃: select 이용 timeout := time.After(10 * time.Millisecond), select { case <- timeout: }**
   		대응: 1) 아무것도 안음, 2) 채널 닫기, 3) 타임아웃 메시지 전송
  2. **공유 메모리: SharedMap 정의 (Mutex 이용) 및 패턴** 
  3. **파이프라인**
  4. 맵리듀스

----
## 6장 에러 처리
### 6.1 에러 타입: type error interface
### 6.2 에러 생성
  1. errors.New() 사용
  2. fmt.Errorf() 패키지 함수 사용: fmt.Errorf("%g", f)
  3. 사용자 정의 에러 타입: Error() string 메서드 가진 타입
  	json.SyntaxError
### 6.3 panic() & recover()
  1. 런타임 에러와 패닉: Must로 시작하는 경우 많음: regexp.MustCompile 
  2. recover(): panic에서 종료절차를 중지시킴
### 6.4 에러 처리
  1. 에러 확인 함수 사용: 반복 코드 줄임: func check(err error) { if err != nill { panic(err) } }
  2. **클로저로 에러 처리: 예제 확인 할 것**

----
## 7장 패키지
### 7.1 커스텀 패키지
  1. 패키지 만들기: $GOPATH/src 아래 리포저장소 기준으로 생성할것 패키지명 = 폴더명
  2. 별칭: _ 빈식별자로 지정가능: DB나 이미지 관련 된 것들 그렇게 하면 좋음
  3. 운영체제에 종속적인 코드 처리
     + (1) runtime.GOOS로 확인후 분기 처리
     + (2) "_운영체제이름" 추가한 소스 생성
  4. 문서화: godoc은 기본적으로 public만 보여줌
### 7.2 서드 파티 패키지: http://godoc.org 에서 확인 가능
	go get -u 패키지 최신으로 유지
### 7.3 Go의 기본 라이브러리: https://golang.org/pkg 에서 확인 가능
  1. 문자열 패키지: strings, bytes, strconv, fmt, unicode, text/template, html/template
  2. 컬렉션: 최대/최소 container/heap, 더블링크드 리스트  container/list, container/ring, database/sql
  3. 파일, 디렉터리, 운영체제 환경 다루기
     + os, bufio, io/ioutil, path, path/filepath, runtime
  	 + encoding, encoding/json, encoding/xml
       + json.Marshal(t), json.Unmarchal(b, &t), 태그 `json:"이름"`
  4. 숫자 연산: math, math/big
  5. 네트워크: net, net/http, net/url, net/rpc, net/smtp
  6. 리플렉션: reflect
  7. 테스트: testing 패키지, go test -v 명령으로 수행, _test.go 파일
  	 + Test로 함수이름 시작하고 *testing.T 매개변수 받음
  8. 기타
     + crypto, os/exec, flag, log, regexp, sort, image,
     + archive/tar, archive/zip, compress/gzip, compress/bzip2, complress/lzw,
  	 + time: time.After()time.Tick()

----
## 8장 나만의 웹 프레임워크 만들기
### 8.1 나만의 웹 프레임워크 만들기
 + 풀스택 프레임워크: Revel, Beego
 + 마이크로 프레임워크: Gin, Goji, Martini
 + 라이브러리/툴킷: Gorilla
### 8.2 첫 번째 웹 애플리케이션
### 8.3 라우터: gorilla/mux, httprouter
### 8.4 컨텍스트: 처리상태를 저장, negroini
### 8.5 미들웨어: DRY 원칙, 로그, 에러, 정적파일, 웹요청정보파싱, 인증...
### 8.6 추상화: 서버 타입 정의해서 
### 8.7 렌더러: encoding/json, render
### 8.8 커스텀 미들웨어: login 인증 처리 등 

----
## 9장 다양한 패키지를 조합하여 마이크로 프레임워크 구성하기
### 9.1 채팅 애플리케이션 만들기: html5 웹소켓 이용
  1. 라우터: gin
  2. 컨텍스트: 처리상태를 저장, negroini "github.com/codegangsta/negroni"
  3. 세션관리: negroni-sessions
  4. 미들웨어: negroni
  5. 렌더러: encoding/json, "github.com/unrolled/render"
  6. 서셜로그인: gomniauth
  7. 웹소켓: gorilla/websocket
  8. 몽고DB: mgo 
### 9.2 웹 서버 구동하기
### 9.3 인증 처리하기
  1. 세션
  2. 로그인
  3. 인증
### 9.4 채팅방과 메시지 처리하기
  1. 몽고DB 환경 구성
  2. 채팅방 관리 기능 구현
  3. 메시지 조회 기능 구현
### 9.5 HTML과 자바스크립트로 클라이언트 화면 만들기
### 9.6 웹소켓 기능 구현하기

----
## 10장 Revel 프레임워크로 블로그 만들기
10.1 Revel 프로젝트 만들기
  1. Revel 설치하기
  2. 블로그 애플리케이션 만들기
  3. 데이터베이스 설정
10.2 Hello World
  1. 웹 서버 시작하기
  2. Hello Revel!
10.3 포스트 기능 만들기
  1. 포스트 모델 만들기
  2. 데이터베이스 초기화
  3. 포스트 컨트롤러 만들기
  4. 포스트 목록 보기
  5. 새 포스트 만들기
  6. 각 포스트 보여주기
  7. 포스트 수정하기
  8. 포스트 삭제하기
10.4 댓글 기능 만들기
  1. 코멘트 모델 만들기
  2. 코멘트 컨트롤러 작성하기
  3. 댓글 작성을 위한 라우팅 규칙 추가하기
  4. 포스트의 Show 페이지에서 댓글 보여주기
10.5 리팩토링
  1. 댓글 목록을 별도의 템플릿으로 만들기
  2. 댓글 작성 폼을 별도의 템플릿으로 만들기
  3. 포스트 작성 폼을 별도의 템플릿으로 만들기
  4. 새로 분리한 템플릿에 댓글 삭제 기능 추가하기
10.6 데이터 처리에 ORM 프레임워크 적용
  1. gorm 초기 설정
  2. gorm으로 데이터 처리
10.7 로그인과 보안
  1. 사용자 모델 추가
  2. 로그인 기능 구현
  3. 기본 계정 생성
  4. 인증 인터셉터 추가
  5. 권한이 있는 사용자만 해당 기능에 접근

