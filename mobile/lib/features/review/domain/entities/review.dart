import 'package:equatable/equatable.dart';

class Review extends Equatable {
  final String id;
  final String placeId;
  final String userId;
  final int rating;
  final String comment;

  const Review({
    required this.id,
    required this.placeId,
    required this.userId,
    required this.rating,
    required this.comment,
  });

  @override
  List<Object?> get props => [id, placeId, userId, rating, comment];
}
