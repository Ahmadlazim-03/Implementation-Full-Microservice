import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:get_it/get_it.dart';

import '../domain/repositories/auth_repository.dart';
import '../domain/usecases/login_user.dart';
import '../domain/usecases/register_user.dart';
import 'auth_notifier.dart';
import 'auth_state.dart';

final authNotifierProvider = StateNotifierProvider<AuthNotifier, AuthState>((ref) {
  final sl = GetIt.instance;
  return AuthNotifier(
    loginUser: sl<LoginUser>(),
    registerUser: sl<RegisterUser>(),
    repository: sl<AuthRepository>(),
  );
});
