require 'openssl'

class WebhookController < ApplicationController
    SIGNATURE_KEY_CONST = "signature-key"
    API_KEY_CONST = "x-api-key"
    API_KEY_VALUE_CONST = "sample-key"
    WEBHOOK_URL_CONST = "http://127.0.0.1:8080/webhook"


    def handle_webhook

        request_timestamp = request.headers['conceal-timestamp']
        request_signature = request.headers['conceal-signature']
        request_api_key = request.headers[API_KEY_CONST]
       
        # Api key validation
        if request_api_key != API_KEY_VALUE_CONST
            logger.info 'Invalid API Key'
            return render json: { error: 'Invalid API Key'}, status: :unauthorized
        end

        #Timestamp validation
        if (!is_timestamp_valid(request_timestamp))
            logger.info 'Invalid Timestamp.'
            return render json: { error: 'Invalid Timestamp.'}, status: :bad_request
        end

        #Signature validaton
        if (!is_signature_valid(request_timestamp, request_signature))
            logger.info 'Missing/Invalid Signature.'
            return render json: { error: 'Invalid Signature.'}, status: :unauthorized
        end

        #Json Body Validation
        begin
            JSON.parse(request.body.read)
        rescue
            return render json: { error: 'Invalid Body'}, status: :bad_request
        end

        # Process the webhook payload
        # ..
        log_request(request)
        # ..
        
        # Return a success response
        logger.info "OK"
        render json: { message: "OK"}
    end

    # Validate timestamp timestamp is in the range of [current_timestamp-60sec, current_timestamp_120sec]
    def is_timestamp_valid(request_timestamp)
        request_timestamp_int = request_timestamp.to_i
        current_timestamp = Time.now.to_i 
        return (request_timestamp_int - current_timestamp > -60000 and request_timestamp_int - current_timestamp < 12000)
    end

    # Validate signature
    def is_signature_valid(request_timestamp, request_signature)
        message = "#{request_timestamp}|#{WEBHOOK_URL_CONST}"
        digester  = OpenSSL::Digest::Digest.new("sha256")
        signature = OpenSSL::HMAC.hexdigest(digester, SIGNATURE_KEY_CONST, message)
        
        return request_signature == signature
    end

    # Log request
    def log_request(request)
        http_envs = {}.tap do |envs|
            request.headers.each do |key, value|
                envs[key] = value if key.downcase.starts_with?('http')
            end
        end
        logger.info "req [" + request.method.inspect +  "] " + request.url.inspect
        logger.info "headers: " + http_envs.inspect
        logger.info "body: " + request.body.read
    end
      
  end