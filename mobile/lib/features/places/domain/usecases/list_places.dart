import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/errors/failures.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/place.dart';
import '../repositories/place_repository.dart';

class ListPlaces implements UseCase<List<Place>, ListPlacesParams> {
  final PlaceRepository repository;
  ListPlaces(this.repository);

  @override
  Future<Either<Failure, List<Place>>> call(ListPlacesParams params) {
    return repository.listPlaces(categoryId: params.categoryId, search: params.search);
  }
}

class ListPlacesParams extends Equatable {
  final String? categoryId;
  final String? search;
  const ListPlacesParams({this.categoryId, this.search});

  @override
  List<Object?> get props => [categoryId, search];
}
