package io.conceal.webhook.spring.example;

import org.apache.commons.codec.digest.HmacUtils;
import org.springframework.stereotype.Service;

@Service
public class MiddleWareService implements MiddleWare {

    private static final String SIGNATURE_KEY_CONSTANT = "signature-key";
    private static final String API_VALUE_CONSTANT = "sample-key";
    private static final String WEB_HOOK_URL = "http://localhost:8080/api-key-signature-protected";

    @Override
    public boolean apiKeyValidator(final String apiKey) {
        if(apiKey == null || !apiKey.equals(API_VALUE_CONSTANT)) {
            return false;
        }

        return true;
    }

    @Override
    public boolean timeStampValidator(final Long timeStamp) {
        final long currentTime = System.currentTimeMillis() / 1000;

        if (timeStamp - currentTime < -60000 ||
            timeStamp - currentTime > 120000) {
            return false;
        }

        return true;
    }

    @Override
    public boolean signatureValidator(final String timeStamp, final String signature) {
        final String messasge = timeStamp + WEB_HOOK_URL;
        final String expextedSignature = HmacUtils.hmacSha256Hex(SIGNATURE_KEY_CONSTANT, messasge);

        if (signature != null && signature.equals(expextedSignature)) {
            return false;
        }

        return true;
    }

    
}