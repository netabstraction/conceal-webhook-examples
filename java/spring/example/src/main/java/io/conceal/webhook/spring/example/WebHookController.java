package io.conceal.webhook.spring.example;

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

@Controller
@EnableAutoConfiguration
@SpringBootApplication
public class WebHookController {

    @Autowired
    private MiddleWareService middleWareService;

    @PostMapping("/api-key-signature-protected")
    @ResponseBody
    public ResponseEntity<String> respondToConceal(
        @RequestBody final String payload,
        @RequestHeader("x-api-key") final String apiKey,
        @RequestHeader("conceal_timestamp") final String timeStamp,
        @RequestHeader("conceal_signature") final String signature ) throws IllegalArgumentException {

            if (!middleWareService.apiKeyValidator(apiKey)) {
                return new ResponseEntity<>("API Key missing/API Key doesnot match", HttpStatus.UNAUTHORIZED);
            }

            if (!middleWareService.timeStampValidator(Long.parseLong(timeStamp))) {
                return new ResponseEntity<>("Invalid Timestamp. Timestamp not in range", HttpStatus.BAD_REQUEST);
            }

            if (!middleWareService.signatureValidator(timeStamp, signature)) {
                return new ResponseEntity<>("Invalid Signature", HttpStatus.BAD_REQUEST);
            }


            return new ResponseEntity<>("SUCCESS", HttpStatus.OK);
    }

    public static void main(String[] args) {
        SpringApplication.run(WebHookController.class, args);
    }
    
}

