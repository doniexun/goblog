<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

</head>

<body>
  <header>
    <h1 class="logo">Welcome to Beego</h1>
    <h3>{{.USER_AGENT}}</h3>
    <div class="description">
      Beego is a simple & powerful Go web framework which is inspired by tornado and sinatra.
    </div>
  </header>

	<div itemscope itemtype="http://schema.org/Blog">
		{{range $k, $v := .list}}
			<article class="post" itemprop="blogPost" itemscope itemtype="http://schema.org/BlogPosting">
				<div class="meta">
					{{$k}}  -- {{$v}}
				</div>
			</article>
		{{end}}
	</div>


  <footer>
    <div class="author">
      Official website:
      <a href="http://{{.Website}}">{{.Website}}</a> /
      Contact me:
      <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
    </div>
  </footer>
  <div class="backdrop"></div>

  <script src="/static/js/reload.min.js"></script>
</body>
</html>
