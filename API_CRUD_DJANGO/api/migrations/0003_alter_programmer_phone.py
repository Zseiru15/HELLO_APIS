# Generated by Django 5.1.3 on 2024-12-03 22:20

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('api', '0002_programmer_phone'),
    ]

    operations = [
        migrations.AlterField(
            model_name='programmer',
            name='phone',
            field=models.CharField(default=None, max_length=10, null=True),
        ),
    ]
