package com.github.itpkg.magnolia.auth.controllers;

import com.github.itpkg.magnolia.auth.services.SettingService;
import cz.jiripinkas.jsitemapgenerator.generator.SitemapGenerator;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;

import javax.annotation.Resource;

/**
 * Created by bandari on 16-10-6.
 */
@Controller("auth.siteController")
@RequestMapping(path = "/site")
public class SiteController {
    @PostMapping("/ping")
    public void ping() {
        SitemapGenerator sg = new SitemapGenerator(settingService.getHome());
        sg.pingGoogle();
        sg.pingBing();
        //todo check permission
//        Map<String, Object> rst = new HashMap<>();
//
//        String home = settingService.getHome();
//        RestTemplate rest = new RestTemplate();
//        HttpHeaders headers = new HttpHeaders();
//        headers.add("Content-Type", "application/json");
//        headers.add("Accept", "*/*");
//        HttpEntity<String> requestEntity = new HttpEntity<>("", headers);
//        for (String url : new String[]{
//                "https://www.google.com",
//                "https://www.baidu.com"}) {
//            ResponseEntity<String> responseEntity = rest.exchange(
//                    String.format("%s/ping?sitemap=%s/sitemap.gz", url, home),
//                    HttpMethod.GET,
//                    requestEntity,
//                    String.class);
//
//            Map<String, Object> rsp = new HashMap<>();
//            rsp.put("status", responseEntity.getStatusCode());
//            rsp.put("body", responseEntity.getBody());
//            rst.put(url, rsp);
//        }
//        return rst;
    }

    @Resource
    SettingService settingService;
}
