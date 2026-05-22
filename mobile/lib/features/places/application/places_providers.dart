import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:get_it/get_it.dart';

import '../domain/usecases/list_categories.dart';
import '../domain/usecases/list_places.dart';
import 'places_notifier.dart';
import 'places_state.dart';

final placesNotifierProvider = StateNotifierProvider<PlacesNotifier, PlacesState>((ref) {
  final sl = GetIt.instance;
  return PlacesNotifier(
    listPlaces: sl<ListPlaces>(),
    listCategories: sl<ListCategories>(),
  );
});
