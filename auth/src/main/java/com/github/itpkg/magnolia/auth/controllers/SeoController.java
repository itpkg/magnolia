package com.github.itpkg.magnolia.auth.controllers;

import com.github.itpkg.magnolia.auth.helpers.TemplateHelper;
import com.github.itpkg.magnolia.auth.services.SettingService;
import cz.jiripinkas.jsitemapgenerator.RssItemBuilder;
import cz.jiripinkas.jsitemapgenerator.WebPageBuilder;
import cz.jiripinkas.jsitemapgenerator.generator.RssGenerator;
import cz.jiripinkas.jsitemapgenerator.generator.SitemapGenerator;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.client.HttpClientErrorException;

import javax.annotation.Resource;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.io.StringReader;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;

/**
 * Created by bandari on 16-10-6.
 */
@Controller("auth.seoController")
public class SeoController {
    @GetMapping(value = "/google{code:\\w+}.html")
    @ResponseBody
    public String getGoogle(@PathVariable String code) {
        String google = settingService.get("google.verify", String.class);
        if (!code.equals(google)) {
            throw new HttpClientErrorException(HttpStatus.NOT_FOUND);
        }
        return String.format("google-site-verification: google%s.html", code);

    }

    @GetMapping("/baidu_verify_{code:\\w+}.html")
    @ResponseBody
    public String getBaidu(@PathVariable String code) {
        String baidu = settingService.get("baidu.verify", String.class);
        if (!code.equals(baidu)) {
            throw new HttpClientErrorException(HttpStatus.NOT_FOUND);
        }
        return code;
    }

    @GetMapping("/robots.txt")
    public void getRobots(HttpServletResponse response) throws IOException {
        Map<String, String> map = new HashMap<>();
        map.put("home", settingService.getHome());
        String robots = settingService.get("robots.txt", String.class);
        if (robots == null) {
            templateHelper.parse("robots.txt", response.getWriter(), map);
        } else {
            templateHelper.parse(new StringReader(robots), response.getWriter(), map);
        }

    }

    //http://www.sitemaps.org/protocol.html
    @GetMapping("/sitemap.xml")
    @ResponseBody
    public String getSitemap() {

        SitemapGenerator sg = new SitemapGenerator(settingService.getHome());

        sg.addPage(new WebPageBuilder().name("index.php")
                .priorityMax().changeFreqNever().lastModNow().build());
        sg.addPage(new WebPageBuilder().name("latest.php").build());
        sg.addPage(new WebPageBuilder().name("contact.php").build());

        return sg.constructSitemapString();

    }

    @GetMapping("/rss.atom")
    @ResponseBody
    public String getRss() {
        RssGenerator rg = new RssGenerator("http://www.topjavablogs",
                "Top Java Blogs", "Best Java Blogs");
        rg.addPage(new RssItemBuilder().pubDate(new Date()).title("News Title")
                .description("News Description").name("page-name").build());
        return rg.constructRss();
    }

    @Resource
    SettingService settingService;
    @Resource
    TemplateHelper templateHelper;
}
