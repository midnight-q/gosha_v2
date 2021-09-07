import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:gosha_app/app/create_app.dart';
import 'package:gosha_app/app/models_list.dart';
import 'package:gosha_app/controller/app_controller.dart';

class BodyScreen extends StatelessWidget {
  BodyScreen({Key? key}) : super(key: key);

  AppController controller = Get.find();

  @override
  Widget build(BuildContext context) {
    return Obx(() {
      if (!this.controller.isPathSelected.value) {
        return Center(
          child: Text("Select project directory"),
        );
      }
      if (!this.controller.isAppExist.value) {
        return CreateAppScreen();
      }

      return Center(
        child: ModelList(),
      );
    });
  }
}
