import 'package:equatable/equatable.dart';

import '../domain/entities/category.dart';
import '../domain/entities/place.dart';

class PlacesState extends Equatable {
  final bool isLoading;
  final List<Place> places;
  final List<Category> categories;
  final String? selectedCategoryId;
  final String? error;

  const PlacesState({
    this.isLoading = false,
    this.places = const [],
    this.categories = const [],
    this.selectedCategoryId,
    this.error,
  });

  PlacesState copyWith({
    bool? isLoading,
    List<Place>? places,
    List<Category>? categories,
    String? selectedCategoryId,
    String? error,
    bool clearCategory = false,
    bool clearError = false,
  }) {
    return PlacesState(
      isLoading: isLoading ?? this.isLoading,
      places: places ?? this.places,
      categories: categories ?? this.categories,
      selectedCategoryId: clearCategory ? null : (selectedCategoryId ?? this.selectedCategoryId),
      error: clearError ? null : (error ?? this.error),
    );
  }

  @override
  List<Object?> get props => [isLoading, places, categories, selectedCategoryId, error];
}
