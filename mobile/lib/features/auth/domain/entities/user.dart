import 'package:equatable/equatable.dart';

/// User adalah Entity di domain layer.
/// Tidak boleh tahu apapun soal JSON, Dio, atau database.
class User extends Equatable {
  final String id;
  final String email;
  final String name;
  final String role;

  const User({
    required this.id,
    required this.email,
    required this.name,
    required this.role,
  });

  bool get isAdmin => role == 'admin';

  @override
  List<Object?> get props => [id, email, name, role];
}
