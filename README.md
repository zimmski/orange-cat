# `orange-cat`

[![build
status](https://api.travis-ci.org/noraesae/orange-cat.svg)](https://travis-ci.org/noraesae/orange-cat)

A Markdown previewer written in Go

## Why orange-cat?

GitHub suggested naming it as `orange-cat`.

> Great repository names are short and memorable. Need inspiration? How about `orange-cat`.

GitHub is always right and I had to obey it.

## Demo

![demo](http://i.imgur.com/qETC9A4.gif)

## How it works

To launch `orange-cat`, simply run the `orange` command.

```
$ orange README.md
Listening :6060 ...
```

Then `orange-cat` will start watching the Markdown file and open a
browser window where the preview of the file will be displayed. You can
also open `http://localhost:6060/some_file.md` manually.

When you modify the file, `orange-cat` watcher will catch the
modification and send the modified data to the browser through a
websocket connection. It means, you don't even need to refresh the page.

To stop it, simply enter `^C`.

## Why another Markdown preview?

I know there're already plenty of Markdown previewers, such as Atom's
Markdown preview package, some Vim plugins and other web-based or desktop
apps.

However, I don't use any modern IDE or editor. I love Vim. There must be
people who love their own prefered editors, like me. I wanted to make a
previewer running offline, with any editor, without any dependency.

This is a binary executable, not a script. We don't need any `gem`,
`npm` or `pip` to use this. How to use is completely up to you.

I sincerely hope you like `orange-cat` :)

## Install

You can download binaries for your environment in
[Releases](https://github.com/noraesae/orange-cat/releases).

If you prefer to build from source, it's also very easy.

```
$ git clone git@github.com:noraesae/orange-cat.git
$ cd orange-cat
$ make build
```

The binary `orange` will be created at `$GOPATH/bin/orange`.

## Custom CSS

`orange-cat` will try to find a custom CSS file from
`~/.orange-cat.css`. If there's no custom CSS file, it'll use a default
CSS style, which shows a similar output to GitHub's one.

## Contribution

I welcome every kind of contribution.

If you have any problem using `orange-cat`, please file an issue in
[Issues](https://github.com/noraesae/orange-cat/issues).

If you'd like to contribute on source, please upload a pull request in
[Pull Requests](https://github.com/noraesae/orange-cat/pulls). Please
don't forget to check if it's gofmt'ed and passes every test before
uploading a new pull request. It can be done with following commands.

```
$ make fmt # gofmt for every source code
$ make test # run Ginkgo test suite
```

If needed, please add a new test case with your patch.

## License

MIT
