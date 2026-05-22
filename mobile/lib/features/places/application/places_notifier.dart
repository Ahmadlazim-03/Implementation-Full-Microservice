import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../../../core/usecase/usecase.dart';
import '../domain/usecases/list_categories.dart';
import '../domain/usecases/list_places.dart';
import 'places_state.dart';

class PlacesNotifier extends StateNotifier<PlacesState> {
  final ListPlaces listPlaces;
  final ListCategories listCategories;

  PlacesNotifier({required this.listPlaces, required this.listCategories})
      : super(const PlacesState());

  Future<void> loadInitial() async {
    state = state.copyWith(isLoading: true, clearError: true);
    final catRes = await listCategories(const NoParams());
    final placesRes = await listPlaces(const ListPlacesParams());

    catRes.fold(
      (f) => state = state.copyWith(error: f.message),
      (cats) => state = state.copyWith(categories: cats),
    );
    placesRes.fold(
      (f) => state = state.copyWith(isLoading: false, error: f.message),
      (places) => state = state.copyWith(isLoading: false, places: places),
    );
  }

  Future<void> filterByCategory(String? categoryId) async {
    state = state.copyWith(isLoading: true, selectedCategoryId: categoryId, clearCategory: categoryId == null);
    final res = await listPlaces(ListPlacesParams(categoryId: categoryId));
    res.fold(
      (f) => state = state.copyWith(isLoading: false, error: f.message),
      (places) => state = state.copyWith(isLoading: false, places: places),
    );
  }
}
