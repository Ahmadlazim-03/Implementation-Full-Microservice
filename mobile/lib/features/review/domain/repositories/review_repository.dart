import 'package:dartz/dartz.dart';

import '../../../../core/errors/failures.dart';
import '../entities/review.dart';

abstract class ReviewRepository {
  Future<Either<Failure, List<Review>>> listByPlace(String placeId);
  Future<Either<Failure, Review>> create({
    required String placeId,
    required String userId,
    required int rating,
    required String comment,
  });
}
