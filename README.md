# Smart-Media



## Overview

교육을 통해 백엔드에 대한 이해도를 높이고 실습을 통해 실무를 익힌다.





# Git & GitHub

## :books: Git

### :bulb:Overview

Git은 Configuration Management Tool 중 하나이다. 소프트웨어 개발에 사용되는 소스코드를 효과적으로 관리할 수 있게 도와주는 무료, 공개 소프트웨어이다. SVN과는 다르다. SVN은 중아 서버에 소스코드와 히스토리를 저장하는 것과는 달리 각각의 PC에 분산하여 저장한다. 이를 분산형 관리 시스템이라고 한다.

### :file_cabinet:Git을 왜 사용할까?

Git은 소스코드의 버젼별로 저장을 하게된다. 따라서 상황에 따라 해당 버젼으로 복구가 가능하다. 또한 여러명이서 동시에 개발이 가능하다. 브랜치(branch)를 나누어 개발한 뒤 다시 머지(merge)를 통해 소스를 합칠 수도 있다.

### :book:용어 정리

`Commit` : 변경사항을 저장하여 새로운 버젼으로 만든다.

`Branch` : Branch에 Commit을 할 수 있다. Branch는 분리될수도 다시 합쳐질 수도 있다.

`Merge` : 두개의 Branch를 다시 하나로 합친다.

`Staging Area` : 변경사항을 commit하기전 준비하는 위치

### :computer:커맨드

- 현재폴더를 Git에서 관리할 수 있도록 설정한다.

```shell
$ git init
```

- 브랜치를 조회한다.

```shell
$ git branch
```

- 'develop'이란 이름의 브랜치를 생성한다.

```shell
$ git branch develop
```

- git 상태 확인

```shell
$ git status
```

- 'test.txt'파일을 `staging area`로 이동한다.

```shell
$ git add test.txt
```

- `staging area`안에 있는 내용을 현재 `branch`에 `commit`한다.

```shell
$ git commit -m "test.txt파일 추가"
```

- git log확인

```shell
$ git log
```

## :books:Github

### :bulb:Overview

Github은 개발자용 SNS이다. git에 저장된 소스를 온라인에 저장하고 공유할 수 있는 웹 호스팅 서비스이다. 또한 협업에 관련하여 다양한 기능들을 제공하여 언제 어디서든 git을 통해 소스코드를 비교하고 수 있다.

### :book:용어 정리

`Repository` : branch, 소스등을 포함하여 git이 관리하는 하나의 프로젝트이다.

`clone` : Repository 내용을 PC로 download한다.

`push` : git소스를 github으로 upload한다.

`pull` : github소스를 git으로 download한다.

`fork` : 다른 Repository를 복사한다.

`pull request` : fork한 Repository의 변경사항을 원본 repository에 적용을 요청한다.

### :computer:커맨드

- github의 내용을 clone한다.

```shell
git clone 저장소 주소
```
