import '../../domain/entities/category.dart';

class CategoryModel extends Category {
  const CategoryModel({required super.id, required super.name, required super.icon});

  factory CategoryModel.fromJson(Map<String, dynamic> json) => CategoryModel(
        id: json['id'] as String,
        name: json['name'] as String,
        icon: json['icon'] as String? ?? '',
      );
}
