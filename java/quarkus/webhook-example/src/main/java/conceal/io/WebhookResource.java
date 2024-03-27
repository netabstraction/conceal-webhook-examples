package conceal.io;

import javax.ws.rs.GET;
import javax.ws.rs.POST;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;
import org.apache.commons.codec.digest.HmacUtils;
import org.jboss.resteasy.reactive.RestHeader;
import javax.ws.rs.core.Response;

@Path("/webhook")
public class WebhookResource {

    @POST
    @Produces(MediaType.TEXT_PLAIN)
    public Response respondToConceal(
        @RestHeader("x-api-key") final String apiKey,
        @RestHeader("conceal-timestamp") final String timeStamp,
        @RestHeader("conceal-signature") final String signature) {
            
            if (!apiKeyValidator(apiKey)) {
                return Response.status(403).entity("API Key missing/API Key doesnot match").build();
            }

            if (!timeStampValidator(Long.parseLong(timeStamp))) {
                return Response.status(401).entity("Invalid Timestamp. Timestamp not in range").build();
            }

            if (!signatureValidator(timeStamp, signature)) {
                return Response.status(403).entity("Invalid Signature").build();
            }

            return Response.status(200).entity("OK").build();
    }

    private static final String SIGNATURE_KEY_CONSTANT = "signature-key";
    private static final String API_VALUE_CONSTANT = "sample-key";
    private static final String WEB_HOOK_URL = "http://localhost:8080/webhook";

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
        final String expectedSignature = HmacUtils.hmacSha256Hex(SIGNATURE_KEY_CONSTANT, messasge);

        if (signature != null && signature.equals(expectedSignature)) {
            return false;
        }

        return true;
    }
}