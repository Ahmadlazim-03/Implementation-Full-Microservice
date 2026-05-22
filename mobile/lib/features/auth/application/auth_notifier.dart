import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../domain/repositories/auth_repository.dart';
import '../domain/usecases/login_user.dart';
import '../domain/usecases/register_user.dart';
import 'auth_state.dart';

/// AuthNotifier adalah application service di presentation side.
/// Memanggil use case, mengubah state berdasarkan hasilnya.
class AuthNotifier extends StateNotifier<AuthState> {
  final LoginUser loginUser;
  final RegisterUser registerUser;
  final AuthRepository repository;

  AuthNotifier({
    required this.loginUser,
    required this.registerUser,
    required this.repository,
  }) : super(const AuthState());

  Future<void> bootstrap() async {
    final res = await repository.getCurrentUser();
    res.fold(
      (f) => state = state.copyWith(errorMessage: f.message),
      (user) => state = state.copyWith(user: user, clearError: true),
    );
  }

  Future<void> login(String email, String password) async {
    state = state.copyWith(isLoading: true, clearError: true);
    final res = await loginUser(LoginParams(email: email, password: password));
    res.fold(
      (f) => state = state.copyWith(isLoading: false, errorMessage: f.message),
      (session) => state = state.copyWith(isLoading: false, user: session.user),
    );
  }

  Future<void> register(String email, String password, String name) async {
    state = state.copyWith(isLoading: true, clearError: true);
    final res = await registerUser(RegisterParams(email: email, password: password, name: name));
    res.fold(
      (f) => state = state.copyWith(isLoading: false, errorMessage: f.message),
      (session) => state = state.copyWith(isLoading: false, user: session.user),
    );
  }

  Future<void> logout() async {
    await repository.logout();
    state = const AuthState();
  }
}
