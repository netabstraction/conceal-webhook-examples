package io.conceal.webhook.spring.example;

import org.springframework.stereotype.Component;

@Component
public interface MiddleWare {

    public boolean apiKeyValidator(String apiKey);

    public boolean timeStampValidator(Long timeStamp);

    public boolean signatureValidator(String timeStamp, String signature);
    
}