require 'openssl'

class WebhookController < ApplicationController
    SIGNATURE_KEY_CONST = "signature-key"
    API_KEY_CONST = "x-api-key"
    API_KEY_VALUE_CONST = "sample-key"
    WEBHOOK_URL_CONST = "http://127.0.0.1:3000/ruby/rails/api-key-signature-protected"


    def index
        puts "Received Request"

        # Api key validation
        if request.headers[API_KEY_CONST] != API_KEY_VALUE_CONST
            return render json: { error: 'Missing/Invalid API Key'}, status: :unauthorized
        end

        #Timestamp validation
        request_timestamp = request.headers['conceal-timestamp']
        request_timestamp_int = request_timestamp.to_i
        current_timestamp = Time.now.to_i 
        if (request_timestamp_int - current_timestamp < -60000 or request_timestamp_int - current_timestamp > 12000)
            return render json: { error: 'Missing/Invalid Timestamp. Timestamp out of range'}, status: :bad_request
        end

        #Signature validaton
        request_signature = request.headers['conceal-signature']
        message = "#{request_timestamp}|#{WEBHOOK_URL_CONST}"
        digester  = OpenSSL::Digest::Digest.new("sha256")
        signature = OpenSSL::HMAC.hexdigest(digester, SIGNATURE_KEY_CONST, message)
        puts request_signature
        puts signature
        puts message
        if request_signature != signature
            return render json: { error: 'Missing/Invalid Signature.'}, status: :unauthorized
        end
        
        render json: { message: '200 OK'}
    end
  end