import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'core/router/app_router.dart';
import 'features/auth/application/auth_providers.dart';
import 'injection.dart';

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await initDependencies();
  runApp(const ProviderScope(child: KampusMapApp()));
}

class KampusMapApp extends ConsumerStatefulWidget {
  const KampusMapApp({super.key});

  @override
  ConsumerState<KampusMapApp> createState() => _KampusMapAppState();
}

class _KampusMapAppState extends ConsumerState<KampusMapApp> {
  @override
  void initState() {
    super.initState();
    Future.microtask(() => ref.read(authNotifierProvider.notifier).bootstrap());
  }

  @override
  Widget build(BuildContext context) {
    final router = ref.watch(appRouterProvider);
    return MaterialApp.router(
      title: 'Kampus Map',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.indigo),
        useMaterial3: true,
      ),
      routerConfig: router,
    );
  }
}
