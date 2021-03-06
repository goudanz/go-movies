package utils

import (
	"github.com/go-redis/redis/v7"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/panjf2000/ants/v2"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 爬取网站的域名
//const host = "http://www.jisudhw.com"

//const host = "http://www.okzy.co"
const host = "http://www.okzyw.com"

// redis key

// 分类key
const CategoriesKey = "categories"

// 电影详情
const moviesDetail = "movies_detail:"

type Categories struct {
	Link string       `json:"link"`
	Name string       `json:"name"`
	Sub  []Categories `json:"sub"`
}

type Movies struct {
	Link      string `json:"link"`
	Name      string `json:"name"`
	Category  string `json:"category"`
	Cover     string `json:"cover"`
	UpdatedAt string `json:"updated_at"`
}

type MoviesDetail struct {
	Link     string                 `json:"link"`
	Name     string                 `json:"name"`
	Cover    string                 `json:"cover"`
	Quality  string                 `json:"quality"`
	Score    string                 `json:"score"`
	KuYun    string                 `json:"ku_yun"`
	CK       string                 `json:"ckm3u8"`
	Download string                 `json:"download"`
	Detail   map[string]interface{} `json:"detail"`
}

var (
	Smutex sync.Mutex
)

func StartSpider() {
	// 获取所有分类
	Categories := SpiderOKCategories()
	for _, v := range Categories {

		cateUrl := v.Link
		// 爬取所有主类下面的商品
		go SpiderOKMovies(cateUrl)
	}
}

// 爬取所有类别
func SpiderOKCategories() []Categories {

	c := colly.NewCollector(
		colly.Async(true),
	)

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.WithTransport(&http.Transport{
		DisableKeepAlives: true,
	})

	retryCount := 0

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(res *colly.Response, err error) {
		log.Println("Something went wrong:", err)
		if retryCount < 3 {
			retryCount += 1
			_retryErr := res.Request.Retry()
			log.Println("retry wrong:", _retryErr)
		}
	})

	// 主类
	Cate := make([]Categories, 0)

	// 导航栏、分类
	c.OnHTML("ul#sddm li", func(e *colly.HTMLElement) {

		categoryLink := e.ChildAttr("a", "href")

		categoryName := e.ChildText("a[onmouseout]")

		// 子类
		SubCate := make([]Categories, 0)

		e.ForEach("div a", func(i int, element *colly.HTMLElement) {

			subCategoryLink := element.Attr("href")
			subCategoryName := element.Text

			_subCate := Categories{
				Link: subCategoryLink,
				Name: subCategoryName,
			}

			if subCategoryName != categoryName {
				// 追加
				Smutex.Lock()
				SubCate = append(SubCate, _subCate)
				Smutex.Unlock()
			}

		})

		// 主类别
		_cate := Categories{
			Link: categoryLink,
			Name: categoryName,
			Sub:  SubCate,
		}

		// 去掉首页、福利、综艺片、解说 链接
		if categoryName != "" && categoryName != "福利片" && categoryName != "综艺片" && categoryName != "解说" {
			// 追加
			Smutex.Lock()
			Cate = append(Cate, _cate)
			Smutex.Unlock()
		}

	})

	// 在OnHTML之后被调用
	c.OnScraped(func(_ *colly.Response) {

		categories, _ := Json.MarshalIndent(Cate, "", " ")

		Smutex.Lock()
		err := RedisDB.Set(CategoriesKey, string(categories), 0).Err()
		Smutex.Unlock()
		log.Println(err)

	})

	visitError := c.Visit(host)

	log.Println(visitError)

	c.Wait()

	return Cate
}

var wg sync.WaitGroup

