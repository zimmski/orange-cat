package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"text/template"
)

func Template(w http.ResponseWriter, filepath string, port int) {
	var style string
	if css, err := CustomCSS(); err == nil {
		style = *css
	} else {
		style = "<style>" + DefaultStyle + "</style>"
	}

	templateStr := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
  <meta charset='UTF-8' />
  <title>%[1]s</title>
  %[3]s
</head>
<body>
  <div id='md' class='markdown-body'></div>
  <script>
    (function () {
      var markdown = document.getElementById("md");
      var conn = new WebSocket("ws://localhost:%[2]d/%[1]s");
      conn.onmessage = function (evt) {
        markdown.innerHTML = evt.data;
      };
    })();
  </script>
</body>`, filepath, port, style)

	var (
		t   *template.Template
		err error
	)

	if t, err = template.New("template").Parse(templateStr); err != nil {
		panic(err)
	}

	if err = t.Execute(w, nil); err != nil {
		panic(err)
	}
}

func CustomCSS() (*string, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	customCSSPath := filepath.Join(usr.HomeDir, ".orange-cat.css")

	if stat, err := os.Stat(customCSSPath); err == nil && stat.Mode().IsRegular() {
		customCSS := "<link rel='stylesheet' href='" + customCSSPath + "' />"
		return &customCSS, nil
	} else {
		return nil, errors.New("No custom CSS")
	}
}

var DefaultStyle string = `
/* github-markdown-css */
.markdown-body {
  overflow: hidden;
  font-family: "Helvetica Neue", Helvetica, "Segoe UI", Arial, freesans, sans-serif;
  font-size: 16px;
  line-height: 1.6;
  word-wrap: break-word
}

.markdown-body>*:first-child {
  margin-top: 0 !important
}

.markdown-body>*:last-child {
  margin-bottom: 0 !important
}

.markdown-body .absent {
  color: #c00
}

.markdown-body .anchor {
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  display: block;
  padding-right: 6px;
  padding-left: 30px;
  margin-left: -30px
}

.markdown-body .anchor:focus {
  outline: none
}

.markdown-body h1, .markdown-body h2, .markdown-body h3, .markdown-body h4,
.markdown-body h5, .markdown-body h6 {
  position: relative;
  margin-top: 1em;
  margin-bottom: 16px;
  font-weight: bold;
  line-height: 1.4
}

.markdown-body h1 .octicon-link, .markdown-body h2 .octicon-link,
.markdown-body h3 .octicon-link, .markdown-body h4 .octicon-link, .markdown-body h5 .octicon-link,
.markdown-body h6 .octicon-link {
  display: none;
  color: #000;
  vertical-align: middle
}

.markdown-body h1:hover .anchor, .markdown-body h2:hover .anchor, .markdown-body h3:hover .anchor, .markdown-body h4:hover .anchor, .markdown-body h5:hover .anchor, .markdown-body h6:hover .anchor {
  padding-left: 8px;
  margin-left: -30px;
  line-height: 1;
  text-decoration: none
}

.markdown-body h1:hover .anchor .octicon-link, .markdown-body h2:hover .anchor .octicon-link, .markdown-body h3:hover .anchor .octicon-link, .markdown-body h4:hover .anchor .octicon-link, .markdown-body h5:hover .anchor .octicon-link, .markdown-body h6:hover .anchor .octicon-link {
  display: inline-block
}

.markdown-body h1 tt, .markdown-body h1 code, .markdown-body h2 tt, .markdown-body h2 code,
.markdown-body h3 tt, .markdown-body h3 code, .markdown-body h4 tt, .markdown-body h4 code,
.markdown-body h5 tt, .markdown-body h5 code, .markdown-body h6 tt, .markdown-body h6 code {
  font-weight: normal;
  font-size: inherit
}

.markdown-body h1 {
  padding-bottom: 0.3em;
  font-size: 2.25em;
  line-height: 1.2;
  border-bottom: 1px solid #eee
}

.markdown-body h2 {
  padding-bottom: 0.3em;
  font-size: 1.75em;
  line-height: 1.225;
  border-bottom: 1px solid #eee
}

.markdown-body h3 {
  font-size: 1.5em;
  line-height: 1.43
}

.markdown-body h4 {
  font-size: 1.25em
}

.markdown-body h5 {
  font-size: 1em
}

.markdown-body h6 {
  font-size: 1em;
  color: #777
}

.markdown-body p, .markdown-body blockquote, .markdown-body ul, .markdown-body ol,
.markdown-body dl, .markdown-body table, .markdown-body pre {
  margin-top: 0;
  margin-bottom: 16px
}

.markdown-body hr {
  height: 4px;
  padding: 0;
  margin: 16px 0;
  background-color: #e7e7e7;
  border: 0 none
}

.markdown-body ul, .markdown-body ol {
  padding-left: 2em
}

.markdown-body ul.no-list, .markdown-body ol.no-list {
  padding: 0;
  list-style-type: none
}

.markdown-body ul ul, .markdown-body ul ol, .markdown-body ol ol, .markdown-body ol ul {
  margin-top: 0;
  margin-bottom: 0
}

.markdown-body li>p {
  margin-top: 16px
}

.markdown-body dl {
  padding: 0
}

.markdown-body dl dt {
  padding: 0;
  margin-top: 16px;
  font-size: 1em;
  font-style: italic;
  font-weight: bold
}

.markdown-body dl dd {
  padding: 0 16px;
  margin-bottom: 16px
}

.markdown-body blockquote {
  margin-left: 0;
  margin-right: 0;
  padding: 0 15px;
  color: #777;
  border-left: 4px solid #ddd
}

.markdown-body blockquote>:first-child {
  margin-top: 0
}

.markdown-body blockquote>:last-child {
  margin-bottom: 0
}

.markdown-body table {
  display: block;
  width: 100%%;
  overflow: auto;
  word-break: normal;
  word-break: keep-all
}

.markdown-body table th {
  font-weight: bold
}

.markdown-body table th, .markdown-body table td {
  padding: 6px 13px;
  border: 1px solid #ddd
}

.markdown-body table tr {
  background-color: #fff;
  border-top: 1px solid #ccc
}

.markdown-body table tr:nth-child(2n) {
  background-color: #f8f8f8
}

.markdown-body img {
  max-width: 100%%;
  -moz-box-sizing: border-box;
  box-sizing: border-box
}

.markdown-body span.frame {
  display: block;
  overflow: hidden
}

.markdown-body span.frame>span {
  display: block;
  float: left;
  width: auto;
  padding: 7px;
  margin: 13px 0 0;
  overflow: hidden;
  border: 1px solid #ddd
}

.markdown-body span.frame span img {
  display: block;
  float: left
}

.markdown-body span.frame span span {
  display: block;
  padding: 5px 0 0;
  clear: both;
  color: #333
}

.markdown-body span.align-center {
  display: block;
  overflow: hidden;
  clear: both
}

.markdown-body span.align-center>span {
  display: block;
  margin: 13px auto 0;
  overflow: hidden;
  text-align: center
}

.markdown-body span.align-center span img {
  margin: 0 auto;
  text-align: center
}

.markdown-body span.align-right {
  display: block;
  overflow: hidden;
  clear: both
}

.markdown-body span.align-right>span {
  display: block;
  margin: 13px 0 0;
  overflow: hidden;
  text-align: right
}

.markdown-body span.align-right span img {
  margin: 0;
  text-align: right
}

.markdown-body span.float-left {
  display: block;
  float: left;
  margin-right: 13px;
  overflow: hidden
}

.markdown-body span.float-left span {
  margin: 13px 0 0
}

.markdown-body span.float-right {
  display: block;
  float: right;
  margin-left: 13px;
  overflow: hidden
}

.markdown-body span.float-right>span {
  display: block;
  margin: 13px auto 0;
  overflow: hidden;
  text-align: right
}

.markdown-body code, .markdown-body tt {
  font-family: Consolas, 'Liberation Mono', Menlo, Courier, monospace;
  padding: 0;
  padding-top: 0.2em;
  padding-bottom: 0.2em;
  margin: 0;
  font-size: 85%%;
  background-color: rgba(0, 0, 0, 0.04);
  border-radius: 3px
}

.markdown-body code:before, .markdown-body code:after, .markdown-body tt:before, .markdown-body tt:after {
  letter-spacing: -0.2em;
  content: "\00a0"
}

.markdown-body code br, .markdown-body tt br {
  display: none
}

.markdown-body del code {
  text-decoration: inherit
}

.markdown-body pre>code {
  padding: 0;
  margin: 0;
  font-size: 100%%;
  word-break: normal;
  white-space: pre;
  background: transparent;
  border: 0
}

.markdown-body .highlight {
  margin-bottom: 16px
}

.markdown-body .highlight pre, .markdown-body pre {
  padding: 16px;
  overflow: auto;
  font-size: 85%%;
  line-height: 1.45;
  background-color: #f7f7f7;
  border-radius: 3px
}

.markdown-body .highlight pre {
  margin-bottom: 0;
  word-break: normal
}

.markdown-body pre {
  word-wrap: normal
}

.markdown-body pre code, .markdown-body pre tt {
  display: inline;
  max-width: initial;
  padding: 0;
  margin: 0;
  overflow: initial;
  line-height: inherit;
  word-wrap: normal;
  background-color: transparent;
  border: 0
}

.markdown-body pre code:before, .markdown-body pre code:after, .markdown-body pre tt:before, .markdown-body pre tt:after {
  content: normal
}

/* custom style */
body {
  padding: 20px 0;
}
.markdown-body {
  max-width: 800px;
  margin: 0 auto;
}
a, a:visited {
  color: #4183c4;
  text-decoration: none;
}
a:hover {
  text-decoration: underline;
}`
