import 'package:dio/dio.dart';

import '../../../../core/constants/api_constants.dart';
import '../../../../core/errors/exceptions.dart';
import '../../../../core/network/dio_client.dart';
import '../models/review_model.dart';

abstract class ReviewRemoteDataSource {
  Future<List<ReviewModel>> listByPlace(String placeId);
  Future<ReviewModel> create({
    required String placeId,
    required String userId,
    required int rating,
    required String comment,
  });
}

class ReviewRemoteDataSourceImpl implements ReviewRemoteDataSource {
  final DioClient client;
  ReviewRemoteDataSourceImpl(this.client);

  @override
  Future<List<ReviewModel>> listByPlace(String placeId) async {
    try {
      final res = await client.dio.get(ApiConstants.reviewsByPlace(placeId));
      final list = res.data['data'] as List;
      return list.map((e) => ReviewModel.fromJson(e as Map<String, dynamic>)).toList();
    } on DioException catch (e) {
      throw ServerException(e.message ?? 'Failed', e.response?.statusCode);
    }
  }

  @override
  Future<ReviewModel> create({
    required String placeId,
    required String userId,
    required int rating,
    required String comment,
  }) async {
    try {
      final res = await client.dio.post(ApiConstants.reviews, data: {
        'place_id': placeId,
        'user_id': userId,
        'rating': rating,
        'comment': comment,
      });
      return ReviewModel.fromJson(res.data['data'] as Map<String, dynamic>);
    } on DioException catch (e) {
      throw ServerException(e.message ?? 'Failed', e.response?.statusCode);
    }
  }
}
