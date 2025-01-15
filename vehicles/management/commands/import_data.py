import csv
from django.core.management.base import BaseCommand
from vehicles.models import Type, Brand, Model

class Command(BaseCommand):
    help = 'Import vehicle data from CSV files'

    def add_arguments(self, parser):
        parser.add_argument('--types', type=str, help='The path to the types CSV file')
        parser.add_argument('--brands', type=str, help='The path to the brands CSV file')
        parser.add_argument('--models', type=str, help='The path to the models CSV file')

    def handle(self, *args, **kwargs):
        if kwargs['types']:
            self.import_types(kwargs['types'])
        if kwargs['brands']:
            self.import_brands(kwargs['brands'])
        if kwargs['models']:
            self.import_models(kwargs['models'])

    def import_types(self, csv_file):
        with open(csv_file, newline='', encoding='utf-8') as file:
            reader = csv.DictReader(file)
            for row in reader:
                type_id = row['id']
                type_name = row['name']
                Type.objects.update_or_create(id=type_id, defaults={'name': type_name})
        self.stdout.write(self.style.SUCCESS('Successfully imported types'))

    def import_brands(self, csv_file):
        with open(csv_file, newline='', encoding='utf-8') as file:
            reader = csv.DictReader(file)
            for row in reader:
                brand_id = row['id']
                brand_name = row['name']
                type_id = row['type_id']
                vehicle_type = Type.objects.get(id=type_id)
                Brand.objects.update_or_create(id=brand_id, defaults={'name': brand_name, 'type': vehicle_type})
        self.stdout.write(self.style.SUCCESS('Successfully imported brands'))

    def import_models(self, csv_file):
        with open(csv_file, newline='', encoding='utf-8') as file:
            reader = csv.DictReader(file)
            for row in reader:
                model_id = row['id']
                model_name = row['name']
                brand_id = row['brand_id']
                type_id = row['type_id']
                brand = Brand.objects.get(id=brand_id)
                vehicle_type = Type.objects.get(id=type_id)
                Model.objects.update_or_create(id=model_id, defaults={'name': model_name, 'brand': brand, 'type': vehicle_type})
        self.stdout.write(self.style.SUCCESS('Successfully imported models'))