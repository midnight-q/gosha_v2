import 'dart:io';

import 'package:filesystem_picker/filesystem_picker.dart';
import 'package:flutter/material.dart';
import 'package:gosha_app/app/app_bar.dart';
import 'package:gosha_app/common/directory.dart';
import 'package:gosha_app/sdk/application.dart';
import 'package:gosha_app/app/body.dart';

class GoshaScaffold extends StatefulWidget {
  final String title;

  const GoshaScaffold({Key? key, required this.title})
      : super(key: key);

  @override
  _GoshaScaffoldState createState() => _GoshaScaffoldState();
}

class _GoshaScaffoldState extends State<GoshaScaffold> {
  Application currentApp = Application.empty();
  String currentPath = "";

  openApp() async {
    String? path = await FilesystemPicker.open(
      rootName: await getHomeDirectory()
      title: 'Select folder with project',
      context: context,
      fsType: FilesystemType.folder,
      pickText: 'Select',
      rootDirectory: Directory(await getHomeDirectory()),
    );
    var app = await getCurrentApp(path == null ? "" : path);
    setState(() {
      currentApp = app;
      if (path != null) {
        currentPath = path;
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: GoshaAppBar(
          name: currentApp.name, title: widget.title, onPressed: openApp),
      body: BodyScreen(currentPath, currentApp),
      floatingActionButton: Visibility(
        child: FloatingActionButton(
          onPressed: () {},
          child: Icon(Icons.add, size: 28),
        ),
        visible: currentApp.isAppExist(),
      ),

    );
  }
}
