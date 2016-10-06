package com.github.itpkg.magnolia.auth.helpers;

import com.github.itpkg.magnolia.auth.services.SettingService;
import com.github.mustachejava.DefaultMustacheFactory;
import com.github.mustachejava.Mustache;
import com.github.mustachejava.MustacheFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import javax.annotation.Resource;
import java.io.IOException;
import java.io.Reader;
import java.io.Writer;

/**
 * Created by flamen on 16-9-19.
 */
@Component("auth.templateHelper")
public class TemplateHelper {
    public <T> void parse(Reader r, Writer w, T t) throws IOException {
        Mustache mustache = factory.compile(r, "");
        mustache.execute(w, t).flush();
    }

    public <T> void parse(String f, Writer w, T t) throws IOException {
        Mustache mustache = factory.compile(f);
        mustache.execute(w, t).flush();
    }

    @PostConstruct
    void init() {
        factory = new DefaultMustacheFactory();
    }

    @Resource
    SettingService settingService;

    @Value("${server.port}")
    int port;

    private MustacheFactory factory;

}
