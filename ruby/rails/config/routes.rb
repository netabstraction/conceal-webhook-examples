Rails.application.routes.draw do
  # Define your application routes per the DSL in https://guides.rubyonrails.org/routing.html

  post "ruby/rails/api-key-signature-protected", to: "webhook#index"
end
