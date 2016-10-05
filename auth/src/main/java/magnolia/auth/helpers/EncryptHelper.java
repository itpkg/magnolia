package magnolia.auth.helpers;

/**
 * Created by bandari on 16-10-5.
 */
public interface EncryptHelper {
    String encrypt(String plain);

    String decrypt(String code);

    String sum(String plain);

    boolean chk(String plain, String code);
}
