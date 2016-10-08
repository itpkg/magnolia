{{define "meta"}}
    <meta charset="utf-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no"/>
    <meta http-equiv="x-ua-compatible" content="ie=edge"/>
    <meta name="description" content="{{T .L "site.description"}}">
    <meta name="author" content="{{T .L "site.author.name"}} {{T .L "site.author.email"}}">

    <title>{{.Title}}-{{T .L "site.sub_title"}}-{{T .L "site.title"}}</title>

    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.4/css/bootstrap.min.css" integrity="sha384-2hfp1SzUoho7/TsGGGDaFdsuuDL0LX2hnUp6VkX3CUQ2K4K+xjboZdsXyp4oUHZj" crossorigin="anonymous"/>
    <link rel="stylesheet" href="/static/css/base.css"/>

    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.0.0/jquery.min.js" integrity="sha384-THPy051/pYDQGanwU6poAc/hOdQxjnOEXzbT+OuUAFqNqFjL+4IGLBgCJC3ZOShY" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/tether/1.2.0/js/tether.min.js" integrity="sha384-Plbmg8JY28KFelvJVai01l8WyZzrYWG825m+cZ0eDDS1f7d/js6ikvy1+X+guPIB" crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.4/js/bootstrap.min.js" integrity="sha384-VjEeINv9OSwtWFLAtmc4JCtEJXXBub00gtSnszmspDLCtC0I4z4nqz7rEFbIZLLU" crossorigin="anonymous"></script>
{{end}}

{{define "nav-bar"}}
  <nav class="navbar navbar-fixed-top navbar-dark bg-inverse">
    <a href="#" class="navbar-brand">{{T .L "site.sub_title"}}</a>
    <ul class="nav navbar-nav">
      <li class="nav-item active">
        <a class="nav-link" href="#">Home <span class="sr-only">(current)</span></a>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="#">Features</a>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="#">Pricing</a>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="#">About</a>
      </li>
    </ul>
  </nav>
{{end}}

{{define "copyright"}}
    <footer>
      <p class="pull-xs-right">
        {{T .L "links.all_languages"}}
        <a href="/?locale=en-US">English</a>
        <a href="/?locale=zh-Hans">简体中文</a>
        <a href="/?locale=zh-Hant">正體中文</a>
      </p>
      <p>
        &copy; {{T .L "site.copyright"}} &middot;
        <a href="#">{{T .L "links.privacy"}}</a> &middot;
        <a href="#">{{T .L "links.terms"}}</a>
      </p>
    </footer>
{{end}}

{{define "header"}}
<!DOCTYPE html>
<html lang="{{.Locale}}">
  <head>
    {{template "meta" .}}
  </head>
  <body>
{{end}}

{{define "footer"}}
  </body>
</html>
{{end}}
