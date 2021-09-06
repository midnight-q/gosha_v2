import 'package:path_provider/path_provider.dart';

Future<String> getHomeDirectory() async {
  var dir = await getApplicationDocumentsDirectory();
  return dir.parent.path;
}
