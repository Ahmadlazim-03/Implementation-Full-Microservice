import 'package:dartz/dartz.dart';

import '../../../../core/errors/failures.dart';
import '../entities/auth_session.dart';
import '../entities/user.dart';

/// Repository sebagai port (abstract) di domain layer.
/// Implementasi konkretnya hidup di infrastructure/.
abstract class AuthRepository {
  Future<Either<Failure, AuthSession>> login({
    required String email,
    required String password,
  });

  Future<Either<Failure, AuthSession>> register({
    required String email,
    required String password,
    required String name,
  });

  Future<Either<Failure, User?>> getCurrentUser();

  Future<Either<Failure, void>> logout();
}
