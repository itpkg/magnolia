package com.github.itpkg.magnolia;

/**
 * Created by flamen on 16-9-19.
 */
public class RedisCachePrefix implements org.springframework.data.redis.cache.RedisCachePrefix {
    @Override
    public byte[] prefix(String cacheName) {
        return String.format("%s%s", prefix, cacheName).getBytes();
    }

    private String prefix;

    public void setPrefix(String prefix) {
        this.prefix = prefix;
    }
}
