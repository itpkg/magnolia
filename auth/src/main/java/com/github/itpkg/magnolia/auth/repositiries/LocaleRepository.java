package com.github.itpkg.magnolia.auth.repositiries;

import com.github.itpkg.magnolia.auth.models.Locale;
import org.springframework.data.repository.CrudRepository;

/**
 * Created by bandari on 16-10-5.
 */
public interface LocaleRepository extends CrudRepository<Locale, Long> {
    Locale findByLangAndCode(String lang, String code);
}
