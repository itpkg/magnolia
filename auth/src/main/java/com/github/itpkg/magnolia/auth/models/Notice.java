package com.github.itpkg.magnolia.auth.models;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Table;

/**
 * Created by bandari on 16-10-5.
 */
@Entity
@Table(name = "notices")
public class Notice extends Model {
    @Column(nullable = false, columnDefinition = "TEXT")
    private String body;

}
