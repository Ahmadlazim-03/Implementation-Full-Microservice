class ApiConstants {
  ApiConstants._();

  // Ganti ke IP server / domain saat deploy.
  // Android emulator: 10.0.2.2 -> localhost host.
  static const String baseUrl = 'http://10.0.2.2:8080/api';

  // Auth
  static const String register = '/auth/register';
  static const String login = '/auth/login';

  // Places
  static const String places = '/places';
  static const String categories = '/categories';

  // Reviews
  static const String reviews = '/reviews';
  static String reviewsByPlace(String placeId) => '/reviews/place/$placeId';
}
