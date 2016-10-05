package magnolia.auth.models;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Table;

/**
 * Created by bandari on 16-10-5.
 */
@Entity
@Table(name = "settings")
public class Setting extends Model {
    @Column(nullable = false, unique = true)
    private String key;
    @Column(nullable = false, columnDefinition = "TEXT")
    private String val;
    @Column(nullable = false)
    private boolean encode;

    public String getKey() {
        return key;
    }

    public void setKey(String key) {
        this.key = key;
    }

    public String getVal() {
        return val;
    }

    public void setVal(String val) {
        this.val = val;
    }

    public boolean isEncode() {
        return encode;
    }

    public void setEncode(boolean encode) {
        this.encode = encode;
    }
}
