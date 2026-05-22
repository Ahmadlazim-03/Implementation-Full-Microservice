import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../../application/places_providers.dart';

class PlacesListPage extends ConsumerStatefulWidget {
  const PlacesListPage({super.key});

  @override
  ConsumerState<PlacesListPage> createState() => _PlacesListPageState();
}

class _PlacesListPageState extends ConsumerState<PlacesListPage> {
  @override
  void initState() {
    super.initState();
    Future.microtask(() => ref.read(placesNotifierProvider.notifier).loadInitial());
  }

  @override
  Widget build(BuildContext context) {
    final state = ref.watch(placesNotifierProvider);
    return Scaffold(
      appBar: AppBar(title: const Text('Kampus Map')),
      body: state.isLoading
          ? const Center(child: CircularProgressIndicator())
          : state.error != null
              ? Center(child: Text(state.error!))
              : Column(
                  children: [
                    SizedBox(
                      height: 50,
                      child: ListView.separated(
                        scrollDirection: Axis.horizontal,
                        padding: const EdgeInsets.all(8),
                        itemCount: state.categories.length + 1,
                        separatorBuilder: (_, _) => const SizedBox(width: 8),
                        itemBuilder: (_, i) {
                          if (i == 0) {
                            return ChoiceChip(
                              label: const Text('Semua'),
                              selected: state.selectedCategoryId == null,
                              onSelected: (_) =>
                                  ref.read(placesNotifierProvider.notifier).filterByCategory(null),
                            );
                          }
                          final cat = state.categories[i - 1];
                          return ChoiceChip(
                            label: Text(cat.name),
                            selected: state.selectedCategoryId == cat.id,
                            onSelected: (_) => ref
                                .read(placesNotifierProvider.notifier)
                                .filterByCategory(cat.id),
                          );
                        },
                      ),
                    ),
                    Expanded(
                      child: ListView.builder(
                        itemCount: state.places.length,
                        itemBuilder: (_, i) {
                          final p = state.places[i];
                          return ListTile(
                            leading: const Icon(Icons.place),
                            title: Text(p.name),
                            subtitle: Text(p.address),
                            trailing: const Icon(Icons.chevron_right),
                            onTap: () {},
                          );
                        },
                      ),
                    ),
                  ],
                ),
    );
  }
}
