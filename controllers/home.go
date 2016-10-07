package controllers

import (
	"html/template"
	"os"

	"github.com/astaxie/beego"
)

//HomeController home controller
type HomeController struct {
	beego.Controller
}

// Index home
//@router / [get]
func (p *HomeController) Index() {
	//TODO
	p.Data["Website"] = "beego.me"
	p.Data["Email"] = "astaxie@gmail.com"
	p.TplName = "index.tpl"
}

//Nginx show nginx.conf
//@router /nginx.conf [get]
func (p *HomeController) Nginx() {
	tpl := `
server {
  listen 80;
  server_name {{.Name}};
  rewrite ^(.*) https://$host$1 permanent;
}

upstream {{.Name}}_prod {
  server localhost:{{.Port}} fail_timeout=0;
}

server {
  listen 443;

  ssl  on;
  ssl_certificate  /etc/ssl/certs/{{.Name}}.crt;
  ssl_certificate_key  /etc/ssl/private/{{.Name}}.key;
  ssl_session_timeout  5m;
  ssl_protocols  SSLv2 SSLv3 TLSv1;
  ssl_ciphers  RC4:HIGH:!aNULL:!MD5;
  ssl_prefer_server_ciphers  on;

  client_max_body_size 4G;
  keepalive_timeout 10;
  proxy_buffers 16 64k;
  proxy_buffer_size 128k;

  server_name {{.Name}};
  charset utf-8;
  index index.html;
  access_log /var/log/nginx/{{.Name}}.access.log;
  error_log /var/log/nginx/{{.Name}}.error.log;

  location /(css|js|fonts|img)/ {
    access_log off;
    expires 1d;

    root {{.Root}}/static;
    try_files $uri;
  }

  location / {
    try_files @backend;
  }

  location @backend {
    proxy_set_header X-Forwarded-Proto https;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header Host $http_host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_redirect off;
    proxy_pass http://{{.Name}}_prod;
    # limit_req zone=one;
  }


}
	`

	t, err := template.New("").Parse(tpl)
	InternalServerError(p, err)
	pwd, err := os.Getwd()
	InternalServerError(p, err)

	port, err := beego.AppConfig.Int("httpport")
	InternalServerError(p, err)
	err = t.Execute(p.Ctx.ResponseWriter, struct {
		Name string
		Port int
		Root string
	}{
		Name: "www.change-me.com", //TODO load from database
		Port: port,
		Root: pwd,
	})
	InternalServerError(p, err)
}
