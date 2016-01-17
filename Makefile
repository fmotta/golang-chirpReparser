GO = go
GMESSAGE := "Default Message"
PREFIX := /usr/local/bin
APP = chirpReparser
LANG=golang

all: ${APP}

$(APP): $(LANG)-$(APP)
	mv $^ $@

$(LANG)-$(APP): main.go
	$(GO) build 

install: $(APP)
	cp $^ ${PREFIX}/.

clean:
	rm -f $(APP) $(LANG)-$(APP)
	
push:
	git commit -m ${GMESSAGE}
	git status | grep 'modified:' | awk 'BEGIN{FS=":";}{print $2;}' | xargs git add
	git push origin master
