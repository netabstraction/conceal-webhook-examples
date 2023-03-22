package io.conceal.webhook.spring.example;

import java.io.Console;

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

    @Autowired
    private MiddleWareService middleWareService;

    @PostMapping("/api-key-signature-protected")
    @ResponseBody
    public ResponseEntity<ConcealResponse> respondToConceal(
        @RequestBody final ConcealRequest payload,
        @RequestHeader("x-api-key") final String apiKey,
        @RequestHeader("conceal-timestamp") final String timeStamp,
        @RequestHeader("conceal-signature") final String signature ) throws IllegalArgumentException {

            if (!middleWareService.apiKeyValidator(apiKey)) {
                return new ResponseEntity<>(new ConcealResponse("API Key missing/API Key doesnot match"), HttpStatus.UNAUTHORIZED);
            }

            if (!middleWareService.timeStampValidator(Long.parseLong(timeStamp))) {
                return new ResponseEntity<>(new ConcealResponse("Invalid Timestamp. Timestamp not in range"), HttpStatus.BAD_REQUEST);
            }

            if (!middleWareService.signatureValidator(timeStamp, signature)) {
                return new ResponseEntity<>(new ConcealResponse("Invalid Signature"), HttpStatus.BAD_REQUEST);
            }

            // DEMO: Print the payload
            return new ResponseEntity<>(new ConcealResponse(payload.toString()), HttpStatus.OK);
    }

    public static void main(String[] args) {
        SpringApplication.run(WebHookController.class, args);
    }
    
}

