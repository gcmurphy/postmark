include $(GOROOT)/src/Make.inc

TARG=postmark
GOFILES=\
	msg.go\
    postmark.go\

include $(GOROOT)/src/Make.pkg
