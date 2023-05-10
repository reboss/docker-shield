.PHONY: all install clean uninstall

LIBDIR=${DESTDIR}/lib/systemd/system
BINDIR=${DESTDIR}/usr/lib/docker

all: docker-shield

GOFILES := $(wildcard ./**/*.go) main.go
docker-shield: $(GOFILES)
		go build -o $@ .

install:
		mkdir -p ${LIBDIR} ${BINDIR}
		install -m 644 systemd/docker-shield.service ${LIBDIR}
		install -m 644 systemd/docker-shield.socket ${LIBDIR}
		install -m 755 docker-shield ${BINDIR}

clean:
		rm -rf docker-shield

uninstall:
		rm -f ${LIBDIR}/docker-shield.service
		rm -f ${LIBDIR}/docker-shield.socket
		rm -f ${BINDIR}/docker-shield
