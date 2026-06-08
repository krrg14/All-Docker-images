from django.shortcuts import render, get_object_or_404
from .models import Watch

SAMPLE_WATCHES = [
    {'id': 1, 'name': 'Aurora Classic', 'description': 'Timeless automatic watch.', 'price': '299.00', 'image_url': ''},
    {'id': 2, 'name': 'Boreal Sport', 'description': 'Durable sport watch.', 'price': '199.00', 'image_url': ''},
    {'id': 3, 'name': 'Crown Heritage', 'description': 'Limited edition heritage piece.', 'price': '999.00', 'image_url': ''},
]

def index(request):
    watches = list(Watch.objects.all())
    if not watches:
        # fallback to sample watches (dicts)
        context = {'watches': SAMPLE_WATCHES}
        return render(request, 'catalog/index.html', context)
    context = {'watches': watches}
    return render(request, 'catalog/index.html', context)

def detail(request, pk):
    try:
        watch = Watch.objects.get(pk=pk)
    except Watch.DoesNotExist:
        # fallback to sample
        watch = next((w for w in SAMPLE_WATCHES if w['id'] == pk), None)
    return render(request, 'catalog/detail.html', {'watch': watch})
