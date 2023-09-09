# My Resume

This resume is a LaTeX template heavily based on [AwesomeCV](https://github.com/posquit0/Awesome-CV). It has been modified slightly to better support pointing to custom hrefs and better break support for long sections.

Feel free to use the code for this but please don't use my resume for anything nefarious.

# Building

## Dependencies

* `xelatex`
* GNU Make

The two Nix packages I have installed are:

``` sh
gnumake
texlive.combined.scheme-full 
```

## Shell

The default make target is to produce the resume.

```
make
```

From there, you can open up and review the PDF. I'll either use Emacs or Okular.

## Dagger and Docker

There is also a Dagger pipeline for generating a resume.

``` sh
dagger run go run ci/main.go
```

---
