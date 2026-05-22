import 'package:dartz/dartz.dart';

import '../../../../core/errors/exceptions.dart';
import '../../../../core/errors/failures.dart';
import '../../domain/entities/review.dart';
import '../../domain/repositories/review_repository.dart';
import '../datasources/review_remote_datasource.dart';

class ReviewRepositoryImpl implements ReviewRepository {
  final ReviewRemoteDataSource remote;
  ReviewRepositoryImpl({required this.remote});

  @override
  Future<Either<Failure, List<Review>>> listByPlace(String placeId) async {
    try {
      final res = await remote.listByPlace(placeId);
      return Right(res);
    } on ServerException catch (e) {
      return Left(ServerFailure(e.message));
    }
  }

  @override
  Future<Either<Failure, Review>> create({
    required String placeId,
    required String userId,
    required int rating,
    required String comment,
  }) async {
    try {
      final res = await remote.create(
        placeId: placeId,
        userId: userId,
        rating: rating,
        comment: comment,
      );
      return Right(res);
    } on ServerException catch (e) {
      return Left(ServerFailure(e.message));
    }
  }
}
