import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/errors/failures.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/auth_session.dart';
import '../repositories/auth_repository.dart';

class LoginUser implements UseCase<AuthSession, LoginParams> {
  final AuthRepository repository;
  LoginUser(this.repository);

  @override
  Future<Either<Failure, AuthSession>> call(LoginParams params) {
    return repository.login(email: params.email, password: params.password);
  }
}

class LoginParams extends Equatable {
  final String email;
  final String password;
  const LoginParams({required this.email, required this.password});

  @override
  List<Object?> get props => [email, password];
}
