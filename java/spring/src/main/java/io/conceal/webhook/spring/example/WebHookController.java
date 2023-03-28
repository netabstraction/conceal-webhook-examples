package io.conceal.webhook.spring.example;

import java.io.Console;

import org.apache.commons.codec.digest.HmacUtils;
import org.apache.commons.logging.Log;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.ResponseBody;

import io.conceal.webhook.spring.example.models.ConcealRequest;
import io.conceal.webhook.spring.example.models.ConcealResponse;

@Controller
@EnableAutoConfiguration
@SpringBootApplication
public class WebHookController {

    @PostMapping("/api-key-signature-protected")
    @ResponseBody
    public ResponseEntity<ConcealResponse> respondToConceal(
        @RequestBody final ConcealRequest payload,
        @RequestHeader("x-api-key") final String apiKey,
        @RequestHeader("conceal-timestamp") final String timeStamp,
        @RequestHeader("conceal-signature") final String signature ) throws IllegalArgumentException {

            if (!apiKeyValidator(apiKey)) {
                return new ResponseEntity<>(new ConcealResponse("API Key missing/API Key doesnot match"), HttpStatus.UNAUTHORIZED);
            }

            if (!timeStampValidator(Long.parseLong(timeStamp))) {
                return new ResponseEntity<>(new ConcealResponse("Invalid Timestamp. Timestamp not in range"), HttpStatus.BAD_REQUEST);
            }

            if (!signatureValidator(timeStamp, signature)) {
                return new ResponseEntity<>(new ConcealResponse("Invalid Signature"), HttpStatus.BAD_REQUEST);
            }

            // DEMO: Print the payload
            return new ResponseEntity<>(new ConcealResponse(payload.toString()), HttpStatus.OK);
    }

    public static void main(String[] args) {
        SpringApplication.run(WebHookController.class, args);
    }
    
    private static final String SIGNATURE_KEY_CONSTANT = "signature-key";
    private static final String API_VALUE_CONSTANT = "sample-key";
    private static final String WEB_HOOK_URL = "http://localhost:8080/api-key-signature-protected";

    private boolean apiKeyValidator(final String apiKey) {
        if (apiKey == null || !apiKey.equals(API_VALUE_CONSTANT)) {
            return false;
        }

        return true;
    }

    private boolean timeStampValidator(final Long timeStamp) {
        final long currentTime = System.currentTimeMillis() / 1000;

        if (timeStamp - currentTime < -60000 ||
                timeStamp - currentTime > 120000) {
            return false;
        }

        return true;
    }

    private boolean signatureValidator(final String timeStamp, final String signature) {
        final String messasge = timeStamp + WEB_HOOK_URL;
        final String expextedSignature = HmacUtils.hmacSha256Hex(SIGNATURE_KEY_CONSTANT, messasge);

        if (signature != null && signature.equals(expextedSignature)) {
            return false;
        }

        return true;
    }
}

