import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/errors/failures.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/auth_session.dart';
import '../repositories/auth_repository.dart';

class RegisterUser implements UseCase<AuthSession, RegisterParams> {
  final AuthRepository repository;
  RegisterUser(this.repository);

  @override
  Future<Either<Failure, AuthSession>> call(RegisterParams params) {
    return repository.register(
      email: params.email,
      password: params.password,
      name: params.name,
    );
  }
}

class RegisterParams extends Equatable {
  final String email;
  final String password;
  final String name;
  const RegisterParams({
    required this.email,
    required this.password,
    required this.name,
  });

  @override
  List<Object?> get props => [email, password, name];
}
