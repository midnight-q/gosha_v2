import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:gosha_app/controller/app_controller.dart';

class GoshaAppBar extends StatelessWidget implements PreferredSizeWidget {
  GoshaAppBar({Key? key, required this.title}) : super(key: key);

  final String title;

  final preferredSize = AppBar().preferredSize;

  final AppController controller = Get.find();

  @override
  Widget build(BuildContext context) {
    return AppBar(
      title: Text(title),
      actions: [
        TextButton(
            onPressed: controller.getNewPath,
            child: Padding(
              padding: EdgeInsets.only(right: 20),
              child: Obx(
                () => Text(
                  controller.getActionText(),
                  style: TextStyle(color: Colors.white, fontSize: 16),
                ),
              ),
            ))
      ],
    );
  }
}
