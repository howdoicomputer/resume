##
# Resume
#
# @file
# @version 0.1

.PHONY: sections
CC = xelatex

resume.pdf: clean
	$(CC) resume.tex
	okular resume.pdf

coverletter.pdf: clean
	$(CC) coverletter.tex

build: clean
	dagger run go run ci/main.go

clean:
	rm -rf *.pdf *.aux *.log
# end
