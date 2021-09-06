import 'package:flutter/material.dart';
import 'package:gosha_app/app/scaffold.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Gosha',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: GoshaScaffold(title: 'Gosha'),
      debugShowCheckedModeBanner: false,
    );
  }
}
