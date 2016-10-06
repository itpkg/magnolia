package com.github.itpkg.magnolia.auth.models;

import com.fasterxml.jackson.annotation.JsonIgnore;

import javax.persistence.*;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;

/**
 * Created by bandari on 16-10-5.
 */
@Entity
@Table(name = "users", indexes = {
        @Index(columnList = "name"),
        @Index(columnList = "providerType"),
        @Index(columnList = "providerId,providerType", unique = true)
})
public class User extends Model {
    public enum Type {
        EMAIL
    }

    @Column(nullable = false)
    private String name;
    @Column(nullable = false, unique = true)
    private String email;
    @Column(nullable = false, unique = true, length = 36)
    private String uid;
    @Column(nullable = false)
    private String logo;
    @Column(nullable = false)
    @JsonIgnore
    private String password;
    @Column(nullable = false)
    @JsonIgnore
    private String providerId;
    @Column(nullable = false)
    @Enumerated(EnumType.STRING)
    private Type providerType;

    @Column(nullable = false)
    private int signInCount;
    private String currentSignInIp;
    private Date currentSignInAt;
    private String lastSignInIp;
    private Date lastSignInAt;

    private Date confirmedAt;
    private Date lockedAt;

    @OneToMany(mappedBy = "user", fetch = FetchType.LAZY)
    @JsonIgnore
    private List<Log> logs;

    public boolean isAvailable() {
        return isConfirmed() && !isLocked();
    }

    public boolean isConfirmed() {
        return confirmedAt != null;
    }

    public boolean isLocked() {
        return lockedAt != null;
    }


    public User() {
        logs = new ArrayList<>();
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getUid() {
        return uid;
    }

    public void setUid(String uid) {
        this.uid = uid;
    }

    public String getLogo() {
        return logo;
    }

    public void setLogo(String logo) {
        this.logo = logo;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public String getProviderId() {
        return providerId;
    }

    public void setProviderId(String providerId) {
        this.providerId = providerId;
    }

    public Type getProviderType() {
        return providerType;
    }

    public void setProviderType(Type providerType) {
        this.providerType = providerType;
    }

    public int getSignInCount() {
        return signInCount;
    }

    public void setSignInCount(int signInCount) {
        this.signInCount = signInCount;
    }

    public String getCurrentSignInIp() {
        return currentSignInIp;
    }

    public void setCurrentSignInIp(String currentSignInIp) {
        this.currentSignInIp = currentSignInIp;
    }

    public Date getCurrentSignInAt() {
        return currentSignInAt;
    }

    public void setCurrentSignInAt(Date currentSignInAt) {
        this.currentSignInAt = currentSignInAt;
    }

    public String getLastSignInIp() {
        return lastSignInIp;
    }

    public void setLastSignInIp(String lastSignInIp) {
        this.lastSignInIp = lastSignInIp;
    }

    public Date getLastSignInAt() {
        return lastSignInAt;
    }

    public void setLastSignInAt(Date lastSignInAt) {
        this.lastSignInAt = lastSignInAt;
    }

    public Date getConfirmedAt() {
        return confirmedAt;
    }

    public void setConfirmedAt(Date confirmedAt) {
        this.confirmedAt = confirmedAt;
    }

    public Date getLockedAt() {
        return lockedAt;
    }

    public void setLockedAt(Date lockedAt) {
        this.lockedAt = lockedAt;
    }

    public List<Log> getLogs() {
        return logs;
    }

    public void setLogs(List<Log> logs) {
        this.logs = logs;
    }
}
