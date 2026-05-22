import 'package:dio/dio.dart';

import '../../../../core/constants/api_constants.dart';
import '../../../../core/errors/exceptions.dart';
import '../../../../core/network/dio_client.dart';
import '../models/user_model.dart';

abstract class AuthRemoteDataSource {
  Future<({String token, UserModel user})> login(String email, String password);
  Future<({String token, UserModel user})> register(String email, String password, String name);
}

class AuthRemoteDataSourceImpl implements AuthRemoteDataSource {
  final DioClient client;
  AuthRemoteDataSourceImpl(this.client);

  @override
  Future<({String token, UserModel user})> login(String email, String password) async {
    try {
      final res = await client.dio.post(ApiConstants.login, data: {
        'email': email,
        'password': password,
      });
      final data = res.data['data'] as Map<String, dynamic>;
      return (
        token: data['token'] as String,
        user: UserModel.fromJson(data['user'] as Map<String, dynamic>),
      );
    } on DioException catch (e) {
      final status = e.response?.statusCode;
      final msg = e.response?.data?['error']?['message'] ?? e.message ?? 'Login failed';
      if (status == 401) throw UnauthorizedException(msg);
      throw ServerException(msg, status);
    }
  }

  @override
  Future<({String token, UserModel user})> register(String email, String password, String name) async {
    try {
      final res = await client.dio.post(ApiConstants.register, data: {
        'email': email,
        'password': password,
        'name': name,
      });
      final data = res.data['data'] as Map<String, dynamic>;
      return (
        token: data['token'] as String,
        user: UserModel.fromJson(data['user'] as Map<String, dynamic>),
      );
    } on DioException catch (e) {
      final msg = e.response?.data?['error']?['message'] ?? e.message ?? 'Register failed';
      throw ServerException(msg, e.response?.statusCode);
    }
  }
}
