import 'dart:convert';

import 'package:http/http.dart' as http;

class HttpMethod {
  bool Find = false;
  bool Create = false;
  bool MultiCreate = false;
  bool Read = false;
  bool Update = false;
  bool MultiUpdate = false;
  bool Delete = false;
  bool MultiDelete = false;
  bool FindOrCreate = false;
  bool UpdateOrCreate = false;

  HttpMethod({
    required this.Find,
    required this.Create,
    required this.MultiCreate,
    required this.Read,
    required this.Update,
    required this.MultiUpdate,
    required this.Delete,
    required this.MultiDelete,
    required this.FindOrCreate,
    required this.UpdateOrCreate,
  });

  HttpMethod.empty() {
    Find = false;
    Create = false;
    MultiCreate = false;
    Read = false;
    Update = false;
    MultiUpdate = false;
    Delete = false;
    MultiDelete = false;
    FindOrCreate = false;
    UpdateOrCreate = false;
  }

  HttpMethod.fromJson(Map<String, dynamic> json) {
    Find = json["Find"];
    Create = json["Create"];
    MultiCreate = json["MultiCreate"];
    Read = json["Read"];
    Update = json["Update"];
    MultiUpdate = json["MultiUpdate"];
    Delete = json["Delete"];
    MultiDelete = json["MultiDelete"];
    FindOrCreate = json["FindOrCreate"];
    UpdateOrCreate = json["UpdateOrCreate"];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = Map<String, dynamic>();
    data["Find"] = Find;
    data["Create"] = Create;
    data["MultiCreate"] = MultiCreate;
    data["Read"] = Read;
    data["Update"] = Update;
    data["MultiUpdate"] = MultiUpdate;
    data["Delete"] = Delete;
    data["MultiDelete"] = MultiDelete;
    data["FindOrCreate"] = FindOrCreate;
    data["UpdateOrCreate"] = UpdateOrCreate;
    return data;
  }
}

class HttpRoute {
  String Find = "";
  String Create = "";
  String MultiCreate = "";
  String Read = "";
  String Update = "";
  String MultiUpdate = "";
  String Delete = "";
  String MultiDelete = "";
  String FindOrCreate = "";
  String UpdateOrCreate = "";

  HttpRoute({
    required this.Find,
    required this.Create,
    required this.MultiCreate,
    required this.Read,
    required this.Update,
    required this.MultiUpdate,
    required this.Delete,
    required this.MultiDelete,
    required this.FindOrCreate,
    required this.UpdateOrCreate,
  });

  HttpRoute.empty() {
    Find = "";
    Create = "";
    MultiCreate = "";
    Read = "";
    Update = "";
    MultiUpdate = "";
    Delete = "";
    MultiDelete = "";
    FindOrCreate = "";
    UpdateOrCreate = "";
  }

  HttpRoute.fromJson(Map<String, dynamic> json) {
    Find = json["Find"].toString();
    Create = json["Create"].toString();
    MultiCreate = json["MultiCreate"].toString();
    Read = json["Read"].toString();
    Update = json["Update"].toString();
    MultiUpdate = json["MultiUpdate"].toString();
    Delete = json["Delete"].toString();
    MultiDelete = json["MultiDelete"].toString();
    FindOrCreate = json["FindOrCreate"].toString();
    UpdateOrCreate = json["UpdateOrCreate"].toString();
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = Map<String, dynamic>();
    data["Find"] = Find;
    data["Create"] = Create;
    data["MultiCreate"] = MultiCreate;
    data["Read"] = Read;
    data["Update"] = Update;
    data["MultiUpdate"] = MultiUpdate;
    data["Delete"] = Delete;
    data["MultiDelete"] = MultiDelete;
    data["FindOrCreate"] = FindOrCreate;
    data["UpdateOrCreate"] = UpdateOrCreate;
    return data;
  }
}

class Field {
  String Name = "";
  String Type = "";
  String CommentType = "";
  String CommentDb = "";
  bool IsTypeField = false;
  bool IsDbField = false;
  String SourceModel = "";
  String ModelName = "";
  bool IsPointer = false;
  bool IsArray = false;

  Field({
    required this.Name,
    required this.Type,
    required this.CommentType,
    required this.CommentDb,
    required this.IsTypeField,
    required this.IsDbField,
    required this.SourceModel,
    required this.ModelName,
    required this.IsPointer,
    required this.IsArray,
  });

  Field.empty() {
    Name = "";
    Type = "";
    CommentType = "";
    CommentDb = "";
    IsTypeField = false;
    IsDbField = false;
    SourceModel = "";
    ModelName = "";
    IsPointer = false;
    IsArray = false;
  }

