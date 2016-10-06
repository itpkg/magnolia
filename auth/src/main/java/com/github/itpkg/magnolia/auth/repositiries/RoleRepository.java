package com.github.itpkg.magnolia.auth.repositiries;

import com.github.itpkg.magnolia.auth.models.Role;
import org.springframework.data.repository.CrudRepository;

/**
 * Created by flamen on 16-9-19.
 */
public interface RoleRepository extends CrudRepository<Role, Long> {
    //    @Query("select t from Role t where t.name = ?1 and t.resourceType = ?2 and t.resourceId = ?3")
    Role findByNameAndResourceTypeAndResourceId(String name, String resourceType, Long resourceId);
}
