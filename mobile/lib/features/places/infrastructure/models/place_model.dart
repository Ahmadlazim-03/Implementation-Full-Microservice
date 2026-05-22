import '../../domain/entities/coordinate.dart';
import '../../domain/entities/place.dart';

class PlaceModel extends Place {
  const PlaceModel({
    required super.id,
    required super.categoryId,
    required super.name,
    required super.location,
    required super.address,
    required super.description,
  });

  factory PlaceModel.fromJson(Map<String, dynamic> json) => PlaceModel(  // ignore: prefer_const_constructors
        id: json['id'] as String,
        categoryId: json['category_id'] as String,
        name: json['name'] as String,
        location: Coordinate(
          (json['latitude'] as num).toDouble(),
          (json['longitude'] as num).toDouble(),
        ),
        address: json['address'] as String? ?? '',
        description: json['description'] as String? ?? '',
      );
}
