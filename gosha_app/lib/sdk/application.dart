import 'dart:convert';
import 'package:http/http.dart' as http;

class Application {
  String name = "";
  String email = "";
  String password = "";
  int databaseType = 0;

  Application({
    required this.name,
    required this.email,
    required this.password,
    required this.databaseType,
  });

  Application.fromJson(Map<String, dynamic> json) {
    name = json["Name"].toString();
    email = json["Email"].toString();
    password = json["Password"].toString();
    databaseType = json["DatabaseType"].toInt();
  }

  Application.empty() {
    name = "";
    email = "";
    password = "";
    databaseType = 0;
  }

  bool isAppExist() {
    if (this.name.length > 0 ) {
      return true;
    }
    return false;
  }
}

class ApplicationFindResponse {
  List<Application?> theList = List.empty();
  int total = 0;

  ApplicationFindResponse({
    required this.theList,
    required this.total,
  });

  ApplicationFindResponse.fromJson(Map<String, dynamic> json) {
    if (json["List"] != null) {
      final v = json["List"];
      final arr0 = <Application>[];
      v.forEach((v) {
        arr0.add(Application.fromJson(v));
      });
      theList = arr0;
    }
    total = json["Total"]?.toInt();
  }
}

Future<ApplicationFindResponse> applicationFind(String path) async {
  final response = await http.get(Uri.parse(
      'http://127.0.0.1:4343/api/v1/application?IsDebug=1&PerPage=1&CurrentPage=1&Pwd=$path'));
  if (response.statusCode == 200) {
    return ApplicationFindResponse.fromJson(jsonDecode(response.body));
  } else {
    throw Exception('Failed to load Applications');
  }
}

Future<Application> getCurrentApp(String path) async {
  if (path.length < 1) {
    return Future.value(Application.empty());
  }
  try {
    var res = await applicationFind(path);
    if (res.theList.length < 1) {
      return Future.value(Application.empty());
    }
    return Future.value(res.theList[0]);
  } catch (_) {
    return Future.value(Application.empty());
  }
}
