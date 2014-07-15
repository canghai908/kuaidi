{{define "navbar"}}
<a class="navbar-brand" href="/">会员</a>
<div>
	<ul class="nav navbar-nav">
		<li {{if .IsHome}}class="active"{{end}}><a href="/">首页</a></li>
		
	</ul>
</div>

<div class="pull-right">
	<ul class="nav navbar-nav">
		{{if .IsLogin}}
		<li><a href="/login?exit=true">退出登录</a></li>
		{{else}}
		{{end}}
	</ul>
</div>
{{end}}