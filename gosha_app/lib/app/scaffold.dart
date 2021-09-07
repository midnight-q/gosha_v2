import 'dart:io';

import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:gosha_app/app/app_bar.dart';
import 'package:gosha_app/controller/app_controller.dart';
import 'package:gosha_app/app/body.dart';

class GoshaScaffold extends StatelessWidget {
  final String title;

  GoshaScaffold({Key? key, required this.title}) : super(key: key);

  AppController controller = Get.put(AppController());

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: GoshaAppBar(title: this.title),
      body: BodyScreen(),
      floatingActionButton: Obx(
        () => Visibility(
          child: FloatingActionButton(
            onPressed: () {},
            child: Icon(Icons.add, size: 28),
          ),
          visible: controller.isAppExist.value,
        ),
      ),
    );
  }
}
