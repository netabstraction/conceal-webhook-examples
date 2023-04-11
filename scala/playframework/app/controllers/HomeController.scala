package controllers

import javax.inject._
import play.api._
import play.api.mvc._

/**
 * This controller creates an `Action` to handle HTTP requests to the
 * application's home page.
 */
@Singleton
class HomeController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {

  var SIGNATURE_KEY_CONSTANT : String = "signature-key"
  var API_VALUE_CONSTANT : String = "sample-key"
  var WEBHOOK_URL_CONSTANT : String = "https://localhost:9000"

  /**
   * Create an Action to render an HTML page.
   *
   * The configuration in the `routes` file means that this method
   * will be called when the application receives a `GET` request with
   * a path of `/`.
   */
  def index() = Action { implicit request: Request[AnyContent] =>
    Ok(views.html.index())
  }

  def webhook() = Action { implicit request: Request[AnyContent] =>
    val apiKey = request.headers.get("x-api-key")
    val timestamp = request.headers.get("conceal-timestamp")
    val signature = request.headers.get("conceal-signature")

    val isApiKeyValid = validateApiKey(apiKey.get)
    val isTimestampValid = validateTimestamp(timestamp.get.toLong)
    val isSignatureValid = validateSignature(signature.get, timestamp.get.toLong)

    if(isApiKeyValid) {
       Forbidden("Invalid API Key")
    }

    if(isTimestampValid) {
     Forbidden("Invalid Timestamp")
    }

    if(isSignatureValid) {
       Forbidden("Invalid Signature")
    }

    Ok(s"Successful")
  }

  def validateSignature(signature: String, timestamp: Long): Boolean = {
    import javax.crypto.Mac
    import javax.crypto.spec.SecretKeySpec
    import javax.xml.bind.DatatypeConverter

      val message = timestamp + WEBHOOK_URL_CONSTANT
      val signingKey = new SecretKeySpec(SIGNATURE_KEY_CONSTANT.getBytes, "HmacSHA256")
      val mac = Mac.getInstance("HmacSHA256")
      mac.init(signingKey)

      val expectedSignature = mac.doFinal(message.getBytes())

      if(signature == null || !signature.equals(DatatypeConverter.printHexBinary(expectedSignature))) {
        return false
      }

      return true
  }

  def validateApiKey(apiKey: String): Boolean = {
      if (apiKey == null || !apiKey.equals(API_VALUE_CONSTANT)) {
        return false
      }
      return true

  }

  def validateTimestamp(timestamp: Long): Boolean = {
    val currentTime = System.currentTimeMillis()

    if (timestamp - currentTime < -60000 ||
        timestamp - currentTime > 120000) {
      return false
    }
    return true
  }
}
