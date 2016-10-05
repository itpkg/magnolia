package magnolia.auth.models;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Index;
import javax.persistence.Table;

/**
 * Created by bandari on 16-10-5.
 */
@Entity
@Table(name = "roles", indexes = {
        @Index(columnList = "resourceId,resourceType,name", unique = true),
        @Index(columnList = "resourceType"),
        @Index(columnList = "name"),
}
)

public class Role extends Model {
    private String resourceType;
    private Long resourceId;
    @Column(updatable = false, nullable = false)
    private String name;

    public String getResourceType() {
        return resourceType;
    }

    public void setResourceType(String resourceType) {
        this.resourceType = resourceType;
    }

    public Long getResourceId() {
        return resourceId;
    }

    public void setResourceId(Long resourceId) {
        this.resourceId = resourceId;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
}
