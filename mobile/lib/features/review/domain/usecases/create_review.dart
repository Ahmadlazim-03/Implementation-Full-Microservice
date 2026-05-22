import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/errors/failures.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/review.dart';
import '../repositories/review_repository.dart';

class CreateReview implements UseCase<Review, CreateReviewParams> {
  final ReviewRepository repository;
  CreateReview(this.repository);

  @override
  Future<Either<Failure, Review>> call(CreateReviewParams params) {
    return repository.create(
      placeId: params.placeId,
      userId: params.userId,
      rating: params.rating,
      comment: params.comment,
    );
  }
}

class CreateReviewParams extends Equatable {
  final String placeId;
  final String userId;
  final int rating;
  final String comment;
  const CreateReviewParams({
    required this.placeId,
    required this.userId,
    required this.rating,
    required this.comment,
  });

  @override
  List<Object?> get props => [placeId, userId, rating, comment];
}
