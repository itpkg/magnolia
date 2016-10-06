Rails.application.routes.draw do
  mount Ops::Engine => "/ops"
end
