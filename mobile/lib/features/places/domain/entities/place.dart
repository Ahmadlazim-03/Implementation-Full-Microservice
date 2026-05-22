import 'package:equatable/equatable.dart';

import 'coordinate.dart';

/// Place adalah Entity utama (aggregate root) di bounded context Places.
class Place extends Equatable {
  final String id;
  final String categoryId;
  final String name;
  final Coordinate location;
  final String address;
  final String description;

  const Place({
    required this.id,
    required this.categoryId,
    required this.name,
    required this.location,
    required this.address,
    required this.description,
  });

  @override
  List<Object?> get props => [id, categoryId, name, location, address, description];
}
