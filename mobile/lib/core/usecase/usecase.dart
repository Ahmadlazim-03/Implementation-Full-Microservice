import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../errors/failures.dart';

/// Kontrak universal untuk use case.
/// `T` hasil sukses, `Params` input.
/// Return `Either<Failure, T>` -> harus handle dua kemungkinan.
abstract class UseCase<T, Params> {
  Future<Either<Failure, T>> call(Params params);
}

class NoParams extends Equatable {
  const NoParams();
  @override
  List<Object?> get props => [];
}
