import 'package:dartz/dartz.dart';

import '../../../../core/errors/failures.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/category.dart';
import '../repositories/place_repository.dart';

class ListCategories implements UseCase<List<Category>, NoParams> {
  final PlaceRepository repository;
  ListCategories(this.repository);

  @override
  Future<Either<Failure, List<Category>>> call(NoParams params) {
    return repository.listCategories();
  }
}
