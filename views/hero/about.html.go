// Code generated by hero.
// source: C:\E\project\golang\src\github_s\go_movies\views\hero\about.html
// DO NOT EDIT!
package template

import (
	"bytes"
	"time"

	"github.com/shiyanhui/hero"
)

func About(show map[string]interface{}, buffer *bytes.Buffer) {
	buffer.WriteString(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head profile="http://gmpg.org/xfn/11">
    <meta charset="UTF-8">
    <meta http-equiv="Content-Type" content="text/html">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <title>GoMovies-电影</title>
    <meta name="keywords" content="GoMovies-电影">
    <meta name="description" content="GoMovies-电影">
    <meta property="wb:webmaster" content="bec25808">
    <meta name="referrer" content="never">
    <link rel="canonical" href="/static/css/css">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, user-scalable=0, minimum-scale=1.0, maximum-scale=1.0">
    <link rel="stylesheet" type="text/css" href="/static/css/kube.css">
    <link rel="stylesheet" type="text/css" href="/static/css/reset.css">
    <link rel="stylesheet" type="text/css" href="/static/css/style.css">
    <link rel="stylesheet" type="text/css" href="//cdn.bootcss.com/font-awesome/4.7.0/css/font-awesome.min.css">
    <link rel="shortcut icon" href="/static/image/favicon.ico" type="image/x-icon" />
    <script>var killIE6ImgUrl = "/skin/66scc/images";</script>
    <!--<link rel="stylesheet" type="text/css" href="/static/css/lightbox.css" />-->
    <!--<script type='text/javascript' src='/static/js/lightbox.min.js'></script>-->
    <!--[if lt IE 9]>
    <script src="/static/js/html5.js"></script>
    <![endif]-->
</head>
`)
	buffer.WriteString(`

<body class="custom-background">

`)
	buffer.WriteString(`
    <div id="head" class="row">
        <div class="container row">
            <div class="row">
                <div id="topbar">
                    <ul id="toolbar" class="menu">

                    </ul>
                </div>
                <div id="rss">
                    <ul>
                        <li><a href="javascript:;"  title="最新更新文字版">不要点击</a>  </li>
                    </ul>
                </div>
            </div>
        </div>
        <div class="clear"></div>
        <div class="container">
            <div id="blogname" class="third"> <a href="/" title="6v电影-新版">
                    <h1>
                        go-movies      </h1>
                    <img src="/static/image/logo_.png" alt="6v电影-新版"></a> </div>

        </div>
        <div class="clear"></div>
    </div>`)
	buffer.WriteString(`<div class="mainmenus container" id="nav_b">
        <div class="mainmenu">
            <div class="topnav">
                <ul id="menus">

                    <li class="menu-item
                    `)
	if show["nav_link"] == "/" {
		hero.EscapeHTML("current_page_item", buffer)
	}
	buffer.WriteString(` ">
                    <a href="/">首页</a>
                    </li>


                    `)
	for _, category := range show["categories"].([]map[string]interface{}) {
		buffer.WriteString(`
                        <li class="menu-item
                        `)
		if category["link"].(string) == show["nav_link"] {
			hero.EscapeHTML("current_page_item", buffer)
		}
		buffer.WriteString(` ">
                    <a href="/?cate=`)
		hero.EscapeHTML(category["link"].(string), buffer)
		buffer.WriteString(` ">`)
		hero.EscapeHTML(category["name"].(string), buffer)
		buffer.WriteString(`</a>
                    </li>
                    `)
	}
	buffer.WriteString(`

                    <li class="menu-item
                    `)
	if show["nav_link"] == "/about" {
		hero.EscapeHTML("current_page_item", buffer)
	}
	buffer.WriteString(` ">
                    <a href="/about"> 关于 </a>
                    </li>

                </ul>
                <div id="select_menu">
                    <select onChange="document.location.href=this.options[this.selectedIndex].value;" id="select-menu-nav">
                        <option value="#">导航菜单</option>
                        `)
	for _, category := range show["categories"].([]map[string]interface{}) {
		buffer.WriteString(`
                        <option   value="/?cate=`)
		hero.EscapeHTML(category["link"].(string), buffer)
		buffer.WriteString(` ">`)
		hero.EscapeHTML(category["name"].(string), buffer)
		buffer.WriteString(`</option>
                        `)
	}
	buffer.WriteString(`

                        <option   value="about"> 关于 </option>
                    </select>
                </div>
            </div>
        </div>
        <div class="clear"></div>
    </div>
`)
	buffer.WriteString(`


	<div class="container">
        		        <div class="row">
			                <div class="subsidiary box">
                    <div class="bulletin fourfifth">
                        <span class="sixth">当前位置：</span>       关于              </div>
                </div>
                    </div>






      	<div class="mainleft" id="content">
			<div class="article_container row  box">
                <h1 >关于</h1>

            	<div class="clear"></div>

            <div class="context">
				<div id="post_content">
                    <p>
                        <img alt="" src="/static/image/xuexi.png">
                        <br>◎本　　站：golang + redis 实现的影站(低级爬虫)。无管理后台。
                        <a href="https://github.com/hezhizheng/go-movies" target="_blank" style="cursor:pointer;text-decoration: underline;color: #5ca9e4">使用参考</a>
                        <br>◎说　　明：bug存在，有精力在慢慢搞（PS：页面加载不正常请重新刷新）。
                        <br>◎界　　面：copy某一个电影网站的，排版请见谅。
                        <br>◎联　　系：<a href="https://www.facebook.com/hezhizheng1026" target="_blank"
                                     aria-label="hezhizheng 的 Facebook 地址">
                            <i class="fa fa-facebook fa-fw" title="Facebook"></i>
                        </a>
                        <a href="https://twitter.com/he_zhizheng" target="_blank"
                           aria-label="hezhizheng 的 Twitter 地址">
                            <i class="fa fa-twitter fa-fw" title="Twitter"></i>
                        </a>
                        <a href="http://weibo.com/u/5675317400" target="_blank"
                           aria-label="hezhizheng 的 Weibo 地址">
                            <i class="fa fa-weibo fa-fw" title="Weibo"></i>
                        </a>
                        <a href="https://github.com/hezhizheng" target="_blank"
                           aria-label="hezhizheng 的 Github 地址">
                            <i class="fa fa-github fa-fw" title="Github"></i>
                        </a>
                        <a href="https://www.instagram.com/dexter_ho_cn" target="_blank"
                           aria-label="hezhizheng 的 Instagram 地址">
                            <i class="fa fa-instagram fa-fw" title="Instagram"></i>
                        </a>
                        <a href="https://hezhizheng.com" target="_blank"
                           aria-label="hezhizheng 的 blog 地址">
                            <i class=" fa fa-bold" title="blog"></i>
                        </a>
                    </p>

                    <hr>

</div>
                &nbsp;               	<div class="clear"></div>



            </div>
		</div>

    	<div>

	</div>

  </div>
    </div>




</body>

`)
	buffer.WriteString(`
    <div class="clear"></div>
    <div id="footer">
        <div class="footnav container">
            <ul id="footnav" class="menu">
            </ul>
        </div>
        <div class="footnav container">
            <ul id="friendlink" class="menu">
            </ul>
        </div>
        <div class="copyright">
            <p> Copyright &copy; 2019- `)
	hero.EscapeHTML(time.Now().Format("2006"), buffer)
	buffer.WriteString(` <a href=""><strong>
                        GoMovies     </strong><a href="javascript:;"></a><br>

            <p class="author"> power by <a href="https://hezhizheng.com" target="_blank" rel="external">hezhizheng.com</a></p>
        </div>
    </div>

    <!--gototop-->
    <div id="tbox"> <a id="gotop" href="javascript:void(0)"></a> </div>


    <script type='text/javascript' src='/static/js/jquery.min-3.8.1.js'></script>
    <script type='text/javascript' src='/static/js/lets-kill-ie6-3.8.1.js'></script>
    <script src="/static/org/layer-v3.1.1/layer/layer.js"></script>

    <script type="text/javascript" src="/static/js/jquery.masonry.js"></script>
    <script type="text/javascript" src="/static/js/loostrive.js"></script>

    <script language="javascript" type="text/javascript">
        (function() {
            var oDiv = document.getElementById("nav_b");
            var H = 0, iE6;
            var Y = oDiv;
            while (Y) {
                H += Y.offsetTop;
                Y = Y.offsetParent;
            };
            iE6 = window.ActiveXObject && !window.XMLHttpRequest;
            if (!iE6) {
                window.onscroll = function() {
                    var s = document.body.scrollTop || document.documentElement.scrollTop;
                    if (s > H) {
                        oDiv.className = "mainmenus container nav_b";
                        if (iE6) {
                            oDiv.style.top = (s - H) + "px";
                        }
                    } else {
                        oDiv.className = "mainmenus container";
                    }
                }
            }
        })();
    </script>


    <script>
        function replaceImg(){
            $("img").each(function () {
                let realImgUrl = $(this).attr("data-url");
                if ( realImgUrl !== "" )
                {
                    $(this).attr("src",$(this).attr("data-url"))
                }
            });
        }
        setTimeout(replaceImg, 1000);
    </script>

    <style>
        .search_m{display: none;}
        @media only screen and (max-width: 640px){
            .search_m{max-width: 360px !important;margin: 0 auto;margin-bottom: 10px;display: block;}
        }
    </style>
`)
	buffer.WriteString(`

</html>
`)

}
