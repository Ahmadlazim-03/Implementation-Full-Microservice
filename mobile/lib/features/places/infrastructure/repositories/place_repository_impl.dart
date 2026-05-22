import 'package:dartz/dartz.dart';

import '../../../../core/errors/exceptions.dart';
import '../../../../core/errors/failures.dart';
import '../../domain/entities/category.dart';
import '../../domain/entities/place.dart';
import '../../domain/repositories/place_repository.dart';
import '../datasources/place_remote_datasource.dart';

class PlaceRepositoryImpl implements PlaceRepository {
  final PlaceRemoteDataSource remote;
  PlaceRepositoryImpl({required this.remote});

  @override
  Future<Either<Failure, List<Place>>> listPlaces({String? categoryId, String? search}) async {
    try {
      final res = await remote.listPlaces(categoryId: categoryId, search: search);
      return Right(res);
    } on ServerException catch (e) {
      return Left(ServerFailure(e.message));
    }
  }

  @override
  Future<Either<Failure, Place>> getPlace(String id) async {
    try {
      final res = await remote.getPlace(id);
      return Right(res);
    } on ServerException catch (e) {
      return Left(ServerFailure(e.message));
    }
  }

  @override
  Future<Either<Failure, List<Category>>> listCategories() async {
    try {
      final res = await remote.listCategories();
      return Right(res);
    } on ServerException catch (e) {
      return Left(ServerFailure(e.message));
    }
  }
}
