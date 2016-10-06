package com.github.itpkg.magnolia.cms.controllers;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;

/**
 * Created by bandari on 16-10-6.
 */
@Controller("cms.articlesController")
@RequestMapping(path = "/cms")
public class ArticlesController {
    @GetMapping("/articles")
    public String getIndex(Model model) {
        return "cms/articles/index";
    }
}
