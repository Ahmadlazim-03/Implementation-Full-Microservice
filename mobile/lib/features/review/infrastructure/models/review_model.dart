import '../../domain/entities/review.dart';

class ReviewModel extends Review {
  const ReviewModel({
    required super.id,
    required super.placeId,
    required super.userId,
    required super.rating,
    required super.comment,
  });

  factory ReviewModel.fromJson(Map<String, dynamic> json) => ReviewModel(
        id: json['id'] as String,
        placeId: json['place_id'] as String,
        userId: json['user_id'] as String,
        rating: json['rating'] as int,
        comment: json['comment'] as String? ?? '',
      );
}
