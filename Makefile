include $(GOROOT)/src/Make.inc

TARG=github.com/gcmurphy/postmark
GOFILES=\
	message.go\
    postmark.go\

include $(GOROOT)/src/Make.pkg
