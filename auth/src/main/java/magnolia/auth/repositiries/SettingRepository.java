package magnolia.auth.repositiries;

import magnolia.auth.models.Setting;
import org.springframework.data.repository.CrudRepository;

/**
 * Created by bandari on 16-10-5.
 */
public interface SettingRepository extends CrudRepository<Setting, Long> {
    Setting findByKey(String key);
}
