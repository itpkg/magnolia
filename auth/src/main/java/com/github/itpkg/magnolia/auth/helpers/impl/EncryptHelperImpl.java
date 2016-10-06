package com.github.itpkg.magnolia.auth.helpers.impl;

import com.github.itpkg.magnolia.auth.helpers.EncryptHelper;
import org.jasypt.util.password.PasswordEncryptor;
import org.jasypt.util.password.StrongPasswordEncryptor;
import org.jasypt.util.text.StrongTextEncryptor;
import org.jasypt.util.text.TextEncryptor;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;

/**
 * Created by bandari on 16-10-5.
 */
@Component("auth.encryptHelper")
public class EncryptHelperImpl implements EncryptHelper {

    @Override
    public String encrypt(String plain) {
        return textEncryptor.encrypt(plain);
    }

    @Override
    public String decrypt(String code) {
        return textEncryptor.decrypt(code);
    }

    @Override
    public String sum(String plain) {
        return passwordEncryptor.encryptPassword(plain);
    }

    @Override
    public boolean chk(String plain, String code) {
        return passwordEncryptor.checkPassword(plain, code);
    }

    @PostConstruct
    void init() {
        passwordEncryptor = new StrongPasswordEncryptor();

        StrongTextEncryptor ste = new StrongTextEncryptor();
        ste.setPassword(secrets);
        textEncryptor = ste;
    }

    @Value("${app.secrets}")
    String secrets;
    private PasswordEncryptor passwordEncryptor;
    private TextEncryptor textEncryptor;

}
