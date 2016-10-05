package magnolia.auth.helpers;

import org.springframework.stereotype.Component;

import java.io.UnsupportedEncodingException;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;

/**
 * Created by flamen on 16-9-19.
 */
@Component("auth.gravatarHelper")
public class GravatarHelper {
    public String logo(String email) throws NoSuchAlgorithmException, UnsupportedEncodingException {
        return String.format(
                "https://www.gravatar.com/avatar/%s",
                toHex(MessageDigest.getInstance("MD5").digest(email.getBytes("CP1252")))
        );
    }

    private String toHex(byte[] data) {
        StringBuilder buf = new StringBuilder();
        for (byte b : data) {
            buf.append(Integer.toHexString((b & 0xFF) | 0x100).substring(1, 3));

        }
        return buf.toString();
    }
}
