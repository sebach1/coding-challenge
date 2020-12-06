from django.shortcuts import redirect
from .models import Film, People
from django.views import generic


class IndexView(generic.ListView):
    template_name = 'jpfilms/index.html'
    context_object_name = 'films_list'

    def get_queryset(self):
        return Film.objects.order_by('-title')[:20]


class DetailView(generic.DetailView):
    model = Film
    template_name = 'jpfilms/detail.html'

def view_404(request, exception=None):
    return redirect('/movies/')
