import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';

import '../../features/auth/application/auth_providers.dart';
import '../../features/auth/presentation/pages/login_page.dart';
import '../../features/auth/presentation/pages/register_page.dart';
import '../../features/places/presentation/pages/places_list_page.dart';

final appRouterProvider = Provider<GoRouter>((ref) {
  return GoRouter(
    initialLocation: '/login',
    redirect: (context, state) {
      final isAuth = ref.read(authNotifierProvider).isAuthenticated;
      final goingToAuth = state.matchedLocation == '/login' || state.matchedLocation == '/register';
      if (!isAuth && !goingToAuth) return '/login';
      if (isAuth && goingToAuth) return '/home';
      return null;
    },
    routes: [
      GoRoute(path: '/login', builder: (_, _) => const LoginPage()),
      GoRoute(path: '/register', builder: (_, _) => const RegisterPage()),
      GoRoute(path: '/home', builder: (_, _) => const PlacesListPage()),
    ],
  );
});
