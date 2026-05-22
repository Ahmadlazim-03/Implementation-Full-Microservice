import 'dart:convert';

import 'package:flutter_secure_storage/flutter_secure_storage.dart';

import '../../../../core/errors/exceptions.dart';
import '../models/user_model.dart';

abstract class AuthLocalDataSource {
  Future<void> saveSession(String token, UserModel user);
  Future<({String token, UserModel user})?> readSession();
  Future<void> clear();
}

class AuthLocalDataSourceImpl implements AuthLocalDataSource {
  final FlutterSecureStorage storage;
  AuthLocalDataSourceImpl(this.storage);

  static const _kToken = 'auth_token';
  static const _kUser = 'auth_user';

  @override
  Future<void> saveSession(String token, UserModel user) async {
    await storage.write(key: _kToken, value: token);
    await storage.write(key: _kUser, value: jsonEncode(user.toJson()));
  }

  @override
  Future<({String token, UserModel user})?> readSession() async {
    final token = await storage.read(key: _kToken);
    final userStr = await storage.read(key: _kUser);
    if (token == null || userStr == null) return null;
    try {
      final user = UserModel.fromJson(jsonDecode(userStr) as Map<String, dynamic>);
      return (token: token, user: user);
    } catch (_) {
      throw CacheException('corrupt session');
    }
  }

  @override
  Future<void> clear() async {
    await storage.delete(key: _kToken);
    await storage.delete(key: _kUser);
  }
}
