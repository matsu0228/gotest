
# プロジェクトごとの定義
# ----------------------------------

# REPO_FULL=$(shell git rev-parse --show-toplevel)
# REPO_DIR=$(shell basename $(shell dirname $(shell dirname $(REPO_FULL))))
# REPO_AUTHOR=$(shell basename $(shell dirname $(REPO_FULL)))
# REPO_NAME=$(shell basename $(REPO_FULL))
CMD_NAME=$(shell basename $(shell pwd))
REPO_DIR=github.com
REPO_AUTHOR=matsu0228
REPO_NAME=gotest
# CMD_NAME=tips

FIRST_GOPATH=$(shell echo $(GOPATH) | awk -F ":" '{ print $1 }')
APP_ROOT=$(abspath ./)
WORK_ROOT=$(FIRST_GOPATH)/src/$(REPO_DIR)/$(REPO_AUTHOR)/${REPO_NAME}
WORK_DIR=$(WORK_ROOT)
# $(CMD_NAME)

# LIST_NOT_TARGET      := infla
# LIST_DIRECTORIES     := ${subst /,,${shell ls -d */}}
# LIST_TARGET_FUNCTONS := ${filter-out ${LIST_NOT_TARGET},${LIST_DIRECTORIES}}

# 単体テスト:  make test-run TEST_NAME=**
# TEST_NAME=TestGetOrder


# define some variables


# コマンド一覧
# ----------------------------------
.PHONY: setup deps build clean debug test test-all

deps:
		cd $(WORK_DIR); dep ensure -v

# build: deps
# 		cd $(WORK_DIR); $(APP_ROOT)/build/crosscompile_for_go.sh -o $(CMD_NAME)

test:
		go test ./tips
		go test ./integrate
		go test ./integrate/repository

# test-run:
# 		go test ** 	-v -run $(TEST_NAME)

test-all:
		go test ./tips
		go test ./integrate
		go test ./integrate/repository -v -tags=integration

setup:
ifndef APP_ROOT
		#未定義の場合
		$(info $(APP_ROOT)"が存在しないため、setup完了していません")
else
		go get -u github.com/golang/dep/cmd/dep
		mkdir -p $(abspath $(WORK_ROOT)/..)
		ln -s $(APP_ROOT)  $(abspath $(WORK_ROOT))
		$(info "gopath配下にsetupしました at $(WORK_ROOT)  with GOPATH=$(GOPATH)")
endif

# TODO: linux/mac向けのバイナリ削除
clean:
		cd $(WORK_DIR); rm -rf *.exe

debug:
		$(info repo:    $(REPO_DIR)/$(REPO_AUTHOR)$(REPO_NAME))
		$(info gopath:  $(FIRST_GOPATH))
		$(info workRoot: $(WORK_ROOT))
		$(info workdir: $(WORK_DIR))
		$(info appRoot:     $(APP_ROOT))
		$(info cmd: $(CMD_NAME))