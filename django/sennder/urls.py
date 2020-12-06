from django.conf.urls import include, url

handler404='jpfilms.views.view_404'

urlpatterns = [
    url(r'^movies/', include('jpfilms.urls')),
]