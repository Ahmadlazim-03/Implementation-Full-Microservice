import 'package:dartz/dartz.dart';

import '../../../../core/errors/failures.dart';
import '../entities/category.dart';
import '../entities/place.dart';

abstract class PlaceRepository {
  Future<Either<Failure, List<Place>>> listPlaces({String? categoryId, String? search});
  Future<Either<Failure, Place>> getPlace(String id);
  Future<Either<Failure, List<Category>>> listCategories();
}
