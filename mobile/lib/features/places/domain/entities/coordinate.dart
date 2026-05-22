import 'package:equatable/equatable.dart';

/// Value Object — immutable + self-validating.
class Coordinate extends Equatable {
  final double latitude;
  final double longitude;

  const Coordinate._(this.latitude, this.longitude);

  factory Coordinate(double lat, double lng) {
    if (lat < -90 || lat > 90 || lng < -180 || lng > 180) {
      throw ArgumentError('invalid coordinate');
    }
    return Coordinate._(lat, lng);
  }

  @override
  List<Object?> get props => [latitude, longitude];
}
