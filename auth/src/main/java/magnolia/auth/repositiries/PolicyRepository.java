package magnolia.auth.repositiries;

import magnolia.auth.models.Policy;
import org.springframework.data.repository.CrudRepository;

/**
 * Created by flamen on 16-9-19.
 */
public interface PolicyRepository extends CrudRepository<Policy, Long> {
    Policy findByUserIdAndRoleId(long userId, long roleId);
}
