import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:get_it/get_it.dart';

import 'core/network/dio_client.dart';
import 'core/network/network_info.dart';

// Auth
import 'features/auth/domain/repositories/auth_repository.dart';
import 'features/auth/domain/usecases/login_user.dart';
import 'features/auth/domain/usecases/register_user.dart';
import 'features/auth/infrastructure/datasources/auth_local_datasource.dart';
import 'features/auth/infrastructure/datasources/auth_remote_datasource.dart';
import 'features/auth/infrastructure/repositories/auth_repository_impl.dart';

// Places
import 'features/places/domain/repositories/place_repository.dart';
import 'features/places/domain/usecases/list_categories.dart';
import 'features/places/domain/usecases/list_places.dart';
import 'features/places/infrastructure/datasources/place_remote_datasource.dart';
import 'features/places/infrastructure/repositories/place_repository_impl.dart';

// Review
import 'features/review/domain/repositories/review_repository.dart';
import 'features/review/domain/usecases/create_review.dart';
import 'features/review/infrastructure/datasources/review_remote_datasource.dart';
import 'features/review/infrastructure/repositories/review_repository_impl.dart';

/// Composition Root — semua wiring DI ada di sini.
/// Domain & application layer tidak pernah tahu get_it.
final sl = GetIt.instance;

Future<void> initDependencies() async {
  // ===== Core =====
  sl.registerLazySingleton(() => const FlutterSecureStorage());
  sl.registerLazySingleton(() => DioClient(sl()));
  sl.registerLazySingleton<NetworkInfo>(() => NetworkInfoImpl());

  // ===== Auth =====
  sl.registerLazySingleton<AuthRemoteDataSource>(() => AuthRemoteDataSourceImpl(sl()));
  sl.registerLazySingleton<AuthLocalDataSource>(() => AuthLocalDataSourceImpl(sl()));
  sl.registerLazySingleton<AuthRepository>(
    () => AuthRepositoryImpl(remote: sl(), local: sl()),
  );
  sl.registerLazySingleton(() => LoginUser(sl()));
  sl.registerLazySingleton(() => RegisterUser(sl()));

  // ===== Places =====
  sl.registerLazySingleton<PlaceRemoteDataSource>(() => PlaceRemoteDataSourceImpl(sl()));
  sl.registerLazySingleton<PlaceRepository>(() => PlaceRepositoryImpl(remote: sl()));
  sl.registerLazySingleton(() => ListPlaces(sl()));
  sl.registerLazySingleton(() => ListCategories(sl()));

  // ===== Review =====
  sl.registerLazySingleton<ReviewRemoteDataSource>(() => ReviewRemoteDataSourceImpl(sl()));
  sl.registerLazySingleton<ReviewRepository>(() => ReviewRepositoryImpl(remote: sl()));
  sl.registerLazySingleton(() => CreateReview(sl()));
}
