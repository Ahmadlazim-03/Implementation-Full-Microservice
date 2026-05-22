import 'package:dio/dio.dart';

import '../../../../core/constants/api_constants.dart';
import '../../../../core/errors/exceptions.dart';
import '../../../../core/network/dio_client.dart';
import '../models/category_model.dart';
import '../models/place_model.dart';

abstract class PlaceRemoteDataSource {
  Future<List<PlaceModel>> listPlaces({String? categoryId, String? search});
  Future<PlaceModel> getPlace(String id);
  Future<List<CategoryModel>> listCategories();
}

class PlaceRemoteDataSourceImpl implements PlaceRemoteDataSource {
  final DioClient client;
  PlaceRemoteDataSourceImpl(this.client);

  @override
  Future<List<PlaceModel>> listPlaces({String? categoryId, String? search}) async {
    try {
      final res = await client.dio.get(ApiConstants.places, queryParameters: {
        if (categoryId != null && categoryId.isNotEmpty) 'category_id': categoryId,
        if (search != null && search.isNotEmpty) 'search': search,
      });
      final list = res.data['data'] as List;
      return list.map((e) => PlaceModel.fromJson(e as Map<String, dynamic>)).toList();
    } on DioException catch (e) {
      throw ServerException(e.message ?? 'Failed to list places', e.response?.statusCode);
    }
  }

  @override
  Future<PlaceModel> getPlace(String id) async {
    try {
      final res = await client.dio.get('${ApiConstants.places}/$id');
      return PlaceModel.fromJson(res.data['data'] as Map<String, dynamic>);
    } on DioException catch (e) {
      throw ServerException(e.message ?? 'Failed to get place', e.response?.statusCode);
    }
  }

  @override
  Future<List<CategoryModel>> listCategories() async {
    try {
      final res = await client.dio.get(ApiConstants.categories);
      final list = res.data['data'] as List;
      return list.map((e) => CategoryModel.fromJson(e as Map<String, dynamic>)).toList();
    } on DioException catch (e) {
      throw ServerException(e.message ?? 'Failed to list categories', e.response?.statusCode);
    }
  }
}