  Field.fromJson(Map<String, dynamic> json) {
    Name = json["Name"].toString();
    Type = json["Type"].toString();
    CommentType = json["CommentType"].toString();
    CommentDb = json["CommentDb"].toString();
    IsTypeField = json["IsTypeField"];
    IsDbField = json["IsDbField"];
    SourceModel = json["SourceModel"].toString();
    ModelName = json["ModelName"].toString();
    IsPointer = json["IsPointer"];
    IsArray = json["IsArray"];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = Map<String, dynamic>();
    data["Name"] = Name;
    data["Type"] = Type;
    data["CommentType"] = CommentType;
    data["CommentDb"] = CommentDb;
    data["IsTypeField"] = IsTypeField;
    data["IsDbField"] = IsDbField;
    data["SourceModel"] = SourceModel;
    data["ModelName"] = ModelName;
    data["IsPointer"] = IsPointer;
    data["IsArray"] = IsArray;
    return data;
  }
}

class Model {
  String Name = "";
  List<Field> Fields = List<Field>.empty();
  bool IsTypeModel = false;
  bool IsDbModel = false;
  String PkType = "";
  String CommentType = "";
  String CommentDb = "";
  Model? Filter = null;
  String TypePath = "";
  String DbPath = "";
  HttpRoute HttpRoutes = HttpRoute.empty();
  HttpMethod HttpMethods = HttpMethod.empty();
  bool IsServiceModel = false;

  Model({
    required this.Name,
    required this.Fields,
    required this.IsTypeModel,
    required this.IsDbModel,
    required this.PkType,
    required this.CommentType,
    required this.CommentDb,
    required this.Filter,
    required this.TypePath,
    required this.DbPath,
    required this.HttpRoutes,
    required this.HttpMethods,
    required this.IsServiceModel,
  });

  Model.empty() {
    Name = "";
    Fields = List<Field>.empty();
    IsTypeModel = false;
    IsDbModel = false;
    PkType = "";
    CommentType = "";
    CommentDb = "";
    Filter = null;
    TypePath = "";
    DbPath = "";
    HttpRoutes = HttpRoute.empty();
    HttpMethods = HttpMethod.empty();
    IsServiceModel = false;
  }

  Model.fromJson(Map<String, dynamic> json) {
    Name = json["Name"].toString();
    if (json["Fields"] != null) {
      final v = json["Fields"];
      final arr0 = <Field>[];
      v.forEach((v) {
        arr0.add(Field.fromJson(v));
      });
      Fields = arr0;
    }
    IsTypeModel = json["IsTypeModel"];
    IsDbModel = json["IsDbModel"];
    PkType = json["PkType"].toString();
    CommentType = json["CommentType"].toString();
    CommentDb = json["CommentDb"].toString();
    Filter = (json["Filter"] != null) ? Model.fromJson(json["Filter"]) : null;
    TypePath = json["TypePath"].toString();
    DbPath = json["DbPath"].toString();
    HttpRoutes = (json["HttpRoutes"] != null)
        ? HttpRoute.fromJson(json["HttpRoutes"])
        : HttpRoute.empty();
    HttpMethods = (json["HttpMethods"] != null)
        ? HttpMethod.fromJson(json["HttpMethods"])
        : HttpMethod.empty();
    IsServiceModel = json["IsServiceModel"];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = Map<String, dynamic>();
    data["Name"] = Name;
    final v = Fields;
    final arr0 = [];
    v.forEach((v) {
      arr0.add(v.toJson());
    });
    data["Fields"] = arr0;
    data["IsTypeModel"] = IsTypeModel;
    data["IsDbModel"] = IsDbModel;
    data["PkType"] = PkType;
    data["CommentType"] = CommentType;
    data["CommentDb"] = CommentDb;
    data["Filter"] = Filter?.toJson();
    data["TypePath"] = TypePath;
    data["DbPath"] = DbPath;
    data["HttpRoutes"] = HttpRoutes.toJson();
    data["HttpMethods"] = HttpMethods.toJson();
    data["IsServiceModel"] = IsServiceModel;
    return data;
  }
}

class ModelFindResponse {
  List<Model> theList = List<Model>.empty();
  int total = 0;

  ModelFindResponse({
    required this.theList,
    required this.total,
  });

  ModelFindResponse.fromJson(Map<String, dynamic> json) {
    if (json["List"] != null) {
      final v = json["List"];
      final arr0 = <Model>[];
      v.forEach((v) {
        arr0.add(Model.fromJson(v));
      });
      theList = arr0;
    }
    total = json["Total"]?.toInt();
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = Map<String, dynamic>();
    final v = theList;
    final arr0 = [];
    v.forEach((v) {
      arr0.add(v.toJson());
    });
    data["List"] = arr0;
    data["Total"] = total;
    return data;
  }
}

Future<ModelFindResponse> modelFind(String path) async {
  final response = await http.get(Uri.parse(
      'http://127.0.0.1:4343/api/v1/model?IsDebug=1&PerPage=1&CurrentPage=1&Pwd=$path'));
  if (response.statusCode == 200) {
    return ModelFindResponse.fromJson(jsonDecode(response.body));
  } else {
    throw Exception('Failed to load Models ${response.statusCode}');
  }

}