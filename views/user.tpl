<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <title>Demo</title>
    <meta charset="utf-8">
</head>
<body>
<h1>网站: {{.Website}}</h1>
{{ if .user }}
用户名: {{.user.Username}}
密码 :{{.user.Password}}
{{else}}
查找不到用户
{{ end }}
    <form method="post" action="/index">
        <input type="submit">
    </form>
</body>
</html>