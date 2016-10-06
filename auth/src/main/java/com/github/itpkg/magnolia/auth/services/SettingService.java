package com.github.itpkg.magnolia.auth.services;

import com.github.itpkg.magnolia.auth.helpers.EncryptHelper;
import com.github.itpkg.magnolia.auth.models.Setting;
import com.github.itpkg.magnolia.auth.repositiries.SettingRepository;
import com.google.gson.Gson;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import javax.annotation.PostConstruct;
import javax.annotation.Resource;

/**
 * Created by bandari on 16-10-5.
 */
@Service("auth.settingService")
public class SettingService {
    public <T> void set(String key, T obj) {
        set(key, obj, false);
    }

    public <T> void set(String key, T obj, boolean encode) {
        Setting s = settingRepository.findByKey(key);
        if (s == null) {
            s = new Setting();
            s.setKey(key);
        }
        s.setEncode(encode);

        String val = gson.toJson(obj);
        if (encode) {
            val = encryptHelper.encrypt(val);
        }
        s.setVal(val);
        settingRepository.save(s);
    }

    public <T> T get(String key, Class<T> clazz) {
        Setting s = settingRepository.findByKey(key);
        if (s == null) {
            return null;
        }
        String val = s.getVal();
        if (s.isEncode()) {
            val = encryptHelper.decrypt(val);
        }
        return gson.fromJson(val, clazz);
    }

    public String getHome() {
        String domain = get("site.domain", String.class);
        if (isProduction()) {
            if (get("site.https?", Boolean.class)) {
                return String.format("https://%s", domain);
            }
            return String.format("http://%s", domain);
        }
        return String.format("http://localhost:%d", port);
    }

    public boolean isProduction() {
        return "production".equals(runMode);
    }


    @PostConstruct
    void init() {
        gson = new Gson();
    }

    @Value("${spring.profiles.active}")
    String runMode;
    @Value("${server.port}")
    int port;
    @Resource
    SettingRepository settingRepository;
    @Resource
    EncryptHelper encryptHelper;

    private Gson gson;

    public String getRunMode() {
        return runMode;
    }
}