// 爬取所有类别的电影
func SpiderOKMovies(cateUrl string) {

	defer ants.Release()

	c := colly.NewCollector(
		colly.Async(true),
	)

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.WithTransport(&http.Transport{
		DisableKeepAlives: true,
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	retryCount := 0
	c.OnError(func(res *colly.Response, err error) {
		log.Println("Something went wrong:", err)
		if retryCount < 3 {
			retryCount += 1
			_retryErr := res.Request.Retry()
			log.Println("retry wrong:", _retryErr)
		}
	})

	var lastPageInt int

	c.OnHTML(".pages input[type=button]", func(e *colly.HTMLElement) {

		lastPageStr := e.Attr("onclick")

		lastPageStrSplit := strings.Split(lastPageStr, ",")[1]

		// 最后一页
		lastPage, _ := strconv.Atoi(strings.Split(lastPageStrSplit, ")")[0])

		lastPageInt = lastPage // todo lastPage

		// todo 有时间在研究一下这个用法
		p, _ := ants.NewPoolWithFunc(100, func(i interface{}) {
			wg.Done()
		})
		defer p.Release()

		for j := 1; j <= lastPageInt; j++ {

			wg.Add(1)
			pageUrl := CategoryToPageUrl(cateUrl, strconv.Itoa(j))

			// todo 使用 goroutine 内存跟cpu消耗太高。 暂时没找到解决方案
			ForeachPage(cateUrl, pageUrl)

			// 完成一个分类删除所有缓存
			if j == lastPageInt {
				go DelAllListCacheKey()
			}
		}
		wg.Wait()
	})

	visitError := c.Visit(host + cateUrl)

	log.Println(visitError)

	c.Wait()
}

// 获取电影详情信息
func ForeachPage(cateUrl string, url string) {

	c := colly.NewCollector(
		colly.Async(true),
	)

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.WithTransport(&http.Transport{
		DisableKeepAlives: true,
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	retryCount := 0
	c.OnError(func(res *colly.Response, err error) {
		log.Println("Something went wrong:", err)
		if retryCount < 3 {
			retryCount += 1
			_retryErr := res.Request.Retry()
			log.Println("retry wrong:", _retryErr)
		}
	})

	// 导航栏、分类
	c.OnHTML(".xing_vb li", func(e *colly.HTMLElement) {

		spanClass := e.ChildAttr("span", "class")

		// 列表数据
		if spanClass == "tt" {
			link := e.ChildAttr("a", "href")
			name := e.ChildText("a")
			category := e.ChildText(".xing_vb5")
			updateAt := e.ChildText(".xing_vb6")

			_movies := Movies{
				Link:      link,
				Name:      name,
				Category:  category,
				UpdatedAt: updateAt,
			}

			// 模板时间
			timeTemplate := "2006-01-02"
			stamp1, _ := time.ParseInLocation(timeTemplate, updateAt, time.Local)

			Smutex.Lock()
			RedisDB.ZAdd("detail_links:id:"+TransformId(cateUrl), &redis.Z{
				Score:  float64(stamp1.Unix()),
				Member: link,
			})
			Smutex.Unlock()

			// 首页香港剧单独存一个集合
			if _movies.Category == "香港剧" {
				Smutex.Lock()
				RedisDB.ZAdd("detail_links:hk", &redis.Z{
					Score:  float64(stamp1.Unix()),
					Member: link,
				})
				Smutex.Unlock()
			}

			// 获取详情
			go MoviesInfo(link)
		}
	})

	visitError := c.Visit(host + url)

	log.Println(visitError)
	log.Println("当前页面")
	log.Println(url)
	c.Wait()
}

func MoviesInfo(url string) MoviesDetail {

	c := colly.NewCollector(
		colly.Async(true),
	)

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.WithTransport(&http.Transport{
		DisableKeepAlives: true,
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	retryCount := 0
	c.OnError(func(res *colly.Response, err error) {
		log.Println("Something went wrong:", err)
		if retryCount < 3 {
			retryCount += 1
			_retryErr := res.Request.Retry()
			log.Println("retry wrong:", _retryErr)
		}
	})

	// 所有电影
	md := MoviesDetail{}

	detail := make(map[string]interface{})

	var kuyunAry []map[string]string

	var ckm3u8Ary []map[string]string

	var downloadAry []map[string]string

	c.OnHTML(".warp", func(e *colly.HTMLElement) {

		cover := e.ChildAttr("div .vodImg>img", "src")
		name := e.ChildText("div .vodh>h2")
		quality := e.ChildText("div .vodh span")
		score := e.ChildText("div .vodh label")

		// 有些页面 1 是 ckm3u8  2 是 kuyun  wtf!

		e.ForEach("div #1 ul li", func(i int, element *colly.HTMLElement) {

			playLink := element.ChildAttr("input", "value")

			Episode := strconv.Itoa(i + 1)

			if strings.Index(playLink, "m3u8") == -1 {
				kuyun := map[string]string{
					"episode":   Episode,
					"play_link": playLink}

				Smutex.Lock()
				kuyunAry = append(kuyunAry, kuyun)
				Smutex.Unlock()
			} else {
				ckm3u8 := map[string]string{
					"episode":   Episode,
					"play_link": playLink}
				Smutex.Lock()
				ckm3u8Ary = append(ckm3u8Ary, ckm3u8)
				Smutex.Unlock()
			}

		})

		e.ForEach("div #2 ul li", func(i int, element *colly.HTMLElement) {

			playLink := element.ChildAttr("input", "value")

			Episode := strconv.Itoa(i + 1)

			if strings.Index(playLink, "m3u8") == -1 {
				kuyun := map[string]string{
					"episode":   Episode,
					"play_link": playLink}

				Smutex.Lock()
				kuyunAry = append(kuyunAry, kuyun)
				Smutex.Unlock()
			} else {
				ckm3u8 := map[string]string{
					"episode":   Episode,
					"play_link": playLink}
				Smutex.Lock()
				ckm3u8Ary = append(ckm3u8Ary, ckm3u8)
				Smutex.Unlock()
			}
		})

		e.ForEach("div #down_1 ul li", func(i int, element *colly.HTMLElement) {

			playLink := element.ChildAttr("input", "value")

			Episode := strconv.Itoa(i + 1)

			download := map[string]string{
				"episode":   Episode,
				"play_link": playLink}

			Smutex.Lock()
			downloadAry = append(downloadAry, download)
			Smutex.Unlock()
		})

		kuyunAryJson, _ := Json.MarshalIndent(kuyunAry, "", " ")
		ckm3u8AryJson, _ := Json.MarshalIndent(ckm3u8Ary, "", " ")
		downloadAryJson, _ := Json.MarshalIndent(downloadAry, "", " ")

		// detail["alias"] = e.ChildText("div .vodinfobox>ul>li:eq(0)") // WTF 不支持这样的选择器
		// xpath 还是靠谱
		// 别名
		c.OnXML("/html/body/div[5]/div[1]/div/div/div[2]/div[2]/ul/li[1]/span", func(e *colly.XMLElement) {
			detail["alias"] = e.Text
		})

		// 导演
		c.OnXML("/html/body/div[5]/div[1]/div/div/div[2]/div[2]/ul/li[2]/span", func(e *colly.XMLElement) {
			detail["director"] = e.Text
		})

		// 主演
		c.OnXML("/html/body/div[5]/div[1]/div/div/div[2]/div[2]/ul/li[3]/span", func(e *colly.XMLElement) {
			detail["starring"] = e.Text
		})

		// 类型
		c.OnXML("/html/body/div[5]/div[1]/div/div/div[2]/div[2]/ul/li[4]/span", func(e *colly.XMLElement) {
			detail["type"] = e.Text
		})

		// 地区
		c.OnXML("/html/body/div[5]/div[1]/div/div/div[2]/div[2]/ul/li[5]/span", func(e *colly.XMLElement) {
			detail["area"] = e.Text
		})

		c.OnXML("/html/body/div[5]/div[1]/div/div/div[2]/div[2]/ul/li[6]/span", func(e *colly.XMLElement) {
			detail["language"] = e.Text
		})

		// 上映时间
		c.OnXML("/html/body/div[5]/div[1]/div/div/div[2]/div[2]/ul/li[7]/span", func(e *colly.XMLElement) {
			detail["released"] = e.Text
		})

		// 片长
		c.OnXML("/html/body/div[5]/div[1]/div/div/div[2]/div[2]/ul/li[8]/span", func(e *colly.XMLElement) {
			detail["length"] = e.Text
		})

		// 更新时间
		c.OnXML("/html/body/div[5]/div[1]/div/div/div[2]/div[2]/ul/li[9]/span", func(e *colly.XMLElement) {
			detail["update"] = e.Text
		})

		// 总播放量
		c.OnXML("/html/body/div[5]/div[1]/div/div/div[2]/div[2]/ul/li[10]/span", func(e *colly.XMLElement) {
			detail["total_playback"] = e.Text
		})

		// 剧情简介
		c.OnXML("/html/body/div[5]/div[3]/div[2]", func(e *colly.XMLElement) {
			detail["vod_play_info"] = e.Text
		})

		if detail["vod_play_info"] == "" || detail["vod_play_info"] == nil {
			c.OnXML("/html/body/div[5]/div[2]/div[2]/text()", func(e *colly.XMLElement) {
				detail["vod_play_info"] = e.Text
			})
		}

		md = MoviesDetail{
			Link:     url,
			Name:     name,
			Cover:    cover,
			Quality:  quality,
			Score:    score,
			Detail:   detail,
			KuYun:    string(kuyunAryJson),
			CK:       string(ckm3u8AryJson),
			Download: string(downloadAryJson),
		}

	})

	// 在OnHTML之后被调用
	c.OnScraped(func(_ *colly.Response) {

		_moviesInfo := make(map[string]interface{})

		_moviesInfo["link"] = md.Link
		_moviesInfo["cover"] = md.Cover
		_moviesInfo["name"] = md.Name
		_moviesInfo["quality"] = md.Quality
		_moviesInfo["score"] = md.Score
		_moviesInfo["kuyun"] = md.KuYun
		_moviesInfo["ckm3u8"] = md.CK
		_moviesInfo["download"] = md.Download

		_detail, _ := Json.MarshalIndent(md.Detail, "", " ")

		_moviesInfo["detail"] = string(_detail)

		if md.Name != "" {
			Smutex.Lock()
			t := RedisDB.HMSet(moviesDetail+url+":movie_name:"+md.Name, _moviesInfo).Err()
			log.Println(t)
			Smutex.Unlock()
		}

	})

	visitError := c.Visit(host + url)

	log.Println(visitError)

	c.Wait()

	return md
}

// /?m=vod-type-id-1.html  => /?m=vod-type-id-1-pg-1
func CategoryToPageUrl(categoryUrl string, page string) string {
	// 主类链接： /?m=vod-type-id-1.html
	// 主类的页面链接 /?m=vod-type-id-1-pg-
	categoryUrlStrSplit := strings.Split(categoryUrl, ".html")[0]

	pageUrl := categoryUrlStrSplit + "-pg-" + page + ".html"

	return pageUrl
}

// 获取url中的链接
func TransformId(Url string) string {
	UrlStrSplit := strings.Split(Url, "-id-")[1]

	return strings.TrimRight(UrlStrSplit, ".html")
}

func DelAllListCacheKey() {

	AllListCacheKey := RedisDB.Keys("movie_lists_key:detail_links:*").Val()

	// 删除已经缓存的数据
	for _, val := range AllListCacheKey {
		RedisDB.Del(val)
	}
}
