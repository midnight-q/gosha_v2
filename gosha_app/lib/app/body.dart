

import 'package:flutter/material.dart';
import 'package:gosha_app/sdk/application.dart';

class BodyScreen extends StatelessWidget {
  const BodyScreen(this.path, this.app, {Key? key}) : super(key: key);

  final Application app;
  final String path;

  @override
  Widget build(BuildContext context) {
    if (this.app.isAppExist()) {
      return Center(
        child: Text("TODO: Models list"),
      );
    }

    return Center(
      child: Text("Create app"),
    );
  }
}
